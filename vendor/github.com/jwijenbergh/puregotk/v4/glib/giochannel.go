// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Specifies the type of function passed to g_io_add_watch() or
// g_io_add_watch_full(), which is called when the requested condition
// on a #GIOChannel is satisfied.
type IOFunc func(*IOChannel, IOCondition, uintptr) bool

// A data structure representing an IO Channel. The fields should be
// considered private and should only be accessed with the following
// functions.
type IOChannel struct {
	RefCount int32

	Funcs *IOFuncs

	Encoding uintptr

	ReadCd uintptr

	WriteCd uintptr

	LineTerm uintptr

	LineTermLen uint

	BufSize uint

	ReadBuf *String

	EncodedReadBuf *String

	WriteBuf *String

	PartialWriteBuf uintptr

	UseBuffer uint

	DoEncode uint

	CloseOnUnref uint

	IsReadable uint

	IsWriteable uint

	IsSeekable uint

	Reserved1 uintptr

	Reserved2 uintptr
}

func (x *IOChannel) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewIOChannelFile func(string, string, **Error) *IOChannel

// Open a file @filename as a #GIOChannel using mode @mode. This
// channel will be closed when the last reference to it is dropped,
// so there is no need to call g_io_channel_close() (though doing
// so will not cause problems, as long as no attempt is made to
// access the channel after it is closed).
func NewIOChannelFile(FilenameVar string, ModeVar string) (*IOChannel, error) {
	var cerr *Error

	cret := xNewIOChannelFile(FilenameVar, ModeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelUnixNew func(int) *IOChannel

// Creates a new #GIOChannel given a file descriptor. On UNIX systems
// this works for plain files, pipes, and sockets.
//
// The returned #GIOChannel has a reference count of 1.
//
// The default encoding for #GIOChannel is UTF-8. If your application
// is reading output from a command using via pipe, you may need to set
// the encoding to the encoding of the current locale (see
// g_get_charset()) with the g_io_channel_set_encoding() function.
// By default, the fd passed will not be closed when the final reference
// to the #GIOChannel data structure is dropped.
//
// If you want to read raw binary data without interpretation, then
// call the g_io_channel_set_encoding() function with %NULL for the
// encoding argument.
//
// This function is available in GLib on Windows, too, but you should
// avoid using it on Windows. The domain of file descriptors and
// sockets overlap. There is no way for GLib to know which one you mean
// in case the argument you pass to this function happens to be both a
// valid file descriptor and socket. If that happens a warning is
// issued, and GLib assumes that it is the file descriptor you mean.
func IOChannelUnixNew(FdVar int) *IOChannel {

	cret := xIOChannelUnixNew(FdVar)
	return cret
}

var xIOChannelClose func(uintptr)

// Close an IO channel. Any pending data to be written will be
// flushed, ignoring errors. The channel will not be freed until the
// last reference is dropped using g_io_channel_unref().
func (x *IOChannel) Close() {

	xIOChannelClose(x.GoPointer())

}

var xIOChannelFlush func(uintptr) IOStatus

// Flushes the write buffer for the GIOChannel.
func (x *IOChannel) Flush() (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelFlush(x.GoPointer())
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelGetBufferCondition func(uintptr) IOCondition

// This function returns a #GIOCondition depending on whether there
// is data to be read/space to write data in the internal buffers in
// the #GIOChannel. Only the flags %G_IO_IN and %G_IO_OUT may be set.
func (x *IOChannel) GetBufferCondition() IOCondition {

	cret := xIOChannelGetBufferCondition(x.GoPointer())
	return cret
}

var xIOChannelGetBufferSize func(uintptr) uint

// Gets the buffer size.
func (x *IOChannel) GetBufferSize() uint {

	cret := xIOChannelGetBufferSize(x.GoPointer())
	return cret
}

var xIOChannelGetBuffered func(uintptr) bool

// Returns whether @channel is buffered.
func (x *IOChannel) GetBuffered() bool {

	cret := xIOChannelGetBuffered(x.GoPointer())
	return cret
}

var xIOChannelGetCloseOnUnref func(uintptr) bool

// Returns whether the file/socket/whatever associated with @channel
// will be closed when @channel receives its final unref and is
// destroyed. The default value of this is %TRUE for channels created
// by g_io_channel_new_file (), and %FALSE for all other channels.
func (x *IOChannel) GetCloseOnUnref() bool {

	cret := xIOChannelGetCloseOnUnref(x.GoPointer())
	return cret
}

var xIOChannelGetEncoding func(uintptr) string

// Gets the encoding for the input/output of the channel.
// The internal encoding is always UTF-8. The encoding %NULL
// makes the channel safe for binary data.
func (x *IOChannel) GetEncoding() string {

	cret := xIOChannelGetEncoding(x.GoPointer())
	return cret
}

var xIOChannelGetFlags func(uintptr) IOFlags

// Gets the current flags for a #GIOChannel, including read-only
// flags such as %G_IO_FLAG_IS_READABLE.
//
// The values of the flags %G_IO_FLAG_IS_READABLE and %G_IO_FLAG_IS_WRITABLE
// are cached for internal use by the channel when it is created.
// If they should change at some later point (e.g. partial shutdown
// of a socket with the UNIX shutdown() function), the user
// should immediately call g_io_channel_get_flags() to update
// the internal values of these flags.
func (x *IOChannel) GetFlags() IOFlags {

	cret := xIOChannelGetFlags(x.GoPointer())
	return cret
}

var xIOChannelGetLineTerm func(uintptr, int) string

// This returns the string that #GIOChannel uses to determine
// where in the file a line break occurs. A value of %NULL
// indicates autodetection.
func (x *IOChannel) GetLineTerm(LengthVar int) string {

	cret := xIOChannelGetLineTerm(x.GoPointer(), LengthVar)
	return cret
}

var xIOChannelInit func(uintptr)

// Initializes a #GIOChannel struct.
//
// This is called by each of the above functions when creating a
// #GIOChannel, and so is not often needed by the application
// programmer (unless you are creating a new type of #GIOChannel).
func (x *IOChannel) Init() {

	xIOChannelInit(x.GoPointer())

}

var xIOChannelRead func(uintptr, string, uint, uint) IOError

// Reads data from a #GIOChannel.
func (x *IOChannel) Read(BufVar string, CountVar uint, BytesReadVar uint) IOError {

	cret := xIOChannelRead(x.GoPointer(), BufVar, CountVar, BytesReadVar)
	return cret
}

var xIOChannelReadChars func(uintptr, uintptr, uint, uint, **Error) IOStatus

// Replacement for g_io_channel_read() with the new API.
func (x *IOChannel) ReadChars(BufVar uintptr, CountVar uint, BytesReadVar uint) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelReadChars(x.GoPointer(), BufVar, CountVar, BytesReadVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelReadLine func(uintptr, string, uint, uint, **Error) IOStatus

// Reads a line, including the terminating character(s),
// from a #GIOChannel into a newly-allocated string.
// @str_return will contain allocated memory if the return
// is %G_IO_STATUS_NORMAL.
func (x *IOChannel) ReadLine(StrReturnVar string, LengthVar uint, TerminatorPosVar uint) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelReadLine(x.GoPointer(), StrReturnVar, LengthVar, TerminatorPosVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelReadLineString func(uintptr, *String, uint, **Error) IOStatus

// Reads a line from a #GIOChannel, using a #GString as a buffer.
func (x *IOChannel) ReadLineString(BufferVar *String, TerminatorPosVar uint) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelReadLineString(x.GoPointer(), BufferVar, TerminatorPosVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelReadToEnd func(uintptr, uintptr, uint, **Error) IOStatus

// Reads all the remaining data from the file.
func (x *IOChannel) ReadToEnd(StrReturnVar uintptr, LengthVar uint) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelReadToEnd(x.GoPointer(), StrReturnVar, LengthVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelReadUnichar func(uintptr, uint32, **Error) IOStatus

// Reads a Unicode character from @channel.
// This function cannot be called on a channel with %NULL encoding.
func (x *IOChannel) ReadUnichar(ThecharVar uint32) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelReadUnichar(x.GoPointer(), ThecharVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelRef func(uintptr) *IOChannel

// Increments the reference count of a #GIOChannel.
func (x *IOChannel) Ref() *IOChannel {

	cret := xIOChannelRef(x.GoPointer())
	return cret
}

var xIOChannelSeek func(uintptr, int64, SeekType) IOError

// Sets the current position in the #GIOChannel, similar to the standard
// library function fseek().
func (x *IOChannel) Seek(OffsetVar int64, TypeVar SeekType) IOError {

	cret := xIOChannelSeek(x.GoPointer(), OffsetVar, TypeVar)
	return cret
}

var xIOChannelSeekPosition func(uintptr, int64, SeekType, **Error) IOStatus

// Replacement for g_io_channel_seek() with the new API.
func (x *IOChannel) SeekPosition(OffsetVar int64, TypeVar SeekType) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelSeekPosition(x.GoPointer(), OffsetVar, TypeVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelSetBufferSize func(uintptr, uint)

// Sets the buffer size.
func (x *IOChannel) SetBufferSize(SizeVar uint) {

	xIOChannelSetBufferSize(x.GoPointer(), SizeVar)

}

var xIOChannelSetBuffered func(uintptr, bool)

// The buffering state can only be set if the channel's encoding
// is %NULL. For any other encoding, the channel must be buffered.
//
// A buffered channel can only be set unbuffered if the channel's
// internal buffers have been flushed. Newly created channels or
// channels which have returned %G_IO_STATUS_EOF
// not require such a flush. For write-only channels, a call to
// g_io_channel_flush () is sufficient. For all other channels,
// the buffers may be flushed by a call to g_io_channel_seek_position ().
// This includes the possibility of seeking with seek type %G_SEEK_CUR
// and an offset of zero. Note that this means that socket-based
// channels cannot be set unbuffered once they have had data
// read from them.
//
// On unbuffered channels, it is safe to mix read and write
// calls from the new and old APIs, if this is necessary for
// maintaining old code.
//
// The default state of the channel is buffered.
func (x *IOChannel) SetBuffered(BufferedVar bool) {

	xIOChannelSetBuffered(x.GoPointer(), BufferedVar)

}

var xIOChannelSetCloseOnUnref func(uintptr, bool)

// Whether to close the channel on the final unref of the #GIOChannel
// data structure. The default value of this is %TRUE for channels
// created by g_io_channel_new_file (), and %FALSE for all other channels.
//
// Setting this flag to %TRUE for a channel you have already closed
// can cause problems when the final reference to the #GIOChannel is dropped.
func (x *IOChannel) SetCloseOnUnref(DoCloseVar bool) {

	xIOChannelSetCloseOnUnref(x.GoPointer(), DoCloseVar)

}

var xIOChannelSetEncoding func(uintptr, string, **Error) IOStatus

// Sets the encoding for the input/output of the channel.
// The internal encoding is always UTF-8. The default encoding
// for the external file is UTF-8.
//
// The encoding %NULL is safe to use with binary data.
//
// The encoding can only be set if one of the following conditions
// is true:
//
// - The channel was just created, and has not been written to or read from yet.
//
// - The channel is write-only.
//
//   - The channel is a file, and the file pointer was just repositioned
//     by a call to g_io_channel_seek_position(). (This flushes all the
//     internal buffers.)
//
// - The current encoding is %NULL or UTF-8.
//
//   - One of the (new API) read functions has just returned %G_IO_STATUS_EOF
//     (or, in the case of g_io_channel_read_to_end(), %G_IO_STATUS_NORMAL).
//
//   - One of the functions g_io_channel_read_chars() or
//     g_io_channel_read_unichar() has returned %G_IO_STATUS_AGAIN or
//     %G_IO_STATUS_ERROR. This may be useful in the case of
//     %G_CONVERT_ERROR_ILLEGAL_SEQUENCE.
//     Returning one of these statuses from g_io_channel_read_line(),
//     g_io_channel_read_line_string(), or g_io_channel_read_to_end()
//     does not guarantee that the encoding can be changed.
//
// Channels which do not meet one of the above conditions cannot call
// g_io_channel_seek_position() with an offset of %G_SEEK_CUR, and, if
// they are "seekable", cannot call g_io_channel_write_chars() after
// calling one of the API "read" functions.
func (x *IOChannel) SetEncoding(EncodingVar string) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelSetEncoding(x.GoPointer(), EncodingVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelSetFlags func(uintptr, IOFlags, **Error) IOStatus

// Sets the (writeable) flags in @channel to (@flags &amp; %G_IO_FLAG_SET_MASK).
func (x *IOChannel) SetFlags(FlagsVar IOFlags) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelSetFlags(x.GoPointer(), FlagsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelSetLineTerm func(uintptr, string, int)

// This sets the string that #GIOChannel uses to determine
// where in the file a line break occurs.
func (x *IOChannel) SetLineTerm(LineTermVar string, LengthVar int) {

	xIOChannelSetLineTerm(x.GoPointer(), LineTermVar, LengthVar)

}

var xIOChannelShutdown func(uintptr, bool, **Error) IOStatus

// Close an IO channel. Any pending data to be written will be
// flushed if @flush is %TRUE. The channel will not be freed until the
// last reference is dropped using g_io_channel_unref().
func (x *IOChannel) Shutdown(FlushVar bool) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelShutdown(x.GoPointer(), FlushVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelUnixGetFd func(uintptr) int

// Returns the file descriptor of the #GIOChannel.
//
// On Windows this function returns the file descriptor or socket of
// the #GIOChannel.
func (x *IOChannel) UnixGetFd() int {

	cret := xIOChannelUnixGetFd(x.GoPointer())
	return cret
}

var xIOChannelUnref func(uintptr)

// Decrements the reference count of a #GIOChannel.
func (x *IOChannel) Unref() {

	xIOChannelUnref(x.GoPointer())

}

var xIOChannelWrite func(uintptr, string, uint, uint) IOError

// Writes data to a #GIOChannel.
func (x *IOChannel) Write(BufVar string, CountVar uint, BytesWrittenVar uint) IOError {

	cret := xIOChannelWrite(x.GoPointer(), BufVar, CountVar, BytesWrittenVar)
	return cret
}

var xIOChannelWriteChars func(uintptr, uintptr, int, uint, **Error) IOStatus

// Replacement for g_io_channel_write() with the new API.
//
// On seekable channels with encodings other than %NULL or UTF-8, generic
// mixing of reading and writing is not allowed. A call to g_io_channel_write_chars ()
// may only be made on a channel from which data has been read in the
// cases described in the documentation for g_io_channel_set_encoding ().
func (x *IOChannel) WriteChars(BufVar uintptr, CountVar int, BytesWrittenVar uint) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelWriteChars(x.GoPointer(), BufVar, CountVar, BytesWrittenVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xIOChannelWriteUnichar func(uintptr, uint32, **Error) IOStatus

// Writes a Unicode character to @channel.
// This function cannot be called on a channel with %NULL encoding.
func (x *IOChannel) WriteUnichar(ThecharVar uint32) (IOStatus, error) {
	var cerr *Error

	cret := xIOChannelWriteUnichar(x.GoPointer(), ThecharVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// A table of functions used to handle different types of #GIOChannel
// in a generic way.
type IOFuncs struct {
}

func (x *IOFuncs) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

const (
	WIN32_MSG_HANDLE int = 19981206
)

// Specifies properties of a #GIOChannel. Some of the flags can only be
// read with g_io_channel_get_flags(), but not changed with
// g_io_channel_set_flags().
type IOFlags int

const (

	// turns on append mode, corresponds to %O_APPEND
	//     (see the documentation of the UNIX open() syscall)
	GIoFlagAppendValue IOFlags = 1
	// turns on nonblocking mode, corresponds to
	//     %O_NONBLOCK/%O_NDELAY (see the documentation of the UNIX open()
	//     syscall)
	GIoFlagNonblockValue IOFlags = 2
	// indicates that the io channel is readable.
	//     This flag cannot be changed.
	GIoFlagIsReadableValue IOFlags = 4
	// indicates that the io channel is writable.
	//     This flag cannot be changed.
	GIoFlagIsWritableValue IOFlags = 8
	// a misspelled version of @G_IO_FLAG_IS_WRITABLE
	//     that existed before the spelling was fixed in GLib 2.30. It is kept
	//     here for compatibility reasons. Deprecated since 2.30
	GIoFlagIsWriteableValue IOFlags = 8
	// indicates that the io channel is seekable,
	//     i.e. that g_io_channel_seek_position() can be used on it.
	//     This flag cannot be changed.
	GIoFlagIsSeekableValue IOFlags = 16
	// the mask that specifies all the valid flags.
	GIoFlagMaskValue IOFlags = 31
	// the mask of the flags that are returned from
	//     g_io_channel_get_flags()
	GIoFlagGetMaskValue IOFlags = 31
	// the mask of the flags that the user can modify
	//     with g_io_channel_set_flags()
	GIoFlagSetMaskValue IOFlags = 3
)

// Error codes returned by #GIOChannel operations.
type IOChannelError int

const (

	// File too large.
	GIoChannelErrorFbigValue IOChannelError = 0
	// Invalid argument.
	GIoChannelErrorInvalValue IOChannelError = 1
	// IO error.
	GIoChannelErrorIoValue IOChannelError = 2
	// File is a directory.
	GIoChannelErrorIsdirValue IOChannelError = 3
	// No space left on device.
	GIoChannelErrorNospcValue IOChannelError = 4
	// No such device or address.
	GIoChannelErrorNxioValue IOChannelError = 5
	// Value too large for defined datatype.
	GIoChannelErrorOverflowValue IOChannelError = 6
	// Broken pipe.
	GIoChannelErrorPipeValue IOChannelError = 7
	// Some other error.
	GIoChannelErrorFailedValue IOChannelError = 8
)

// #GIOError is only used by the deprecated functions
// g_io_channel_read(), g_io_channel_write(), and g_io_channel_seek().
type IOError int

const (

	// no error
	GIoErrorNoneValue IOError = 0
	// an EAGAIN error occurred
	GIoErrorAgainValue IOError = 1
	// an EINVAL error occurred
	GIoErrorInvalValue IOError = 2
	// another error occurred
	GIoErrorUnknownValue IOError = 3
)

// Statuses returned by most of the #GIOFuncs functions.
type IOStatus int

const (

	// An error occurred.
	GIoStatusErrorValue IOStatus = 0
	// Success.
	GIoStatusNormalValue IOStatus = 1
	// End of file.
	GIoStatusEofValue IOStatus = 2
	// Resource temporarily unavailable.
	GIoStatusAgainValue IOStatus = 3
)

// An enumeration specifying the base position for a
// g_io_channel_seek_position() operation.
type SeekType int

const (

	// the current position in the file.
	GSeekCurValue SeekType = 0
	// the start of the file.
	GSeekSetValue SeekType = 1
	// the end of the file.
	GSeekEndValue SeekType = 2
)

var xIoAddWatch func(*IOChannel, IOCondition, uintptr, uintptr) uint

// Adds the #GIOChannel into the default main loop context
// with the default priority.
func IoAddWatch(ChannelVar *IOChannel, ConditionVar IOCondition, FuncVar *IOFunc, UserDataVar uintptr) uint {

	cret := xIoAddWatch(ChannelVar, ConditionVar, NewCallback(FuncVar), UserDataVar)
	return cret
}

var xIoAddWatchFull func(*IOChannel, int, IOCondition, uintptr, uintptr, uintptr) uint

// Adds the #GIOChannel into the default main loop context
// with the given priority.
//
// This internally creates a main loop source using g_io_create_watch()
// and attaches it to the main loop context with g_source_attach().
// You can do these steps manually if you need greater control.
func IoAddWatchFull(ChannelVar *IOChannel, PriorityVar int, ConditionVar IOCondition, FuncVar *IOFunc, UserDataVar uintptr, NotifyVar *DestroyNotify) uint {

	cret := xIoAddWatchFull(ChannelVar, PriorityVar, ConditionVar, NewCallback(FuncVar), UserDataVar, NewCallback(NotifyVar))
	return cret
}

var xIoChannelErrorFromErrno func(int) IOChannelError

// Converts an `errno` error number to a #GIOChannelError.
func IoChannelErrorFromErrno(EnVar int) IOChannelError {

	cret := xIoChannelErrorFromErrno(EnVar)
	return cret
}

var xIoCreateWatch func(*IOChannel, IOCondition) *Source

// Creates a #GSource that's dispatched when @condition is met for the
// given @channel. For example, if condition is %G_IO_IN, the source will
// be dispatched when there's data available for reading.
//
// The callback function invoked by the #GSource should be added with
// g_source_set_callback(), but it has type #GIOFunc (not #GSourceFunc).
//
// g_io_add_watch() is a simpler interface to this same functionality, for
// the case where you want to add the source to the default main loop context
// at the default priority.
//
// On Windows, polling a #GSource created to watch a channel for a socket
// puts the socket in non-blocking mode. This is a side-effect of the
// implementation and unavoidable.
func IoCreateWatch(ChannelVar *IOChannel, ConditionVar IOCondition) *Source {

	cret := xIoCreateWatch(ChannelVar, ConditionVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xIoAddWatch, lib, "g_io_add_watch")
	core.PuregoSafeRegister(&xIoAddWatchFull, lib, "g_io_add_watch_full")
	core.PuregoSafeRegister(&xIoChannelErrorFromErrno, lib, "g_io_channel_error_from_errno")
	core.PuregoSafeRegister(&xIoCreateWatch, lib, "g_io_create_watch")

	core.PuregoSafeRegister(&xNewIOChannelFile, lib, "g_io_channel_new_file")
	core.PuregoSafeRegister(&xIOChannelUnixNew, lib, "g_io_channel_unix_new")

	core.PuregoSafeRegister(&xIOChannelClose, lib, "g_io_channel_close")
	core.PuregoSafeRegister(&xIOChannelFlush, lib, "g_io_channel_flush")
	core.PuregoSafeRegister(&xIOChannelGetBufferCondition, lib, "g_io_channel_get_buffer_condition")
	core.PuregoSafeRegister(&xIOChannelGetBufferSize, lib, "g_io_channel_get_buffer_size")
	core.PuregoSafeRegister(&xIOChannelGetBuffered, lib, "g_io_channel_get_buffered")
	core.PuregoSafeRegister(&xIOChannelGetCloseOnUnref, lib, "g_io_channel_get_close_on_unref")
	core.PuregoSafeRegister(&xIOChannelGetEncoding, lib, "g_io_channel_get_encoding")
	core.PuregoSafeRegister(&xIOChannelGetFlags, lib, "g_io_channel_get_flags")
	core.PuregoSafeRegister(&xIOChannelGetLineTerm, lib, "g_io_channel_get_line_term")
	core.PuregoSafeRegister(&xIOChannelInit, lib, "g_io_channel_init")
	core.PuregoSafeRegister(&xIOChannelRead, lib, "g_io_channel_read")
	core.PuregoSafeRegister(&xIOChannelReadChars, lib, "g_io_channel_read_chars")
	core.PuregoSafeRegister(&xIOChannelReadLine, lib, "g_io_channel_read_line")
	core.PuregoSafeRegister(&xIOChannelReadLineString, lib, "g_io_channel_read_line_string")
	core.PuregoSafeRegister(&xIOChannelReadToEnd, lib, "g_io_channel_read_to_end")
	core.PuregoSafeRegister(&xIOChannelReadUnichar, lib, "g_io_channel_read_unichar")
	core.PuregoSafeRegister(&xIOChannelRef, lib, "g_io_channel_ref")
	core.PuregoSafeRegister(&xIOChannelSeek, lib, "g_io_channel_seek")
	core.PuregoSafeRegister(&xIOChannelSeekPosition, lib, "g_io_channel_seek_position")
	core.PuregoSafeRegister(&xIOChannelSetBufferSize, lib, "g_io_channel_set_buffer_size")
	core.PuregoSafeRegister(&xIOChannelSetBuffered, lib, "g_io_channel_set_buffered")
	core.PuregoSafeRegister(&xIOChannelSetCloseOnUnref, lib, "g_io_channel_set_close_on_unref")
	core.PuregoSafeRegister(&xIOChannelSetEncoding, lib, "g_io_channel_set_encoding")
	core.PuregoSafeRegister(&xIOChannelSetFlags, lib, "g_io_channel_set_flags")
	core.PuregoSafeRegister(&xIOChannelSetLineTerm, lib, "g_io_channel_set_line_term")
	core.PuregoSafeRegister(&xIOChannelShutdown, lib, "g_io_channel_shutdown")
	core.PuregoSafeRegister(&xIOChannelUnixGetFd, lib, "g_io_channel_unix_get_fd")
	core.PuregoSafeRegister(&xIOChannelUnref, lib, "g_io_channel_unref")
	core.PuregoSafeRegister(&xIOChannelWrite, lib, "g_io_channel_write")
	core.PuregoSafeRegister(&xIOChannelWriteChars, lib, "g_io_channel_write_chars")
	core.PuregoSafeRegister(&xIOChannelWriteUnichar, lib, "g_io_channel_write_unichar")

}
