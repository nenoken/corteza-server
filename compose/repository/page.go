package repository

import (
	"context"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/titpetric/factory"

	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/rh"
)

type (
	PageRepository interface {
		With(ctx context.Context, db *factory.DB) PageRepository

		FindByID(namespaceID, pageID uint64) (*types.Page, error)
		FindByHandle(namespaceID uint64, handle string) (*types.Page, error)
		FindByModuleID(namespaceID, moduleID uint64) (*types.Page, error)
		Find(filter types.PageFilter) (set types.PageSet, f types.PageFilter, err error)

		Create(mod *types.Page) (*types.Page, error)
		Update(mod *types.Page) (*types.Page, error)
		DeleteByID(namespaceID, pageID uint64) error

		Reorder(namespaceID, selfID uint64, pageIDs []uint64) error
	}

	page struct {
		*repository
	}
)

const (
	ErrPageNotFound        = repositoryError("PageNotFound")
	ErrPageHandleNotUnique = repositoryError("PageHandleNotUnique")
)

func Page(ctx context.Context, db *factory.DB) PageRepository {
	return (&page{}).With(ctx, db)
}

func (r page) With(ctx context.Context, db *factory.DB) PageRepository {
	return &page{
		repository: r.repository.With(ctx, db),
	}
}

func (r page) table() string {
	return "compose_page"
}

func (r page) columns() []string {
	return []string{
		"id",
		"rel_namespace",
		"self_id",
		"rel_module",
		"handle",
		"title",
		"blocks",
		"description",
		"visible",
		"weight",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func (r page) query() squirrel.SelectBuilder {
	return squirrel.
		Select(r.columns()...).
		From(r.table()).
		Where("deleted_at IS NULL")
}

func (r page) FindByID(namespaceID, pageID uint64) (*types.Page, error) {
	return r.findOneBy(namespaceID, "id", pageID)
}

func (r page) FindByHandle(namespaceID uint64, handle string) (*types.Page, error) {
	return r.findOneBy(namespaceID, "handle", handle)
}

func (r page) FindByModuleID(namespaceID, moduleID uint64) (*types.Page, error) {
	return r.findOneBy(namespaceID, "rel_module", moduleID)
}

func (r page) findOneBy(namespaceID uint64, field string, value interface{}) (*types.Page, error) {
	var (
		p = &types.Page{}

		q = r.query().
			Where(squirrel.Eq{field: value, "rel_namespace": namespaceID})

		err = rh.FetchOne(r.db(), q, p)
	)

	if err != nil {
		return nil, err
	} else if p.ID == 0 {
		return nil, ErrPageNotFound
	}

	return p, nil
}

func (r page) Find(filter types.PageFilter) (set types.PageSet, f types.PageFilter, err error) {
	f = filter

	if f.Sort == "" {
		f.Sort = "id ASC"
	}

	query := r.query()

	if filter.NamespaceID > 0 {
		query = query.Where("rel_namespace = ?", filter.NamespaceID)
	}

	if filter.ParentID > 0 {
		query = query.Where("self_id = ?", filter.ParentID)
	} else if filter.Root {
		query = query.Where("self_id = 0")
	}

	if f.Handle != "" {
		query = query.Where("LOWER(handle) = ?", strings.ToLower(f.Handle))
	}

	if f.Query != "" {
		q := "%" + strings.ToLower(f.Query) + "%"
		query = query.Where(squirrel.Or{
			squirrel.Like{"LOWER(title)": q},
			squirrel.Like{"LOWER(description)": q},
		})
	}

	if f.IsReadable != nil {
		query = query.Where(f.IsReadable)
	}

	var orderBy []string
	if orderBy, err = rh.ParseOrder(f.Sort, r.columns()...); err != nil {
		return
	} else {
		query = query.OrderBy(orderBy...)
	}

	if f.Count, err = rh.Count(r.db(), query); err != nil || f.Count == 0 {
		return
	}

	return set, f, rh.FetchPaged(r.db(), query, f.Page, f.PerPage, &set)
}

func (r page) Reorder(namespaceID, parentID uint64, pageIDs []uint64) error {
	var (
		pageMap = map[uint64]bool{}
		filter  = types.PageFilter{NamespaceID: namespaceID, ParentID: parentID}
	)

	if pages, _, err := r.Find(filter); err != nil {
		return nil
	} else {
		for _, page := range pages {
			pageMap[page.ID] = true
		}
	}
	weight := 1
	db := r.db()
	// honor parameter first
	for _, pageID := range pageIDs {
		if pageMap[pageID] {
			pageMap[pageID] = false
			if _, err := db.Exec("UPDATE compose_page SET weight = ? WHERE id = ? AND self_id = ?", weight, pageID, parentID); err != nil {
				return err
			}
			weight++
		}
	}
	for pageID, update := range pageMap {
		if update {
			if _, err := db.Exec("UPDATE compose_page SET weight = ? WHERE id = ? AND self_id = ?", weight, pageID, parentID); err != nil {
				return err
			}
			weight++
		}
	}
	return nil
}

func (r page) Create(mod *types.Page) (*types.Page, error) {
	mod.ID = factory.Sonyflake.NextID()
	rh.SetCurrentTimeRounded(&mod.CreatedAt)
	mod.UpdatedAt = nil

	return mod, r.db().Insert(r.table(), mod)
}

func (r page) Update(mod *types.Page) (*types.Page, error) {
	rh.SetCurrentTimeRounded(&mod.UpdatedAt)

	return mod, r.db().Update(r.table(), mod, "id")
}

func (r page) DeleteByID(namespaceID, pageID uint64) error {
	_, err := r.db().Exec(
		"UPDATE "+r.table()+" SET deleted_at = NOW() WHERE rel_namespace = ? AND id = ?",
		namespaceID,
		pageID,
	)

	return err
}
