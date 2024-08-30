// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Type of the callback passed to gtk_tree_model_foreach() to
// iterate over the rows in a tree model.
type TreeModelForeachFunc func(uintptr, *TreePath, *TreeIter, uintptr) bool

// The `GtkTreeIter` is the primary structure
// for accessing a `GtkTreeModel`. Models are expected to put a unique
// integer in the @stamp member, and put
// model-specific data in the three @user_data
// members.
type TreeIter struct {
	Stamp int

	UserData uintptr

	UserData2 uintptr

	UserData3 uintptr
}

func (x *TreeIter) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xTreeIterCopy func(uintptr) *TreeIter

// Creates a dynamically allocated tree iterator as a copy of @iter.
//
// This function is not intended for use in applications,
// because you can just copy the structs by value
// (`GtkTreeIter new_iter = iter;`).
// You must free this iter with gtk_tree_iter_free().
func (x *TreeIter) Copy() *TreeIter {

	cret := xTreeIterCopy(x.GoPointer())
	return cret
}

var xTreeIterFree func(uintptr)

// Frees an iterator that has been allocated by gtk_tree_iter_copy().
//
// This function is mainly used for language bindings.
func (x *TreeIter) Free() {

	xTreeIterFree(x.GoPointer())

}

type TreeModelIface struct {
	GIface uintptr
}

