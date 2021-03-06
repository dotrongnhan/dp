package main

import (
	"fmt"
	"prototype/prototype"
)

func main() {
	file1 := &prototype.File{Name: "File 1"}
	file2 := &prototype.File{Name: "File 2"}
	file3 := &prototype.File{Name: "File 3"}
	folder1 := &prototype.Folder{
		Children: []prototype.InterfaceNode{file1},
		Name: "Folder 1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.InterfaceNode{folder1, file2, file3},
		Name: "Folder 2",
	}

	fmt.Println("\n Printing for Folder 2")
	folder2.Print("    ")
	cloneFolder := folder2.Clone()
	fmt.Println("\n Printing for clone Folder 2")
	cloneFolder.Print("    ")
}
