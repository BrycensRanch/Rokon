// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type BookmarkListClass struct {
	ParentClass uintptr
}

func (x *BookmarkListClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkBookmarkList` is a list model that wraps `GBookmarkFile`.
//
// It presents a `GListModel` and fills it asynchronously with the
// `GFileInfo`s returned from that function.
//
// The `GFileInfo`s in the list have some attributes in the recent
// namespace added: `recent::private` (boolean) and `recent:applications`
// (stringv).
type BookmarkList struct {
	gobject.Object
}

func BookmarkListNewFromInternalPtr(ptr uintptr) *BookmarkList {
	cls := &BookmarkList{}
	cls.Ptr = ptr
	return cls
}

var xNewBookmarkList func(string, string) uintptr

// Creates a new `GtkBookmarkList` with the given @attributes.
func NewBookmarkList(FilenameVar string, AttributesVar string) *BookmarkList {
	var cls *BookmarkList

	cret := xNewBookmarkList(FilenameVar, AttributesVar)

	if cret == 0 {
		return nil
	}
	cls = &BookmarkList{}
	cls.Ptr = cret
	return cls
}

var xBookmarkListGetAttributes func(uintptr) string

// Gets the attributes queried on the children.
func (x *BookmarkList) GetAttributes() string {

	cret := xBookmarkListGetAttributes(x.GoPointer())
	return cret
}

var xBookmarkListGetFilename func(uintptr) string

// Returns the filename of the bookmark file that
// this list is loading.
func (x *BookmarkList) GetFilename() string {

	cret := xBookmarkListGetFilename(x.GoPointer())
	return cret
}

var xBookmarkListGetIoPriority func(uintptr) int

// Gets the IO priority to use while loading file.
func (x *BookmarkList) GetIoPriority() int {

	cret := xBookmarkListGetIoPriority(x.GoPointer())
	return cret
}

var xBookmarkListIsLoading func(uintptr) bool

// Returns %TRUE if the files are currently being loaded.
//
// Files will be added to @self from time to time while loading is
// going on. The order in which are added is undefined and may change
// in between runs.
func (x *BookmarkList) IsLoading() bool {

	cret := xBookmarkListIsLoading(x.GoPointer())
	return cret
}

var xBookmarkListSetAttributes func(uintptr, string)

// Sets the @attributes to be enumerated and starts the enumeration.
//
// If @attributes is %NULL, no attributes will be queried, but a list
// of `GFileInfo`s will still be created.
func (x *BookmarkList) SetAttributes(AttributesVar string) {

	xBookmarkListSetAttributes(x.GoPointer(), AttributesVar)

}

var xBookmarkListSetIoPriority func(uintptr, int)

// Sets the IO priority to use while loading files.
//
// The default IO priority is %G_PRIORITY_DEFAULT.
func (x *BookmarkList) SetIoPriority(IoPriorityVar int) {

	xBookmarkListSetIoPriority(x.GoPointer(), IoPriorityVar)

}

func (c *BookmarkList) GoPointer() uintptr {
	return c.Ptr
}

func (c *BookmarkList) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// See also: g_list_model_get_n_items()
func (x *BookmarkList) GetItem(PositionVar uint) uintptr {

	cret := gio.XGListModelGetItem(x.GoPointer(), PositionVar)
	return cret
}

// Gets the type of the items in @list.
//
// All items returned from g_list_model_get_item() are of the type
// returned by this function, or a subtype, or if the type is an
// interface, they are an implementation of that interface.
//
// The item type of a #GListModel can not change during the life of the
// model.
func (x *BookmarkList) GetItemType() []interface{} {

	cret := gio.XGListModelGetItemType(x.GoPointer())
	return cret
}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
func (x *BookmarkList) GetNItems() uint {

	cret := gio.XGListModelGetNItems(x.GoPointer())
	return cret
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// This function is meant to be used by language bindings in place
// of g_list_model_get_item().
//
// See also: g_list_model_get_n_items()
func (x *BookmarkList) GetObject(PositionVar uint) *gobject.Object {
	var cls *gobject.Object

	cret := gio.XGListModelGetObject(x.GoPointer(), PositionVar)

	if cret == 0 {
		return nil
	}
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

// Emits the #GListModel::items-changed signal on @list.
//
// This function should only be called by classes implementing
// #GListModel. It has to be called after the internal representation
// of @list has been updated, because handlers connected to this signal
// might query the new state of the list.
//
// Implementations must only make changes to the model (as visible to
// its consumer) in places that will not cause problems for that
// consumer.  For models that are driven directly by a write API (such
// as #GListStore), changes can be reported in response to uses of that
// API.  For models that represent remote data, changes should only be
// made from a fresh mainloop dispatch.  It is particularly not
// permitted to make changes in response to a call to the #GListModel
// consumer API.
//
// Stated another way: in general, it is assumed that code making a
// series of accesses to the model via the API, without returning to the
// mainloop, and without calling other code, will continue to view the
// same contents of the model.
func (x *BookmarkList) ItemsChanged(PositionVar uint, RemovedVar uint, AddedVar uint) {

	gio.XGListModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBookmarkList, lib, "gtk_bookmark_list_new")

	core.PuregoSafeRegister(&xBookmarkListGetAttributes, lib, "gtk_bookmark_list_get_attributes")
	core.PuregoSafeRegister(&xBookmarkListGetFilename, lib, "gtk_bookmark_list_get_filename")
	core.PuregoSafeRegister(&xBookmarkListGetIoPriority, lib, "gtk_bookmark_list_get_io_priority")
	core.PuregoSafeRegister(&xBookmarkListIsLoading, lib, "gtk_bookmark_list_is_loading")
	core.PuregoSafeRegister(&xBookmarkListSetAttributes, lib, "gtk_bookmark_list_set_attributes")
	core.PuregoSafeRegister(&xBookmarkListSetIoPriority, lib, "gtk_bookmark_list_set_io_priority")

}
