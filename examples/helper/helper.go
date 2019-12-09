package helper

import (
	"fmt"
)

type SomeStruct struct {
	Foo int     `json:"foo"`
	Bar float32 `json:"bar"`
}

func (s SomeStruct) FooBar() string {
	return fmt.Sprintf("%d, %.2f", s.Foo, s.Bar)
}
