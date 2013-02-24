// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Portions copyright 2013 James Pirruccello <james@carbocation.com>

// Package binarytree implements a binary tree, derived from the container/list code
//
// To iterate over a tree (where tree is a *Tree):
//	for e := tree.Root(); e != nil; e = e.Left() {
//		// do something with e.Value
//	}
//
package binarytree

import (
//    "fmt"
)

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
	root *Element
	len         int
}

// Init initializes or clears a Tree.
func (tree *Tree) Init() *Tree {
	tree.root = nil
	tree.len = 0
	return tree
}

// Root returns the first element in the tree.
func (tree *Tree) Root() *Element { return tree.root }

// Len returns the number of elements in the tree.
func (tree *Tree) Len() int { return tree.len }

// New returns an initialized tree.
func New(firstElementValue interface{}) *Tree {
    /*
    Create a new tree.
    Require that its first element be passed.
    */
    tree := new(Tree)
    
    e := &Element{nil, nil, nil, tree, firstElementValue}

    tree.root, tree.len = e, 1

    return tree
}

// PushLeft inserts a left node's value under the specified Element and returns a new Element containing the value.
func (e *Element) PushLeft(value interface{}) *Element {
    return e.pushElement("left", value)
}

// PushRight inserts a right node's value under the specified Element and returns a new Element containing the value.
func (e *Element) PushRight(value interface{}) *Element {
    return e.pushElement("right", value)
}

func (e *Element) pushElement(direction string, value interface{}) *Element {
    // Update the tree
    tree := e.tree
    tree.len += 1

    // Create the new element el, record that the parent is e
    el := &Element{e, nil, nil, tree, value}

    //TODO
    // SWITCH/CASE to append to left or right depending on correct output
    // Update the parent element e, record that its left value is el
    switch {
    case direction == "left":
        e.left = el
    
    case direction == "right":
        e.right = el
    }

    // Return the new element
    return el
}

func Walk(el *Element, ch chan string) {
    if el == nil {
        return
    }

    ch <- "<block" + el.Value.(string) + ">"
    ch <- el.Value.(string)
    Walk(el.left, ch)
    Walk(el.right, ch)
    ch <- "</block" + el.Value.(string) + ">"
}

func Walker(e *Element) <-chan string {
	ch := make(chan string)
	go func() {
		Walk(e, ch)
		close(ch)
	}()
	return ch
}










// PushFront inserts the value at the front of the tree and returns a new Element containing the value.
func (tree *Tree) PushFront(value interface{}) *Element {
	e := &Element{nil, nil, nil, tree, value}
	tree.insertFront(e)
	return e
}

// InsertBefore inserts the value immediately before mark and returns a new Element containing the value.
func (tree *Tree) InsertBefore(value interface{}, mark *Element) *Element {
	if mark.tree != tree {
		return nil
	}
	e := &Element{nil, nil, nil, tree, value}
	tree.insertBefore(e, mark)
	return e
}

// InsertAfter inserts the value immediately after mark and returns a new Element containing the value.
func (tree *Tree) InsertAfter(value interface{}, mark *Element) *Element {
	if mark.tree != tree {
		return nil
	}
	e := &Element{nil, nil, nil, tree, value}
	tree.insertAfter(e, mark)
	return e
}




// remove the element from the tree, but do not clear the Element's tree field.
// This is so that other Tree methods may use remove when relocating Elements
// without needing to restore the tree field.
func (tree *Tree) remove(e *Element) {
	if e.tree != tree {
		return
	}
	if e.parent == nil {
		tree.root = e.left
	} else {
		e.parent.left = e.left
	}
	if e.left == nil {
		//tree.back = e.parent
	} else {
		e.left.parent = e.parent
	}

	e.parent = nil
	e.left = nil
	tree.len--
}

func (tree *Tree) insertBefore(e *Element, mark *Element) {
	if mark.parent == nil {
		// new front of the tree
		tree.root = e
	} else {
		mark.parent.left = e
	}
	e.parent = mark.parent
	mark.parent = e
	e.left = mark
	tree.len++
}

func (tree *Tree) insertAfter(e *Element, mark *Element) {
	if mark.left == nil {
		// new back of the tree
		//tree.back = e
	} else {
		mark.left.parent = e
	}
	e.left = mark.left
	mark.left = e
	e.parent = mark
	tree.len++
}

func (tree *Tree) insertFront(e *Element) {
	if tree.root == nil {
		// empty tree
		tree.root = e
		e.parent, e.left = nil, nil
		tree.len = 1
		return
	}
	tree.insertBefore(e, tree.root)
}
