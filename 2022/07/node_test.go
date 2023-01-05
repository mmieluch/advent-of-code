package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Dir(t *testing.T) {
	p := Dir{name: "foo"}
	d := &Dir{
		name:   "bar",
		parent: &p,
		children: []Node{
			&Dir{name: "baz"},
			&File{name: "fandango", size: 123},
		},
	}

	assert.Equal(t, "bar", d.Name())
	assert.Equal(t, &p, d.Parent())
	assert.Equal(t, d.children, d.Children())
	assert.Equal(t, uint(123), d.Size())
	assert.True(t, d.IsDir())
	assert.False(t, d.IsFile())
}

func Test_Dir_AddChild(t *testing.T) {
	d1 := &Dir{name: "foo"}
	assert.Empty(t, d1.Children())

	d1.AddChild(&Dir{name: "child dir"})
	d1.AddChild(&File{name: "child file"})

	cc := d1.Children()
	assert.Len(t, cc, 2)
	assert.Equal(t, "child dir", cc[0].Name())
	assert.Equal(t, "child file", cc[1].Name())

	d2 := &Dir{
		name: "bar",
		children: []Node{
			&Dir{name: "child dir 2"},
			&File{name: "child file 2"},
		},
	}
	d2.AddChild(&Dir{name: "child dir 3"})

	cc2 := d2.Children()
	assert.Len(t, cc2, 3)
	assert.Equal(t, "child dir 2", cc2[0].Name())
	assert.Equal(t, "child file 2", cc2[1].Name())
	assert.Equal(t, "child dir 3", cc2[2].Name())
}

func Test_File(t *testing.T) {
	p := Dir{name: "Gaia"}
	f := &File{
		name:   "foo.txt",
		parent: &p,
		size:   456,
	}
	assert.Equal(t, "foo.txt", f.Name())
	assert.Equal(t, &p, f.Parent())
	assert.Equal(t, uint(456), f.Size())
	assert.True(t, f.IsFile())
	assert.False(t, f.IsDir())
}
