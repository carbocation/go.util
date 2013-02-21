package binarytree

import (
    "container/list"
)

type Element struct {
    list.Element "We incorporate all parts of the list.Element"

    parent *Element "In addition, we create a parent element"
}

// Prev returns the parent list element or nil.
func (e *Element) Parent() *Element { return e.parent }