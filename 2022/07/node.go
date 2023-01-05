package main

type Node interface {
	Name() string
	Parent() Parent
	IsDir() bool
	IsFile() bool
	Size() uint
}

type Parent interface {
	Children() []Node
	AddChild(Node)
}

type Dir struct {
	name     string
	parent   Parent
	children []Node
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) Parent() Parent {
	return d.parent
}

func (d *Dir) IsDir() bool {
	return true
}

func (d *Dir) IsFile() bool {
	return false
}

func (d *Dir) Children() []Node {
	return d.children
}

func (d *Dir) AddChild(child Node) {
	d.children = append(d.children, child)
}

func (d *Dir) Size() uint {
	total := uint(0)

	for _, c := range d.Children() {
		total += c.Size()
	}

	return total
}

type File struct {
	name   string
	parent Parent
	size   uint
}

func (f File) Name() string {
	return f.name
}

func (f File) Parent() Parent {
	return f.parent
}

func (f File) IsDir() bool {
	return false
}

func (f File) IsFile() bool {
	return true
}

func (f File) Size() uint {
	return f.size
}
