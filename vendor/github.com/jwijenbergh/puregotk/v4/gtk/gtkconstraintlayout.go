// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

type ConstraintLayoutChildClass struct {
	ParentClass uintptr
}

func (x *ConstraintLayoutChildClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type ConstraintLayoutClass struct {
	ParentClass uintptr
}

func (x *ConstraintLayoutClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A layout manager using constraints to describe relations between widgets.
//
// `GtkConstraintLayout` is a layout manager that uses relations between
// widget attributes, expressed via [class@Gtk.Constraint] instances, to
// measure and allocate widgets.
//
// ### How do constraints work
//
// Constraints are objects defining the relationship between attributes
// of a widget; you can read the description of the [class@Gtk.Constraint]
// class to have a more in depth definition.
//
// By taking multiple constraints and applying them to the children of
// a widget using `GtkConstraintLayout`, it's possible to describe
// complex layout policies; each constraint applied to a child or to the parent
// widgets contributes to the full description of the layout, in terms of
// parameters for resolving the value of each attribute.
//
// It is important to note that a layout is defined by the totality of
// constraints; removing a child, or a constraint, from an existing layout
// without changing the remaining constraints may result in an unstable
// or unsolvable layout.
//
// Constraints have an implicit "reading order"; you should start describing
// each edge of each child, as well as their relationship with the parent
// container, from the top left (or top right, in RTL languages), horizontally
// first, and then vertically.
//
// A constraint-based layout with too few constraints can become "unstable",
// that is: have more than one solution. The behavior of an unstable layout
// is undefined.
//
// A constraint-based layout with conflicting constraints may be unsolvable,
// and lead to an unstable layout. You can use the [property@Gtk.Constraint:strength]
// property of [class@Gtk.Constraint] to "nudge" the layout towards a solution.
//
// ### GtkConstraintLayout as GtkBuildable
//
// `GtkConstraintLayout` implements the [iface@Gtk.Buildable] interface and
// has a custom "constraints" element which allows describing constraints in
// a [class@Gtk.Builder] UI file.
//
// An example of a UI definition fragment specifying a constraint:
//
// ```xml
//
//	&lt;object class="GtkConstraintLayout"&gt;
//	  &lt;constraints&gt;
//	    &lt;constraint target="button" target-attribute="start"
//	                relation="eq"
//	                source="super" source-attribute="start"
//	                constant="12"
//	                strength="required" /&gt;
//	    &lt;constraint target="button" target-attribute="width"
//	                relation="ge"
//	                constant="250"
//	                strength="strong" /&gt;
//	  &lt;/constraints&gt;
//	&lt;/object&gt;
//
// ```
//
// The definition above will add two constraints to the GtkConstraintLayout:
//
//   - a required constraint between the leading edge of "button" and
//     the leading edge of the widget using the constraint layout, plus
//     12 pixels
//   - a strong, constant constraint making the width of "button" greater
//     than, or equal to 250 pixels
//
// The "target" and "target-attribute" attributes are required.
//
// The "source" and "source-attribute" attributes of the "constraint"
// element are optional; if they are not specified, the constraint is
// assumed to be a constant.
//
// The "relation" attribute is optional; if not specified, the constraint
// is assumed to be an equality.
//
// The "strength" attribute is optional; if not specified, the constraint
// is assumed to be required.
//
// The "source" and "target" attributes can be set to "super" to indicate
// that the constraint target is the widget using the GtkConstraintLayout.
//
// There can be "constant" and "multiplier" attributes.
//
// Additionally, the "constraints" element can also contain a description
// of the `GtkConstraintGuides` used by the layout:
//
// ```xml
//
//	&lt;constraints&gt;
//	  &lt;guide min-width="100" max-width="500" name="hspace"/&gt;
//	  &lt;guide min-height="64" nat-height="128" name="vspace" strength="strong"/&gt;
//	&lt;/constraints&gt;
//
// ```
//
// The "guide" element has the following optional attributes:
//
//   - "min-width", "nat-width", and "max-width", describe the minimum,
//     natural, and maximum width of the guide, respectively
//   - "min-height", "nat-height", and "max-height", describe the minimum,
//     natural, and maximum height of the guide, respectively
//   - "strength" describes the strength of the constraint on the natural
//     size of the guide; if not specified, the constraint is assumed to
//     have a medium strength
//   - "name" describes a name for the guide, useful when debugging
//
// ### Using the Visual Format Language
//
// Complex constraints can be described using a compact syntax called VFL,
// or *Visual Format Language*.
//
// The Visual Format Language describes all the constraints on a row or
// column, typically starting from the leading edge towards the trailing
// one. Each element of the layout is composed by "views", which identify
// a [iface@Gtk.ConstraintTarget].
//
// For instance:
//
// ```
//
//	[button]-[textField]
//
// ```
//
// Describes a constraint that binds the trailing edge of "button" to the
// leading edge of "textField", leaving a default space between the two.
//
// Using VFL is also possible to specify predicates that describe constraints
// on attributes like width and height:
//
// ```
//
//	// Width must be greater than, or equal to 50
//	[button(&gt;=50)]
//
//	// Width of button1 must be equal to width of button2
//	[button1(==button2)]
//
// ```
//
// The default orientation for a VFL description is horizontal, unless
// otherwise specified:
//
// ```
//
//	// horizontal orientation, default attribute: width
//	H:[button(&gt;=150)]
//
//	// vertical orientation, default attribute: height
//	V:[button1(==button2)]
//
// ```
//
// It's also possible to specify multiple predicates, as well as their
// strength:
//
// ```
//
//	// minimum width of button must be 150
//	// natural width of button can be 250
//	[button(&gt;=150@required, ==250@medium)]
//
// ```
//
// Finally, it's also possible to use simple arithmetic operators:
//
// ```
//
//	// width of button1 must be equal to width of button2
//	// divided by 2 plus 12
//	[button1(button2 / 2 + 12)]
//
// ```
type ConstraintLayout struct {
	LayoutManager
}

func ConstraintLayoutNewFromInternalPtr(ptr uintptr) *ConstraintLayout {
	cls := &ConstraintLayout{}
	cls.Ptr = ptr
	return cls
}

var xNewConstraintLayout func() uintptr

// Creates a new `GtkConstraintLayout` layout manager.
func NewConstraintLayout() *ConstraintLayout {
	var cls *ConstraintLayout

	cret := xNewConstraintLayout()

	if cret == 0 {
		return nil
	}
	cls = &ConstraintLayout{}
	cls.Ptr = cret
	return cls
}

var xConstraintLayoutAddConstraint func(uintptr, uintptr)

// Adds a constraint to the layout manager.
//
// The [property@Gtk.Constraint:source] and [property@Gtk.Constraint:target]
// properties of `constraint` can be:
//
//   - set to `NULL` to indicate that the constraint refers to the
//     widget using `layout`
//   - set to the [class@Gtk.Widget] using `layout`
//   - set to a child of the [class@Gtk.Widget] using `layout`
//   - set to a [class@Gtk.ConstraintGuide] that is part of `layout`
//
// The @layout acquires the ownership of @constraint after calling
// this function.
func (x *ConstraintLayout) AddConstraint(ConstraintVar *Constraint) {

	xConstraintLayoutAddConstraint(x.GoPointer(), ConstraintVar.GoPointer())

}

var xConstraintLayoutAddConstraintsFromDescription func(uintptr, uintptr, uint, int, int, **glib.Error, string, ...interface{}) *glib.List

// Creates a list of constraints from a VFL description.
//
// This function is a convenience wrapper around
// [method@Gtk.ConstraintLayout.add_constraints_from_descriptionv], using
// variadic arguments to populate the view/target map.
func (x *ConstraintLayout) AddConstraintsFromDescription(LinesVar uintptr, NLinesVar uint, HspacingVar int, VspacingVar int, ErrorVar **glib.Error, FirstViewVar string, varArgs ...interface{}) *glib.List {

	cret := xConstraintLayoutAddConstraintsFromDescription(x.GoPointer(), LinesVar, NLinesVar, HspacingVar, VspacingVar, ErrorVar, FirstViewVar, varArgs...)
	return cret
}

var xConstraintLayoutAddConstraintsFromDescriptionv func(uintptr, uintptr, uint, int, int, *glib.HashTable, **glib.Error) *glib.List

// Creates a list of constraints from a VFL description.
//
// The Visual Format Language, VFL, is based on Apple's AutoLayout [VFL](https://developer.apple.com/library/content/documentation/UserExperience/Conceptual/AutolayoutPG/VisualFormatLanguage.html).
//
// The `views` dictionary is used to match [iface@Gtk.ConstraintTarget]
// instances to the symbolic view name inside the VFL.
//
// The VFL grammar is:
//
// ```
//
//	     &lt;visualFormatString&gt; = (&lt;orientation&gt;)?
//	                            (&lt;superview&gt;&lt;connection&gt;)?
//	                            &lt;view&gt;(&lt;connection&gt;&lt;view&gt;)*
//	                            (&lt;connection&gt;&lt;superview&gt;)?
//	            &lt;orientation&gt; = 'H' | 'V'
//	              &lt;superview&gt; = '|'
//	             &lt;connection&gt; = '' | '-' &lt;predicateList&gt; '-' | '-'
//	          &lt;predicateList&gt; = &lt;simplePredicate&gt; | &lt;predicateListWithParens&gt;
//	        &lt;simplePredicate&gt; = &lt;metricName&gt; | &lt;positiveNumber&gt;
//	&lt;predicateListWithParens&gt; = '(' &lt;predicate&gt; (',' &lt;predicate&gt;)* ')'
//	              &lt;predicate&gt; = (&lt;relation&gt;)? &lt;objectOfPredicate&gt; (&lt;operatorList&gt;)? ('@' &lt;priority&gt;)?
//	               &lt;relation&gt; = '==' | '&lt;=' | '&gt;='
//	      &lt;objectOfPredicate&gt; = &lt;constant&gt; | &lt;viewName&gt; | ('.' &lt;attributeName&gt;)?
//	               &lt;priority&gt; = &lt;positiveNumber&gt; | 'required' | 'strong' | 'medium' | 'weak'
//	               &lt;constant&gt; = &lt;number&gt;
//	           &lt;operatorList&gt; = (&lt;multiplyOperator&gt;)? (&lt;addOperator&gt;)?
//	       &lt;multiplyOperator&gt; = [ '*' | '/' ] &lt;positiveNumber&gt;
//	            &lt;addOperator&gt; = [ '+' | '-' ] &lt;positiveNumber&gt;
//	               &lt;viewName&gt; = [A-Za-z_]([A-Za-z0-9_]*) // A C identifier
//	             &lt;metricName&gt; = [A-Za-z_]([A-Za-z0-9_]*) // A C identifier
//	          &lt;attributeName&gt; = 'top' | 'bottom' | 'left' | 'right' | 'width' | 'height' |
//	                            'start' | 'end' | 'centerX' | 'centerY' | 'baseline'
//	         &lt;positiveNumber&gt; // A positive real number parseable by g_ascii_strtod()
//	                 &lt;number&gt; // A real number parseable by g_ascii_strtod()
//
// ```
//
// **Note**: The VFL grammar used by GTK is slightly different than the one
// defined by Apple, as it can use symbolic values for the constraint's
// strength instead of numeric values; additionally, GTK allows adding
// simple arithmetic operations inside predicates.
//
// Examples of VFL descriptions are:
//
// ```
//
//	// Default spacing
//	[button]-[textField]
//
//	// Width constraint
//	[button(&gt;=50)]
//
//	// Connection to super view
//	|-50-[purpleBox]-50-|
//
//	// Vertical layout
//	V:[topField]-10-[bottomField]
//
//	// Flush views
//	[maroonView][blueView]
//
//	// Priority
//	[button(100@strong)]
//
//	// Equal widths
//	[button1(==button2)]
//
//	// Multiple predicates
//	[flexibleButton(&gt;=70,&lt;=100)]
//
//	// A complete line of layout
//	|-[find]-[findNext]-[findField(&gt;=20)]-|
//
//	// Operators
//	[button1(button2 / 3 + 50)]
//
//	// Named attributes
//	[button1(==button2.height)]
//
// ```
func (x *ConstraintLayout) AddConstraintsFromDescriptionv(LinesVar uintptr, NLinesVar uint, HspacingVar int, VspacingVar int, ViewsVar *glib.HashTable) (*glib.List, error) {
	var cerr *glib.Error

	cret := xConstraintLayoutAddConstraintsFromDescriptionv(x.GoPointer(), LinesVar, NLinesVar, HspacingVar, VspacingVar, ViewsVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xConstraintLayoutAddGuide func(uintptr, uintptr)

// Adds a guide to `layout`.
//
// A guide can be used as the source or target of constraints,
// like a widget, but it is not visible.
//
// The `layout` acquires the ownership of `guide` after calling
// this function.
func (x *ConstraintLayout) AddGuide(GuideVar *ConstraintGuide) {

	xConstraintLayoutAddGuide(x.GoPointer(), GuideVar.GoPointer())

}

var xConstraintLayoutObserveConstraints func(uintptr) uintptr

// Returns a `GListModel` to track the constraints that are
// part of the layout.
//
// Calling this function will enable extra internal bookkeeping
// to track constraints and emit signals on the returned listmodel.
// It may slow down operations a lot.
//
// Applications should try hard to avoid calling this function
// because of the slowdowns.
func (x *ConstraintLayout) ObserveConstraints() *gio.ListModelBase {
	var cls *gio.ListModelBase

	cret := xConstraintLayoutObserveConstraints(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gio.ListModelBase{}
	cls.Ptr = cret
	return cls
}

var xConstraintLayoutObserveGuides func(uintptr) uintptr

// Returns a `GListModel` to track the guides that are
// part of the layout.
//
// Calling this function will enable extra internal bookkeeping
// to track guides and emit signals on the returned listmodel.
// It may slow down operations a lot.
//
// Applications should try hard to avoid calling this function
// because of the slowdowns.
func (x *ConstraintLayout) ObserveGuides() *gio.ListModelBase {
	var cls *gio.ListModelBase

	cret := xConstraintLayoutObserveGuides(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gio.ListModelBase{}
	cls.Ptr = cret
	return cls
}

var xConstraintLayoutRemoveAllConstraints func(uintptr)

// Removes all constraints from the layout manager.
func (x *ConstraintLayout) RemoveAllConstraints() {

	xConstraintLayoutRemoveAllConstraints(x.GoPointer())

}

var xConstraintLayoutRemoveConstraint func(uintptr, uintptr)

// Removes `constraint` from the layout manager,
// so that it no longer influences the layout.
func (x *ConstraintLayout) RemoveConstraint(ConstraintVar *Constraint) {

	xConstraintLayoutRemoveConstraint(x.GoPointer(), ConstraintVar.GoPointer())

}

var xConstraintLayoutRemoveGuide func(uintptr, uintptr)

// Removes `guide` from the layout manager,
// so that it no longer influences the layout.
func (x *ConstraintLayout) RemoveGuide(GuideVar *ConstraintGuide) {

	xConstraintLayoutRemoveGuide(x.GoPointer(), GuideVar.GoPointer())

}

func (c *ConstraintLayout) GoPointer() uintptr {
	return c.Ptr
}

func (c *ConstraintLayout) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ConstraintLayout) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// `GtkLayoutChild` subclass for children in a `GtkConstraintLayout`.
type ConstraintLayoutChild struct {
	LayoutChild
}

func ConstraintLayoutChildNewFromInternalPtr(ptr uintptr) *ConstraintLayoutChild {
	cls := &ConstraintLayoutChild{}
	cls.Ptr = ptr
	return cls
}

func (c *ConstraintLayoutChild) GoPointer() uintptr {
	return c.Ptr
}

func (c *ConstraintLayoutChild) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewConstraintLayout, lib, "gtk_constraint_layout_new")

	core.PuregoSafeRegister(&xConstraintLayoutAddConstraint, lib, "gtk_constraint_layout_add_constraint")
	core.PuregoSafeRegister(&xConstraintLayoutAddConstraintsFromDescription, lib, "gtk_constraint_layout_add_constraints_from_description")
	core.PuregoSafeRegister(&xConstraintLayoutAddConstraintsFromDescriptionv, lib, "gtk_constraint_layout_add_constraints_from_descriptionv")
	core.PuregoSafeRegister(&xConstraintLayoutAddGuide, lib, "gtk_constraint_layout_add_guide")
	core.PuregoSafeRegister(&xConstraintLayoutObserveConstraints, lib, "gtk_constraint_layout_observe_constraints")
	core.PuregoSafeRegister(&xConstraintLayoutObserveGuides, lib, "gtk_constraint_layout_observe_guides")
	core.PuregoSafeRegister(&xConstraintLayoutRemoveAllConstraints, lib, "gtk_constraint_layout_remove_all_constraints")
	core.PuregoSafeRegister(&xConstraintLayoutRemoveConstraint, lib, "gtk_constraint_layout_remove_constraint")
	core.PuregoSafeRegister(&xConstraintLayoutRemoveGuide, lib, "gtk_constraint_layout_remove_guide")

}
