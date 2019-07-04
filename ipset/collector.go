package ipset

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// FileMap is a map to classify the IPSet files
type FileMap map[string][]*IPSet

// IPSet file representation
type IPSet struct {
	Path string
	Name string
	Size int64
	Ext  string
}

// TraverseIPSetDir Walks recursively in the ipset volume gathering the interested files
func TraverseIPSetDir(dir string) (*FileMap, error) {
	log.Infof("Traversing the dir %s...", dir)
	fileMap := make(FileMap)

	currParent := ""
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "Failed to walk recursively to the dir")
		}

		if info.IsDir() {
			dirPath := filepath.Dir(path)
			parent := filepath.Base(dirPath)

			name := info.Name()
			if (name != "errors" && name != "history") && (parent != "errors" && parent != "history") {
				currParent = name
				return nil
			}
		}

		ext := filepath.Ext(path)
		if ext != ".netset" && ext != ".ipset" {
			return nil
		}

		if currParent == "" {
			return nil
		}

		fileMap[currParent] = append(fileMap[currParent], &IPSet{
			Path: path,
			Name: info.Name(),
			Size: info.Size(),
			Ext:  ext,
		})

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "Failed to explore the IPSet list directory")
	}

	return &fileMap, nil
}

// ReadIPSetFiles reads the IPSet files concurrently, sending the formatted IPs to the ipset-api
func ReadIPSetFiles(fileMap *FileMap) error {
	wg := sync.WaitGroup{}

	wg.Add(len(*fileMap))

	for k, v := range *fileMap {
		fmt.Printf("Key: %s \n", k)
		go func(ipset []*IPSet) {
			fmt.Printf("List len: %d \n", len(ipset))
			for _, item := range ipset {
				item.Name = fmt.Sprintf(item.Name, "+")
			}
			wg.Done()
		}(v)
	}

	fmt.Printf("Waiting\n")
	wg.Wait()

	fmt.Printf("Done")

	return nil
}
