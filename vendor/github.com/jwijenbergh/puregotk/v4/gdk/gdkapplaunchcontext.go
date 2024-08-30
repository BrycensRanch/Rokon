// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// `GdkAppLaunchContext` handles launching an application in a graphical context.
//
// It is an implementation of `GAppLaunchContext` that provides startup
// notification and allows to launch applications on a specific workspace.
//
// ## Launching an application
//
// ```c
// GdkAppLaunchContext *context;
//
// context = gdk_display_get_app_launch_context (display);
//
// gdk_app_launch_context_set_timestamp (gdk_event_get_time (event));
//
// if (!g_app_info_launch_default_for_uri ("http://www.gtk.org", context, &amp;error))
//
//	g_warning ("Launching failed: %s\n", error-&gt;message);
//
// g_object_unref (context);
// ```
type AppLaunchContext struct {
	gio.AppLaunchContext
}

func AppLaunchContextNewFromInternalPtr(ptr uintptr) *AppLaunchContext {
	cls := &AppLaunchContext{}
	cls.Ptr = ptr
	return cls
}

var xAppLaunchContextGetDisplay func(uintptr) uintptr

// Gets the `GdkDisplay` that @context is for.
func (x *AppLaunchContext) GetDisplay() *Display {
	var cls *Display

	cret := xAppLaunchContextGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Display{}
	cls.Ptr = cret
	return cls
}

var xAppLaunchContextSetDesktop func(uintptr, int)

// Sets the workspace on which applications will be launched.
//
// This only works when running under a window manager that
// supports multiple workspaces, as described in the
// [Extended Window Manager Hints](http://www.freedesktop.org/Standards/wm-spec).
// Specifically this sets the `_NET_WM_DESKTOP` property described
// in that spec.
//
// This only works when using the X11 backend.
//
// When the workspace is not specified or @desktop is set to -1,
// it is up to the window manager to pick one, typically it will
// be the current workspace.
func (x *AppLaunchContext) SetDesktop(DesktopVar int) {

	xAppLaunchContextSetDesktop(x.GoPointer(), DesktopVar)

}

var xAppLaunchContextSetIcon func(uintptr, uintptr)

// Sets the icon for applications that are launched with this
// context.
//
// Window Managers can use this information when displaying startup
// notification.
//
// See also [method@Gdk.AppLaunchContext.set_icon_name].
func (x *AppLaunchContext) SetIcon(IconVar gio.Icon) {

	xAppLaunchContextSetIcon(x.GoPointer(), IconVar.GoPointer())

}

var xAppLaunchContextSetIconName func(uintptr, string)

// Sets the icon for applications that are launched with this context.
//
// The @icon_name will be interpreted in the same way as the Icon field
// in desktop files. See also [method@Gdk.AppLaunchContext.set_icon].
//
// If both @icon and @icon_name are set, the @icon_name takes priority.
// If neither @icon or @icon_name is set, the icon is taken from either
// the file that is passed to launched application or from the `GAppInfo`
// for the launched application itself.
func (x *AppLaunchContext) SetIconName(IconNameVar string) {

	xAppLaunchContextSetIconName(x.GoPointer(), IconNameVar)

}

var xAppLaunchContextSetTimestamp func(uintptr, uint32)

// Sets the timestamp of @context.
//
// The timestamp should ideally be taken from the event that
// triggered the launch.
//
// Window managers can use this information to avoid moving the
// focus to the newly launched application when the user is busy
// typing in another window. This is also known as 'focus stealing
// prevention'.
func (x *AppLaunchContext) SetTimestamp(TimestampVar uint32) {

	xAppLaunchContextSetTimestamp(x.GoPointer(), TimestampVar)

}

func (c *AppLaunchContext) GoPointer() uintptr {
	return c.Ptr
}

func (c *AppLaunchContext) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAppLaunchContextGetDisplay, lib, "gdk_app_launch_context_get_display")
	core.PuregoSafeRegister(&xAppLaunchContextSetDesktop, lib, "gdk_app_launch_context_set_desktop")
	core.PuregoSafeRegister(&xAppLaunchContextSetIcon, lib, "gdk_app_launch_context_set_icon")
	core.PuregoSafeRegister(&xAppLaunchContextSetIconName, lib, "gdk_app_launch_context_set_icon_name")
	core.PuregoSafeRegister(&xAppLaunchContextSetTimestamp, lib, "gdk_app_launch_context_set_timestamp")

}
