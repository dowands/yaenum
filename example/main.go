package main

import (
	"fmt"
	"yet-another-enum/enum"
	"yet-another-enum/example/type_enum"
)

func main() {
	//check
	on := type_enum.EnumList.On
	val, _ := enum.ValueOf(type_enum.EnumList, "on")
	if on != val {
		fmt.Printf("type is mismatch %s %s\n", on.String(), val.String())
	} else {
		fmt.Println("type is match")
	}
}
