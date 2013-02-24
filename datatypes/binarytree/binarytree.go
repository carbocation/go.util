// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2013 James Pirruccello <james@carbocation.com>
// Use of this source code is governed by the MIT license

// Package binarytree implements a binary tree, derived from the container/list code
//
// To iterate over a tree (where tree is a *Tree):
//	for e := tree.Root(); e != nil; e = e.Left() {
//		// do something with e.Value
//	}
//
package binarytree

// Tree is an element in the tree.
type Tree struct {
	// Parent, left, and right pointers in the tree of elements.
	// The root of the tree has parent = nil
    // Each leaf of a tree has left = nil.
	parent, left, right *Tree

	// The contents of this tree element.
	Value interface{}
}

func New(value interface{}) *Tree {
    return &Tree{Value: value}
}

// Left returns the left tree element or nil.
func (e *Tree) Left() *Tree { return e.left }

// Right returns the right tree element or nil.
func (e *Tree) Right() *Tree { return e.right }

// Parent returns the parent tree element or nil.
func (e *Tree) Parent() *Tree { return e.parent }

// PushLeft inserts a left node's value under the specified Tree and returns a new Tree containing the value.
func (e *Tree) PushLeft(value interface{}) *Tree {
    return e.pushTree("left", value)
}

// PushRight inserts a right node's value under the specified Tree and returns a new Tree containing the value.
func (e *Tree) PushRight(value interface{}) *Tree {
    return e.pushTree("right", value)
}

// Wrapped 
func (e *Tree) pushTree(direction string, value interface{}) *Tree {
    // Create the new element el, record that the parent is e
    el := &Tree{e, nil, nil, value}

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

// Traverse the tree and send output to a channel
func Walk(el *Tree, ch chan interface{}) {
    if el == nil {
        return
    }

    ch <- el.Value
    Walk(el.left, ch)
    Walk(el.right, ch)
}

// Wrapper for the Walk method that creates a string channel for output
func Walker(e *Tree) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		Walk(e, ch)
		close(ch)
	}()
	return ch
}
