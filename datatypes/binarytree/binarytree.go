// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Portions copyright 2013 James Pirruccello <james@carbocation.com>

// Package binarytree implements a binary tree, derived from the container/list code
//
// To iterate over a tree (where l is a *Tree):
//	for e := l.Front(); e != nil; e = e.Left() {
//		// do something with e.Value
//	}
//
package binarytree

// Element is an element in the tree.
type Element struct {
	// Parent, left, and right pointers in the tree of elements.
	// The root of the tree has parent = nil
    // Each leaf of a tree has left = nil.
	parent, left, right *Element

	// The tree to which this element belongs.
	tree *Tree

	// The contents of this tree element.
	Value interface{}
}

// Left returns the left tree element or nil.
func (e *Element) Left() *Element { return e.left }

// Right returns the right tree element or nil.
func (e *Element) Right() *Element { return e.right }

// Parent returns the parent tree element or nil.
func (e *Element) Parent() *Element { return e.parent }

// Tree represents a doubly linked tree.
// The zero value for Tree is an empty tree ready to use.
type Tree struct {
	front, back *Element
	len         int
}

// Init initializes or clears a Tree.
func (l *Tree) Init() *Tree {
	l.front = nil
	l.back = nil
	l.len = 0
	return l
}

// New returns an initialized tree.
func New() *Tree { return new(Tree) }

// Front returns the first element in the tree.
func (l *Tree) Front() *Element { return l.front }

// Back returns the last element in the tree.
func (l *Tree) Back() *Element { return l.back }

// Remove removes the element from the tree
// and returns its Value.
func (l *Tree) Remove(e *Element) interface{} {
	l.remove(e)
	e.tree = nil // do what remove does not
	return e.Value
}

// remove the element from the tree, but do not clear the Element's tree field.
// This is so that other Tree methods may use remove when relocating Elements
// without needing to restore the tree field.
func (l *Tree) remove(e *Element) {
	if e.tree != l {
		return
	}
	if e.parent == nil {
		l.front = e.left
	} else {
		e.parent.left = e.left
	}
	if e.left == nil {
		l.back = e.parent
	} else {
		e.left.parent = e.parent
	}

	e.parent = nil
	e.left = nil
	l.len--
}

func (l *Tree) insertBefore(e *Element, mark *Element) {
	if mark.parent == nil {
		// new front of the tree
		l.front = e
	} else {
		mark.parent.left = e
	}
	e.parent = mark.parent
	mark.parent = e
	e.left = mark
	l.len++
}

func (l *Tree) insertAfter(e *Element, mark *Element) {
	if mark.left == nil {
		// new back of the tree
		l.back = e
	} else {
		mark.left.parent = e
	}
	e.left = mark.left
	mark.left = e
	e.parent = mark
	l.len++
}

func (l *Tree) insertFront(e *Element) {
	if l.front == nil {
		// empty tree
		l.front, l.back = e, e
		e.parent, e.left = nil, nil
		l.len = 1
		return
	}
	l.insertBefore(e, l.front)
}

func (l *Tree) insertBack(e *Element) {
	if l.back == nil {
		// empty tree
		l.front, l.back = e, e
		e.parent, e.left = nil, nil
		l.len = 1
		return
	}
	l.insertAfter(e, l.back)
}

// PushFront inserts the value at the front of the tree and returns a new Element containing the value.
func (l *Tree) PushFront(value interface{}) *Element {
	e := &Element{nil, nil, nil, l, value}
	l.insertFront(e)
	return e
}

// PushBack inserts the value at the back of the tree and returns a new Element containing the value.
func (l *Tree) PushBack(value interface{}) *Element {
	e := &Element{nil, nil, nil, l, value}
	l.insertBack(e)
	return e
}

// InsertBefore inserts the value immediately before mark and returns a new Element containing the value.
func (l *Tree) InsertBefore(value interface{}, mark *Element) *Element {
	if mark.tree != l {
		return nil
	}
	e := &Element{nil, nil, nil, l, value}
	l.insertBefore(e, mark)
	return e
}

// InsertAfter inserts the value immediately after mark and returns a new Element containing the value.
func (l *Tree) InsertAfter(value interface{}, mark *Element) *Element {
	if mark.tree != l {
		return nil
	}
	e := &Element{nil, nil, nil, l, value}
	l.insertAfter(e, mark)
	return e
}

// MoveToFront moves the element to the front of the tree.
func (l *Tree) MoveToFront(e *Element) {
	if e.tree != l || l.front == e {
		return
	}
	l.remove(e)
	l.insertFront(e)
}

// MoveToBack moves the element to the back of the tree.
func (l *Tree) MoveToBack(e *Element) {
	if e.tree != l || l.back == e {
		return
	}
	l.remove(e)
	l.insertBack(e)
}

// Len returns the number of elements in the tree.
func (l *Tree) Len() int { return l.len }

// PushBackTree inserts each element of ol at the back of the tree.
func (l *Tree) PushBackTree(ol *Tree) {
	last := ol.Back()
	for e := ol.Front(); e != nil; e = e.Left() {
		l.PushBack(e.Value)
		if e == last {
			break
		}
	}
}

// PushFrontTree inserts each element of ol at the front of the tree. The ordering of the passed tree is preserved.
func (l *Tree) PushFrontTree(ol *Tree) {
	first := ol.Front()
	for e := ol.Back(); e != nil; e = e.Parent() {
		l.PushFront(e.Value)
		if e == first {
			break
		}
	}
}
