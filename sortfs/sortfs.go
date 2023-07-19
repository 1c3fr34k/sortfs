package sortfs

import (
	"fmt"
	"os"
	"path/filepath"
)

type sortFs struct {
	rootPath       string
	files          map[string]string
	folders        map[string]string
	extensionNames []string
	extensionPaths []string
}

func New(root string) *sortFs {
	return &sortFs{rootPath: root}
}

func (s *sortFs) Sort() {
	s.files, s.folders = getFolderAndFilePaths(s.rootPath)
	s.extensionNames = getDestinctFileExtensions(s.files)
	s.extensionPaths = createFoldersForExtensions(s.rootPath, s.extensionNames)
	moveFilesToExtensionFolder(s.files, s.extensionPaths)

}

func getFolderAndFilePaths(root string) (files, folders map[string]string) {
	files = make(map[string]string)
	folders = make(map[string]string)
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path != root {
			if filepath.Dir(path) != root {
				return filepath.SkipDir
			} else {
				if entry.IsDir() {
					folders[entry.Name()] = path
				} else {
					files[entry.Name()] = path
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return files, folders
}

func getDestinctFileExtensions(files map[string]string) (extensions []string) {
	for _, value := range files {
		ext_exists := false
		ext := filepath.Ext(value)

		for _, value2 := range extensions {
			if value2 == ext {
				ext_exists = true

			}
		}

		if !ext_exists {
			extensions = append(extensions, ext)
		}

	}

	return extensions
}

func createFoldersForExtensions(root string, extensions []string) (extensionPaths []string) {
	for _, value := range extensions {
		if value != "" {
			err := os.MkdirAll(root+"\\"+value, os.ModePerm)

			if err != nil {
				fmt.Println(err)
			}
			extensionPaths = append(extensionPaths, root+"\\"+value)
		}
	}
	return extensionPaths
}

func generateNewFolderName(fileName, extensionPath string) (newFolderName string) {
	newFolderName = filepath.Join(extensionPath, fileName)
	return newFolderName
}

func moveFilesToExtensionFolder(files map[string]string, extensionPaths []string) {
	for kFile, vFile := range files {
		for _, extension := range extensionPaths {
			if filepath.Ext(extension) == filepath.Ext(vFile) {
				err := os.Rename(vFile, generateNewFolderName(kFile, extension))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
