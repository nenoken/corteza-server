// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package system

var Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x18\x00	\x000000_access_control.yamlUT\x05\x00\x01\x80Cm8allow:\n  everyone:\n    system:user:\n      - read\n\n    system:application:\n      - read\n\n    system:role:\n      - read\n\n  admins:\n    system:\n      - access\n      - grant\n      - settings.read\n      - settings.manage\n      - organisation.create\n      - application.create\n      - user.create\n      - role.create\n\n    system:application:\n      - read\n      - update\n      - delete\n\n    system:user:\n      - read\n      - update\n      - suspend\n      - unsuspend\n      - delete\n      - unmask.email\n      - unmask.name\n\n    system:role:\n      - read\n      - update\n      - delete\n      - members.manage\nPK\x07\x08\x1c\x8e\xd6\x97W\x02\x00\x00W\x02\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x12\x00	\x000100_settings.yamlUT\x05\x00\x01\x80Cm8settings:\n  privacy.mask.email: true\n  privacy.mask.name: true\n\n  general.mail.logo: data:image/png;base64,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\n  general.mail.header.en: |-\n    <div style=\"width:100%;min-height:100%;margin:0;padding:0;color:#3a393c;font-size:12px;line-height:18px;font-family:Verdana,Arial,sans-serif\">\n      <table width=\"100%\" align=\"center\" style=\"width:100%;height:100%;border-collapse:collapse;border:0;padding:60px\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\" summary=\"\">\n        <tbody>\n          <tr>\n            <td valign=\"top\" align=\"center\" style=\"padding: 20px 0;\">\n              <table width=\"800\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\">\n                <tbody>\n                  <tr>\n                    <td width=\"800\" bgcolor=\"#ffffff\" style=\"color:#3a393c;font-size:14px;line-height:20px;font-family:Helvetica Neue,Helvetica,Arial,sans-serif;text-align:left\">\n                      <table width=\"800\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\">\n                        <tbody>\n                          <tr style=\"background-color:#ffffff;height:50px;\">\n                            <td style=\"border-bottom:2px solid #568ba2;\">\n                              <a href=\"{{ .BaseURL }}\" style=\"text-decoration:none\" target=\"_blank\">\n                                <img src=\"{{ .Logo }}\" style=\"display: block;margin: 0 auto;padding: 10px;\">\n                              </a>\n                            </td>\n                          </tr>\n                          <tr>\n                            <td width=\"800\" style=\"padding:40px 30px\">\n\n  general.mail.footer.en: |-\n    </td>\n                          </tr>\n                          <tr>\n                            <td style=\"padding:30px;border-top: 1px solid #F3F3F5\">\n                              <p>If you have any questions, please contact <a href=\"mailto:{{ .SignatureEmail }}\" style=\"color:#568ba2;\">{{ .SignatureEmail }}</a>.</p>\n                              <p>Kind regards, <br>\n                              {{ .SignatureName }}</p>\n                            </td>\n                          </tr>\n                        </tbody>\n                      </table>\n                    </td>\n                  </tr>\n                </tbody>\n              </table>\n            </td>\n          </tr>\n        </tbody>\n      </table>\n    </div>\n\n  auth.mail.email-confirmation.subject.en: Confirm your email address\n  auth.mail.email-confirmation.body.en: |-\n    {{.EmailHeaderEn}}\n      <h2 style=\"color: #568ba2;text-align: center;\">Confirm your email address</h2>\n      <p>Hello,</p>\n      <p>Follow <a href=\"{{ .URL }}\" style=\"color:#568ba2;\">this link</a> to confirm your email address.</p>\n      <p>You will be logged-in after successful confirmation.</p>\n    {{.EmailFooterEn}}\n\n  auth.mail.password-reset.subject.en: Reset your password\n  auth.mail.password-reset.body.en: |-\n    {{.EmailHeaderEn}}\n      <h2 style=\"color: #568ba2;text-align: center;\">Reset your password</h2>\n      <p>Hello,</p>\n      <p>Follow <a href=\"{{ .URL }}\" style=\"color:#568ba2;\">this link</a> and reset your password.</p>\n      <p>You will be logged-in after successful reset.</p>\n    {{.EmailFooterEn}}\nPK\x07\x08k\xabF\x93\xebE\x00\x00\xebE\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x1c\x8e\xd6\x97W\x02\x00\x00W\x02\x00\x00\x18\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x000000_access_control.yamlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(k\xabF\x93\xebE\x00\x00\xebE\x00\x00\x12\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xa6\x02\x00\x000100_settings.yamlUT\x05\x00\x01\x80Cm8PK\x05\x06\x00\x00\x00\x00\x02\x00\x02\x00\x98\x00\x00\x00\xdaH\x00\x00\x00\x00"
