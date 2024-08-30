// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for #GSocketControlMessage.
type SocketControlMessageClass struct {
	ParentClass uintptr
}

func (x *SocketControlMessageClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type SocketControlMessagePrivate struct {
}

func (x *SocketControlMessagePrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A #GSocketControlMessage is a special-purpose utility message that
// can be sent to or received from a #GSocket. These types of
// messages are often called "ancillary data".
//
// The message can represent some sort of special instruction to or
// information from the socket or can represent a special kind of
// transfer to the peer (for example, sending a file descriptor over
// a UNIX socket).
//
// These messages are sent with g_socket_send_message() and received
// with g_socket_receive_message().
//
// To extend the set of control message that can be sent, subclass this
// class and override the get_size, get_level, get_type and serialize
// methods.
//
// To extend the set of control messages that can be received, subclass
// this class and implement the deserialize method. Also, make sure your
// class is registered with the GType typesystem before calling
// g_socket_receive_message() to read such a message.
type SocketControlMessage struct {
	gobject.Object
}

func SocketControlMessageNewFromInternalPtr(ptr uintptr) *SocketControlMessage {
	cls := &SocketControlMessage{}
	cls.Ptr = ptr
	return cls
}

var xSocketControlMessageGetLevel func(uintptr) int

// Returns the "level" (i.e. the originating protocol) of the control message.
// This is often SOL_SOCKET.
func (x *SocketControlMessage) GetLevel() int {

	cret := xSocketControlMessageGetLevel(x.GoPointer())
	return cret
}

var xSocketControlMessageGetMsgType func(uintptr) int

// Returns the protocol specific type of the control message.
// For instance, for UNIX fd passing this would be SCM_RIGHTS.
func (x *SocketControlMessage) GetMsgType() int {

	cret := xSocketControlMessageGetMsgType(x.GoPointer())
	return cret
}

var xSocketControlMessageGetSize func(uintptr) uint

// Returns the space required for the control message, not including
// headers or alignment.
func (x *SocketControlMessage) GetSize() uint {

	cret := xSocketControlMessageGetSize(x.GoPointer())
	return cret
}

var xSocketControlMessageSerialize func(uintptr, uintptr)

// Converts the data in the message to bytes placed in the
// message.
//
// @data is guaranteed to have enough space to fit the size
// returned by g_socket_control_message_get_size() on this
// object.
func (x *SocketControlMessage) Serialize(DataVar uintptr) {

	xSocketControlMessageSerialize(x.GoPointer(), DataVar)

}

func (c *SocketControlMessage) GoPointer() uintptr {
	return c.Ptr
}

func (c *SocketControlMessage) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

var xSocketControlMessageDeserialize func(int, int, uint, uintptr) uintptr

// Tries to deserialize a socket control message of a given
// @level and @type. This will ask all known (to GType) subclasses
// of #GSocketControlMessage if they can understand this kind
// of message and if so deserialize it into a #GSocketControlMessage.
//
// If there is no implementation for this kind of control message, %NULL
// will be returned.
func SocketControlMessageDeserialize(LevelVar int, TypeVar int, SizeVar uint, DataVar uintptr) *SocketControlMessage {
	var cls *SocketControlMessage

	cret := xSocketControlMessageDeserialize(LevelVar, TypeVar, SizeVar, DataVar)

	if cret == 0 {
		return nil
	}
	cls = &SocketControlMessage{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSocketControlMessageGetLevel, lib, "g_socket_control_message_get_level")
	core.PuregoSafeRegister(&xSocketControlMessageGetMsgType, lib, "g_socket_control_message_get_msg_type")
	core.PuregoSafeRegister(&xSocketControlMessageGetSize, lib, "g_socket_control_message_get_size")
	core.PuregoSafeRegister(&xSocketControlMessageSerialize, lib, "g_socket_control_message_serialize")

	core.PuregoSafeRegister(&xSocketControlMessageDeserialize, lib, "g_socket_control_message_deserialize")

}
