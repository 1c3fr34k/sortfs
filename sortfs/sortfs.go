package sortfs

import (
	"fmt"
	"os"
	"path/filepath"
)

type sortFs struct {
	rootPath       string
	fileNames      []string
	filePaths      []string
	folderNames    []string
	folderPaths    []string
	extensionNames []string
	extensionPaths []string
}

func New(root string) *sortFs {
	return &sortFs{rootPath: root}
}

func (s *sortFs) Sort() {
	s.filePaths, s.fileNames, _ = getFolderAndFilePaths(s.rootPath)
	s.extensionNames = getDestinctFileExtensions(s.filePaths)

	fmt.Println(getDestinctFileExtensions(s.filePaths))

	s.extensionPaths = createFoldersForExtensions(s.rootPath, s.extensionNames)
	// moveFilesToExtensionFolder(files, extensionPaths)

}

func getFolderAndFilePaths(root string) (filePaths, fileNames, folderPaths []string) {
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path != root {
			if filepath.Dir(path) != root {
				return filepath.SkipDir
			} else {
				if entry.IsDir() {
					folderPaths = append(folderPaths, filepath.Join(path))
				} else {
					filePaths = append(filePaths, path)
					fileNames = append(fileNames, entry.Name())
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, nil, nil
	}

	return filePaths, fileNames, folderPaths
}

func getDestinctFileExtensions(files []string) (extensions []string) {
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
	for _, file := range fileName {
		for _, ext := range extensionPath {
			if filepath.Ext(string(file)) == filepath.Ext(string(ext)) {
				newFolderName = filepath.Join(string(ext), string(file))
			}
		}
	}
	return newFolderName
}

func moveFilesToExtensionFolder(filePaths []string, extensionPaths []string) {

	for _, file := range filePaths {
		for _, extension := range extensionPaths {
			if extension == filepath.Ext(file) {
				err := os.Rename(file, extension)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
