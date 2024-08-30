// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type GestureSwipeClass struct {
}

func (x *GestureSwipeClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkGestureSwipe` is a `GtkGesture` for swipe gestures.
//
// After a press/move/.../move/release sequence happens, the
// [signal@Gtk.GestureSwipe::swipe] signal will be emitted,
// providing the velocity and directionality of the sequence
// at the time it was lifted.
//
// If the velocity is desired in intermediate points,
// [method@Gtk.GestureSwipe.get_velocity] can be called in a
// [signal@Gtk.Gesture::update] handler.
//
// All velocities are reported in pixels/sec units.
type GestureSwipe struct {
	GestureSingle
}

func GestureSwipeNewFromInternalPtr(ptr uintptr) *GestureSwipe {
	cls := &GestureSwipe{}
	cls.Ptr = ptr
	return cls
}

var xNewGestureSwipe func() uintptr

// Returns a newly created `GtkGesture` that recognizes swipes.
func NewGestureSwipe() *GestureSwipe {
	var cls *GestureSwipe

	cret := xNewGestureSwipe()

	if cret == 0 {
		return nil
	}
	cls = &GestureSwipe{}
	cls.Ptr = cret
	return cls
}

var xGestureSwipeGetVelocity func(uintptr, float64, float64) bool

// Gets the current velocity.
//
// If the gesture is recognized, this function returns %TRUE and fills
// in @velocity_x and @velocity_y with the recorded velocity, as per the
// last events processed.
func (x *GestureSwipe) GetVelocity(VelocityXVar float64, VelocityYVar float64) bool {

	cret := xGestureSwipeGetVelocity(x.GoPointer(), VelocityXVar, VelocityYVar)
	return cret
}

func (c *GestureSwipe) GoPointer() uintptr {
	return c.Ptr
}

func (c *GestureSwipe) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the recognized gesture is finished.
//
// Velocity and direction are a product of previously recorded events.
func (x *GestureSwipe) ConnectSwipe(cb *func(GestureSwipe, float64, float64)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "swipe", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, VelocityXVarp float64, VelocityYVarp float64) {
		fa := GestureSwipe{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, VelocityXVarp, VelocityYVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "swipe", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewGestureSwipe, lib, "gtk_gesture_swipe_new")

	core.PuregoSafeRegister(&xGestureSwipeGetVelocity, lib, "gtk_gesture_swipe_get_velocity")

}
