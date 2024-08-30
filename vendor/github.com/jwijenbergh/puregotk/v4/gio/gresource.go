// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// #GStaticResource is an opaque data structure and can only be accessed
// using the following functions.
type StaticResource struct {
	Data byte

	DataLen uint

	Resource *Resource

	Next *StaticResource

	Padding uintptr
}

func (x *StaticResource) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xStaticResourceFini func(uintptr)

// Finalized a GResource initialized by g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources]
// and is not typically used by other code.
func (x *StaticResource) Fini() {

	xStaticResourceFini(x.GoPointer())

}

var xStaticResourceGetResource func(uintptr) *Resource

// Gets the GResource that was registered by a call to g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources]
// and is not typically used by other code.
func (x *StaticResource) GetResource() *Resource {

	cret := xStaticResourceGetResource(x.GoPointer())
	return cret
}

var xStaticResourceInit func(uintptr)

// Initializes a GResource from static data using a
// GStaticResource.
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources]
// and is not typically used by other code.
func (x *StaticResource) Init() {

	xStaticResourceInit(x.GoPointer())

}

var xResourceLoad func(string, **glib.Error) *Resource

// Loads a binary resource bundle and creates a #GResource representation of it, allowing
// you to query it for data.
//
// If you want to use this resource in the global resource namespace you need
// to register it with g_resources_register().
//
// If @filename is empty or the data in it is corrupt,
// %G_RESOURCE_ERROR_INTERNAL will be returned. If @filename doesn’t exist, or
// there is an error in reading it, an error from g_mapped_file_new() will be
// returned.
func ResourceLoad(FilenameVar string) (*Resource, error) {
	var cerr *glib.Error

	cret := xResourceLoad(FilenameVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xResourcesEnumerateChildren func(string, ResourceLookupFlags, **glib.Error) []string

// Returns all the names of children at the specified @path in the set of
// globally registered resources.
// The return result is a %NULL terminated list of strings which should
// be released with g_strfreev().
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesEnumerateChildren(PathVar string, LookupFlagsVar ResourceLookupFlags) ([]string, error) {
	var cerr *glib.Error

	cret := xResourcesEnumerateChildren(PathVar, LookupFlagsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xResourcesGetInfo func(string, ResourceLookupFlags, uint, uint32, **glib.Error) bool

// Looks for a file at the specified @path in the set of
// globally registered resources and if found returns information about it.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesGetInfo(PathVar string, LookupFlagsVar ResourceLookupFlags, SizeVar uint, FlagsVar uint32) (bool, error) {
	var cerr *glib.Error

	cret := xResourcesGetInfo(PathVar, LookupFlagsVar, SizeVar, FlagsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xResourcesLookupData func(string, ResourceLookupFlags, **glib.Error) *glib.Bytes

// Looks for a file at the specified @path in the set of
// globally registered resources and returns a #GBytes that
// lets you directly access the data in memory.
//
// The data is always followed by a zero byte, so you
// can safely use the data as a C string. However, that byte
// is not included in the size of the GBytes.
//
// For uncompressed resource files this is a pointer directly into
// the resource bundle, which is typically in some readonly data section
// in the program binary. For compressed files we allocate memory on
// the heap and automatically uncompress the data.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesLookupData(PathVar string, LookupFlagsVar ResourceLookupFlags) (*glib.Bytes, error) {
	var cerr *glib.Error

	cret := xResourcesLookupData(PathVar, LookupFlagsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xResourcesOpenStream func(string, ResourceLookupFlags, **glib.Error) uintptr

// Looks for a file at the specified @path in the set of
// globally registered resources and returns a #GInputStream
// that lets you read the data.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesOpenStream(PathVar string, LookupFlagsVar ResourceLookupFlags) (*InputStream, error) {
	var cls *InputStream
	var cerr *glib.Error

	cret := xResourcesOpenStream(PathVar, LookupFlagsVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &InputStream{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xResourcesRegister func(*Resource)

// Registers the resource with the process-global set of resources.
// Once a resource is registered the files in it can be accessed
// with the global resource lookup functions like g_resources_lookup_data().
func ResourcesRegister(ResourceVar *Resource) {

	xResourcesRegister(ResourceVar)

}

var xResourcesUnregister func(*Resource)

// Unregisters the resource from the process-global set of resources.
func ResourcesUnregister(ResourceVar *Resource) {

	xResourcesUnregister(ResourceVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xResourceLoad, lib, "g_resource_load")
	core.PuregoSafeRegister(&xResourcesEnumerateChildren, lib, "g_resources_enumerate_children")
	core.PuregoSafeRegister(&xResourcesGetInfo, lib, "g_resources_get_info")
	core.PuregoSafeRegister(&xResourcesLookupData, lib, "g_resources_lookup_data")
	core.PuregoSafeRegister(&xResourcesOpenStream, lib, "g_resources_open_stream")
	core.PuregoSafeRegister(&xResourcesRegister, lib, "g_resources_register")
	core.PuregoSafeRegister(&xResourcesUnregister, lib, "g_resources_unregister")

	core.PuregoSafeRegister(&xStaticResourceFini, lib, "g_static_resource_fini")
	core.PuregoSafeRegister(&xStaticResourceGetResource, lib, "g_static_resource_get_resource")
	core.PuregoSafeRegister(&xStaticResourceInit, lib, "g_static_resource_init")

}
