// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type RangeClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *RangeClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkRange` is the common base class for widgets which visualize an
// adjustment.
//
// Widgets that are derived from `GtkRange` include
// [class@Gtk.Scale] and [class@Gtk.Scrollbar].
//
// Apart from signals for monitoring the parameters of the adjustment,
// `GtkRange` provides properties and methods for setting a
// “fill level” on range widgets. See [method@Gtk.Range.set_fill_level].
type Range struct {
	Widget
}

func RangeNewFromInternalPtr(ptr uintptr) *Range {
	cls := &Range{}
	cls.Ptr = ptr
	return cls
}

var xRangeGetAdjustment func(uintptr) uintptr

// Get the adjustment which is the “model” object for `GtkRange`.
func (x *Range) GetAdjustment() *Adjustment {
	var cls *Adjustment

	cret := xRangeGetAdjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

var xRangeGetFillLevel func(uintptr) float64

// Gets the current position of the fill level indicator.
func (x *Range) GetFillLevel() float64 {

	cret := xRangeGetFillLevel(x.GoPointer())
	return cret
}

var xRangeGetFlippable func(uintptr) bool

// Gets whether the `GtkRange` respects text direction.
//
// See [method@Gtk.Range.set_flippable].
func (x *Range) GetFlippable() bool {

	cret := xRangeGetFlippable(x.GoPointer())
	return cret
}

var xRangeGetInverted func(uintptr) bool

// Gets whether the range is inverted.
//
// See [method@Gtk.Range.set_inverted].
func (x *Range) GetInverted() bool {

	cret := xRangeGetInverted(x.GoPointer())
	return cret
}

var xRangeGetRangeRect func(uintptr, *gdk.Rectangle)

// This function returns the area that contains the range’s trough,
// in coordinates relative to @range's origin.
//
// This function is useful mainly for `GtkRange` subclasses.
func (x *Range) GetRangeRect(RangeRectVar *gdk.Rectangle) {

	xRangeGetRangeRect(x.GoPointer(), RangeRectVar)

}

var xRangeGetRestrictToFillLevel func(uintptr) bool

// Gets whether the range is restricted to the fill level.
func (x *Range) GetRestrictToFillLevel() bool {

	cret := xRangeGetRestrictToFillLevel(x.GoPointer())
	return cret
}

var xRangeGetRoundDigits func(uintptr) int

// Gets the number of digits to round the value to when
// it changes.
//
// See [signal@Gtk.Range::change-value].
func (x *Range) GetRoundDigits() int {

	cret := xRangeGetRoundDigits(x.GoPointer())
	return cret
}

var xRangeGetShowFillLevel func(uintptr) bool

// Gets whether the range displays the fill level graphically.
func (x *Range) GetShowFillLevel() bool {

	cret := xRangeGetShowFillLevel(x.GoPointer())
	return cret
}

var xRangeGetSliderRange func(uintptr, int, int)

// This function returns sliders range along the long dimension,
// in widget-&gt;window coordinates.
//
// This function is useful mainly for `GtkRange` subclasses.
func (x *Range) GetSliderRange(SliderStartVar int, SliderEndVar int) {

	xRangeGetSliderRange(x.GoPointer(), SliderStartVar, SliderEndVar)

}

var xRangeGetSliderSizeFixed func(uintptr) bool

// This function is useful mainly for `GtkRange` subclasses.
//
// See [method@Gtk.Range.set_slider_size_fixed].
func (x *Range) GetSliderSizeFixed() bool {

	cret := xRangeGetSliderSizeFixed(x.GoPointer())
	return cret
}

var xRangeGetValue func(uintptr) float64

// Gets the current value of the range.
func (x *Range) GetValue() float64 {

	cret := xRangeGetValue(x.GoPointer())
	return cret
}

var xRangeSetAdjustment func(uintptr, uintptr)

// Sets the adjustment to be used as the “model” object for the `GtkRange`
//
// The adjustment indicates the current range value, the minimum and
// maximum range values, the step/page increments used for keybindings
// and scrolling, and the page size.
//
// The page size is normally 0 for `GtkScale` and nonzero for `GtkScrollbar`,
// and indicates the size of the visible area of the widget being scrolled.
// The page size affects the size of the scrollbar slider.
func (x *Range) SetAdjustment(AdjustmentVar *Adjustment) {

	xRangeSetAdjustment(x.GoPointer(), AdjustmentVar.GoPointer())

}

var xRangeSetFillLevel func(uintptr, float64)

// Set the new position of the fill level indicator.
//
// The “fill level” is probably best described by its most prominent
// use case, which is an indicator for the amount of pre-buffering in
// a streaming media player. In that use case, the value of the range
// would indicate the current play position, and the fill level would
// be the position up to which the file/stream has been downloaded.
//
// This amount of prebuffering can be displayed on the range’s trough
// and is themeable separately from the trough. To enable fill level
// display, use [method@Gtk.Range.set_show_fill_level]. The range defaults
// to not showing the fill level.
//
// Additionally, it’s possible to restrict the range’s slider position
// to values which are smaller than the fill level. This is controlled
// by [method@Gtk.Range.set_restrict_to_fill_level] and is by default
// enabled.
func (x *Range) SetFillLevel(FillLevelVar float64) {

	xRangeSetFillLevel(x.GoPointer(), FillLevelVar)

}

var xRangeSetFlippable func(uintptr, bool)

// Sets whether the `GtkRange` respects text direction.
//
// If a range is flippable, it will switch its direction
// if it is horizontal and its direction is %GTK_TEXT_DIR_RTL.
//
// See [method@Gtk.Widget.get_direction].
func (x *Range) SetFlippable(FlippableVar bool) {

	xRangeSetFlippable(x.GoPointer(), FlippableVar)

}

var xRangeSetIncrements func(uintptr, float64, float64)

// Sets the step and page sizes for the range.
//
// The step size is used when the user clicks the `GtkScrollbar`
// arrows or moves a `GtkScale` via arrow keys. The page size
// is used for example when moving via Page Up or Page Down keys.
func (x *Range) SetIncrements(StepVar float64, PageVar float64) {

	xRangeSetIncrements(x.GoPointer(), StepVar, PageVar)

}

var xRangeSetInverted func(uintptr, bool)

// Sets whether to invert the range.
//
// Ranges normally move from lower to higher values as the
// slider moves from top to bottom or left to right. Inverted
// ranges have higher values at the top or on the right rather
// than on the bottom or left.
func (x *Range) SetInverted(SettingVar bool) {

	xRangeSetInverted(x.GoPointer(), SettingVar)

}

var xRangeSetRange func(uintptr, float64, float64)

// Sets the allowable values in the `GtkRange`.
//
// The range value is clamped to be between @min and @max.
// (If the range has a non-zero page size, it is clamped
// between @min and @max - page-size.)
func (x *Range) SetRange(MinVar float64, MaxVar float64) {

	xRangeSetRange(x.GoPointer(), MinVar, MaxVar)

}

var xRangeSetRestrictToFillLevel func(uintptr, bool)

// Sets whether the slider is restricted to the fill level.
//
// See [method@Gtk.Range.set_fill_level] for a general description
// of the fill level concept.
func (x *Range) SetRestrictToFillLevel(RestrictToFillLevelVar bool) {

	xRangeSetRestrictToFillLevel(x.GoPointer(), RestrictToFillLevelVar)

}

var xRangeSetRoundDigits func(uintptr, int)

// Sets the number of digits to round the value to when
// it changes.
//
// See [signal@Gtk.Range::change-value].
func (x *Range) SetRoundDigits(RoundDigitsVar int) {

	xRangeSetRoundDigits(x.GoPointer(), RoundDigitsVar)

}

var xRangeSetShowFillLevel func(uintptr, bool)

// Sets whether a graphical fill level is show on the trough.
//
// See [method@Gtk.Range.set_fill_level] for a general description
// of the fill level concept.
func (x *Range) SetShowFillLevel(ShowFillLevelVar bool) {

	xRangeSetShowFillLevel(x.GoPointer(), ShowFillLevelVar)

}

var xRangeSetSliderSizeFixed func(uintptr, bool)

// Sets whether the range’s slider has a fixed size, or a size that
// depends on its adjustment’s page size.
//
// This function is useful mainly for `GtkRange` subclasses.
func (x *Range) SetSliderSizeFixed(SizeFixedVar bool) {

	xRangeSetSliderSizeFixed(x.GoPointer(), SizeFixedVar)

}

var xRangeSetValue func(uintptr, float64)

// Sets the current value of the range.
//
// If the value is outside the minimum or maximum range values,
// it will be clamped to fit inside them. The range emits the
// [signal@Gtk.Range::value-changed] signal if the value changes.
func (x *Range) SetValue(ValueVar float64) {

	xRangeSetValue(x.GoPointer(), ValueVar)

}

func (c *Range) GoPointer() uintptr {
	return c.Ptr
}

func (c *Range) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted before clamping a value, to give the application a
// chance to adjust the bounds.
func (x *Range) ConnectAdjustBounds(cb *func(Range, float64)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "adjust-bounds", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ValueVarp float64) {
		fa := Range{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ValueVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "adjust-bounds", cbRefPtr)
}

// Emitted when a scroll action is performed on a range.
//
// It allows an application to determine the type of scroll event
// that occurred and the resultant new value. The application can
// handle the event itself and return %TRUE to prevent further
// processing. Or, by returning %FALSE, it can pass the event to
// other handlers until the default GTK handler is reached.
//
// The value parameter is unrounded. An application that overrides
// the ::change-value signal is responsible for clamping the value
// to the desired number of decimal digits; the default GTK
// handler clamps the value based on [property@Gtk.Range:round-digits].
func (x *Range) ConnectChangeValue(cb *func(Range, ScrollType, float64) bool) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "change-value", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ScrollVarp ScrollType, ValueVarp float64) bool {
		fa := Range{}
		fa.Ptr = clsPtr
		cbFn := *cb

		return cbFn(fa, ScrollVarp, ValueVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "change-value", cbRefPtr)
}

