// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type PasswordEntryBufferClass struct {
	ParentClass uintptr
}

func (x *PasswordEntryBufferClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `GtkEntryBuffer` that locks the underlying memory to prevent it
// from being swapped to disk.
//
// `GtkPasswordEntry` uses a `GtkPasswordEntryBuffer`.
type PasswordEntryBuffer struct {
	EntryBuffer
}

func PasswordEntryBufferNewFromInternalPtr(ptr uintptr) *PasswordEntryBuffer {
	cls := &PasswordEntryBuffer{}
	cls.Ptr = ptr
	return cls
}

var xNewPasswordEntryBuffer func() uintptr

// Creates a new `GtkEntryBuffer` using secure memory allocations.
func NewPasswordEntryBuffer() *PasswordEntryBuffer {
	var cls *PasswordEntryBuffer

	cret := xNewPasswordEntryBuffer()

	if cret == 0 {
		return nil
	}
	cls = &PasswordEntryBuffer{}
	cls.Ptr = cret
	return cls
}

func (c *PasswordEntryBuffer) GoPointer() uintptr {
	return c.Ptr
}

func (c *PasswordEntryBuffer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPasswordEntryBuffer, lib, "gtk_password_entry_buffer_new")

}
