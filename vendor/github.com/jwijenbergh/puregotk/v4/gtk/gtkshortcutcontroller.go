// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ShortcutControllerClass struct {
}

func (x *ShortcutControllerClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkShortcutController` is an event controller that manages shortcuts.
//
// Most common shortcuts are using this controller implicitly, e.g. by
// adding a mnemonic underline to a `GtkLabel`, or by installing a key
// binding using [method@Gtk.WidgetClass.add_binding], or by adding accelerators
// to global actions using [method@Gtk.Application.set_accels_for_action].
//
// But it is possible to create your own shortcut controller, and add
// shortcuts to it.
//
// `GtkShortcutController` implements `GListModel` for querying the
// shortcuts that have been added to it.
//
// # GtkShortcutController as a GtkBuildable
//
// `GtkShortcutControllers` can be creates in ui files to set up
// shortcuts in the same place as the widgets.
//
// An example of a UI definition fragment with `GtkShortcutController`:
// ```xml
//
//	&lt;object class='GtkButton'&gt;
//	  &lt;child&gt;
//	    &lt;object class='GtkShortcutController'&gt;
//	      &lt;property name='scope'&gt;managed&lt;/property&gt;
//	      &lt;child&gt;
//	        &lt;object class='GtkShortcut'&gt;
//	          &lt;property name='trigger'&gt;&amp;lt;Control&amp;gt;k&lt;/property&gt;
//	          &lt;property name='action'&gt;activate&lt;/property&gt;
//	        &lt;/object&gt;
//	      &lt;/child&gt;
//	    &lt;/object&gt;
//	  &lt;/child&gt;
//	&lt;/object&gt;
//
// ```
//
// This example creates a [class@Gtk.ActivateAction] for triggering the
// `activate` signal of the `GtkButton`. See [ctor@Gtk.ShortcutAction.parse_string]
// for the syntax for other kinds of `GtkShortcutAction`. See
// [ctor@Gtk.ShortcutTrigger.parse_string] to learn more about the syntax
// for triggers.
type ShortcutController struct {
	EventController
}

func ShortcutControllerNewFromInternalPtr(ptr uintptr) *ShortcutController {
	cls := &ShortcutController{}
	cls.Ptr = ptr
	return cls
}

var xNewShortcutController func() uintptr

// Creates a new shortcut controller.
func NewShortcutController() *ShortcutController {
	var cls *ShortcutController

	cret := xNewShortcutController()

	if cret == 0 {
		return nil
	}
	cls = &ShortcutController{}
	cls.Ptr = cret
	return cls
}

var xNewShortcutControllerForModel func(uintptr) uintptr

// Creates a new shortcut controller that takes its shortcuts from
// the given list model.
//
// A controller created by this function does not let you add or
// remove individual shortcuts using the shortcut controller api,
// but you can change the contents of the model.
func NewShortcutControllerForModel(ModelVar gio.ListModel) *ShortcutController {
	var cls *ShortcutController

	cret := xNewShortcutControllerForModel(ModelVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &ShortcutController{}
	cls.Ptr = cret
	return cls
}

var xShortcutControllerAddShortcut func(uintptr, uintptr)

// Adds @shortcut to the list of shortcuts handled by @self.
//
// If this controller uses an external shortcut list, this
// function does nothing.
func (x *ShortcutController) AddShortcut(ShortcutVar *Shortcut) {

	xShortcutControllerAddShortcut(x.GoPointer(), ShortcutVar.GoPointer())

}

var xShortcutControllerGetMnemonicsModifiers func(uintptr) gdk.ModifierType

// Gets the mnemonics modifiers for when this controller activates its shortcuts.
func (x *ShortcutController) GetMnemonicsModifiers() gdk.ModifierType {

	cret := xShortcutControllerGetMnemonicsModifiers(x.GoPointer())
	return cret
}

var xShortcutControllerGetScope func(uintptr) ShortcutScope

// Gets the scope for when this controller activates its shortcuts.
//
// See [method@Gtk.ShortcutController.set_scope] for details.
func (x *ShortcutController) GetScope() ShortcutScope {

	cret := xShortcutControllerGetScope(x.GoPointer())
	return cret
}

var xShortcutControllerRemoveShortcut func(uintptr, uintptr)

// Removes @shortcut from the list of shortcuts handled by @self.
//
// If @shortcut had not been added to @controller or this controller
// uses an external shortcut list, this function does nothing.
func (x *ShortcutController) RemoveShortcut(ShortcutVar *Shortcut) {

	xShortcutControllerRemoveShortcut(x.GoPointer(), ShortcutVar.GoPointer())

}

var xShortcutControllerSetMnemonicsModifiers func(uintptr, gdk.ModifierType)

// Sets the controller to use the given modifier for mnemonics.
//
// The mnemonics modifiers determines which modifiers need to be pressed to allow
// activation of shortcuts with mnemonics triggers.
//
// GTK normally uses the Alt modifier for mnemonics, except in `GtkPopoverMenu`s,
// where mnemonics can be triggered without any modifiers. It should be very
// rarely necessary to change this, and doing so is likely to interfere with
// other shortcuts.
//
// This value is only relevant for local shortcut controllers. Global and managed
// shortcut controllers will have their shortcuts activated from other places which
// have their own modifiers for activating mnemonics.
func (x *ShortcutController) SetMnemonicsModifiers(ModifiersVar gdk.ModifierType) {

	xShortcutControllerSetMnemonicsModifiers(x.GoPointer(), ModifiersVar)

}

var xShortcutControllerSetScope func(uintptr, ShortcutScope)

// Sets the controller to have the given @scope.
//
// The scope allows shortcuts to be activated outside of the normal
// event propagation. In particular, it allows installing global
// keyboard shortcuts that can be activated even when a widget does
// not have focus.
//
// With %GTK_SHORTCUT_SCOPE_LOCAL, shortcuts will only be activated
// when the widget has focus.
func (x *ShortcutController) SetScope(ScopeVar ShortcutScope) {

	xShortcutControllerSetScope(x.GoPointer(), ScopeVar)

}

func (c *ShortcutController) GoPointer() uintptr {
	return c.Ptr
}

func (c *ShortcutController) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// See also: g_list_model_get_n_items()
func (x *ShortcutController) GetItem(PositionVar uint) uintptr {

	cret := gio.XGListModelGetItem(x.GoPointer(), PositionVar)
	return cret
}

// Gets the type of the items in @list.
//
// All items returned from g_list_model_get_item() are of the type
// returned by this function, or a subtype, or if the type is an
// interface, they are an implementation of that interface.
//
// The item type of a #GListModel can not change during the life of the
// model.
func (x *ShortcutController) GetItemType() []interface{} {

	cret := gio.XGListModelGetItemType(x.GoPointer())
	return cret
}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
func (x *ShortcutController) GetNItems() uint {

	cret := gio.XGListModelGetNItems(x.GoPointer())
	return cret
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// This function is meant to be used by language bindings in place
// of g_list_model_get_item().
//
// See also: g_list_model_get_n_items()
func (x *ShortcutController) GetObject(PositionVar uint) *gobject.Object {
	var cls *gobject.Object

	cret := gio.XGListModelGetObject(x.GoPointer(), PositionVar)

	if cret == 0 {
		return nil
	}
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

// Emits the #GListModel::items-changed signal on @list.
//
// This function should only be called by classes implementing
// #GListModel. It has to be called after the internal representation
// of @list has been updated, because handlers connected to this signal
// might query the new state of the list.
//
// Implementations must only make changes to the model (as visible to
// its consumer) in places that will not cause problems for that
// consumer.  For models that are driven directly by a write API (such
// as #GListStore), changes can be reported in response to uses of that
// API.  For models that represent remote data, changes should only be
// made from a fresh mainloop dispatch.  It is particularly not
// permitted to make changes in response to a call to the #GListModel
// consumer API.
//
// Stated another way: in general, it is assumed that code making a
// series of accesses to the model via the API, without returning to the
// mainloop, and without calling other code, will continue to view the
// same contents of the model.
func (x *ShortcutController) ItemsChanged(PositionVar uint, RemovedVar uint, AddedVar uint) {

	gio.XGListModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ShortcutController) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewShortcutController, lib, "gtk_shortcut_controller_new")
	core.PuregoSafeRegister(&xNewShortcutControllerForModel, lib, "gtk_shortcut_controller_new_for_model")

	core.PuregoSafeRegister(&xShortcutControllerAddShortcut, lib, "gtk_shortcut_controller_add_shortcut")
	core.PuregoSafeRegister(&xShortcutControllerGetMnemonicsModifiers, lib, "gtk_shortcut_controller_get_mnemonics_modifiers")
	core.PuregoSafeRegister(&xShortcutControllerGetScope, lib, "gtk_shortcut_controller_get_scope")
	core.PuregoSafeRegister(&xShortcutControllerRemoveShortcut, lib, "gtk_shortcut_controller_remove_shortcut")
	core.PuregoSafeRegister(&xShortcutControllerSetMnemonicsModifiers, lib, "gtk_shortcut_controller_set_mnemonics_modifiers")
	core.PuregoSafeRegister(&xShortcutControllerSetScope, lib, "gtk_shortcut_controller_set_scope")

}
