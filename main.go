package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"

	//"path/filepath"
	//"strings"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTree1(out, path, printFiles, 0, "")
}

func dirTree1(out io.Writer, path string, printFiles bool, level int, prefix string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	//dirs := make([]os.FileInfo, len(files))
	//for _, f := range files {
	//	if f.IsDir() {
	//		dirs = append(dirs, f)
	//	}
	//}

	sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	for i, f := range files {
		x := prefix + "├───"
		if i == len(files)-1 {
			x = prefix + "└───"
		}

		if f.IsDir() {
			s := fmt.Sprintf("%v%v\n", x, f.Name())

			_, err := out.Write([]byte(prefix + s))
			if err != nil {
				return err
			}

			var nextPrefix string
			if i == len(files)-1 {
				nextPrefix = prefix + "\t"
			} else {
				nextPrefix = prefix + "│\t"
			}

			err = dirTree1(out, path+"/"+f.Name(), printFiles, level+1, nextPrefix)
			if err != nil {
				return err
			}
		} else if printFiles {
			size := f.Size()
			sizeS := fmt.Sprintf("%vb", size)
			if size == 0 {
				sizeS = "empty"
			}

			s := fmt.Sprintf("%v%v (%v)\n", x, f.Name(), sizeS)

			_, err := out.Write([]byte(prefix + s))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
