package main

import (
	"fmt"
	"github.com/dowands/yaenum"
	"github.com/dowands/yaenum/example/type_enum"
)

func main() {
	//check
	on := type_enum.EnumList.On
	val, _ := yaenum.ValueOf(type_enum.EnumList, "on")
	if on != val {
		fmt.Printf("type is mismatch %s %s\n", on.String(), val.String())
	} else {
		fmt.Println("type is match")
	}

}
