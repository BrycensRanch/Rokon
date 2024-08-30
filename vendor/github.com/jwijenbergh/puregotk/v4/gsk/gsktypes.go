// Package gsk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gsk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/graphene"
)

// `GskTransform` is an object to describe transform matrices.
//
// Unlike `graphene_matrix_t`, `GskTransform` retains the steps in how
// a transform was constructed, and allows inspecting them. It is modeled
// after the way CSS describes transforms.
//
// `GskTransform` objects are immutable and cannot be changed after creation.
// This means code can safely expose them as properties of objects without
// having to worry about others changing them.
type Transform struct {
}

func (x *Transform) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewTransform func() *Transform

func NewTransform() *Transform {

	cret := xNewTransform()
	return cret
}

var xTransformEqual func(uintptr, *Transform) bool

// Checks two transforms for equality.
func (x *Transform) Equal(SecondVar *Transform) bool {

	cret := xTransformEqual(x.GoPointer(), SecondVar)
	return cret
}

var xTransformGetCategory func(uintptr) TransformCategory

// Returns the category this transform belongs to.
func (x *Transform) GetCategory() TransformCategory {

	cret := xTransformGetCategory(x.GoPointer())
	return cret
}

var xTransformInvert func(uintptr) *Transform

// Inverts the given transform.
//
// If @self is not invertible, %NULL is returned.
// Note that inverting %NULL also returns %NULL, which is
// the correct inverse of %NULL. If you need to differentiate
// between those cases, you should check @self is not %NULL
// before calling this function.
func (x *Transform) Invert() *Transform {

	cret := xTransformInvert(x.GoPointer())
	return cret
}

var xTransformMatrix func(uintptr, *graphene.Matrix) *Transform

// Multiplies @next with the given @matrix.
func (x *Transform) Matrix(MatrixVar *graphene.Matrix) *Transform {

	cret := xTransformMatrix(x.GoPointer(), MatrixVar)
	return cret
}

var xTransformPerspective func(uintptr, float32) *Transform

// Applies a perspective projection transform.
//
// This transform scales points in X and Y based on their Z value,
// scaling points with positive Z values away from the origin, and
// those with negative Z values towards the origin. Points
// on the z=0 plane are unchanged.
func (x *Transform) Perspective(DepthVar float32) *Transform {

	cret := xTransformPerspective(x.GoPointer(), DepthVar)
	return cret
}

var xTransformPrint func(uintptr, *glib.String)

// Converts @self into a human-readable string representation suitable
// for printing.
//
// The result of this function can later be parsed with
// [func@Gsk.Transform.parse].
func (x *Transform) Print(StringVar *glib.String) {

	xTransformPrint(x.GoPointer(), StringVar)

}

var xTransformRef func(uintptr) *Transform

// Acquires a reference on the given `GskTransform`.
func (x *Transform) Ref() *Transform {

	cret := xTransformRef(x.GoPointer())
	return cret
}

var xTransformRotate func(uintptr, float32) *Transform

// Rotates @next @angle degrees in 2D - or in 3D-speak, around the z axis.
func (x *Transform) Rotate(AngleVar float32) *Transform {

	cret := xTransformRotate(x.GoPointer(), AngleVar)
	return cret
}

var xTransformRotate3d func(uintptr, float32, *graphene.Vec3) *Transform

// Rotates @next @angle degrees around @axis.
//
// For a rotation in 2D space, use [method@Gsk.Transform.rotate]
func (x *Transform) Rotate3d(AngleVar float32, AxisVar *graphene.Vec3) *Transform {

	cret := xTransformRotate3d(x.GoPointer(), AngleVar, AxisVar)
	return cret
}

var xTransformScale func(uintptr, float32, float32) *Transform

// Scales @next in 2-dimensional space by the given factors.
//
// Use [method@Gsk.Transform.scale_3d] to scale in all 3 dimensions.
func (x *Transform) Scale(FactorXVar float32, FactorYVar float32) *Transform {

	cret := xTransformScale(x.GoPointer(), FactorXVar, FactorYVar)
	return cret
}

var xTransformScale3d func(uintptr, float32, float32, float32) *Transform

// Scales @next by the given factors.
func (x *Transform) Scale3d(FactorXVar float32, FactorYVar float32, FactorZVar float32) *Transform {

	cret := xTransformScale3d(x.GoPointer(), FactorXVar, FactorYVar, FactorZVar)
	return cret
}

