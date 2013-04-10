package binarytree

import (
	"bytes"
	"html/template"
	"strings"
	"testing"
)

var e *Tree
var expected string

// Populate our unbalanced binary tree and set our expected output when traversed
func init() {
	e = New("Root")
	e.PushLeft("Root:Left")
	e.PushRight("Root:Right")
	e.Left().PushLeft("Root:Left:Left")
	e.Left().PushRight("Root:Left:Right")
	e.Left().Right().PushRight("Root:Left:Right:Right")

	expected = "Root::Root:Left::Root:Left:Left::Root:Left:Right::Root:Left:Right:Right::Root:Right::"
}

// In this test, we are validating an unbalanced binary tree
func TestTree(t *testing.T) {

	//Get a channel with our output and print it out line by line
	x := Walker(e)
	val := ""
	for y := range x {
		val += y.(string) + "::"
	}

	if val != expected {
		t.Errorf("Unexpected value. Received %s, expected %s.", val, expected)
	}
}

//This test derives from Rob Pike's example on the go-nuts mailing list:
// https://groups.google.com/d/msg/golang-nuts/wk5IyGzHQf8/BeWnz82qEEwJ
func TestTreeTemplate(t *testing.T) {
	const treeTemplate = `
        {{define "tree"}}
                {{.Value}}::
                {{with .Left}}
                        {{template "tree" .}}
                {{end}}
                {{with .Right}}
                        {{template "tree" .}}
                {{end}}
        {{end}}
	`

	// Build and parse the set (of one element).
	set := template.Must(template.New("tree").Parse(treeTemplate)) // ALWAYS CHECK ERRORS!!!!!

	var b bytes.Buffer
	// Use set.Execute, starting with the "tree" template.
	// To see the output directly, instead of &b use os.Stdout.
	err := set.Execute(&b, e) // ALWAYS CHECK ERRORS!!!!!
	if err != nil {
		t.Fatal("exec error:", err)
	}
	// This hoo-hah is to make the comparison easy and clear.
	stripSpace := func(r rune) rune {
		if r == '\t' || r == ' ' || r == '\n' {
			return -1
		}
		return r
	}
	result := strings.Map(stripSpace, b.String())
	//const expect = "[1[2[3[4]][5[6]]][7[8[9]][10[11]]]]"
	const expect = "Root::Root:Left::Root:Left:Left::Root:Left:Right::Root:Left:Right:Right::Root:Right::"
	if result != expect {
		t.Errorf("expected %q got %q", expect, result)
	}
}
