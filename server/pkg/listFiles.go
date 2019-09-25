package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ListFiles .
func ListFiles() map[string]map[string]int64 {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(dir)

	subDirToSkip := "cert"

	list := make(map[string]map[string]int64)

	fmt.Println("On Unix:")
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		if !info.IsDir() {
			fmt.Println(info.Name())

			list[info.Name()] = map[string]int64{path: info.Size()}
			//return nil
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
		//return "err"
	}

	return list
}
