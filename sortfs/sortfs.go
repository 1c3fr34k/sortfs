package sortfs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type sortFs struct {
	rootPath       string
	files          map[string]string
	folders        map[string]string
	extensionNames []string
	extensionPaths []string
}

// Init a new sortFs struct
func New(root string) (*sortFs, error) {
	if !pathDoesExist(root) {
		return nil, errors.New("path does not exist")

	}
	return &sortFs{rootPath: root}, nil

}

// Entry point
func (s *sortFs) Sort() {
	s.files, s.folders = getFolderAndFilePaths(s.rootPath)
	s.extensionNames = getDestinctFileExtensions(s.files)
	s.extensionPaths = createFoldersForExtensions(s.rootPath, s.extensionNames)
	moveFilesToExtensionFolder(s.files, s.extensionPaths)

}

// CLI
func CLI() {
	var pathinput string
	fmt.Println("Enter the Path you want to sort:")
	fmt.Scanln(&pathinput)

	rootpath, err := New(pathinput)

	if err != nil {
		panic(err)
	}

	start := time.Now()
	rootpath.Sort()
	elapsed := time.Since(start)
	fmt.Printf("\nElapsed time: %s", elapsed)
}

// Returns folders and files non-recursively as a map
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

// Returns extensions of files in rootpath as a slice
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

// Creates the extension folders and returns a slice containing the Folderpaths
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

// generateNewPathName generates a new folder name for a file based on the extension path.
// It takes in two parameters: fileName which is the name of the file, and extensionPath which is the path of the extension folder.
// It iterates through the extension folder and generates a new folder name using the format "i__fileName" where i is an integer.
// It returns the new folder name as a string.
func generateNewPathName(fileName, extensionPath string) (newFolderName string) {
	newFolderName = filepath.Join(extensionPath, fileName)

	if !pathDoesExist(newFolderName) {
		return newFolderName
	}

	for i := 1; ; i++ {
		newFolderName = filepath.Join(extensionPath, fmt.Sprintf("%d__%s", i, fileName))
		if !pathDoesExist(newFolderName) {
			return newFolderName
		}
	}
}

// pathDoesExist checks if a path already exists.
// It takes in a string parameter path which is the path to check.
// It returns a boolean value doesExist which is true if the path exists and false if it does not.
func pathDoesExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

// moveFilesToExtensionFolder moves files to their respective extension folders.
// It takes in two parameters: files which is a map of file names and their paths, and extensionPaths which is a slice of extension folder paths.
// It iterates through each file in the files map and checks if the file extension matches any of the extensions in the extensionPaths slice.
// If a match is found, it generates a new folder name using the generateNewFolderName function and moves the file to the new folder.
// It returns nothing.
func moveFilesToExtensionFolder(files map[string]string, extensionPaths []string) {
	for kFile, vFile := range files {
		for _, extension := range extensionPaths {
			if filepath.Ext(extension) == filepath.Ext(vFile) {
				oldFile := vFile
				newFile := generateNewPathName(kFile, extension)
				err := os.Rename(oldFile, newFile)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
