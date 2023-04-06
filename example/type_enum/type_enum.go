package type_enum

import (
	"github.com/dowands/yaenum"
)

type Enum struct {
	Every *yaenum.Instance[Enum] `enum:"every"`
	On    *yaenum.Instance[Enum] `enum:"on"`
}

var EnumList = yaenum.Init[Enum](&Enum{})