// Virtual function that moves the slider.
//
// Used for keybindings.
func (x *Range) ConnectMoveSlider(cb *func(Range, ScrollType)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "move-slider", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, StepVarp ScrollType) {
		fa := Range{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, StepVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "move-slider", cbRefPtr)
}

// Emitted when the range value changes.
func (x *Range) ConnectValueChanged(cb *func(Range)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "value-changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Range{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "value-changed", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Range) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Range) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Range) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Range) ResetState(StateVar AccessibleState) {

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
func (x *Range) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Range) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Range) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Range) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Range) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Range) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Range) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *Range) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *Range) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xRangeGetAdjustment, lib, "gtk_range_get_adjustment")
	core.PuregoSafeRegister(&xRangeGetFillLevel, lib, "gtk_range_get_fill_level")
	core.PuregoSafeRegister(&xRangeGetFlippable, lib, "gtk_range_get_flippable")
	core.PuregoSafeRegister(&xRangeGetInverted, lib, "gtk_range_get_inverted")
	core.PuregoSafeRegister(&xRangeGetRangeRect, lib, "gtk_range_get_range_rect")
	core.PuregoSafeRegister(&xRangeGetRestrictToFillLevel, lib, "gtk_range_get_restrict_to_fill_level")
	core.PuregoSafeRegister(&xRangeGetRoundDigits, lib, "gtk_range_get_round_digits")
	core.PuregoSafeRegister(&xRangeGetShowFillLevel, lib, "gtk_range_get_show_fill_level")
	core.PuregoSafeRegister(&xRangeGetSliderRange, lib, "gtk_range_get_slider_range")
	core.PuregoSafeRegister(&xRangeGetSliderSizeFixed, lib, "gtk_range_get_slider_size_fixed")
	core.PuregoSafeRegister(&xRangeGetValue, lib, "gtk_range_get_value")
	core.PuregoSafeRegister(&xRangeSetAdjustment, lib, "gtk_range_set_adjustment")
	core.PuregoSafeRegister(&xRangeSetFillLevel, lib, "gtk_range_set_fill_level")
	core.PuregoSafeRegister(&xRangeSetFlippable, lib, "gtk_range_set_flippable")
	core.PuregoSafeRegister(&xRangeSetIncrements, lib, "gtk_range_set_increments")
	core.PuregoSafeRegister(&xRangeSetInverted, lib, "gtk_range_set_inverted")
	core.PuregoSafeRegister(&xRangeSetRange, lib, "gtk_range_set_range")
	core.PuregoSafeRegister(&xRangeSetRestrictToFillLevel, lib, "gtk_range_set_restrict_to_fill_level")
	core.PuregoSafeRegister(&xRangeSetRoundDigits, lib, "gtk_range_set_round_digits")
	core.PuregoSafeRegister(&xRangeSetShowFillLevel, lib, "gtk_range_set_show_fill_level")
	core.PuregoSafeRegister(&xRangeSetSliderSizeFixed, lib, "gtk_range_set_slider_size_fixed")
	core.PuregoSafeRegister(&xRangeSetValue, lib, "gtk_range_set_value")

}
