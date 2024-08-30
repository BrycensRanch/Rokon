// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A set of functions used to perform memory allocation. The same #GMemVTable must
// be used for all allocations in the same program; a call to g_mem_set_vtable(),
// if it exists, should be prior to any use of GLib.
//
// This functions related to this has been deprecated in 2.46, and no longer work.
type MemVTable struct {
}

func (x *MemVTable) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xAlignedAlloc func(uint, uint, uint) uintptr

// This function is similar to g_malloc(), allocating (@n_blocks * @n_block_bytes)
// bytes, but care is taken to align the allocated memory to with the given
// alignment value. Additionally, it will detect possible overflow during
// multiplication.
//
// Aligned memory allocations returned by this function can only be
// freed using g_aligned_free().
func AlignedAlloc(NBlocksVar uint, NBlockBytesVar uint, AlignmentVar uint) uintptr {

	cret := xAlignedAlloc(NBlocksVar, NBlockBytesVar, AlignmentVar)
	return cret
}

var xAlignedAlloc0 func(uint, uint, uint) uintptr

// This function is similar to g_aligned_alloc(), but it will
// also clear the allocated memory before returning it.
func AlignedAlloc0(NBlocksVar uint, NBlockBytesVar uint, AlignmentVar uint) uintptr {

	cret := xAlignedAlloc0(NBlocksVar, NBlockBytesVar, AlignmentVar)
	return cret
}

var xAlignedFree func(uintptr)

// Frees the memory allocated by g_aligned_alloc().
func AlignedFree(MemVar uintptr) {

	xAlignedFree(MemVar)

}

var xClearPointer func(uintptr, uintptr)

// Clears a reference to a variable.
//
// @pp must not be %NULL.
//
// If the reference is %NULL then this function does nothing.
// Otherwise, the variable is destroyed using @destroy and the
// pointer is set to %NULL.
//
// A macro is also included that allows this function to be used without
// pointer casts. This will mask any warnings about incompatible function types
// or calling conventions, so you must ensure that your @destroy function is
// compatible with being called as `GDestroyNotify` using the standard calling
// convention for the platform that GLib was compiled for; otherwise the program
// will experience undefined behaviour.
func ClearPointer(PpVar uintptr, DestroyVar *DestroyNotify) {

	xClearPointer(PpVar, NewCallback(DestroyVar))

}

var xFree func(uintptr)

// Frees the memory pointed to by @mem.
//
// If @mem is %NULL it simply returns, so there is no need to check @mem
// against %NULL before calling this function.
func Free(MemVar uintptr) {

	xFree(MemVar)

}

var xMalloc func(uint) uintptr

// Allocates @n_bytes bytes of memory.
// If @n_bytes is 0 it returns %NULL.
func Malloc(NBytesVar uint) uintptr {

	cret := xMalloc(NBytesVar)
	return cret
}

var xMalloc0 func(uint) uintptr

// Allocates @n_bytes bytes of memory, initialized to 0's.
// If @n_bytes is 0 it returns %NULL.
func Malloc0(NBytesVar uint) uintptr {

	cret := xMalloc0(NBytesVar)
	return cret
}

var xMalloc0N func(uint, uint) uintptr

// This function is similar to g_malloc0(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
func Malloc0N(NBlocksVar uint, NBlockBytesVar uint) uintptr {

	cret := xMalloc0N(NBlocksVar, NBlockBytesVar)
	return cret
}

var xMallocN func(uint, uint) uintptr

// This function is similar to g_malloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
func MallocN(NBlocksVar uint, NBlockBytesVar uint) uintptr {

	cret := xMallocN(NBlocksVar, NBlockBytesVar)
	return cret
}

var xMemIsSystemMalloc func() bool

// Checks whether the allocator used by g_malloc() is the system's
// malloc implementation. If it returns %TRUE memory allocated with
// malloc() can be used interchangeably with memory allocated using g_malloc().
// This function is useful for avoiding an extra copy of allocated memory returned
// by a non-GLib-based API.
func MemIsSystemMalloc() bool {

	cret := xMemIsSystemMalloc()
	return cret
}

var xMemProfile func()

// GLib used to support some tools for memory profiling, but this
// no longer works. There are many other useful tools for memory
// profiling these days which can be used instead.
func MemProfile() {

	xMemProfile()

}

var xMemSetVtable func(*MemVTable)

// This function used to let you override the memory allocation function.
// However, its use was incompatible with the use of global constructors
// in GLib and GIO, because those use the GLib allocators before main is
// reached. Therefore this function is now deprecated and is just a stub.
func MemSetVtable(VtableVar *MemVTable) {

	xMemSetVtable(VtableVar)

}

var xRealloc func(uintptr, uint) uintptr

