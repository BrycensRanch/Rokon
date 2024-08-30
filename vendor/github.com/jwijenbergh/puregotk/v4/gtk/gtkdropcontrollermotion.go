// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type DropControllerMotionClass struct {
}

func (x *DropControllerMotionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkDropControllerMotion` is an event controller tracking
// the pointer during Drag-and-Drop operations.
//
// It is modeled after [class@Gtk.EventControllerMotion] so if you
// have used that, this should feel really familiar.
//
// This controller is not able to accept drops, use [class@Gtk.DropTarget]
// for that purpose.
type DropControllerMotion struct {
	EventController
}

func DropControllerMotionNewFromInternalPtr(ptr uintptr) *DropControllerMotion {
	cls := &DropControllerMotion{}
	cls.Ptr = ptr
	return cls
}

var xNewDropControllerMotion func() uintptr

// Creates a new event controller that will handle pointer motion
// events during drag and drop.
func NewDropControllerMotion() *DropControllerMotion {
	var cls *DropControllerMotion

	cret := xNewDropControllerMotion()

	if cret == 0 {
		return nil
	}
	cls = &DropControllerMotion{}
	cls.Ptr = cret
	return cls
}

var xDropControllerMotionContainsPointer func(uintptr) bool

// Returns if a Drag-and-Drop operation is within the widget
// @self or one of its children.
func (x *DropControllerMotion) ContainsPointer() bool {

	cret := xDropControllerMotionContainsPointer(x.GoPointer())
	return cret
}

var xDropControllerMotionGetDrop func(uintptr) uintptr

// Returns the `GdkDrop` of a current Drag-and-Drop operation
// over the widget of @self.
func (x *DropControllerMotion) GetDrop() *gdk.Drop {
	var cls *gdk.Drop

	cret := xDropControllerMotionGetDrop(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Drop{}
	cls.Ptr = cret
	return cls
}

var xDropControllerMotionIsPointer func(uintptr) bool

// Returns if a Drag-and-Drop operation is within the widget
// @self, not one of its children.
func (x *DropControllerMotion) IsPointer() bool {

	cret := xDropControllerMotionIsPointer(x.GoPointer())
	return cret
}

func (c *DropControllerMotion) GoPointer() uintptr {
	return c.Ptr
}

func (c *DropControllerMotion) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Signals that the pointer has entered the widget.
func (x *DropControllerMotion) ConnectEnter(cb *func(DropControllerMotion, float64, float64)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "enter", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) {
		fa := DropControllerMotion{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, XVarp, YVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "enter", cbRefPtr)
}

// Signals that the pointer has left the widget.
func (x *DropControllerMotion) ConnectLeave(cb *func(DropControllerMotion)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "leave", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := DropControllerMotion{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "leave", cbRefPtr)
}

// Emitted when the pointer moves inside the widget.
func (x *DropControllerMotion) ConnectMotion(cb *func(DropControllerMotion, float64, float64)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "motion", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) {
		fa := DropControllerMotion{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, XVarp, YVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "motion", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDropControllerMotion, lib, "gtk_drop_controller_motion_new")

	core.PuregoSafeRegister(&xDropControllerMotionContainsPointer, lib, "gtk_drop_controller_motion_contains_pointer")
	core.PuregoSafeRegister(&xDropControllerMotionGetDrop, lib, "gtk_drop_controller_motion_get_drop")
	core.PuregoSafeRegister(&xDropControllerMotionIsPointer, lib, "gtk_drop_controller_motion_is_pointer")

}
