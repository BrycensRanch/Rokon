// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/cairo"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

var xRenderActivity func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders an activity indicator (such as in `GtkSpinner`).
// The state %GTK_STATE_FLAG_CHECKED determines whether there is
// activity going on.
func RenderActivity(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderActivity(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xRenderArrow func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders an arrow pointing to @angle.
//
// Typical arrow rendering at 0, 1⁄2 π;, π; and 3⁄2 π:
//
// ![](arrows.png)
func RenderArrow(ContextVar *StyleContext, CrVar *cairo.Context, AngleVar float64, XVar float64, YVar float64, SizeVar float64) {

	xRenderArrow(ContextVar.GoPointer(), CrVar, AngleVar, XVar, YVar, SizeVar)

}

var xRenderBackground func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders the background of an element.
//
// Typical background rendering, showing the effect of
// `background-image`, `border-width` and `border-radius`:
//
// ![](background.png)
func RenderBackground(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderBackground(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xRenderCheck func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders a checkmark (as in a `GtkCheckButton`).
//
// The %GTK_STATE_FLAG_CHECKED state determines whether the check is
// on or off, and %GTK_STATE_FLAG_INCONSISTENT determines whether it
// should be marked as undefined.
//
// Typical checkmark rendering:
//
// ![](checks.png)
func RenderCheck(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderCheck(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xRenderExpander func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders an expander (as used in `GtkTreeView` and `GtkExpander`) in the area
// defined by @x, @y, @width, @height. The state %GTK_STATE_FLAG_CHECKED
// determines whether the expander is collapsed or expanded.
//
// Typical expander rendering:
//
// ![](expanders.png)
func RenderExpander(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderExpander(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xRenderFocus func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders a focus indicator on the rectangle determined by @x, @y, @width, @height.
//
// Typical focus rendering:
//
// ![](focus.png)
func RenderFocus(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderFocus(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xRenderFrame func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders a frame around the rectangle defined by @x, @y, @width, @height.
//
// Examples of frame rendering, showing the effect of `border-image`,
// `border-color`, `border-width`, `border-radius` and junctions:
//
// ![](frames.png)
func RenderFrame(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderFrame(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xRenderHandle func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders a handle (as in `GtkPaned` and `GtkWindow`’s resize grip),
// in the rectangle determined by @x, @y, @width, @height.
//
// Handles rendered for the paned and grip classes:
//
// ![](handles.png)
func RenderHandle(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderHandle(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xRenderIcon func(uintptr, *cairo.Context, uintptr, float64, float64)

// Renders the icon in @texture at the specified @x and @y coordinates.
//
// This function will render the icon in @texture at exactly its size,
// regardless of scaling factors, which may not be appropriate when
// drawing on displays with high pixel densities.
func RenderIcon(ContextVar *StyleContext, CrVar *cairo.Context, TextureVar *gdk.Texture, XVar float64, YVar float64) {

	xRenderIcon(ContextVar.GoPointer(), CrVar, TextureVar.GoPointer(), XVar, YVar)

}

var xRenderLayout func(uintptr, *cairo.Context, float64, float64, uintptr)

// Renders @layout on the coordinates @x, @y
func RenderLayout(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, LayoutVar *pango.Layout) {

	xRenderLayout(ContextVar.GoPointer(), CrVar, XVar, YVar, LayoutVar.GoPointer())

}

var xRenderLine func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders a line from (x0, y0) to (x1, y1).
func RenderLine(ContextVar *StyleContext, CrVar *cairo.Context, X0Var float64, Y0Var float64, X1Var float64, Y1Var float64) {

	xRenderLine(ContextVar.GoPointer(), CrVar, X0Var, Y0Var, X1Var, Y1Var)

}

var xRenderOption func(uintptr, *cairo.Context, float64, float64, float64, float64)

// Renders an option mark (as in a radio button), the %GTK_STATE_FLAG_CHECKED
// state will determine whether the option is on or off, and
// %GTK_STATE_FLAG_INCONSISTENT whether it should be marked as undefined.
//
// Typical option mark rendering:
//
// ![](options.png)
func RenderOption(ContextVar *StyleContext, CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xRenderOption(ContextVar.GoPointer(), CrVar, XVar, YVar, WidthVar, HeightVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xRenderActivity, lib, "gtk_render_activity")
	core.PuregoSafeRegister(&xRenderArrow, lib, "gtk_render_arrow")
	core.PuregoSafeRegister(&xRenderBackground, lib, "gtk_render_background")
	core.PuregoSafeRegister(&xRenderCheck, lib, "gtk_render_check")
	core.PuregoSafeRegister(&xRenderExpander, lib, "gtk_render_expander")
	core.PuregoSafeRegister(&xRenderFocus, lib, "gtk_render_focus")
	core.PuregoSafeRegister(&xRenderFrame, lib, "gtk_render_frame")
	core.PuregoSafeRegister(&xRenderHandle, lib, "gtk_render_handle")
	core.PuregoSafeRegister(&xRenderIcon, lib, "gtk_render_icon")
	core.PuregoSafeRegister(&xRenderLayout, lib, "gtk_render_layout")
	core.PuregoSafeRegister(&xRenderLine, lib, "gtk_render_line")
	core.PuregoSafeRegister(&xRenderOption, lib, "gtk_render_option")

}
