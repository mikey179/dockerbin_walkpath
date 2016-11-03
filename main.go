package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Let's see what's inside...")
	filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() {
			fmt.Printf("Directory: %s", filepath.ToSlash(path))
		} else {
			fmt.Printf("File: %s", filepath.ToSlash(path))
		}

		fmt.Println("")
		return nil
	})
}