var xTransformSkew func(uintptr, float32, float32) *Transform

// Applies a skew transform.
func (x *Transform) Skew(SkewXVar float32, SkewYVar float32) *Transform {

	cret := xTransformSkew(x.GoPointer(), SkewXVar, SkewYVar)
	return cret
}

var xTransformTo2d func(uintptr, float32, float32, float32, float32, float32, float32)

// Converts a `GskTransform` to a 2D transformation matrix.
//
// @self must be a 2D transformation. If you are not
// sure, use gsk_transform_get_category() &gt;=
// %GSK_TRANSFORM_CATEGORY_2D to check.
//
// The returned values have the following layout:
//
// ```
//
//	| xx yx |   |  a  b  0 |
//	| xy yy | = |  c  d  0 |
//	| dx dy |   | tx ty  1 |
//
// ```
//
// This function can be used to convert between a `GskTransform`
// and a matrix type from other 2D drawing libraries, in particular
// Cairo.
func (x *Transform) To2d(OutXxVar float32, OutYxVar float32, OutXyVar float32, OutYyVar float32, OutDxVar float32, OutDyVar float32) {

	xTransformTo2d(x.GoPointer(), OutXxVar, OutYxVar, OutXyVar, OutYyVar, OutDxVar, OutDyVar)

}

var xTransformTo2dComponents func(uintptr, float32, float32, float32, float32, float32, float32, float32)

// Converts a `GskTransform` to 2D transformation factors.
//
// To recreate an equivalent transform from the factors returned
// by this function, use
//
//	gsk_transform_skew (
//	    gsk_transform_scale (
//	        gsk_transform_rotate (
//	            gsk_transform_translate (NULL, &amp;GRAPHENE_POINT_T (dx, dy)),
//	            angle),
//	        scale_x, scale_y),
//	    skew_x, skew_y)
//
// @self must be a 2D transformation. If you are not sure, use
//
//	gsk_transform_get_category() &gt;= %GSK_TRANSFORM_CATEGORY_2D
//
// to check.
func (x *Transform) To2dComponents(OutSkewXVar float32, OutSkewYVar float32, OutScaleXVar float32, OutScaleYVar float32, OutAngleVar float32, OutDxVar float32, OutDyVar float32) {

	xTransformTo2dComponents(x.GoPointer(), OutSkewXVar, OutSkewYVar, OutScaleXVar, OutScaleYVar, OutAngleVar, OutDxVar, OutDyVar)

}

var xTransformToAffine func(uintptr, float32, float32, float32, float32)

// Converts a `GskTransform` to 2D affine transformation factors.
//
// To recreate an equivalent transform from the factors returned
// by this function, use
//
//	gsk_transform_scale (gsk_transform_translate (NULL,
//	                                              &amp;GRAPHENE_POINT_T (dx, dy)),
//	                     sx, sy)
//
// @self must be a 2D affine transformation. If you are not
// sure, use
//
//	gsk_transform_get_category() &gt;= %GSK_TRANSFORM_CATEGORY_2D_AFFINE
//
// to check.
func (x *Transform) ToAffine(OutScaleXVar float32, OutScaleYVar float32, OutDxVar float32, OutDyVar float32) {

	xTransformToAffine(x.GoPointer(), OutScaleXVar, OutScaleYVar, OutDxVar, OutDyVar)

}

var xTransformToMatrix func(uintptr, *graphene.Matrix)

// Computes the actual value of @self and stores it in @out_matrix.
//
// The previous value of @out_matrix will be ignored.
func (x *Transform) ToMatrix(OutMatrixVar *graphene.Matrix) {

	xTransformToMatrix(x.GoPointer(), OutMatrixVar)

}

var xTransformToString func(uintptr) string

// Converts a matrix into a string that is suitable for printing.
//
// The resulting string can be parsed with [func@Gsk.Transform.parse].
//
// This is a wrapper around [method@Gsk.Transform.print].
func (x *Transform) ToString() string {

	cret := xTransformToString(x.GoPointer())
	return cret
}

var xTransformToTranslate func(uintptr, float32, float32)

// Converts a `GskTransform` to a translation operation.
//
// @self must be a 2D transformation. If you are not
// sure, use
//
//	gsk_transform_get_category() &gt;= %GSK_TRANSFORM_CATEGORY_2D_TRANSLATE
//
// to check.
func (x *Transform) ToTranslate(OutDxVar float32, OutDyVar float32) {

	xTransformToTranslate(x.GoPointer(), OutDxVar, OutDyVar)

}

