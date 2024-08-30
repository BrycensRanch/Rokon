// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type StringListClass struct {
	ParentClass uintptr
}

func (x *StringListClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type StringObjectClass struct {
	ParentClass uintptr
}

func (x *StringObjectClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkStringList` is a list model that wraps an array of strings.
//
// The objects in the model have a "string" property.
//
// `GtkStringList` is well-suited for any place where you would
// typically use a `char*[]`, but need a list model.
//
// # GtkStringList as GtkBuildable
//
// The `GtkStringList` implementation of the `GtkBuildable` interface
// supports adding items directly using the &lt;items&gt; element and
// specifying &lt;item&gt; elements for each item. Each &lt;item&gt; element
// supports the regular translation attributes “translatable”,
// “context” and “comments”.
//
// Here is a UI definition fragment specifying a `GtkStringList`
//
// ```xml
// &lt;object class="GtkStringList"&gt;
//
//	&lt;items&gt;
//	  &lt;item translatable="yes"&gt;Factory&lt;/item&gt;
//	  &lt;item translatable="yes"&gt;Home&lt;/item&gt;
//	  &lt;item translatable="yes"&gt;Subway&lt;/item&gt;
//	&lt;/items&gt;
//
// &lt;/object&gt;
// ```
type StringList struct {
	gobject.Object
}

func StringListNewFromInternalPtr(ptr uintptr) *StringList {
	cls := &StringList{}
	cls.Ptr = ptr
	return cls
}

var xNewStringList func(uintptr) uintptr

// Creates a new `GtkStringList` with the given @strings.
func NewStringList(StringsVar uintptr) *StringList {
	var cls *StringList

	cret := xNewStringList(StringsVar)

	if cret == 0 {
		return nil
	}
	cls = &StringList{}
	cls.Ptr = cret
	return cls
}

var xStringListAppend func(uintptr, string)

// Appends @string to @self.
//
// The @string will be copied. See
// [method@Gtk.StringList.take] for a way to avoid that.
func (x *StringList) Append(StringVar string) {

	xStringListAppend(x.GoPointer(), StringVar)

}

var xStringListGetString func(uintptr, uint) string

// Gets the string that is at @position in @self.
//
// If @self does not contain @position items, %NULL is returned.
//
// This function returns the const char *. To get the
// object wrapping it, use g_list_model_get_item().
func (x *StringList) GetString(PositionVar uint) string {

	cret := xStringListGetString(x.GoPointer(), PositionVar)
	return cret
}

var xStringListRemove func(uintptr, uint)

// Removes the string at @position from @self.
//
// @position must be smaller than the current
// length of the list.
func (x *StringList) Remove(PositionVar uint) {

	xStringListRemove(x.GoPointer(), PositionVar)

}

var xStringListSplice func(uintptr, uint, uint, uintptr)

// Changes @self by removing @n_removals strings and adding @additions
// to it.
//
// This function is more efficient than [method@Gtk.StringList.append]
// and [method@Gtk.StringList.remove], because it only emits the
// ::items-changed signal once for the change.
//
// This function copies the strings in @additions.
//
// The parameters @position and @n_removals must be correct (ie:
// @position + @n_removals must be less than or equal to the length
// of the list at the time this function is called).
func (x *StringList) Splice(PositionVar uint, NRemovalsVar uint, AdditionsVar uintptr) {

	xStringListSplice(x.GoPointer(), PositionVar, NRemovalsVar, AdditionsVar)

}

var xStringListTake func(uintptr, string)

// Adds @string to self at the end, and takes
// ownership of it.
//
// This variant of [method@Gtk.StringList.append]
// is convenient for formatting strings:
//
// ```c
// gtk_string_list_take (self, g_strdup_print ("%d dollars", lots));
// ```
func (x *StringList) Take(StringVar string) {

	xStringListTake(x.GoPointer(), StringVar)

}

func (c *StringList) GoPointer() uintptr {
	return c.Ptr
}

func (c *StringList) SetGoPointer(ptr uintptr) {
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
func (x *StringList) GetItem(PositionVar uint) uintptr {

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
func (x *StringList) GetItemType() []interface{} {

	cret := gio.XGListModelGetItemType(x.GoPointer())
	return cret
}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
func (x *StringList) GetNItems() uint {

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
func (x *StringList) GetObject(PositionVar uint) *gobject.Object {
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
func (x *StringList) ItemsChanged(PositionVar uint, RemovedVar uint, AddedVar uint) {

	gio.XGListModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *StringList) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// `GtkStringObject` is the type of items in a `GtkStringList`.
//
// A `GtkStringObject` is a wrapper around a `const char*`; it has
// a [property@Gtk.StringObject:string] property.
type StringObject struct {
	gobject.Object
}

func StringObjectNewFromInternalPtr(ptr uintptr) *StringObject {
	cls := &StringObject{}
	cls.Ptr = ptr
	return cls
}

var xNewStringObject func(string) uintptr

// Wraps a string in an object for use with `GListModel`.
func NewStringObject(StringVar string) *StringObject {
	var cls *StringObject

	cret := xNewStringObject(StringVar)

	if cret == 0 {
		return nil
	}
	cls = &StringObject{}
	cls.Ptr = cret
	return cls
}

var xStringObjectGetString func(uintptr) string

// Returns the string contained in a `GtkStringObject`.
func (x *StringObject) GetString() string {

	cret := xStringObjectGetString(x.GoPointer())
	return cret
}

func (c *StringObject) GoPointer() uintptr {
	return c.Ptr
}

func (c *StringObject) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewStringList, lib, "gtk_string_list_new")

	core.PuregoSafeRegister(&xStringListAppend, lib, "gtk_string_list_append")
	core.PuregoSafeRegister(&xStringListGetString, lib, "gtk_string_list_get_string")
	core.PuregoSafeRegister(&xStringListRemove, lib, "gtk_string_list_remove")
	core.PuregoSafeRegister(&xStringListSplice, lib, "gtk_string_list_splice")
	core.PuregoSafeRegister(&xStringListTake, lib, "gtk_string_list_take")

	core.PuregoSafeRegister(&xNewStringObject, lib, "gtk_string_object_new")

	core.PuregoSafeRegister(&xStringObjectGetString, lib, "gtk_string_object_get_string")

}
