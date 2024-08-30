// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// `GtkAppChooserWidget` is a widget for selecting applications.
//
// It is the main building block for [class@Gtk.AppChooserDialog].
// Most applications only need to use the latter; but you can use
// this widget as part of a larger widget if you have special needs.
//
// `GtkAppChooserWidget` offers detailed control over what applications
// are shown, using the
// [property@Gtk.AppChooserWidget:show-default],
// [property@Gtk.AppChooserWidget:show-recommended],
// [property@Gtk.AppChooserWidget:show-fallback],
// [property@Gtk.AppChooserWidget:show-other] and
// [property@Gtk.AppChooserWidget:show-all] properties. See the
// [iface@Gtk.AppChooser] documentation for more information about these
// groups of applications.
//
// To keep track of the selected application, use the
// [signal@Gtk.AppChooserWidget::application-selected] and
// [signal@Gtk.AppChooserWidget::application-activated] signals.
//
// # CSS nodes
//
// `GtkAppChooserWidget` has a single CSS node with name appchooser.
type AppChooserWidget struct {
	Widget
}

func AppChooserWidgetNewFromInternalPtr(ptr uintptr) *AppChooserWidget {
	cls := &AppChooserWidget{}
	cls.Ptr = ptr
	return cls
}

var xNewAppChooserWidget func(string) uintptr

// Creates a new `GtkAppChooserWidget` for applications
// that can handle content of the given type.
func NewAppChooserWidget(ContentTypeVar string) *AppChooserWidget {
	var cls *AppChooserWidget

	cret := xNewAppChooserWidget(ContentTypeVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &AppChooserWidget{}
	cls.Ptr = cret
	return cls
}

var xAppChooserWidgetGetDefaultText func(uintptr) string

// Returns the text that is shown if there are not applications
// that can handle the content type.
func (x *AppChooserWidget) GetDefaultText() string {

	cret := xAppChooserWidgetGetDefaultText(x.GoPointer())
	return cret
}

var xAppChooserWidgetGetShowAll func(uintptr) bool

// Gets whether the app chooser should show all applications
// in a flat list.
func (x *AppChooserWidget) GetShowAll() bool {

	cret := xAppChooserWidgetGetShowAll(x.GoPointer())
	return cret
}

var xAppChooserWidgetGetShowDefault func(uintptr) bool

// Gets whether the app chooser should show the default handler
// for the content type in a separate section.
func (x *AppChooserWidget) GetShowDefault() bool {

	cret := xAppChooserWidgetGetShowDefault(x.GoPointer())
	return cret
}

var xAppChooserWidgetGetShowFallback func(uintptr) bool

// Gets whether the app chooser should show related applications
// for the content type in a separate section.
func (x *AppChooserWidget) GetShowFallback() bool {

	cret := xAppChooserWidgetGetShowFallback(x.GoPointer())
	return cret
}

var xAppChooserWidgetGetShowOther func(uintptr) bool

// Gets whether the app chooser should show applications
// which are unrelated to the content type.
func (x *AppChooserWidget) GetShowOther() bool {

	cret := xAppChooserWidgetGetShowOther(x.GoPointer())
	return cret
}

var xAppChooserWidgetGetShowRecommended func(uintptr) bool

// Gets whether the app chooser should show recommended applications
// for the content type in a separate section.
func (x *AppChooserWidget) GetShowRecommended() bool {

	cret := xAppChooserWidgetGetShowRecommended(x.GoPointer())
	return cret
}

var xAppChooserWidgetSetDefaultText func(uintptr, string)

// Sets the text that is shown if there are not applications
// that can handle the content type.
func (x *AppChooserWidget) SetDefaultText(TextVar string) {

	xAppChooserWidgetSetDefaultText(x.GoPointer(), TextVar)

}

var xAppChooserWidgetSetShowAll func(uintptr, bool)

// Sets whether the app chooser should show all applications
// in a flat list.
func (x *AppChooserWidget) SetShowAll(SettingVar bool) {

	xAppChooserWidgetSetShowAll(x.GoPointer(), SettingVar)

}

var xAppChooserWidgetSetShowDefault func(uintptr, bool)

// Sets whether the app chooser should show the default handler
// for the content type in a separate section.
func (x *AppChooserWidget) SetShowDefault(SettingVar bool) {

	xAppChooserWidgetSetShowDefault(x.GoPointer(), SettingVar)

}

var xAppChooserWidgetSetShowFallback func(uintptr, bool)

// Sets whether the app chooser should show related applications
// for the content type in a separate section.
func (x *AppChooserWidget) SetShowFallback(SettingVar bool) {

	xAppChooserWidgetSetShowFallback(x.GoPointer(), SettingVar)

}

var xAppChooserWidgetSetShowOther func(uintptr, bool)

// Sets whether the app chooser should show applications
// which are unrelated to the content type.
func (x *AppChooserWidget) SetShowOther(SettingVar bool) {

	xAppChooserWidgetSetShowOther(x.GoPointer(), SettingVar)

}

var xAppChooserWidgetSetShowRecommended func(uintptr, bool)

// Sets whether the app chooser should show recommended applications
// for the content type in a separate section.
func (x *AppChooserWidget) SetShowRecommended(SettingVar bool) {

	xAppChooserWidgetSetShowRecommended(x.GoPointer(), SettingVar)

}

func (c *AppChooserWidget) GoPointer() uintptr {
	return c.Ptr
}

func (c *AppChooserWidget) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when an application item is activated from the widget's list.
//
// This usually happens when the user double clicks an item, or an item
// is selected and the user presses one of the keys Space, Shift+Space,
// Return or Enter.
func (x *AppChooserWidget) ConnectApplicationActivated(cb *func(AppChooserWidget, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "application-activated", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ApplicationVarp uintptr) {
		fa := AppChooserWidget{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ApplicationVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "application-activated", cbRefPtr)
}

// Emitted when an application item is selected from the widget's list.
func (x *AppChooserWidget) ConnectApplicationSelected(cb *func(AppChooserWidget, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "application-selected", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ApplicationVarp uintptr) {
		fa := AppChooserWidget{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ApplicationVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "application-selected", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *AppChooserWidget) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *AppChooserWidget) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *AppChooserWidget) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *AppChooserWidget) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *AppChooserWidget) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AppChooserWidget) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *AppChooserWidget) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AppChooserWidget) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *AppChooserWidget) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AppChooserWidget) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Returns the currently selected application.
