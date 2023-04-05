/*
	Компоновщик — это структурный паттерн, который позволяет создавать дерево объектов
	и работать с ним так же, как и с единичным объектом.

	Компоновщик давно стал синонимом всех задач, связанных с построением дерева объектов.
	Все операции компоновщика основаны на рекурсии и «суммировании» результатов на
	ветвях дерева.
*/
package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}

// Searching recursively for keyword rose in folder Folder2
// Searching for keyword rose in file File2
// Searching for keyword rose in file File3
// Searching recursively for keyword rose in folder Folder1
// Searching for keyword rose in file File1