// Reallocates the memory pointed to by @mem, so that it now has space for
// @n_bytes bytes of memory. It returns the new address of the memory, which may
// have been moved. @mem may be %NULL, in which case it's considered to
// have zero-length. @n_bytes may be 0, in which case %NULL will be returned
// and @mem will be freed unless it is %NULL.
func Realloc(MemVar uintptr, NBytesVar uint) uintptr {

	cret := xRealloc(MemVar, NBytesVar)
	return cret
}

var xReallocN func(uintptr, uint, uint) uintptr

// This function is similar to g_realloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
func ReallocN(MemVar uintptr, NBlocksVar uint, NBlockBytesVar uint) uintptr {

	cret := xReallocN(MemVar, NBlocksVar, NBlockBytesVar)
	return cret
}

var xTryMalloc func(uint) uintptr

// Attempts to allocate @n_bytes, and returns %NULL on failure.
// Contrast with g_malloc(), which aborts the program on failure.
func TryMalloc(NBytesVar uint) uintptr {

	cret := xTryMalloc(NBytesVar)
	return cret
}

var xTryMalloc0 func(uint) uintptr

// Attempts to allocate @n_bytes, initialized to 0's, and returns %NULL on
// failure. Contrast with g_malloc0(), which aborts the program on failure.
func TryMalloc0(NBytesVar uint) uintptr {

	cret := xTryMalloc0(NBytesVar)
	return cret
}

var xTryMalloc0N func(uint, uint) uintptr

// This function is similar to g_try_malloc0(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
func TryMalloc0N(NBlocksVar uint, NBlockBytesVar uint) uintptr {

	cret := xTryMalloc0N(NBlocksVar, NBlockBytesVar)
	return cret
}

var xTryMallocN func(uint, uint) uintptr

// This function is similar to g_try_malloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
func TryMallocN(NBlocksVar uint, NBlockBytesVar uint) uintptr {

	cret := xTryMallocN(NBlocksVar, NBlockBytesVar)
	return cret
}

var xTryRealloc func(uintptr, uint) uintptr

// Attempts to realloc @mem to a new size, @n_bytes, and returns %NULL
// on failure. Contrast with g_realloc(), which aborts the program
// on failure.
//
// If @mem is %NULL, behaves the same as g_try_malloc().
func TryRealloc(MemVar uintptr, NBytesVar uint) uintptr {

	cret := xTryRealloc(MemVar, NBytesVar)
	return cret
}

var xTryReallocN func(uintptr, uint, uint) uintptr

// This function is similar to g_try_realloc(), allocating (@n_blocks * @n_block_bytes) bytes,
// but care is taken to detect possible overflow during multiplication.
func TryReallocN(MemVar uintptr, NBlocksVar uint, NBlockBytesVar uint) uintptr {

	cret := xTryReallocN(MemVar, NBlocksVar, NBlockBytesVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xAlignedAlloc, lib, "g_aligned_alloc")
	core.PuregoSafeRegister(&xAlignedAlloc0, lib, "g_aligned_alloc0")
	core.PuregoSafeRegister(&xAlignedFree, lib, "g_aligned_free")
	core.PuregoSafeRegister(&xClearPointer, lib, "g_clear_pointer")
	core.PuregoSafeRegister(&xFree, lib, "g_free")
	core.PuregoSafeRegister(&xMalloc, lib, "g_malloc")
	core.PuregoSafeRegister(&xMalloc0, lib, "g_malloc0")
	core.PuregoSafeRegister(&xMalloc0N, lib, "g_malloc0_n")
	core.PuregoSafeRegister(&xMallocN, lib, "g_malloc_n")
	core.PuregoSafeRegister(&xMemIsSystemMalloc, lib, "g_mem_is_system_malloc")
	core.PuregoSafeRegister(&xMemProfile, lib, "g_mem_profile")
	core.PuregoSafeRegister(&xMemSetVtable, lib, "g_mem_set_vtable")
	core.PuregoSafeRegister(&xRealloc, lib, "g_realloc")
	core.PuregoSafeRegister(&xReallocN, lib, "g_realloc_n")
	core.PuregoSafeRegister(&xTryMalloc, lib, "g_try_malloc")
	core.PuregoSafeRegister(&xTryMalloc0, lib, "g_try_malloc0")
	core.PuregoSafeRegister(&xTryMalloc0N, lib, "g_try_malloc0_n")
	core.PuregoSafeRegister(&xTryMallocN, lib, "g_try_malloc_n")
	core.PuregoSafeRegister(&xTryRealloc, lib, "g_try_realloc")
	core.PuregoSafeRegister(&xTryReallocN, lib, "g_try_realloc_n")

}
