// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// Interface for implementing operations for mountable volumes.
type VolumeIface struct {
	GIface uintptr
}

func (x *VolumeIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// The #GVolume interface represents user-visible objects that can be
// mounted. Note, when porting from GnomeVFS, #GVolume is the moral
// equivalent of #GnomeVFSDrive.
//
// Mounting a #GVolume instance is an asynchronous operation. For more
// information about asynchronous operations, see #GAsyncResult and
// #GTask. To mount a #GVolume, first call g_volume_mount() with (at
// least) the #GVolume instance, optionally a #GMountOperation object
// and a #GAsyncReadyCallback.
//
// Typically, one will only want to pass %NULL for the
// #GMountOperation if automounting all volumes when a desktop session
// starts since it's not desirable to put up a lot of dialogs asking
// for credentials.
//
// The callback will be fired when the operation has resolved (either
// with success or failure), and a #GAsyncResult instance will be
// passed to the callback.  That callback should then call
// g_volume_mount_finish() with the #GVolume instance and the
// #GAsyncResult data to see if the operation was completed
// successfully.  If an @error is present when g_volume_mount_finish()
// is called, then it will be filled with any error information.
//
// ## Volume Identifiers # {#volume-identifier}
//
// It is sometimes necessary to directly access the underlying
// operating system object behind a volume (e.g. for passing a volume
// to an application via the commandline). For this purpose, GIO
// allows to obtain an 'identifier' for the volume. There can be
// different kinds of identifiers, such as Hal UDIs, filesystem labels,
// traditional Unix devices (e.g. `/dev/sda2`), UUIDs. GIO uses predefined
// strings as names for the different kinds of identifiers:
// %G_VOLUME_IDENTIFIER_KIND_UUID, %G_VOLUME_IDENTIFIER_KIND_LABEL, etc.
// Use g_volume_get_identifier() to obtain an identifier for a volume.
//
// Note that %G_VOLUME_IDENTIFIER_KIND_HAL_UDI will only be available
// when the gvfs hal volume monitor is in use. Other volume monitors
// will generally be able to provide the %G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE
// identifier, which can be used to obtain a hal device by means of
// libhal_manager_find_device_string_match().
type Volume interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	CanEject() bool
	CanMount() bool
	Eject(FlagsVar MountUnmountFlags, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr)
	EjectFinish(ResultVar AsyncResult) bool
	EjectWithOperation(FlagsVar MountUnmountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr)
	EjectWithOperationFinish(ResultVar AsyncResult) bool
	EnumerateIdentifiers() []string
	GetActivationRoot() *FileBase
	GetDrive() *DriveBase
	GetIcon() *IconBase
	GetIdentifier(KindVar string) string
	GetMount() *MountBase
	GetName() string
	GetSortKey() string
	GetSymbolicIcon() *IconBase
	GetUuid() string
	Mount(FlagsVar MountMountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr)
	MountFinish(ResultVar AsyncResult) bool
	ShouldAutomount() bool
}
type VolumeBase struct {
	Ptr uintptr
}

