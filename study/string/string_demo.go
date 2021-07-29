package string

import (
	"fmt"
	"strings"
	"unicode"
)

func FieldsDemo(){
	// 按照空格分割
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	// 自定义分割
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("We are humans. We are social animals.", func(r rune) bool {
		return unicode.IsSpace(r) || r == '.'
	}))
}
