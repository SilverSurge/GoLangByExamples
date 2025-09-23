package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating file ...")
	f, err := os.Create(p)
	if err != nil {
		ps := "createFile:" + err.Error()
		panic(ps)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing to file ...")
	fmt.Fprintln(f, "writing to file using Fprintfln")
}

func closeFile(f *os.File) {
	fmt.Println("closing file ...")
	err := f.Close()

	if err != nil {
		panic("closeFile:" + err.Error())
	}
}
