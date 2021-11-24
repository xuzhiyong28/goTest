package string

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func FieldsDemo() {
	// 按照空格分割
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	// 自定义分割
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("We are humans. We are social animals.", func(r rune) bool {
		return unicode.IsSpace(r) || r == '.'
	}))
}

const BLOG = "http://www.flysnow.org/"

func initStrings(N int) []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

func initStringi(N int) []interface{} {
	s := make([]interface{}, N)
	for i := 0; i < N; i++ {
		s[i] = BLOG
	}
	return s
}

// +号方式拼接
func StringPlus(p []string) string {
	var s string
	l := len(p)
	for i := 0; i < l; i++ {
		s += p[i]
	}
	return s
}

// fmt.Sprint方式拼接
func StringFmt(p []interface{}) string {
	return fmt.Sprint(p...)
}


// strings.Join方式拼接
func StringJoin(p []string) string {
	return strings.Join(p, "")
}

// bytes.Buffer方式拼接
func StringBuffer(p []string) string {
	var b bytes.Buffer
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

// strings.builder方式拼接
func StringBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}
