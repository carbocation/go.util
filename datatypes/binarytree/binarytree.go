// Copyright 2013 James Pirruccello <james@carbocation.com>
// All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package binarytree implements a binary tree
//
// To iterate over a tree (where tree is a *Tree), 
// use the Walk() function or a similar method.
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

// Set the left tree element or nil
func (e *Tree) SetLeft(replacement *Tree) { e.left = replacement }

// Set the right tree element or nil
func (e *Tree) SetRight(replacement *Tree) { e.right = replacement }

// Set the parent tree element or nil
func (e *Tree) SetParent(replacement *Tree) { e.parent = replacement }

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

// Find the root node of the entire tree
func (el *Tree) Root() *Tree {
	if el.Parent() == nil {
		return el
	}
	
	return el.Parent().Root()
}

func (el *Tree) Len() int {
	r := el.Root()
	
	return r.count()
}

func (el *Tree) count() int {
	if el == nil {
		return 0
	}
	
	return 1 + el.Left().count() + el.Right().count()
}
/*	
	ch := make(chan interface{})
	go func() {
		el.Traverse(ch)
		close(ch)
	}()
	return ch
}
*/

func (el *Tree) Traverse(ch chan interface{}) {
	if el == nil {
		return
	}
	
	ch <- el.Value
	el.parent.Traverse(ch) //Also traverse up the parent
	el.left.Traverse(ch)
	el.right.Traverse(ch)
} 

// Walk the tree and send output to a channel
// Note that this is simply one reference implementation. For some applications, 
// e.g., threaded discussion, something like this would be implemented in the view.
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
