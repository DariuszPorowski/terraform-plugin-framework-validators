package stringvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/orange-cloudavenue/terraform-plugin-framework-validators/stringvalidator/common"
)

/*
IsURN
returns a validator which ensures that the configured attribute
value is a valid URN.

Null (unconfigured) and unknown (known after apply) values are skipped.
*/
func IsURN() validator.String {
	return &common.RegexValidator{
		Desc:         "must be a valid URN",
		Regex:        `(?m)urn:[A-Za-z0-9][A-Za-z0-9-]{0,31}:([A-Za-z0-9()+,\-.:=@;$_!*']|%[0-9A-Fa-f]{2})+`,
		ErrorSummary: "Failed to parse URN",
		ErrorDetail:  "This value is not a valid URN",
	}
}