func (x *VolumeBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *VolumeBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Checks if a volume can be ejected.
func (x *VolumeBase) CanEject() bool {

	cret := XGVolumeCanEject(x.GoPointer())
	return cret
}

// Checks if a volume can be mounted.
func (x *VolumeBase) CanMount() bool {

	cret := XGVolumeCanMount(x.GoPointer())
	return cret
}

// Ejects a volume. This is an asynchronous operation, and is
// finished by calling g_volume_eject_finish() with the @volume
// and #GAsyncResult returned in the @callback.
func (x *VolumeBase) Eject(FlagsVar MountUnmountFlags, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGVolumeEject(x.GoPointer(), FlagsVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Finishes ejecting a volume. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
func (x *VolumeBase) EjectFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGVolumeEjectFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Ejects a volume. This is an asynchronous operation, and is
// finished by calling g_volume_eject_with_operation_finish() with the @volume
// and #GAsyncResult data returned in the @callback.
func (x *VolumeBase) EjectWithOperation(FlagsVar MountUnmountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGVolumeEjectWithOperation(x.GoPointer(), FlagsVar, MountOperationVar.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Finishes ejecting a volume. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
func (x *VolumeBase) EjectWithOperationFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGVolumeEjectWithOperationFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Gets the kinds of [identifiers][volume-identifier] that @volume has.
// Use g_volume_get_identifier() to obtain the identifiers themselves.
func (x *VolumeBase) EnumerateIdentifiers() []string {

	cret := XGVolumeEnumerateIdentifiers(x.GoPointer())
	return cret
}

// Gets the activation root for a #GVolume if it is known ahead of
// mount time. Returns %NULL otherwise. If not %NULL and if @volume
// is mounted, then the result of g_mount_get_root() on the
// #GMount object obtained from g_volume_get_mount() will always
// either be equal or a prefix of what this function returns. In
// other words, in code
//
// |[&lt;!-- language="C" --&gt;
//
//	GMount *mount;
//	GFile *mount_root
//	GFile *volume_activation_root;
//
//	mount = g_volume_get_mount (volume); // mounted, so never NULL
//	mount_root = g_mount_get_root (mount);
//	volume_activation_root = g_volume_get_activation_root (volume); // assume not NULL
//
// ]|
// then the expression
// |[&lt;!-- language="C" --&gt;
//
//	(g_file_has_prefix (volume_activation_root, mount_root) ||
//	 g_file_equal (volume_activation_root, mount_root))
//
// ]|
// will always be %TRUE.
//
// Activation roots are typically used in #GVolumeMonitor
// implementations to find the underlying mount to shadow, see
// g_mount_is_shadowed() for more details.
func (x *VolumeBase) GetActivationRoot() *FileBase {
	var cls *FileBase

	cret := XGVolumeGetActivationRoot(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &FileBase{}
	cls.Ptr = cret
	return cls
}

// Gets the drive for the @volume.
func (x *VolumeBase) GetDrive() *DriveBase {
	var cls *DriveBase

	cret := XGVolumeGetDrive(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &DriveBase{}
	cls.Ptr = cret
	return cls
}

// Gets the icon for @volume.
func (x *VolumeBase) GetIcon() *IconBase {
	var cls *IconBase

	cret := XGVolumeGetIcon(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

// Gets the identifier of the given kind for @volume.
// See the [introduction][volume-identifier] for more
// information about volume identifiers.
func (x *VolumeBase) GetIdentifier(KindVar string) string {

	cret := XGVolumeGetIdentifier(x.GoPointer(), KindVar)
	return cret
}

// Gets the mount for the @volume.
func (x *VolumeBase) GetMount() *MountBase {
	var cls *MountBase

	cret := XGVolumeGetMount(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &MountBase{}
	cls.Ptr = cret
	return cls
}

// Gets the name of @volume.
func (x *VolumeBase) GetName() string {

	cret := XGVolumeGetName(x.GoPointer())
	return cret
}

// Gets the sort key for @volume, if any.
func (x *VolumeBase) GetSortKey() string {

	cret := XGVolumeGetSortKey(x.GoPointer())
	return cret
}

// Gets the symbolic icon for @volume.
func (x *VolumeBase) GetSymbolicIcon() *IconBase {
	var cls *IconBase

	cret := XGVolumeGetSymbolicIcon(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &IconBase{}
	cls.Ptr = cret
	return cls
}

// Gets the UUID for the @volume. The reference is typically based on
// the file system UUID for the volume in question and should be
// considered an opaque string. Returns %NULL if there is no UUID
// available.
func (x *VolumeBase) GetUuid() string {

	cret := XGVolumeGetUuid(x.GoPointer())
	return cret
}

// Mounts a volume. This is an asynchronous operation, and is
// finished by calling g_volume_mount_finish() with the @volume
// and #GAsyncResult returned in the @callback.
func (x *VolumeBase) Mount(FlagsVar MountMountFlags, MountOperationVar *MountOperation, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGVolumeMount(x.GoPointer(), FlagsVar, MountOperationVar.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Finishes mounting a volume. If any errors occurred during the operation,
// @error will be set to contain the errors and %FALSE will be returned.
//
// If the mount operation succeeded, g_volume_get_mount() on @volume
// is guaranteed to return the mount right after calling this
// function; there's no need to listen for the 'mount-added' signal on
// #GVolumeMonitor.
func (x *VolumeBase) MountFinish(ResultVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGVolumeMountFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Returns whether the volume should be automatically mounted.
func (x *VolumeBase) ShouldAutomount() bool {

	cret := XGVolumeShouldAutomount(x.GoPointer())
	return cret
}

var XGVolumeCanEject func(uintptr) bool
var XGVolumeCanMount func(uintptr) bool
var XGVolumeEject func(uintptr, MountUnmountFlags, uintptr, uintptr, uintptr)
var XGVolumeEjectFinish func(uintptr, uintptr, **glib.Error) bool
var XGVolumeEjectWithOperation func(uintptr, MountUnmountFlags, uintptr, uintptr, uintptr, uintptr)
var XGVolumeEjectWithOperationFinish func(uintptr, uintptr, **glib.Error) bool
var XGVolumeEnumerateIdentifiers func(uintptr) []string
var XGVolumeGetActivationRoot func(uintptr) uintptr
var XGVolumeGetDrive func(uintptr) uintptr
var XGVolumeGetIcon func(uintptr) uintptr
var XGVolumeGetIdentifier func(uintptr, string) string
var XGVolumeGetMount func(uintptr) uintptr
var XGVolumeGetName func(uintptr) string
var XGVolumeGetSortKey func(uintptr) string
var XGVolumeGetSymbolicIcon func(uintptr) uintptr
var XGVolumeGetUuid func(uintptr) string
var XGVolumeMount func(uintptr, MountMountFlags, uintptr, uintptr, uintptr, uintptr)
var XGVolumeMountFinish func(uintptr, uintptr, **glib.Error) bool
var XGVolumeShouldAutomount func(uintptr) bool

const (
	// The string used to obtain the volume class with g_volume_get_identifier().
	//
	// Known volume classes include `device`, `network`, and `loop`. Other
	// classes may be added in the future.
	//
	// This is intended to be used by applications to classify #GVolume
	// instances into different sections - for example a file manager or
	// file chooser can use this information to show `network` volumes under
	// a "Network" heading and `device` volumes under a "Devices" heading.
	VOLUME_IDENTIFIER_KIND_CLASS string = "class"
	// The string used to obtain a Hal UDI with g_volume_get_identifier().
	VOLUME_IDENTIFIER_KIND_HAL_UDI string = "hal-udi"
	// The string used to obtain a filesystem label with g_volume_get_identifier().
	VOLUME_IDENTIFIER_KIND_LABEL string = "label"
	// The string used to obtain a NFS mount with g_volume_get_identifier().
	VOLUME_IDENTIFIER_KIND_NFS_MOUNT string = "nfs-mount"
	// The string used to obtain a Unix device path with g_volume_get_identifier().
	VOLUME_IDENTIFIER_KIND_UNIX_DEVICE string = "unix-device"
	// The string used to obtain a UUID with g_volume_get_identifier().
	VOLUME_IDENTIFIER_KIND_UUID string = "uuid"
)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGVolumeCanEject, lib, "g_volume_can_eject")
	core.PuregoSafeRegister(&XGVolumeCanMount, lib, "g_volume_can_mount")
	core.PuregoSafeRegister(&XGVolumeEject, lib, "g_volume_eject")
	core.PuregoSafeRegister(&XGVolumeEjectFinish, lib, "g_volume_eject_finish")
	core.PuregoSafeRegister(&XGVolumeEjectWithOperation, lib, "g_volume_eject_with_operation")
	core.PuregoSafeRegister(&XGVolumeEjectWithOperationFinish, lib, "g_volume_eject_with_operation_finish")
	core.PuregoSafeRegister(&XGVolumeEnumerateIdentifiers, lib, "g_volume_enumerate_identifiers")
	core.PuregoSafeRegister(&XGVolumeGetActivationRoot, lib, "g_volume_get_activation_root")
	core.PuregoSafeRegister(&XGVolumeGetDrive, lib, "g_volume_get_drive")
	core.PuregoSafeRegister(&XGVolumeGetIcon, lib, "g_volume_get_icon")
	core.PuregoSafeRegister(&XGVolumeGetIdentifier, lib, "g_volume_get_identifier")
	core.PuregoSafeRegister(&XGVolumeGetMount, lib, "g_volume_get_mount")
	core.PuregoSafeRegister(&XGVolumeGetName, lib, "g_volume_get_name")
	core.PuregoSafeRegister(&XGVolumeGetSortKey, lib, "g_volume_get_sort_key")
	core.PuregoSafeRegister(&XGVolumeGetSymbolicIcon, lib, "g_volume_get_symbolic_icon")
	core.PuregoSafeRegister(&XGVolumeGetUuid, lib, "g_volume_get_uuid")
	core.PuregoSafeRegister(&XGVolumeMount, lib, "g_volume_mount")
	core.PuregoSafeRegister(&XGVolumeMountFinish, lib, "g_volume_mount_finish")
	core.PuregoSafeRegister(&XGVolumeShouldAutomount, lib, "g_volume_should_automount")

}
