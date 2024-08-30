// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Flags describing the seat capabilities.
type SeatCapabilities int

const (

	// No input capabilities
	SeatCapabilityNoneValue SeatCapabilities = 0
	// The seat has a pointer (e.g. mouse)
	SeatCapabilityPointerValue SeatCapabilities = 1
	// The seat has touchscreen(s) attached
	SeatCapabilityTouchValue SeatCapabilities = 2
	// The seat has drawing tablet(s) attached
	SeatCapabilityTabletStylusValue SeatCapabilities = 4
	// The seat has keyboard(s) attached
	SeatCapabilityKeyboardValue SeatCapabilities = 8
	// The seat has drawing tablet pad(s) attached
	SeatCapabilityTabletPadValue SeatCapabilities = 16
	// The union of all pointing capabilities
	SeatCapabilityAllPointingValue SeatCapabilities = 7
	// The union of all capabilities
	SeatCapabilityAllValue SeatCapabilities = 15
)

// The `GdkSeat` object represents a collection of input devices
// that belong to a user.
type Seat struct {
	gobject.Object
}

func SeatNewFromInternalPtr(ptr uintptr) *Seat {
	cls := &Seat{}
	cls.Ptr = ptr
	return cls
}

var xSeatGetCapabilities func(uintptr) SeatCapabilities

// Returns the capabilities this `GdkSeat` currently has.
func (x *Seat) GetCapabilities() SeatCapabilities {

	cret := xSeatGetCapabilities(x.GoPointer())
	return cret
}

var xSeatGetDevices func(uintptr, SeatCapabilities) *glib.List

// Returns the devices that match the given capabilities.
func (x *Seat) GetDevices(CapabilitiesVar SeatCapabilities) *glib.List {

	cret := xSeatGetDevices(x.GoPointer(), CapabilitiesVar)
	return cret
}

var xSeatGetDisplay func(uintptr) uintptr

// Returns the `GdkDisplay` this seat belongs to.
func (x *Seat) GetDisplay() *Display {
	var cls *Display

	cret := xSeatGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Display{}
	cls.Ptr = cret
	return cls
}

var xSeatGetKeyboard func(uintptr) uintptr

// Returns the device that routes keyboard events.
func (x *Seat) GetKeyboard() *Device {
	var cls *Device

	cret := xSeatGetKeyboard(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Device{}
	cls.Ptr = cret
	return cls
}

var xSeatGetPointer func(uintptr) uintptr

// Returns the device that routes pointer events.
func (x *Seat) GetPointer() *Device {
	var cls *Device

	cret := xSeatGetPointer(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Device{}
	cls.Ptr = cret
	return cls
}

var xSeatGetTools func(uintptr) *glib.List

// Returns all `GdkDeviceTools` that are known to the application.
func (x *Seat) GetTools() *glib.List {

	cret := xSeatGetTools(x.GoPointer())
	return cret
}

func (c *Seat) GoPointer() uintptr {
	return c.Ptr
}

func (c *Seat) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a new input device is related to this seat.
func (x *Seat) ConnectDeviceAdded(cb *func(Seat, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "device-added", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, DeviceVarp uintptr) {
		fa := Seat{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, DeviceVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "device-added", cbRefPtr)
}

// Emitted when an input device is removed (e.g. unplugged).
func (x *Seat) ConnectDeviceRemoved(cb *func(Seat, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "device-removed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, DeviceVarp uintptr) {
		fa := Seat{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, DeviceVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "device-removed", cbRefPtr)
}

// Emitted whenever a new tool is made known to the seat.
//
// The tool may later be assigned to a device (i.e. on
// proximity with a tablet). The device will emit the
// [signal@Gdk.Device::tool-changed] signal accordingly.
//
// A same tool may be used by several devices.
func (x *Seat) ConnectToolAdded(cb *func(Seat, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "tool-added", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ToolVarp uintptr) {
		fa := Seat{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ToolVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "tool-added", cbRefPtr)
}

// Emitted whenever a tool is no longer known to this @seat.
func (x *Seat) ConnectToolRemoved(cb *func(Seat, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "tool-removed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ToolVarp uintptr) {
		fa := Seat{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ToolVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "tool-removed", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSeatGetCapabilities, lib, "gdk_seat_get_capabilities")
	core.PuregoSafeRegister(&xSeatGetDevices, lib, "gdk_seat_get_devices")
	core.PuregoSafeRegister(&xSeatGetDisplay, lib, "gdk_seat_get_display")
	core.PuregoSafeRegister(&xSeatGetKeyboard, lib, "gdk_seat_get_keyboard")
	core.PuregoSafeRegister(&xSeatGetPointer, lib, "gdk_seat_get_pointer")
	core.PuregoSafeRegister(&xSeatGetTools, lib, "gdk_seat_get_tools")

}
