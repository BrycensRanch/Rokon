// Package gdkpixbuf was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdkpixbuf

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type PixbufLoaderClass struct {
	ParentClass uintptr
}

func (x *PixbufLoaderClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Incremental image loader.
//
// `GdkPixbufLoader` provides a way for applications to drive the
// process of loading an image, by letting them send the image data
// directly to the loader instead of having the loader read the data
// from a file. Applications can use this functionality instead of
// `gdk_pixbuf_new_from_file()` or `gdk_pixbuf_animation_new_from_file()`
// when they need to parse image data in small chunks. For example,
// it should be used when reading an image from a (potentially) slow
// network connection, or when loading an extremely large file.
//
// To use `GdkPixbufLoader` to load an image, create a new instance,
// and call [method@GdkPixbuf.PixbufLoader.write] to send the data
// to it. When done, [method@GdkPixbuf.PixbufLoader.close] should be
// called to end the stream and finalize everything.
//
// The loader will emit three important signals throughout the process:
//
//   - [signal@GdkPixbuf.PixbufLoader::size-prepared] will be emitted as
//     soon as the image has enough information to determine the size of
//     the image to be used. If you want to scale the image while loading
//     it, you can call [method@GdkPixbuf.PixbufLoader.set_size] in
//     response to this signal.
//   - [signal@GdkPixbuf.PixbufLoader::area-prepared] will be emitted as
//     soon as the pixbuf of the desired has been allocated. You can obtain
//     the `GdkPixbuf` instance by calling [method@GdkPixbuf.PixbufLoader.get_pixbuf].
//     If you want to use it, simply acquire a reference to it. You can
//     also call `gdk_pixbuf_loader_get_pixbuf()` later to get the same
//     pixbuf.
//   - [signal@GdkPixbuf.PixbufLoader::area-updated] will be emitted every
//     time a region is updated. This way you can update a partially
//     completed image. Note that you do not know anything about the
//     completeness of an image from the updated area. For example, in an
//     interlaced image you will need to make several passes before the
//     image is done loading.
//
// ## Loading an animation
//
// Loading an animation is almost as easy as loading an image. Once the
// first [signal@GdkPixbuf.PixbufLoader::area-prepared] signal has been
// emitted, you can call [method@GdkPixbuf.PixbufLoader.get_animation] to
// get the [class@GdkPixbuf.PixbufAnimation] instance, and then call
// and [method@GdkPixbuf.PixbufAnimation.get_iter] to get a
// [class@GdkPixbuf.PixbufAnimationIter] to retrieve the pixbuf for the
// desired time stamp.
type PixbufLoader struct {
	gobject.Object
}

func PixbufLoaderNewFromInternalPtr(ptr uintptr) *PixbufLoader {
	cls := &PixbufLoader{}
	cls.Ptr = ptr
	return cls
}

var xNewPixbufLoader func() uintptr

// Creates a new pixbuf loader object.
func NewPixbufLoader() *PixbufLoader {
	var cls *PixbufLoader

	cret := xNewPixbufLoader()

	if cret == 0 {
		return nil
	}
	cls = &PixbufLoader{}
	cls.Ptr = cret
	return cls
}

var xNewPixbufLoaderWithMimeType func(string, **glib.Error) uintptr

// Creates a new pixbuf loader object that always attempts to parse
// image data as if it were an image of MIME type @mime_type, instead of
// identifying the type automatically.
//
// This function is useful if you want an error if the image isn't the
// expected MIME type; for loading image formats that can't be reliably
// identified by looking at the data; or if the user manually forces a
// specific MIME type.
//
// The list of supported mime types depends on what image loaders
// are installed, but typically "image/png", "image/jpeg", "image/gif",
// "image/tiff" and "image/x-xpixmap" are among the supported mime types.
// To obtain the full list of supported mime types, call
// gdk_pixbuf_format_get_mime_types() on each of the #GdkPixbufFormat
// structs returned by gdk_pixbuf_get_formats().
func NewPixbufLoaderWithMimeType(MimeTypeVar string) (*PixbufLoader, error) {
	var cls *PixbufLoader
	var cerr *glib.Error

	cret := xNewPixbufLoaderWithMimeType(MimeTypeVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &PixbufLoader{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewPixbufLoaderWithType func(string, **glib.Error) uintptr

// Creates a new pixbuf loader object that always attempts to parse
// image data as if it were an image of type @image_type, instead of
// identifying the type automatically.
//
// This function is useful if you want an error if the image isn't the
// expected type; for loading image formats that can't be reliably
// identified by looking at the data; or if the user manually forces
// a specific type.
//
// The list of supported image formats depends on what image loaders
// are installed, but typically "png", "jpeg", "gif", "tiff" and
// "xpm" are among the supported formats. To obtain the full list of
// supported image formats, call gdk_pixbuf_format_get_name() on each
// of the #GdkPixbufFormat structs returned by gdk_pixbuf_get_formats().
func NewPixbufLoaderWithType(ImageTypeVar string) (*PixbufLoader, error) {
	var cls *PixbufLoader
	var cerr *glib.Error

	cret := xNewPixbufLoaderWithType(ImageTypeVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &PixbufLoader{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xPixbufLoaderClose func(uintptr) bool

// Informs a pixbuf loader that no further writes with
// gdk_pixbuf_loader_write() will occur, so that it can free its
// internal loading structures.
//
// This function also tries to parse any data that hasn't yet been parsed;
// if the remaining data is partial or corrupt, an error will be returned.
//
// If `FALSE` is returned, `error` will be set to an error from the
// `GDK_PIXBUF_ERROR` or `G_FILE_ERROR` domains.
//
// If you're just cancelling a load rather than expecting it to be finished,
// passing `NULL` for `error` to ignore it is reasonable.
//
// Remember that this function does not release a reference on the loader, so
// you will need to explicitly release any reference you hold.
func (x *PixbufLoader) Close() (bool, error) {
	var cerr *glib.Error

	cret := xPixbufLoaderClose(x.GoPointer())
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xPixbufLoaderGetAnimation func(uintptr) uintptr

// Queries the #GdkPixbufAnimation that a pixbuf loader is currently creating.
//
// In general it only makes sense to call this function after the
// [signal@GdkPixbuf.PixbufLoader::area-prepared] signal has been emitted by
// the loader.
//
// If the loader doesn't have enough bytes yet, and hasn't emitted the `area-prepared`
// signal, this function will return `NULL`.
func (x *PixbufLoader) GetAnimation() *PixbufAnimation {
	var cls *PixbufAnimation

	cret := xPixbufLoaderGetAnimation(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &PixbufAnimation{}
	cls.Ptr = cret
	return cls
}

var xPixbufLoaderGetFormat func(uintptr) *PixbufFormat

// Obtains the available information about the format of the
// currently loading image file.
func (x *PixbufLoader) GetFormat() *PixbufFormat {

	cret := xPixbufLoaderGetFormat(x.GoPointer())
	return cret
}

var xPixbufLoaderGetPixbuf func(uintptr) uintptr

// Queries the #GdkPixbuf that a pixbuf loader is currently creating.
//
// In general it only makes sense to call this function after the
// [signal@GdkPixbuf.PixbufLoader::area-prepared] signal has been
// emitted by the loader; this means that enough data has been read
// to know the size of the image that will be allocated.
//
// If the loader has not received enough data via gdk_pixbuf_loader_write(),
// then this function returns `NULL`.
//
// The returned pixbuf will be the same in all future calls to the loader,
// so if you want to keep using it, you should acquire a reference to it.
//
// Additionally, if the loader is an animation, it will return the "static
// image" of the animation (see gdk_pixbuf_animation_get_static_image()).
func (x *PixbufLoader) GetPixbuf() *Pixbuf {
	var cls *Pixbuf

	cret := xPixbufLoaderGetPixbuf(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Pixbuf{}
	cls.Ptr = cret
	return cls
}

var xPixbufLoaderSetSize func(uintptr, int, int)

// Causes the image to be scaled while it is loaded.
//
// The desired image size can be determined relative to the original
// size of the image by calling gdk_pixbuf_loader_set_size() from a
// signal handler for the ::size-prepared signal.
//
// Attempts to set the desired image size  are ignored after the
// emission of the ::size-prepared signal.
func (x *PixbufLoader) SetSize(WidthVar int, HeightVar int) {

	xPixbufLoaderSetSize(x.GoPointer(), WidthVar, HeightVar)

}

var xPixbufLoaderWrite func(uintptr, uintptr, uint, **glib.Error) bool

// Parses the next `count` bytes in the given image buffer.
func (x *PixbufLoader) Write(BufVar uintptr, CountVar uint) (bool, error) {
	var cerr *glib.Error

	cret := xPixbufLoaderWrite(x.GoPointer(), BufVar, CountVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xPixbufLoaderWriteBytes func(uintptr, *glib.Bytes, **glib.Error) bool

// Parses the next contents of the given image buffer.
func (x *PixbufLoader) WriteBytes(BufferVar *glib.Bytes) (bool, error) {
	var cerr *glib.Error

	cret := xPixbufLoaderWriteBytes(x.GoPointer(), BufferVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func (c *PixbufLoader) GoPointer() uintptr {
	return c.Ptr
}

func (c *PixbufLoader) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal is emitted when the pixbuf loader has allocated the
// pixbuf in the desired size.
//
// After this signal is emitted, applications can call
// gdk_pixbuf_loader_get_pixbuf() to fetch the partially-loaded
// pixbuf.
func (x *PixbufLoader) ConnectAreaPrepared(cb *func(PixbufLoader)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "area-prepared", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := PixbufLoader{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "area-prepared", cbRefPtr)
}

// This signal is emitted when a significant area of the image being
// loaded has been updated.
//
// Normally it means that a complete scanline has been read in, but
// it could be a different area as well.
//
// Applications can use this signal to know when to repaint
// areas of an image that is being loaded.
func (x *PixbufLoader) ConnectAreaUpdated(cb *func(PixbufLoader, int, int, int, int)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "area-updated", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, XVarp int, YVarp int, WidthVarp int, HeightVarp int) {
		fa := PixbufLoader{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, XVarp, YVarp, WidthVarp, HeightVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "area-updated", cbRefPtr)
}

// This signal is emitted when gdk_pixbuf_loader_close() is called.
//
// It can be used by different parts of an application to receive
// notification when an image loader is closed by the code that
// drives it.
func (x *PixbufLoader) ConnectClosed(cb *func(PixbufLoader)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "closed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := PixbufLoader{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "closed", cbRefPtr)
}

// This signal is emitted when the pixbuf loader has been fed the
// initial amount of data that is required to figure out the size
// of the image that it will create.
//
// Applications can call gdk_pixbuf_loader_set_size() in response
// to this signal to set the desired size to which the image
// should be scaled.
func (x *PixbufLoader) ConnectSizePrepared(cb *func(PixbufLoader, int, int)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "size-prepared", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, WidthVarp int, HeightVarp int) {
		fa := PixbufLoader{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, WidthVarp, HeightVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "size-prepared", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDKPIXBUF"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPixbufLoader, lib, "gdk_pixbuf_loader_new")
	core.PuregoSafeRegister(&xNewPixbufLoaderWithMimeType, lib, "gdk_pixbuf_loader_new_with_mime_type")
	core.PuregoSafeRegister(&xNewPixbufLoaderWithType, lib, "gdk_pixbuf_loader_new_with_type")

	core.PuregoSafeRegister(&xPixbufLoaderClose, lib, "gdk_pixbuf_loader_close")
	core.PuregoSafeRegister(&xPixbufLoaderGetAnimation, lib, "gdk_pixbuf_loader_get_animation")
	core.PuregoSafeRegister(&xPixbufLoaderGetFormat, lib, "gdk_pixbuf_loader_get_format")
	core.PuregoSafeRegister(&xPixbufLoaderGetPixbuf, lib, "gdk_pixbuf_loader_get_pixbuf")
	core.PuregoSafeRegister(&xPixbufLoaderSetSize, lib, "gdk_pixbuf_loader_set_size")
	core.PuregoSafeRegister(&xPixbufLoaderWrite, lib, "gdk_pixbuf_loader_write")
	core.PuregoSafeRegister(&xPixbufLoaderWriteBytes, lib, "gdk_pixbuf_loader_write_bytes")

}
