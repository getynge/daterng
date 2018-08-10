package main

import (
	"os"
	"time"
	"fmt"
	"io/ioutil"
	"path"
	"math/rand"
)

var mindate = time.Date(2015, time.September, 20, 0, 0, 0, 0, time.UTC)
var maxdiff = time.Since(mindate).Nanoseconds()

func handleDir(dir string) (err error){
	contents, err := ioutil.ReadDir(dir)

	for _, file := range contents {
		joined := path.Join(dir, file.Name())
		if file.Mode().IsDir() {
			fmt.Printf("Recursively rewriting directory %s\n", joined)
			if err = handleDir(joined); err != nil {
				fmt.Printf("There were errors editing %s will continue\n", joined)
			} else {
				fmt.Printf("Successfully edited all the timestamps in directory %s\n", joined)
			}
		} else {
			if err = handleFile(joined); err != nil {
				fmt.Printf("Failed to edit stamp of file %s, will continue\n", joined)
			} else {
				fmt.Println("Success")
			}
		}
	}
	return
}

func handleFile(path string) error{
	diff := time.Duration(rand.Int63n(maxdiff))
	stamp := mindate.Add(diff)
	fmt.Printf("Changing timestamp of file %s to %s\n", path, stamp)
	return os.Chtimes(path, stamp, stamp)
}

func main(){
	dirs := os.Args[1:]

	for _, dir := range dirs {
		file, err := os.Stat(dir)
		if err != nil {
			fmt.Printf("Error: File or directory %s does not exist\n", dir)
			continue
		}

		if file.Mode().IsDir() {
			if err := handleDir(dir); err != nil {
				fmt.Printf("There were errors editing %s will continue\n", dir)
			} else {
				fmt.Printf("Successfully edited all the timestamps in directory %s\n", dir)
			}
		} else {
			if err = handleFile(dir); err != nil {
				fmt.Printf("Failed to edit stamp of file %s, will continue\n", dir)
			} else {
				fmt.Println("Success")
			}
		}
	}
}