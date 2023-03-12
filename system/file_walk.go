package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var imgSuffix = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
	".tiff": true,
	".eps":  true,
}

func printUsage() {
	fmt.Printf(`Find images

Usage:
	%s [path to find]
`, os.Args[0])
}

func traverse() error {
	if len(os.Args) == 1 {
		printUsage()

		return errors.New("no arguments was passed")
	}
	root := os.Args[1]

	err := filepath.Walk(root,
		func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				if info.Name() == "_build" {
					return filepath.SkipDir
				}

				return nil
			}

			extension := strings.ToLower(filepath.Ext(info.Name()))
			if imgSuffix[extension] {
				rel, err := filepath.Rel(root, path)
				if err != nil {
					return nil
				}

				fmt.Printf("rel: %v\n", rel)
			}

			return nil
		})

	if err != nil {
		return fmt.Errorf("failed to traverse filepath: %w", err)
	}

	return nil
}
