// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The virtual function table for #GMemoryMonitor.
type MemoryMonitorInterface struct {
	GIface uintptr
}

func (x *MemoryMonitorInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GMemoryMonitor will monitor system memory and suggest to the application
// when to free memory so as to leave more room for other applications.
// It is implemented on Linux using the [Low Memory Monitor](https://gitlab.freedesktop.org/hadess/low-memory-monitor/)
// ([API documentation](https://hadess.pages.freedesktop.org/low-memory-monitor/)).
//
// There is also an implementation for use inside Flatpak sandboxes.
//
// Possible actions to take when the signal is received are:
//
//   - Free caches
//   - Save files that haven't been looked at in a while to disk, ready to be reopened when needed
//   - Run a garbage collection cycle
//   - Try and compress fragmented allocations
//   - Exit on idle if the process has no reason to stay around
//   - Call [`malloc_trim(3)`](man:malloc_trim) to return cached heap pages to
//     the kernel (if supported by your libc)
//
// Note that some actions may not always improve system performance, and so
// should be profiled for your application. `malloc_trim()`, for example, may
// make future heap allocations slower (due to releasing cached heap pages back
// to the kernel).
//
// See #GMemoryMonitorWarningLevel for details on the various warning levels.
//
// |[&lt;!-- language="C" --&gt;
// static void
// warning_cb (GMemoryMonitor *m, GMemoryMonitorWarningLevel level)
//
//	{
//	  g_debug ("Warning level: %d", level);
//	  if (warning_level &gt; G_MEMORY_MONITOR_WARNING_LEVEL_LOW)
//	    drop_caches ();
//	}
//
// static GMemoryMonitor *
// monitor_low_memory (void)
//
//	{
//	  GMemoryMonitor *m;
//	  m = g_memory_monitor_dup_default ();
//	  g_signal_connect (G_OBJECT (m), "low-memory-warning",
//	                    G_CALLBACK (warning_cb), NULL);
//	  return m;
//	}
//
// ]|
//
// Don't forget to disconnect the #GMemoryMonitor::low-memory-warning
// signal, and unref the #GMemoryMonitor itself when exiting.
type MemoryMonitor interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
}
type MemoryMonitorBase struct {
	Ptr uintptr
}

func (x *MemoryMonitorBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *MemoryMonitorBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

const (
	// Extension point for memory usage monitoring functionality.
	// See [Extending GIO][extending-gio].
	MEMORY_MONITOR_EXTENSION_POINT_NAME string = "gio-memory-monitor"
)

var xMemoryMonitorDupDefault func() uintptr

// Gets a reference to the default #GMemoryMonitor for the system.
func MemoryMonitorDupDefault() *MemoryMonitorBase {
	var cls *MemoryMonitorBase

	cret := xMemoryMonitorDupDefault()

	if cret == 0 {
		return nil
	}
	cls = &MemoryMonitorBase{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xMemoryMonitorDupDefault, lib, "g_memory_monitor_dup_default")

}
