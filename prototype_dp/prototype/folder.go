package prototype

import "fmt"

type Folder struct {
	Children []InterfaceNode
	Name     string
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, i := range f.Children {
		i.Print(s + s)
	}
}

func (f *Folder) Clone() InterfaceNode {
	cloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildren []InterfaceNode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