func (x *TreeModelIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// An opaque structure representing a path to a row in a model.
type TreePath struct {
}

func (x *TreePath) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewTreePath func() *TreePath

// Creates a new `GtkTreePath`
// This refers to a row.
func NewTreePath() *TreePath {

	cret := xNewTreePath()
	return cret
}

var xNewTreePathFirst func() *TreePath

// Creates a new `GtkTreePath`.
//
// The string representation of this path is “0”.
func NewTreePathFirst() *TreePath {

	cret := xNewTreePathFirst()
	return cret
}

var xNewTreePathFromIndices func(int, ...interface{}) *TreePath

// Creates a new path with @first_index and @varargs as indices.
func NewTreePathFromIndices(FirstIndexVar int, varArgs ...interface{}) *TreePath {

	cret := xNewTreePathFromIndices(FirstIndexVar, varArgs...)
	return cret
}

var xNewTreePathFromIndicesv func(uintptr, uint) *TreePath

// Creates a new path with the given @indices array of @length.
func NewTreePathFromIndicesv(IndicesVar uintptr, LengthVar uint) *TreePath {

	cret := xNewTreePathFromIndicesv(IndicesVar, LengthVar)
	return cret
}

var xNewTreePathFromString func(string) *TreePath

// Creates a new `GtkTreePath` initialized to @path.
//
// @path is expected to be a colon separated list of numbers.
// For example, the string “10:4:0” would create a path of depth
// 3 pointing to the 11th child of the root node, the 5th
// child of that 11th child, and the 1st child of that 5th child.
// If an invalid path string is passed in, %NULL is returned.
func NewTreePathFromString(PathVar string) *TreePath {

	cret := xNewTreePathFromString(PathVar)
	return cret
}

var xTreePathAppendIndex func(uintptr, int)

// Appends a new index to a path.
//
// As a result, the depth of the path is increased.
func (x *TreePath) AppendIndex(IndexVar int) {

	xTreePathAppendIndex(x.GoPointer(), IndexVar)

}

var xTreePathCompare func(uintptr, *TreePath) int

// Compares two paths.
//
// If @a appears before @b in a tree, then -1 is returned.
// If @b appears before @a, then 1 is returned.
// If the two nodes are equal, then 0 is returned.
func (x *TreePath) Compare(BVar *TreePath) int {

	cret := xTreePathCompare(x.GoPointer(), BVar)
	return cret
}

var xTreePathCopy func(uintptr) *TreePath

// Creates a new `GtkTreePath` as a copy of @path.
func (x *TreePath) Copy() *TreePath {

	cret := xTreePathCopy(x.GoPointer())
	return cret
}

var xTreePathDown func(uintptr)

// Moves @path to point to the first child of the current path.
func (x *TreePath) Down() {

	xTreePathDown(x.GoPointer())

}

var xTreePathFree func(uintptr)

// Frees @path. If @path is %NULL, it simply returns.
func (x *TreePath) Free() {

	xTreePathFree(x.GoPointer())

}

var xTreePathGetDepth func(uintptr) int

// Returns the current depth of @path.
func (x *TreePath) GetDepth() int {

	cret := xTreePathGetDepth(x.GoPointer())
	return cret
}

var xTreePathGetIndices func(uintptr) int

// Returns the current indices of @path.
//
// This is an array of integers, each representing a node in a tree.
// This value should not be freed.
//
// The length of the array can be obtained with gtk_tree_path_get_depth().
func (x *TreePath) GetIndices() int {

	cret := xTreePathGetIndices(x.GoPointer())
	return cret
}

var xTreePathGetIndicesWithDepth func(uintptr, int) uintptr

// Returns the current indices of @path.
//
// This is an array of integers, each representing a node in a tree.
// It also returns the number of elements in the array.
// The array should not be freed.
func (x *TreePath) GetIndicesWithDepth(DepthVar int) uintptr {

	cret := xTreePathGetIndicesWithDepth(x.GoPointer(), DepthVar)
	return cret
}

var xTreePathIsAncestor func(uintptr, *TreePath) bool

// Returns %TRUE if @descendant is a descendant of @path.
func (x *TreePath) IsAncestor(DescendantVar *TreePath) bool {

	cret := xTreePathIsAncestor(x.GoPointer(), DescendantVar)
	return cret
}

var xTreePathIsDescendant func(uintptr, *TreePath) bool

// Returns %TRUE if @path is a descendant of @ancestor.
func (x *TreePath) IsDescendant(AncestorVar *TreePath) bool {

	cret := xTreePathIsDescendant(x.GoPointer(), AncestorVar)
	return cret
}

var xTreePathNext func(uintptr)

// Moves the @path to point to the next node at the current depth.
func (x *TreePath) Next() {

	xTreePathNext(x.GoPointer())

}

var xTreePathPrependIndex func(uintptr, int)

// Prepends a new index to a path.
//
// As a result, the depth of the path is increased.
func (x *TreePath) PrependIndex(IndexVar int) {

	xTreePathPrependIndex(x.GoPointer(), IndexVar)

}

var xTreePathPrev func(uintptr) bool

// Moves the @path to point to the previous node at the
// current depth, if it exists.
func (x *TreePath) Prev() bool {

	cret := xTreePathPrev(x.GoPointer())
	return cret
}

var xTreePathToString func(uintptr) string

// Generates a string representation of the path.
//
// This string is a “:” separated list of numbers.
// For example, “4:10:0:3” would be an acceptable
// return value for this string. If the path has
// depth 0, %NULL is returned.
func (x *TreePath) ToString() string {

	cret := xTreePathToString(x.GoPointer())
	return cret
}

var xTreePathUp func(uintptr) bool

// Moves the @path to point to its parent node, if it has a parent.
func (x *TreePath) Up() bool {

	cret := xTreePathUp(x.GoPointer())
	return cret
}

// A GtkTreeRowReference tracks model changes so that it always refers to the
// same row (a `GtkTreePath` refers to a position, not a fixed row). Create a
// new GtkTreeRowReference with gtk_tree_row_reference_new().
type TreeRowReference struct {
}

func (x *TreeRowReference) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewTreeRowReference func(uintptr, *TreePath) *TreeRowReference

// Creates a row reference based on @path.
//
// This reference will keep pointing to the node pointed to
// by @path, so long as it exists. Any changes that occur on @model are
// propagated, and the path is updated appropriately. If
// @path isn’t a valid path in @model, then %NULL is returned.
func NewTreeRowReference(ModelVar TreeModel, PathVar *TreePath) *TreeRowReference {

	cret := xNewTreeRowReference(ModelVar.GoPointer(), PathVar)
	return cret
}

var xNewTreeRowReferenceProxy func(uintptr, uintptr, *TreePath) *TreeRowReference

// You do not need to use this function.
//
// Creates a row reference based on @path.
//
// This reference will keep pointing to the node pointed to
// by @path, so long as it exists. If @path isn’t a valid
// path in @model, then %NULL is returned. However, unlike
// references created with gtk_tree_row_reference_new(), it
// does not listen to the model for changes. The creator of
// the row reference must do this explicitly using
// gtk_tree_row_reference_inserted(), gtk_tree_row_reference_deleted(),
// gtk_tree_row_reference_reordered().
//
// These functions must be called exactly once per proxy when the
// corresponding signal on the model is emitted. This single call
// updates all row references for that proxy. Since built-in GTK
// objects like `GtkTreeView` already use this mechanism internally,
// using them as the proxy object will produce unpredictable results.
// Further more, passing the same object as @model and @proxy
// doesn’t work for reasons of internal implementation.
//
// This type of row reference is primarily meant by structures that
// need to carefully monitor exactly when a row reference updates
// itself, and is not generally needed by most applications.
func NewTreeRowReferenceProxy(ProxyVar *gobject.Object, ModelVar TreeModel, PathVar *TreePath) *TreeRowReference {

	cret := xNewTreeRowReferenceProxy(ProxyVar.GoPointer(), ModelVar.GoPointer(), PathVar)
	return cret
}

var xTreeRowReferenceCopy func(uintptr) *TreeRowReference

// Copies a `GtkTreeRowReference`.
func (x *TreeRowReference) Copy() *TreeRowReference {

	cret := xTreeRowReferenceCopy(x.GoPointer())
	return cret
}

var xTreeRowReferenceFree func(uintptr)

// Free’s @reference. @reference may be %NULL
func (x *TreeRowReference) Free() {

	xTreeRowReferenceFree(x.GoPointer())

}

var xTreeRowReferenceGetModel func(uintptr) uintptr

// Returns the model that the row reference is monitoring.
func (x *TreeRowReference) GetModel() *TreeModelBase {
	var cls *TreeModelBase

	cret := xTreeRowReferenceGetModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TreeModelBase{}
	cls.Ptr = cret
	return cls
}

var xTreeRowReferenceGetPath func(uintptr) *TreePath

// Returns a path that the row reference currently points to,
// or %NULL if the path pointed to is no longer valid.
func (x *TreeRowReference) GetPath() *TreePath {

	cret := xTreeRowReferenceGetPath(x.GoPointer())
	return cret
}

var xTreeRowReferenceValid func(uintptr) bool

// Returns %TRUE if the @reference is non-%NULL and refers to
// a current valid path.
func (x *TreeRowReference) Valid() bool {

	cret := xTreeRowReferenceValid(x.GoPointer())
	return cret
}

// The tree interface used by GtkTreeView
//
// The `GtkTreeModel` interface defines a generic tree interface for
// use by the `GtkTreeView` widget. It is an abstract interface, and
// is designed to be usable with any appropriate data structure. The
// programmer just has to implement this interface on their own data
// type for it to be viewable by a `GtkTreeView` widget.
//
// The model is represented as a hierarchical tree of strongly-typed,
// columned data. In other words, the model can be seen as a tree where
// every node has different values depending on which column is being
// queried. The type of data found in a column is determined by using
// the GType system (ie. %G_TYPE_INT, %GTK_TYPE_BUTTON, %G_TYPE_POINTER,
// etc). The types are homogeneous per column across all nodes. It is
// important to note that this interface only provides a way of examining
// a model and observing changes. The implementation of each individual
// model decides how and if changes are made.
//
// In order to make life simpler for programmers who do not need to
// write their own specialized model, two generic models are provided
// — the `GtkTreeStore` and the `GtkListStore`. To use these, the
// developer simply pushes data into these models as necessary. These
// models provide the data structure as well as all appropriate tree
// interfaces. As a result, implementing drag and drop, sorting, and
// storing data is trivial. For the vast majority of trees and lists,
// these two models are sufficient.
//
// Models are accessed on a node/column level of granularity. One can
// query for the value of a model at a certain node and a certain
// column on that node. There are two structures used to reference a
// particular node in a model. They are the [struct@Gtk.TreePath] and
// the [struct@Gtk.TreeIter] (“iter” is short for iterator). Most of the
// interface consists of operations on a [struct@Gtk.TreeIter].
//
// A path is essentially a potential node. It is a location on a model
// that may or may not actually correspond to a node on a specific
// model. A [struct@Gtk.TreePath] can be converted into either an
// array of unsigned integers or a string. The string form is a list
// of numbers separated by a colon. Each number refers to the offset
// at that level. Thus, the path `0` refers to the root
// node and the path `2:4` refers to the fifth child of
// the third node.
//
// By contrast, a [struct@Gtk.TreeIter] is a reference to a specific node on
// a specific model. It is a generic struct with an integer and three
// generic pointers. These are filled in by the model in a model-specific
// way. One can convert a path to an iterator by calling
// gtk_tree_model_get_iter(). These iterators are the primary way
// of accessing a model and are similar to the iterators used by
// `GtkTextBuffer`. They are generally statically allocated on the
// stack and only used for a short time. The model interface defines
// a set of operations using them for navigating the model.
//
// It is expected that models fill in the iterator with private data.
// For example, the `GtkListStore` model, which is internally a simple
// linked list, stores a list node in one of the pointers. The
// `GtkTreeModel`Sort stores an array and an offset in two of the
// pointers. Additionally, there is an integer field. This field is
// generally filled with a unique stamp per model. This stamp is for
// catching errors resulting from using invalid iterators with a model.
//
// The lifecycle of an iterator can be a little confusing at first.
// Iterators are expected to always be valid for as long as the model
// is unchanged (and doesn’t emit a signal). The model is considered
// to own all outstanding iterators and nothing needs to be done to
// free them from the user’s point of view. Additionally, some models
// guarantee that an iterator is valid for as long as the node it refers
// to is valid (most notably the `GtkTreeStore` and `GtkListStore`).
// Although generally uninteresting, as one always has to allow for
// the case where iterators do not persist beyond a signal, some very
// important performance enhancements were made in the sort model.
// As a result, the %GTK_TREE_MODEL_ITERS_PERSIST flag was added to
// indicate this behavior.
//
// To help show some common operation of a model, some examples are
// provided. The first example shows three ways of getting the iter at
// the location `3:2:5`. While the first method shown is
// easier, the second is much more common, as you often get paths from
// callbacks.
//
// ## Acquiring a `GtkTreeIter`
//
// ```c
// // Three ways of getting the iter pointing to the location
// GtkTreePath *path;
// GtkTreeIter iter;
// GtkTreeIter parent_iter;
//
// // get the iterator from a string
// gtk_tree_model_get_iter_from_string (model,
//
//	&amp;iter,
//	"3:2:5");
//
// // get the iterator from a path
// path = gtk_tree_path_new_from_string ("3:2:5");
// gtk_tree_model_get_iter (model, &amp;iter, path);
// gtk_tree_path_free (path);
//
// // walk the tree to find the iterator
// gtk_tree_model_iter_nth_child (model, &amp;iter,
//
//	NULL, 3);
//
// parent_iter = iter;
// gtk_tree_model_iter_nth_child (model, &amp;iter,
//
//	&amp;parent_iter, 2);
//
// parent_iter = iter;
// gtk_tree_model_iter_nth_child (model, &amp;iter,
//
//	&amp;parent_iter, 5);
//
// ```
//
// This second example shows a quick way of iterating through a list
// and getting a string and an integer from each row. The
// populate_model() function used below is not
// shown, as it is specific to the `GtkListStore`. For information on
// how to write such a function, see the `GtkListStore` documentation.
//
// ## Reading data from a `GtkTreeModel`
//
// ```c
// enum
//
//	{
//	  STRING_COLUMN,
//	  INT_COLUMN,
//	  N_COLUMNS
//	};
//
// ...
//
// GtkTreeModel *list_store;
// GtkTreeIter iter;
// gboolean valid;
// int row_count = 0;
//
// // make a new list_store
// list_store = gtk_list_store_new (N_COLUMNS,
//
//	G_TYPE_STRING,
//	G_TYPE_INT);
//
// // Fill the list store with data
// populate_model (list_store);
//
// // Get the first iter in the list, check it is valid and walk
// // through the list, reading each row.
//
// valid = gtk_tree_model_get_iter_first (list_store,
//
//	&amp;iter);
//
// while (valid)
//
//	{
//	  char *str_data;
//	  int    int_data;
//
//	  // Make sure you terminate calls to gtk_tree_model_get() with a “-1” value
//	  gtk_tree_model_get (list_store, &amp;iter,
//	                      STRING_COLUMN, &amp;str_data,
//	                      INT_COLUMN, &amp;int_data,
//	                      -1);
//
//	  // Do something with the data
//	  g_print ("Row %d: (%s,%d)\n",
//	           row_count, str_data, int_data);
//	  g_free (str_data);
//
//	  valid = gtk_tree_model_iter_next (list_store,
//	                                    &amp;iter);
//	  row_count++;
//	}
//
// ```
//
// The `GtkTreeModel` interface contains two methods for reference
// counting: gtk_tree_model_ref_node() and gtk_tree_model_unref_node().
// These two methods are optional to implement. The reference counting
// is meant as a way for views to let models know when nodes are being
// displayed. `GtkTreeView` will take a reference on a node when it is
// visible, which means the node is either in the toplevel or expanded.
// Being displayed does not mean that the node is currently directly
// visible to the user in the viewport. Based on this reference counting
// scheme a caching model, for example, can decide whether or not to cache
// a node based on the reference count. A file-system based model would
// not want to keep the entire file hierarchy in memory, but just the
// folders that are currently expanded in every current view.
//
// When working with reference counting, the following rules must be taken
// into account:
//
//   - Never take a reference on a node without owning a reference on its parent.
//     This means that all parent nodes of a referenced node must be referenced
//     as well.
//
//   - Outstanding references on a deleted node are not released. This is not
//     possible because the node has already been deleted by the time the
//     row-deleted signal is received.
//
//   - Models are not obligated to emit a signal on rows of which none of its
//     siblings are referenced. To phrase this differently, signals are only
//     required for levels in which nodes are referenced. For the root level
//     however, signals must be emitted at all times (however the root level
//     is always referenced when any view is attached).
type TreeModel interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	FilterNew(RootVar *TreePath) *TreeModelBase
	Foreach(FuncVar *TreeModelForeachFunc, UserDataVar uintptr)
	Get(IterVar *TreeIter, varArgs ...interface{})
	GetColumnType(IndexVar int) []interface{}
	GetFlags() TreeModelFlags
	GetIter(IterVar *TreeIter, PathVar *TreePath) bool
	GetIterFirst(IterVar *TreeIter) bool
	GetIterFromString(IterVar *TreeIter, PathStringVar string) bool
	GetNColumns() int
	GetPath(IterVar *TreeIter) *TreePath
	GetStringFromIter(IterVar *TreeIter) string
	GetValist(IterVar *TreeIter, VarArgsVar []interface{})
	GetValue(IterVar *TreeIter, ColumnVar int, ValueVar *gobject.Value)
	IterChildren(IterVar *TreeIter, ParentVar *TreeIter) bool
	IterHasChild(IterVar *TreeIter) bool
	IterNChildren(IterVar *TreeIter) int
	IterNext(IterVar *TreeIter) bool
	IterNthChild(IterVar *TreeIter, ParentVar *TreeIter, NVar int) bool
	IterParent(IterVar *TreeIter, ChildVar *TreeIter) bool
	IterPrevious(IterVar *TreeIter) bool
	RefNode(IterVar *TreeIter)
	RowChanged(PathVar *TreePath, IterVar *TreeIter)
	RowDeleted(PathVar *TreePath)
	RowHasChildToggled(PathVar *TreePath, IterVar *TreeIter)
	RowInserted(PathVar *TreePath, IterVar *TreeIter)
	RowsReordered(PathVar *TreePath, IterVar *TreeIter, NewOrderVar int)
	RowsReorderedWithLength(PathVar *TreePath, IterVar *TreeIter, NewOrderVar uintptr, LengthVar int)
	UnrefNode(IterVar *TreeIter)
}
type TreeModelBase struct {
	Ptr uintptr
}

