package main

import (
	"fmt"
	"os"
	"strings"
	"log"
	"path/filepath"
	"strconv"
)

//CheckError logs error
func checkError(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("You must specify a search parameter")
		os.Exit(0)
	}

	x := 0
    searchDir := "C:/"
    fileList := []string{}
	
    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		file := filepath.Base(path)
		if strings.Contains(strings.ToLower(file), strings.ToLower(os.Args[1])) {
			fileList = append(fileList, path)
		}
		if x % 20000 == 0 {
			fmt.Print(".")	
		}
		x++
        return nil
    })
	
	checkError(err)
	
	fmt.Println()
	fmt.Println()
    fmt.Println("Found: (" + strconv.Itoa(len(fileList)) + ") files")
	for _, file := range fileList {
        fmt.Println(file)
    }
}