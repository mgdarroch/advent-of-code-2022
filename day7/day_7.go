package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node interface {
	GetName() string
	GetSize() int
	GetParent() *Directory
}

type File struct {
	name   string
	size   int
	parent *Directory
}

func (f *File) GetParent() *Directory {
	return f.parent
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

type Directory struct {
	name     string
	size     int
	children []Node
	parent   *Directory
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
		size: 0,
	}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	totalSize := d.size
	for _, child := range d.children {
		totalSize += child.GetSize()
	}
	return totalSize
}

func (d *Directory) GetParent() *Directory {
	return d.parent
}

func (d *Directory) AddChild(child Node) {
	d.children = append(d.children, child)
	if dir, ok := child.(*Directory); ok {
		dir.parent = d
	}
}

func GetDirectoriesWithSizeAtMost(root *Directory, size int) []*Directory {
	results := []*Directory{}
	var traverse func(node *Directory)
	traverse = func(node *Directory) {
		if node.GetSize() <= size {
			results = append(results, node)
		}
		for _, child := range node.children {
			if dir, ok := child.(*Directory); ok {
				traverse(dir)
			}
		}
	}
	traverse(root)
	return results
}

func PrintFileSizes(root *Directory) {
	var traverse func(node *Directory)
	traverse = func(node *Directory) {
		for _, child := range node.children {
			if dir, ok := child.(*Directory); ok {
				traverse(dir)
			} else {
				fmt.Printf("%s: %d\n", child.GetName(), child.GetSize())
			}
		}
	}

	traverse(root)
}

func PrintChildTypes(root *Directory) {
	var traverse func(node *Directory)
	traverse = func(node *Directory) {
		for _, child := range node.children {
			if dir, ok := child.(*Directory); ok {
				fmt.Printf("%s: directory\n", dir.GetName())
				traverse(dir)
			} else {
				fmt.Printf("%s: file\n", child.GetName())
			}
		}
	}
	traverse(root)
}

func getDirectories(root *Directory) []*Directory {
	list := []*Directory{root}
	var traverse func(node *Directory)
	traverse = func(node *Directory) {
		for _, child := range node.children {
			if dir, ok := child.(*Directory); ok {
				list = append(list, dir)
				traverse(dir)
			} else {
				continue
			}
		}
	}
	traverse(root)
	return list
}

func PrintFileChildren(root *Directory) {
	var traverse func(node *Directory)
	traverse = func(node *Directory) {
		for _, child := range node.children {
			if dir, ok := child.(*Directory); ok {
				fmt.Printf("%s: directory - size:%d\n", dir.GetName(), dir.GetSize())
				traverse(dir)
			} else {
				fmt.Printf("%s: file - size: %d\n", child.GetName(), child.GetSize())
			}
		}
	}
	traverse(root)
}

func PrintTree(root *Directory) {
	var traverse func(node *Directory, indent int)
	traverse = func(node *Directory, indent int) {
		fmt.Printf("%*sdir %s - size:%d\n", indent, "", node.GetName(), node.GetSize())
		for _, child := range node.children {
			if dir, ok := child.(*Directory); ok {
				traverse(dir, indent+2)
			} else {
				fmt.Printf("%*sfile %s - size:%d\n", indent+2, "", child.GetName(), child.GetSize())
			}
		}
	}
	traverse(root, 0)
}

func countNodes(root Node) int {
	queue := []Node{root}
	nodeCount := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodeCount++
		if dir, ok := node.(*Directory); ok {
			queue = append(queue, dir.children...)
		}
	}
	return nodeCount
}

func getFullPath(start *Directory) string {
	if start.GetName() == "/" {
		return "/"
	}

	fullPathString := ""
	var traverse func(node *Directory)
	traverse = func(node *Directory) {
		if node.GetName() == "/" {
			fullPathString = "/" + fullPathString
			return
		} else {
			fullPathString = node.GetName() + "/" + fullPathString
			traverse(node.GetParent())
		}
	}

	traverse(start)
	return fullPathString
}

func main() {

	f, err := os.Open("day7/input_d7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	root := NewDirectory("/")
	directories := map[string]*Directory{
		"/": root,
	}
	var currentDir *Directory
	fullPathString := ""
	fileCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			parts = parts[1:]
		}

		switch parts[0] {
		case "cd":
			if parts[1] == ".." {
				if currentDir.parent != nil {
					x := len(fullPathString) - len(currentDir.GetName()) - 1
					fullPathString = fullPathString[:x]
					currentDir = currentDir.parent
				}
			} else if parts[1] == "/" {
				currentDir = root
				fullPathString = "/"
			} else {
				fullPathString = fullPathString + parts[1]
				dir, ok := directories[fullPathString]
				if !ok {
					dir = NewDirectory(parts[1])
					currentDir.AddChild(dir)
					fullPathString = getFullPath(dir)
					directories[fullPathString] = dir
				} else {
				}
				currentDir = dir
				continue
			}
		case "ls":
			continue
		case "dir":
			tmpFullPath := fullPathString + parts[1]
			dir, ok := directories[fullPathString]
			if !ok {
				dir = NewDirectory(parts[1])
				currentDir.AddChild(dir)
				tmpFullPath = getFullPath(currentDir) + parts[1]
				directories[tmpFullPath] = dir
			}
			continue
		default:
			size, _ := strconv.Atoi(parts[0])
			file := &File{
				name: parts[1],
				size: size,
			}
			currentDir.AddChild(file)
			fileCount++
		}
	}

	totalEntities := fileCount + len(directories)
	fmt.Println("Total Entities:", totalEntities)
	fmt.Println("Node Count:", countNodes(root))
	directoriesLessThan := GetDirectoriesWithSizeAtMost(root, 100000)
	sum := 0
	for _, v := range directoriesLessThan {
		sum += v.GetSize()
	}
	fmt.Println("Sum:", sum)

	dirList := getDirectories(root)
	smallestDir := findSmallestDirectoryToDelete(root, dirList)

	fmt.Println("Smallest Dir Size:", smallestDir.GetSize())

}

func findSmallestDirectoryToDelete(root *Directory, dirList []*Directory) *Directory {
	max := 70000000
	required := 30000000
	space := max - root.GetSize()
	smallestDir := root
	for _, v := range dirList {
		fmt.Printf("Directory: %s Size: %d\n", v.GetName(), v.GetSize())
		dirSize := v.GetSize()
		tmp := space + dirSize
		fmt.Println("Space + dirSize =", tmp)
		if tmp > required {
			if dirSize < smallestDir.GetSize() {
				smallestDir = v
			}
		}
	}
	return smallestDir
}