func (x *TreeModelBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *TreeModelBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Creates a new `GtkTreeModel`, with @child_model as the child_model
// and @root as the virtual root.
func (x *TreeModelBase) FilterNew(RootVar *TreePath) *TreeModelBase {
	var cls *TreeModelBase

	cret := XGtkTreeModelFilterNew(x.GoPointer(), RootVar)

	if cret == 0 {
		return nil
	}
	cls = &TreeModelBase{}
	cls.Ptr = cret
	return cls
}

// Calls @func on each node in model in a depth-first fashion.
//
// If @func returns %TRUE, then the tree ceases to be walked,
// and gtk_tree_model_foreach() returns.
func (x *TreeModelBase) Foreach(FuncVar *TreeModelForeachFunc, UserDataVar uintptr) {

	XGtkTreeModelForeach(x.GoPointer(), glib.NewCallback(FuncVar), UserDataVar)

}

// Gets the value of one or more cells in the row referenced by @iter.
//
// The variable argument list should contain integer column numbers,
// each column number followed by a place to store the value being
// retrieved.  The list is terminated by a -1. For example, to get a
// value from column 0 with type %G_TYPE_STRING, you would
// write: `gtk_tree_model_get (model, iter, 0, &amp;place_string_here, -1)`,
// where `place_string_here` is a #gchararray
// to be filled with the string.
//
// Returned values with type %G_TYPE_OBJECT have to be unreferenced,
// values with type %G_TYPE_STRING or %G_TYPE_BOXED have to be freed.
// Other values are passed by value.
func (x *TreeModelBase) Get(IterVar *TreeIter, varArgs ...interface{}) {

	XGtkTreeModelGet(x.GoPointer(), IterVar, varArgs...)

}

// Returns the type of the column.
func (x *TreeModelBase) GetColumnType(IndexVar int) []interface{} {

	cret := XGtkTreeModelGetColumnType(x.GoPointer(), IndexVar)
	return cret
}

// Returns a set of flags supported by this interface.
//
// The flags are a bitwise combination of `GtkTreeModel`Flags.
// The flags supported should not change during the lifetime
// of the @tree_model.
func (x *TreeModelBase) GetFlags() TreeModelFlags {

	cret := XGtkTreeModelGetFlags(x.GoPointer())
	return cret
}

// Sets @iter to a valid iterator pointing to @path.
//
// If @path does not exist, @iter is set to an invalid
// iterator and %FALSE is returned.
func (x *TreeModelBase) GetIter(IterVar *TreeIter, PathVar *TreePath) bool {

	cret := XGtkTreeModelGetIter(x.GoPointer(), IterVar, PathVar)
	return cret
}

// Initializes @iter with the first iterator in the tree
// (the one at the path "0").
//
// Returns %FALSE if the tree is empty, %TRUE otherwise.
func (x *TreeModelBase) GetIterFirst(IterVar *TreeIter) bool {

	cret := XGtkTreeModelGetIterFirst(x.GoPointer(), IterVar)
	return cret
}

// Sets @iter to a valid iterator pointing to @path_string, if it
// exists.
//
// Otherwise, @iter is left invalid and %FALSE is returned.
func (x *TreeModelBase) GetIterFromString(IterVar *TreeIter, PathStringVar string) bool {

	cret := XGtkTreeModelGetIterFromString(x.GoPointer(), IterVar, PathStringVar)
	return cret
}

// Returns the number of columns supported by @tree_model.
func (x *TreeModelBase) GetNColumns() int {

	cret := XGtkTreeModelGetNColumns(x.GoPointer())
	return cret
}

// Returns a newly-created `GtkTreePath` referenced by @iter.
//
// This path should be freed with gtk_tree_path_free().
func (x *TreeModelBase) GetPath(IterVar *TreeIter) *TreePath {

	cret := XGtkTreeModelGetPath(x.GoPointer(), IterVar)
	return cret
}

// Generates a string representation of the iter.
//
// This string is a “:” separated list of numbers.
// For example, “4:10:0:3” would be an acceptable
// return value for this string.
func (x *TreeModelBase) GetStringFromIter(IterVar *TreeIter) string {

	cret := XGtkTreeModelGetStringFromIter(x.GoPointer(), IterVar)
	return cret
}

// Gets the value of one or more cells in the row referenced by @iter.
//
// See [method@Gtk.TreeModel.get], this version takes a va_list
// for language bindings to use.
func (x *TreeModelBase) GetValist(IterVar *TreeIter, VarArgsVar []interface{}) {

	XGtkTreeModelGetValist(x.GoPointer(), IterVar, VarArgsVar)

}

// Initializes and sets @value to that at @column.
//
// When done with @value, g_value_unset() needs to be called
// to free any allocated memory.
func (x *TreeModelBase) GetValue(IterVar *TreeIter, ColumnVar int, ValueVar *gobject.Value) {

	XGtkTreeModelGetValue(x.GoPointer(), IterVar, ColumnVar, ValueVar)

}

// Sets @iter to point to the first child of @parent.
//
// If @parent has no children, %FALSE is returned and @iter is
// set to be invalid. @parent will remain a valid node after this
// function has been called.
//
// If @parent is %NULL returns the first node, equivalent to
// `gtk_tree_model_get_iter_first (tree_model, iter);`
func (x *TreeModelBase) IterChildren(IterVar *TreeIter, ParentVar *TreeIter) bool {

	cret := XGtkTreeModelIterChildren(x.GoPointer(), IterVar, ParentVar)
	return cret
}

// Returns %TRUE if @iter has children, %FALSE otherwise.
func (x *TreeModelBase) IterHasChild(IterVar *TreeIter) bool {

	cret := XGtkTreeModelIterHasChild(x.GoPointer(), IterVar)
	return cret
}

// Returns the number of children that @iter has.
//
// As a special case, if @iter is %NULL, then the number
// of toplevel nodes is returned.
func (x *TreeModelBase) IterNChildren(IterVar *TreeIter) int {

	cret := XGtkTreeModelIterNChildren(x.GoPointer(), IterVar)
	return cret
}

// Sets @iter to point to the node following it at the current level.
//
// If there is no next @iter, %FALSE is returned and @iter is set
// to be invalid.
func (x *TreeModelBase) IterNext(IterVar *TreeIter) bool {

	cret := XGtkTreeModelIterNext(x.GoPointer(), IterVar)
	return cret
}

// Sets @iter to be the child of @parent, using the given index.
//
// The first index is 0. If @n is too big, or @parent has no children,
// @iter is set to an invalid iterator and %FALSE is returned. @parent
// will remain a valid node after this function has been called. As a
// special case, if @parent is %NULL, then the @n-th root node
// is set.
func (x *TreeModelBase) IterNthChild(IterVar *TreeIter, ParentVar *TreeIter, NVar int) bool {

	cret := XGtkTreeModelIterNthChild(x.GoPointer(), IterVar, ParentVar, NVar)
	return cret
}

// Sets @iter to be the parent of @child.
//
// If @child is at the toplevel, and doesn’t have a parent, then
// @iter is set to an invalid iterator and %FALSE is returned.
// @child will remain a valid node after this function has been
// called.
//
// @iter will be initialized before the lookup is performed, so @child
// and @iter cannot point to the same memory location.
func (x *TreeModelBase) IterParent(IterVar *TreeIter, ChildVar *TreeIter) bool {

	cret := XGtkTreeModelIterParent(x.GoPointer(), IterVar, ChildVar)
	return cret
}

// Sets @iter to point to the previous node at the current level.
//
// If there is no previous @iter, %FALSE is returned and @iter is
// set to be invalid.
func (x *TreeModelBase) IterPrevious(IterVar *TreeIter) bool {

	cret := XGtkTreeModelIterPrevious(x.GoPointer(), IterVar)
	return cret
}

// Lets the tree ref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons.
//
// This function is primarily meant as a way for views to let
// caching models know when nodes are being displayed (and hence,
// whether or not to cache that node). Being displayed means a node
// is in an expanded branch, regardless of whether the node is currently
// visible in the viewport. For example, a file-system based model
// would not want to keep the entire file-hierarchy in memory,
// just the sections that are currently being displayed by
// every current view.
//
// A model should be expected to be able to get an iter independent
// of its reffed state.
func (x *TreeModelBase) RefNode(IterVar *TreeIter) {

	XGtkTreeModelRefNode(x.GoPointer(), IterVar)

}

// Emits the ::row-changed signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-changed].
func (x *TreeModelBase) RowChanged(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowChanged(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::row-deleted signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-deleted].
//
// This should be called by models after a row has been removed.
// The location pointed to by @path should be the location that
// the row previously was at. It may not be a valid location anymore.
//
// Nodes that are deleted are not unreffed, this means that any
// outstanding references on the deleted node should not be released.
func (x *TreeModelBase) RowDeleted(PathVar *TreePath) {

	XGtkTreeModelRowDeleted(x.GoPointer(), PathVar)

}

// Emits the ::row-has-child-toggled signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-has-child-toggled].
//
// This should be called by models after the child
// state of a node changes.
func (x *TreeModelBase) RowHasChildToggled(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowHasChildToggled(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::row-inserted signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-inserted].
func (x *TreeModelBase) RowInserted(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowInserted(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::rows-reordered signal on @tree_model.
//
// See [signal@Gtk.TreeModel::rows-reordered].
//
// This should be called by models when their rows have been
// reordered.
func (x *TreeModelBase) RowsReordered(PathVar *TreePath, IterVar *TreeIter, NewOrderVar int) {

	XGtkTreeModelRowsReordered(x.GoPointer(), PathVar, IterVar, NewOrderVar)

}

// Emits the ::rows-reordered signal on @tree_model.
//
// See [signal@Gtk.TreeModel::rows-reordered].
//
// This should be called by models when their rows have been
// reordered.
func (x *TreeModelBase) RowsReorderedWithLength(PathVar *TreePath, IterVar *TreeIter, NewOrderVar uintptr, LengthVar int) {

	XGtkTreeModelRowsReorderedWithLength(x.GoPointer(), PathVar, IterVar, NewOrderVar, LengthVar)

}

// Lets the tree unref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons. For more information on what
// this means, see gtk_tree_model_ref_node().
//
// Please note that nodes that are deleted are not unreffed.
func (x *TreeModelBase) UnrefNode(IterVar *TreeIter) {

	XGtkTreeModelUnrefNode(x.GoPointer(), IterVar)

}

var XGtkTreeModelFilterNew func(uintptr, *TreePath) uintptr
var XGtkTreeModelForeach func(uintptr, uintptr, uintptr)
var XGtkTreeModelGet func(uintptr, *TreeIter, ...interface{})
var XGtkTreeModelGetColumnType func(uintptr, int) []interface{}
var XGtkTreeModelGetFlags func(uintptr) TreeModelFlags
var XGtkTreeModelGetIter func(uintptr, *TreeIter, *TreePath) bool
var XGtkTreeModelGetIterFirst func(uintptr, *TreeIter) bool
var XGtkTreeModelGetIterFromString func(uintptr, *TreeIter, string) bool
var XGtkTreeModelGetNColumns func(uintptr) int
var XGtkTreeModelGetPath func(uintptr, *TreeIter) *TreePath
var XGtkTreeModelGetStringFromIter func(uintptr, *TreeIter) string
var XGtkTreeModelGetValist func(uintptr, *TreeIter, []interface{})
var XGtkTreeModelGetValue func(uintptr, *TreeIter, int, *gobject.Value)
var XGtkTreeModelIterChildren func(uintptr, *TreeIter, *TreeIter) bool
var XGtkTreeModelIterHasChild func(uintptr, *TreeIter) bool
var XGtkTreeModelIterNChildren func(uintptr, *TreeIter) int
var XGtkTreeModelIterNext func(uintptr, *TreeIter) bool
var XGtkTreeModelIterNthChild func(uintptr, *TreeIter, *TreeIter, int) bool
var XGtkTreeModelIterParent func(uintptr, *TreeIter, *TreeIter) bool
var XGtkTreeModelIterPrevious func(uintptr, *TreeIter) bool
var XGtkTreeModelRefNode func(uintptr, *TreeIter)
var XGtkTreeModelRowChanged func(uintptr, *TreePath, *TreeIter)
var XGtkTreeModelRowDeleted func(uintptr, *TreePath)
var XGtkTreeModelRowHasChildToggled func(uintptr, *TreePath, *TreeIter)
var XGtkTreeModelRowInserted func(uintptr, *TreePath, *TreeIter)
var XGtkTreeModelRowsReordered func(uintptr, *TreePath, *TreeIter, int)
var XGtkTreeModelRowsReorderedWithLength func(uintptr, *TreePath, *TreeIter, uintptr, int)
var XGtkTreeModelUnrefNode func(uintptr, *TreeIter)

// These flags indicate various properties of a `GtkTreeModel`.
//
// They are returned by [method@Gtk.TreeModel.get_flags], and must be
// static for the lifetime of the object. A more complete description
// of %GTK_TREE_MODEL_ITERS_PERSIST can be found in the overview of
// this section.
type TreeModelFlags int

const (

	// iterators survive all signals
	//   emitted by the tree
	TreeModelItersPersistValue TreeModelFlags = 1
	// the model is a list only, and never
	//   has children
	TreeModelListOnlyValue TreeModelFlags = 2
)

var xTreeRowReferenceDeleted func(uintptr, *TreePath)

// Lets a set of row reference created by
// gtk_tree_row_reference_new_proxy() know that the
// model emitted the ::row-deleted signal.
func TreeRowReferenceDeleted(ProxyVar *gobject.Object, PathVar *TreePath) {

	xTreeRowReferenceDeleted(ProxyVar.GoPointer(), PathVar)

}

var xTreeRowReferenceInserted func(uintptr, *TreePath)

// Lets a set of row reference created by
// gtk_tree_row_reference_new_proxy() know that the
// model emitted the ::row-inserted signal.
func TreeRowReferenceInserted(ProxyVar *gobject.Object, PathVar *TreePath) {

	xTreeRowReferenceInserted(ProxyVar.GoPointer(), PathVar)

}

var xTreeRowReferenceReordered func(uintptr, *TreePath, *TreeIter, uintptr)

// Lets a set of row reference created by
// gtk_tree_row_reference_new_proxy() know that the
// model emitted the ::rows-reordered signal.
func TreeRowReferenceReordered(ProxyVar *gobject.Object, PathVar *TreePath, IterVar *TreeIter, NewOrderVar uintptr) {

	xTreeRowReferenceReordered(ProxyVar.GoPointer(), PathVar, IterVar, NewOrderVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xTreeRowReferenceDeleted, lib, "gtk_tree_row_reference_deleted")
	core.PuregoSafeRegister(&xTreeRowReferenceInserted, lib, "gtk_tree_row_reference_inserted")
	core.PuregoSafeRegister(&xTreeRowReferenceReordered, lib, "gtk_tree_row_reference_reordered")

	core.PuregoSafeRegister(&xTreeIterCopy, lib, "gtk_tree_iter_copy")
	core.PuregoSafeRegister(&xTreeIterFree, lib, "gtk_tree_iter_free")

	core.PuregoSafeRegister(&xNewTreePath, lib, "gtk_tree_path_new")
	core.PuregoSafeRegister(&xNewTreePathFirst, lib, "gtk_tree_path_new_first")
	core.PuregoSafeRegister(&xNewTreePathFromIndices, lib, "gtk_tree_path_new_from_indices")
	core.PuregoSafeRegister(&xNewTreePathFromIndicesv, lib, "gtk_tree_path_new_from_indicesv")
	core.PuregoSafeRegister(&xNewTreePathFromString, lib, "gtk_tree_path_new_from_string")

	core.PuregoSafeRegister(&xTreePathAppendIndex, lib, "gtk_tree_path_append_index")
	core.PuregoSafeRegister(&xTreePathCompare, lib, "gtk_tree_path_compare")
	core.PuregoSafeRegister(&xTreePathCopy, lib, "gtk_tree_path_copy")
	core.PuregoSafeRegister(&xTreePathDown, lib, "gtk_tree_path_down")
	core.PuregoSafeRegister(&xTreePathFree, lib, "gtk_tree_path_free")
	core.PuregoSafeRegister(&xTreePathGetDepth, lib, "gtk_tree_path_get_depth")
	core.PuregoSafeRegister(&xTreePathGetIndices, lib, "gtk_tree_path_get_indices")
	core.PuregoSafeRegister(&xTreePathGetIndicesWithDepth, lib, "gtk_tree_path_get_indices_with_depth")
	core.PuregoSafeRegister(&xTreePathIsAncestor, lib, "gtk_tree_path_is_ancestor")
	core.PuregoSafeRegister(&xTreePathIsDescendant, lib, "gtk_tree_path_is_descendant")
	core.PuregoSafeRegister(&xTreePathNext, lib, "gtk_tree_path_next")
	core.PuregoSafeRegister(&xTreePathPrependIndex, lib, "gtk_tree_path_prepend_index")
	core.PuregoSafeRegister(&xTreePathPrev, lib, "gtk_tree_path_prev")
	core.PuregoSafeRegister(&xTreePathToString, lib, "gtk_tree_path_to_string")
	core.PuregoSafeRegister(&xTreePathUp, lib, "gtk_tree_path_up")

	core.PuregoSafeRegister(&xNewTreeRowReference, lib, "gtk_tree_row_reference_new")
	core.PuregoSafeRegister(&xNewTreeRowReferenceProxy, lib, "gtk_tree_row_reference_new_proxy")

	core.PuregoSafeRegister(&xTreeRowReferenceCopy, lib, "gtk_tree_row_reference_copy")
	core.PuregoSafeRegister(&xTreeRowReferenceFree, lib, "gtk_tree_row_reference_free")
	core.PuregoSafeRegister(&xTreeRowReferenceGetModel, lib, "gtk_tree_row_reference_get_model")
	core.PuregoSafeRegister(&xTreeRowReferenceGetPath, lib, "gtk_tree_row_reference_get_path")
	core.PuregoSafeRegister(&xTreeRowReferenceValid, lib, "gtk_tree_row_reference_valid")

	core.PuregoSafeRegister(&XGtkTreeModelFilterNew, lib, "gtk_tree_model_filter_new")
	core.PuregoSafeRegister(&XGtkTreeModelForeach, lib, "gtk_tree_model_foreach")
	core.PuregoSafeRegister(&XGtkTreeModelGet, lib, "gtk_tree_model_get")
	core.PuregoSafeRegister(&XGtkTreeModelGetColumnType, lib, "gtk_tree_model_get_column_type")
	core.PuregoSafeRegister(&XGtkTreeModelGetFlags, lib, "gtk_tree_model_get_flags")
	core.PuregoSafeRegister(&XGtkTreeModelGetIter, lib, "gtk_tree_model_get_iter")
	core.PuregoSafeRegister(&XGtkTreeModelGetIterFirst, lib, "gtk_tree_model_get_iter_first")
	core.PuregoSafeRegister(&XGtkTreeModelGetIterFromString, lib, "gtk_tree_model_get_iter_from_string")
	core.PuregoSafeRegister(&XGtkTreeModelGetNColumns, lib, "gtk_tree_model_get_n_columns")
	core.PuregoSafeRegister(&XGtkTreeModelGetPath, lib, "gtk_tree_model_get_path")
	core.PuregoSafeRegister(&XGtkTreeModelGetStringFromIter, lib, "gtk_tree_model_get_string_from_iter")
	core.PuregoSafeRegister(&XGtkTreeModelGetValist, lib, "gtk_tree_model_get_valist")
	core.PuregoSafeRegister(&XGtkTreeModelGetValue, lib, "gtk_tree_model_get_value")
	core.PuregoSafeRegister(&XGtkTreeModelIterChildren, lib, "gtk_tree_model_iter_children")
	core.PuregoSafeRegister(&XGtkTreeModelIterHasChild, lib, "gtk_tree_model_iter_has_child")
	core.PuregoSafeRegister(&XGtkTreeModelIterNChildren, lib, "gtk_tree_model_iter_n_children")
	core.PuregoSafeRegister(&XGtkTreeModelIterNext, lib, "gtk_tree_model_iter_next")
	core.PuregoSafeRegister(&XGtkTreeModelIterNthChild, lib, "gtk_tree_model_iter_nth_child")
	core.PuregoSafeRegister(&XGtkTreeModelIterParent, lib, "gtk_tree_model_iter_parent")
	core.PuregoSafeRegister(&XGtkTreeModelIterPrevious, lib, "gtk_tree_model_iter_previous")
	core.PuregoSafeRegister(&XGtkTreeModelRefNode, lib, "gtk_tree_model_ref_node")
	core.PuregoSafeRegister(&XGtkTreeModelRowChanged, lib, "gtk_tree_model_row_changed")
	core.PuregoSafeRegister(&XGtkTreeModelRowDeleted, lib, "gtk_tree_model_row_deleted")
	core.PuregoSafeRegister(&XGtkTreeModelRowHasChildToggled, lib, "gtk_tree_model_row_has_child_toggled")
	core.PuregoSafeRegister(&XGtkTreeModelRowInserted, lib, "gtk_tree_model_row_inserted")
	core.PuregoSafeRegister(&XGtkTreeModelRowsReordered, lib, "gtk_tree_model_rows_reordered")
	core.PuregoSafeRegister(&XGtkTreeModelRowsReorderedWithLength, lib, "gtk_tree_model_rows_reordered_with_length")
	core.PuregoSafeRegister(&XGtkTreeModelUnrefNode, lib, "gtk_tree_model_unref_node")

}
