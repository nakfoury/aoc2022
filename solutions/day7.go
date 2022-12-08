package solutions

import (
	"bufio"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	CMD_RE = regexp.MustCompile(`\$ (.+) (.+)*`)
)

func init() {
	DAY_TO_FUNC[7] = Day7
}

type File interface {
	getSize() int
	isDir() bool
	getParent() File
}

type DataFile struct {
	name   string
	size   int
	parent File
}

func (d DataFile) isDir() bool {
	return false
}

func (d DataFile) getSize() int {
	return d.size
}

func (d DataFile) getParent() File {
	return d.parent
}

type Dir struct {
	name   string
	files  []File
	parent File
}

func (d Dir) isDir() bool {
	return true
}

func (d Dir) getSize() int {
	totalSize := 0
	for _, f := range d.files {
		totalSize += f.getSize()
	}
	return totalSize
}

func (d Dir) getParent() File {
	return d.parent
}

func (d *Dir) mkdir(name string) *Dir {
	newDir := &Dir{
		name:   name,
		parent: d,
	}
	d.files = append(d.files, newDir)
	return newDir
}

func (d *Dir) touch(name string, size int) {
	d.files = append(d.files, DataFile{
		name:   name,
		size:   size,
		parent: d,
	})
}

type FileSystem struct {
	root    *Dir
	workdir *Dir
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		root: &Dir{
			name: "/",
		},
	}
}

func (f *FileSystem) ls(line string) string {
	s := strings.Split(line, " ")
	if s[0] == "dir" {
		f.workdir.mkdir(s[1])
	} else {
		size, _ := strconv.Atoi(s[0])
		f.workdir.touch(s[1], size)
	}
	return ""
}

func (f *FileSystem) cd(dir string) string {
	if dir == "/" {
		f.workdir = f.root
	} else if dir == ".." {
		parent := f.workdir.getParent()
		if parent != nil {
			f.workdir = parent.(*Dir)
		} else {
			parent = Dir{
				name: dir,
			}
			f.workdir.parent = parent
			f.workdir = parent.(*Dir)
		}
	} else {
		f.workdir = f.workdir.mkdir(dir)
	}

	return ""
}

func Day7(part2 bool, inp *bufio.Scanner) any {
	result := 0
	fs := NewFileSystem()

	listing := false
	for inp.Scan() {
		line := inp.Text()

		if listing && line[0] != '$' {
			fs.ls(line)
			continue
		} else {
			listing = false
		}

		if line[:4] == "$ cd" {
			fs.cd(line[5:])
		} else if line[:4] == "$ ls" {
			listing = true
		}
	}

	for _, f := range fs.root.files {
		if f.getSize() <= 100000 {
			result += f.getSize()
		}
	}

	if !part2 {
		result = expand(fs.root)
	} else {
		target := 30000000 - (70000000 - fs.root.getSize())
		result = smallest(fs.root, target)
	}

	return result
}

func expand(f File) int {
	subresult := 0
	if f.isDir() {
		if f.getSize() <= 100000 {
			subresult += f.getSize()
		}
		for _, child := range f.(*Dir).files {
			subresult += expand(child)
		}
	}
	return subresult
}

func smallest(f File, target int) int {
	subresult := math.MaxInt
	if f.isDir() {
		if f.getSize() >= target {
			subresult = min(subresult, f.getSize())
		}
		for _, child := range f.(*Dir).files {
			subresult = min(subresult, smallest(child, target))
		}
	}
	return subresult
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
