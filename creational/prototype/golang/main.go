/*
	Прототип — это порождающий паттерн, который позволяет копировать объекты любой
	сложности без привязки к их конкретным классам.

	Все классы-прототипы имеют общий интерфейс. Поэтому вы можете копировать объекты,
	не обращая внимания на их конкретные типы и всегда быть уверены, что получите
	точную копию. Клонирование совершается самим объектом-прототипом, что позволяет
	ему скопировать значения всех полей, даже приватных.
*/
package main

import "fmt"

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder2")
	cloneFolder.print("  ")
}

// Printing hierarchy for Folder2
//  Folder2
//    Folder1
//        File1
//    File2
//    File3
//
// Printing hierarchy for clone Folder2
//  Folder2_clone
//    Folder1_clone
//        File1_clone
//    File2_clone
//    File3_clone
