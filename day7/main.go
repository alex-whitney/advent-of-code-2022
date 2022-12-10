package main

import (
	"strconv"
	"strings"

	"github.com/alex-whitney/advent-of-code-2022/lib"
)

type file struct {
	name string
	size int
}

func newFile(line string) (*file, error) {
	parts := strings.Split(line, " ")
	size, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	return &file{
		name: parts[1],
		size: size,
	}, nil
}

type directory struct {
	name        string
	files       []*file
	directories []*directory
	parent      *directory
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:        name,
		parent:      parent,
		files:       make([]*file, 0),
		directories: make([]*directory, 0),
	}
}

func (d *directory) getOrCreateDirectory(name string) *directory {
	for _, d := range d.directories {
		if d.name == name {
			return d
		}
	}

	dir := newDirectory(name, d)
	d.directories = append(d.directories, dir)
	return dir
}

func (d *directory) size() int {
	s := 0
	for _, file := range d.files {
		s += file.size
	}
	return s
}

func (d *directory) recursiveSize() int {
	s := d.size()
	for _, subDir := range d.directories {
		s += subDir.recursiveSize()
	}
	return s
}

type Today struct {
	root *directory
}

func (d *Today) Init(input string) error {
	raw, err := lib.ReadStringFile(input)
	if err != nil {
		return err
	}

	d.root = newDirectory("", nil)
	cd := d.root

	for _, line := range raw {
		if strings.HasPrefix(line, "$ cd ") {
			instr := strings.TrimLeft(line, "$ cd ")

			if instr == "/" {
				// going to cheat a little, only the first instruction is an absolute path
			} else if instr == ".." {
				cd = cd.parent
			} else {
				newDir := cd.getOrCreateDirectory(instr)
				cd = newDir
			}
		} else if line == "$ ls" {
			// ignore - CD and LS are the only commands. anything else is an output from ls
		} else if strings.HasPrefix(line, "dir") {
			dirName := strings.TrimLeft(line, "dir ")
			cd.getOrCreateDirectory(dirName)
		} else {
			file, err := newFile(line)
			if err != nil {
				return err
			}
			cd.files = append(cd.files, file)
		}
	}

	return nil
}

func (d *Today) Part1() (string, error) {
	threshold := 100000
	sum := 0

	directories := []*directory{d.root}

	for len(directories) > 0 {
		dir := directories[0]
		directories = directories[1:]

		size := dir.recursiveSize()
		if size <= threshold {
			sum += size
		}

		if len(dir.directories) > 0 {
			directories = append(directories, dir.directories...)
		}
	}

	return strconv.Itoa(sum), nil
}

func (d *Today) Part2() (string, error) {
	diskSize := 70000000
	requiredSpace := 30000000
	rootSize := d.root.recursiveSize()

	directories := []*directory{d.root}

	minFreeSpace := diskSize
	deletedSize := 0

	for len(directories) > 0 {
		dir := directories[0]
		directories = directories[1:]

		size := dir.recursiveSize()
		freedSpace := diskSize - rootSize + size
		if freedSpace > requiredSpace && freedSpace < minFreeSpace {
			minFreeSpace = freedSpace
			deletedSize = size
		}

		if len(dir.directories) > 0 {
			directories = append(directories, dir.directories...)
		}
	}

	return strconv.Itoa(deletedSize), nil
}

func main() {
	day := &Today{}
	lib.Run(day)
}
