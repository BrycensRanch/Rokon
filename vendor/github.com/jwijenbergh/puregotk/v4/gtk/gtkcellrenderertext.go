// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type CellRendererTextClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *CellRendererTextClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Renders text in a cell
//
// A `GtkCellRendererText` renders a given text in its cell, using the font, color and
// style information provided by its properties. The text will be ellipsized if it is
// too long and the `GtkCellRendererText:ellipsize` property allows it.
//
// If the `GtkCellRenderer:mode` is %GTK_CELL_RENDERER_MODE_EDITABLE,
// the `GtkCellRendererText` allows to edit its text using an entry.
type CellRendererText struct {
	CellRenderer
}

func CellRendererTextNewFromInternalPtr(ptr uintptr) *CellRendererText {
	cls := &CellRendererText{}
	cls.Ptr = ptr
	return cls
}

var xNewCellRendererText func() uintptr

// Creates a new `GtkCellRendererText`. Adjust how text is drawn using
// object properties. Object properties can be
// set globally (with g_object_set()). Also, with `GtkTreeViewColumn`,
// you can bind a property to a value in a `GtkTreeModel`. For example,
// you can bind the “text” property on the cell renderer to a string
// value in the model, thus rendering a different string in each row
// of the `GtkTreeView`.
func NewCellRendererText() *CellRendererText {
	var cls *CellRendererText

	cret := xNewCellRendererText()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellRendererText{}
	cls.Ptr = cret
	return cls
}

var xCellRendererTextSetFixedHeightFromFont func(uintptr, int)

// Sets the height of a renderer to explicitly be determined by the “font” and
// “y_pad” property set on it.  Further changes in these properties do not
// affect the height, so they must be accompanied by a subsequent call to this
// function.  Using this function is inflexible, and should really only be used
// if calculating the size of a cell is too slow (ie, a massive number of cells
// displayed).  If @number_of_rows is -1, then the fixed height is unset, and
// the height is determined by the properties again.
func (x *CellRendererText) SetFixedHeightFromFont(NumberOfRowsVar int) {

	xCellRendererTextSetFixedHeightFromFont(x.GoPointer(), NumberOfRowsVar)

}

func (c *CellRendererText) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellRendererText) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal is emitted after @renderer has been edited.
//
// It is the responsibility of the application to update the model
// and store @new_text at the position indicated by @path.
func (x *CellRendererText) ConnectEdited(cb *func(CellRendererText, string, string)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "edited", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, PathVarp string, NewTextVarp string) {
		fa := CellRendererText{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, PathVarp, NewTextVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "edited", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCellRendererText, lib, "gtk_cell_renderer_text_new")

	core.PuregoSafeRegister(&xCellRendererTextSetFixedHeightFromFont, lib, "gtk_cell_renderer_text_set_fixed_height_from_font")

}