var xTransformTransform func(uintptr, *Transform) *Transform

// Applies all the operations from @other to @next.
func (x *Transform) Transform(OtherVar *Transform) *Transform {

	cret := xTransformTransform(x.GoPointer(), OtherVar)
	return cret
}

var xTransformTransformBounds func(uintptr, *graphene.Rect, *graphene.Rect)

// Transforms a `graphene_rect_t` using the given transform @self.
//
// The result is the bounding box containing the coplanar quad.
func (x *Transform) TransformBounds(RectVar *graphene.Rect, OutRectVar *graphene.Rect) {

	xTransformTransformBounds(x.GoPointer(), RectVar, OutRectVar)

}

var xTransformTransformPoint func(uintptr, *graphene.Point, *graphene.Point)

// Transforms a `graphene_point_t` using the given transform @self.
func (x *Transform) TransformPoint(PointVar *graphene.Point, OutPointVar *graphene.Point) {

	xTransformTransformPoint(x.GoPointer(), PointVar, OutPointVar)

}

var xTransformTranslate func(uintptr, *graphene.Point) *Transform

// Translates @next in 2-dimensional space by @point.
func (x *Transform) Translate(PointVar *graphene.Point) *Transform {

	cret := xTransformTranslate(x.GoPointer(), PointVar)
	return cret
}

var xTransformTranslate3d func(uintptr, *graphene.Point3D) *Transform

// Translates @next by @point.
func (x *Transform) Translate3d(PointVar *graphene.Point3D) *Transform {

	cret := xTransformTranslate3d(x.GoPointer(), PointVar)
	return cret
}

var xTransformUnref func(uintptr)

// Releases a reference on the given `GskTransform`.
//
// If the reference was the last, the resources associated to the @self are
// freed.
func (x *Transform) Unref() {

	xTransformUnref(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GSK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTransform, lib, "gsk_transform_new")

	core.PuregoSafeRegister(&xTransformEqual, lib, "gsk_transform_equal")
	core.PuregoSafeRegister(&xTransformGetCategory, lib, "gsk_transform_get_category")
	core.PuregoSafeRegister(&xTransformInvert, lib, "gsk_transform_invert")
	core.PuregoSafeRegister(&xTransformMatrix, lib, "gsk_transform_matrix")
	core.PuregoSafeRegister(&xTransformPerspective, lib, "gsk_transform_perspective")
	core.PuregoSafeRegister(&xTransformPrint, lib, "gsk_transform_print")
	core.PuregoSafeRegister(&xTransformRef, lib, "gsk_transform_ref")
	core.PuregoSafeRegister(&xTransformRotate, lib, "gsk_transform_rotate")
	core.PuregoSafeRegister(&xTransformRotate3d, lib, "gsk_transform_rotate_3d")
	core.PuregoSafeRegister(&xTransformScale, lib, "gsk_transform_scale")
	core.PuregoSafeRegister(&xTransformScale3d, lib, "gsk_transform_scale_3d")
	core.PuregoSafeRegister(&xTransformSkew, lib, "gsk_transform_skew")
	core.PuregoSafeRegister(&xTransformTo2d, lib, "gsk_transform_to_2d")
	core.PuregoSafeRegister(&xTransformTo2dComponents, lib, "gsk_transform_to_2d_components")
	core.PuregoSafeRegister(&xTransformToAffine, lib, "gsk_transform_to_affine")
	core.PuregoSafeRegister(&xTransformToMatrix, lib, "gsk_transform_to_matrix")
	core.PuregoSafeRegister(&xTransformToString, lib, "gsk_transform_to_string")
	core.PuregoSafeRegister(&xTransformToTranslate, lib, "gsk_transform_to_translate")
	core.PuregoSafeRegister(&xTransformTransform, lib, "gsk_transform_transform")
	core.PuregoSafeRegister(&xTransformTransformBounds, lib, "gsk_transform_transform_bounds")
	core.PuregoSafeRegister(&xTransformTransformPoint, lib, "gsk_transform_transform_point")
	core.PuregoSafeRegister(&xTransformTranslate, lib, "gsk_transform_translate")
	core.PuregoSafeRegister(&xTransformTranslate3d, lib, "gsk_transform_translate_3d")
	core.PuregoSafeRegister(&xTransformUnref, lib, "gsk_transform_unref")

}
