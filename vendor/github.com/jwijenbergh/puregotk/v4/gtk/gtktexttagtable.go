// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A function used with gtk_text_tag_table_foreach(),
// to iterate over every `GtkTextTag` inside a `GtkTextTagTable`.
type TextTagTableForeach func(uintptr, uintptr)

// The collection of tags in a `GtkTextBuffer`
//
// You may wish to begin by reading the
// [text widget conceptual overview](section-text-widget.html),
// which gives an overview of all the objects and data types
// related to the text widget and how they work together.
//
// # GtkTextTagTables as GtkBuildable
//
// The `GtkTextTagTable` implementation of the `GtkBuildable` interface
// supports adding tags by specifying “tag” as the “type” attribute
// of a &lt;child&gt; element.
//
// An example of a UI definition fragment specifying tags:
// ```xml
// &lt;object class="GtkTextTagTable"&gt;
//
//	&lt;child type="tag"&gt;
//	  &lt;object class="GtkTextTag"/&gt;
//	&lt;/child&gt;
//
// &lt;/object&gt;
// ```
type TextTagTable struct {
	gobject.Object
}

func TextTagTableNewFromInternalPtr(ptr uintptr) *TextTagTable {
	cls := &TextTagTable{}
	cls.Ptr = ptr
	return cls
}

var xNewTextTagTable func() uintptr

// Creates a new `GtkTextTagTable`.
//
// The table contains no tags by default.
func NewTextTagTable() *TextTagTable {
	var cls *TextTagTable

	cret := xNewTextTagTable()

	if cret == 0 {
		return nil
	}
	cls = &TextTagTable{}
	cls.Ptr = cret
	return cls
}

var xTextTagTableAdd func(uintptr, uintptr) bool

// Add a tag to the table.
//
// The tag is assigned the highest priority in the table.
//
// @tag must not be in a tag table already, and may not have
// the same name as an already-added tag.
func (x *TextTagTable) Add(TagVar *TextTag) bool {

	cret := xTextTagTableAdd(x.GoPointer(), TagVar.GoPointer())
	return cret
}

var xTextTagTableForeach func(uintptr, uintptr, uintptr)

// Calls @func on each tag in @table, with user data @data.
//
// Note that the table may not be modified while iterating
// over it (you can’t add/remove tags).
func (x *TextTagTable) Foreach(FuncVar *TextTagTableForeach, DataVar uintptr) {

	xTextTagTableForeach(x.GoPointer(), glib.NewCallback(FuncVar), DataVar)

}

var xTextTagTableGetSize func(uintptr) int

// Returns the size of the table (number of tags)
func (x *TextTagTable) GetSize() int {

	cret := xTextTagTableGetSize(x.GoPointer())
	return cret
}

var xTextTagTableLookup func(uintptr, string) uintptr

// Look up a named tag.
func (x *TextTagTable) Lookup(NameVar string) *TextTag {
	var cls *TextTag

	cret := xTextTagTableLookup(x.GoPointer(), NameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TextTag{}
	cls.Ptr = cret
	return cls
}

var xTextTagTableRemove func(uintptr, uintptr)

// Remove a tag from the table.
//
// If a `GtkTextBuffer` has @table as its tag table, the tag is
// removed from the buffer. The table’s reference to the tag is
// removed, so the tag will end up destroyed if you don’t have
// a reference to it.
func (x *TextTagTable) Remove(TagVar *TextTag) {

	xTextTagTableRemove(x.GoPointer(), TagVar.GoPointer())

}

func (c *TextTagTable) GoPointer() uintptr {
	return c.Ptr
}

func (c *TextTagTable) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted every time a new tag is added in the `GtkTextTagTable`.
func (x *TextTagTable) ConnectTagAdded(cb *func(TextTagTable, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "tag-added", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, TagVarp uintptr) {
		fa := TextTagTable{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, TagVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "tag-added", cbRefPtr)
}

// Emitted every time a tag in the `GtkTextTagTable` changes.
func (x *TextTagTable) ConnectTagChanged(cb *func(TextTagTable, uintptr, bool)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "tag-changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, TagVarp uintptr, SizeChangedVarp bool) {
		fa := TextTagTable{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, TagVarp, SizeChangedVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "tag-changed", cbRefPtr)
}

// Emitted every time a tag is removed from the `GtkTextTagTable`.
//
// The @tag is still valid by the time the signal is emitted, but
// it is not associated with a tag table any more.
func (x *TextTagTable) ConnectTagRemoved(cb *func(TextTagTable, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "tag-removed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, TagVarp uintptr) {
		fa := TextTagTable{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, TagVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "tag-removed", cbRefPtr)
}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *TextTagTable) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTextTagTable, lib, "gtk_text_tag_table_new")

	core.PuregoSafeRegister(&xTextTagTableAdd, lib, "gtk_text_tag_table_add")
	core.PuregoSafeRegister(&xTextTagTableForeach, lib, "gtk_text_tag_table_foreach")
	core.PuregoSafeRegister(&xTextTagTableGetSize, lib, "gtk_text_tag_table_get_size")
	core.PuregoSafeRegister(&xTextTagTableLookup, lib, "gtk_text_tag_table_lookup")
	core.PuregoSafeRegister(&xTextTagTableRemove, lib, "gtk_text_tag_table_remove")

}
