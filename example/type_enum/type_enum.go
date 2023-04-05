package type_enum

import (
	"github.com/dowands/yaenum"
)

type enumList struct {
	Every *yaenum.Instance `enum:"every"`
	On    *yaenum.Instance `enum:"on"`
}

var EnumList = yaenum.Init(&enumList{})
