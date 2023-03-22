// Package main is the entry point of this repository
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func ls() {
	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if d.Name()[0] == '.' {
				return fs.SkipDir
			}

			return nil
		}

		files = append(files, path)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("files: %v\n", files)
}

func walk() {
	filepath.Walk("root", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		return nil
	})
}
