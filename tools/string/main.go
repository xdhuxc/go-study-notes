package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

const (
	// PathParameterKind = indicator of Request parameter type "path"
	PathParameterKind = iota

	// QueryParameterKind = indicator of Request parameter type "query"
	QueryParameterKind

	// BodyParameterKind = indicator of Request parameter type "body"
	BodyParameterKind

	// HeaderParameterKind = indicator of Request parameter type "header"
	HeaderParameterKind

	// FormParameterKind = indicator of Request parameter type "form"
	FormParameterKind

	// CollectionFormatCSV comma separated values `foo,bar`
	CollectionFormatCSV = CollectionFormat("csv")

	// CollectionFormatSSV space separated values `foo bar`
	CollectionFormatSSV = CollectionFormat("ssv")

	// CollectionFormatTSV tab separated values `foo\tbar`
	CollectionFormatTSV = "tsv"

	// CollectionFormatPipes pipe separated values `foo|bar`
	CollectionFormatPipes = CollectionFormat("pipes")

	// CollectionFormatMulti corresponds to multiple parameter instances instead of multiple values for a single
	// instance `foo=bar&foo=baz`. This is valid only for QueryParameters and FormParameters
	CollectionFormatMulti = CollectionFormat("multi")
)

type CollectionFormat string

func (cf CollectionFormat) String() string {
	return string(cf)
}

func main() {
	s := strings.Repeat("abc", 3)
	b := Str2Bytes(s)
	r := Bytes2Str(b)

	fmt.Println(b, r)

	fmt.Println(CollectionFormatTSV)
}
