package main

import "composite/composite"

func main() {
	file1 := &composite.File{Name: "file1"}
	file2 := &composite.File{Name: "file2"}
	file3 := &composite.File{Name: "file3"}

	folder1 := &composite.Folder{Name: "Folder1"}
	folder1.Add(file1)
	folder2 := &composite.Folder{Name: "Folder2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
