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
	list := make(map[string]map[string]int64)

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
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