func (x *AppChooserWidget) GetAppInfo() *gio.AppInfoBase {
	var cls *gio.AppInfoBase

	cret := XGtkAppChooserGetAppInfo(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gio.AppInfoBase{}
	cls.Ptr = cret
	return cls
}

// Returns the content type for which the `GtkAppChooser`
// shows applications.
func (x *AppChooserWidget) GetContentType() string {

	cret := XGtkAppChooserGetContentType(x.GoPointer())
	return cret
}

// Reloads the list of applications.
func (x *AppChooserWidget) Refresh() {

	XGtkAppChooserRefresh(x.GoPointer())

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *AppChooserWidget) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewAppChooserWidget, lib, "gtk_app_chooser_widget_new")

	core.PuregoSafeRegister(&xAppChooserWidgetGetDefaultText, lib, "gtk_app_chooser_widget_get_default_text")
	core.PuregoSafeRegister(&xAppChooserWidgetGetShowAll, lib, "gtk_app_chooser_widget_get_show_all")
	core.PuregoSafeRegister(&xAppChooserWidgetGetShowDefault, lib, "gtk_app_chooser_widget_get_show_default")
	core.PuregoSafeRegister(&xAppChooserWidgetGetShowFallback, lib, "gtk_app_chooser_widget_get_show_fallback")
	core.PuregoSafeRegister(&xAppChooserWidgetGetShowOther, lib, "gtk_app_chooser_widget_get_show_other")
	core.PuregoSafeRegister(&xAppChooserWidgetGetShowRecommended, lib, "gtk_app_chooser_widget_get_show_recommended")
	core.PuregoSafeRegister(&xAppChooserWidgetSetDefaultText, lib, "gtk_app_chooser_widget_set_default_text")
	core.PuregoSafeRegister(&xAppChooserWidgetSetShowAll, lib, "gtk_app_chooser_widget_set_show_all")
	core.PuregoSafeRegister(&xAppChooserWidgetSetShowDefault, lib, "gtk_app_chooser_widget_set_show_default")
	core.PuregoSafeRegister(&xAppChooserWidgetSetShowFallback, lib, "gtk_app_chooser_widget_set_show_fallback")
	core.PuregoSafeRegister(&xAppChooserWidgetSetShowOther, lib, "gtk_app_chooser_widget_set_show_other")
	core.PuregoSafeRegister(&xAppChooserWidgetSetShowRecommended, lib, "gtk_app_chooser_widget_set_show_recommended")

}
