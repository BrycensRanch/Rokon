// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xQsortWithData func(uintptr, int, uint, uintptr, uintptr)

// This is just like the standard C qsort() function, but
// the comparison routine accepts a user data argument.
//
// This is guaranteed to be a stable sort since version 2.32.
func QsortWithData(PbaseVar uintptr, TotalElemsVar int, SizeVar uint, CompareFuncVar *CompareDataFunc, UserDataVar uintptr) {

	xQsortWithData(PbaseVar, TotalElemsVar, SizeVar, NewCallback(CompareFuncVar), UserDataVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xQsortWithData, lib, "g_qsort_with_data")

}
