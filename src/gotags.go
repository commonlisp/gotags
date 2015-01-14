package main

import (
	"fmt"
	"os"
	"bytes"
	"go/parser"
	"go/token"
	)

func Foo() {}
func bar() {}

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("Usage:")
	}
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, os.Args[1], func (inf os.FileInfo) bool { return true },
					0)
	if err != nil {
		fmt.Printf("Error %s %s\n", err.Error(), os.Args[1])
		return
	}
	for _, pkg := range pkgs {
		for fn, f := range pkg.Files {
			var buf bytes.Buffer
			for scopeItem, obj := range f.Scope.Objects {
				buf.WriteString(fmt.Sprintf("%s,%s,%d\n", scopeItem, fn, obj.Pos()))
			}
			fmt.Printf("\x0c\n%s,%d\n", fn, buf.Len())
		}
	}
}

