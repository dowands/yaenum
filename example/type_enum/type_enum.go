package type_enum

import (
	enum "yet-another-enum/enum"
)

type enumList struct {
	Every *enum.Instance `enum:"every"`
	On    *enum.Instance `enum:"on"`
}

var EnumList = enum.Init(&enumList{})
