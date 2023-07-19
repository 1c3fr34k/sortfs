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

// Init a new sortFs struct
func New(root string) *sortFs {
	return &sortFs{rootPath: root}
}

// Entry point
func (s *sortFs) Sort() {
	s.files, s.folders = getFolderAndFilePaths(s.rootPath)
	s.extensionNames = getDestinctFileExtensions(s.files)
	s.extensionPaths = createFoldersForExtensions(s.rootPath, s.extensionNames)
	moveFilesToExtensionFolder(s.files, s.extensionPaths)

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

// generateNewFolderName generates a new folder name by joining the extension path and the file name.
// It takes in two string parameters: fileName and extensionPath.
// It returns a string value newFolderName.
func generateNewFolderName(fileName, extensionPath string) (newFolderName string) {
	newFolderName = filepath.Join(extensionPath, fileName)
	return newFolderName
}

// checkForExistingFile checks if a file already exists in the given path.
// It takes in a string parameter newPath which is the path to check.
// It returns a boolean value doesExist which is true if the file exists and false otherwise.
func checkForExistingFile(newPath string) (doesExist bool) {
	_, err := os.Stat(newPath)
	if err == nil {
		return true
	} else {
		return false
	}
}

// moveFilesToExtensionFolder moves files to their respective extension folders.
// It takes in two parameters: files which is a map of file names and their paths, and extensionPaths which is a slice of extension folder paths.
// It iterates through each file in the files map and checks if the file extension matches any of the extensions in the extensionPaths slice.
// If a match is found, it generates a new folder name using the generateNewFolderName function and checks if the file already exists in the new path using the checkForExistingFile function.
// If the file does not exist, it renames the file to the new path using the os.Rename function.
func moveFilesToExtensionFolder(files map[string]string, extensionPaths []string) {
	for kFile, vFile := range files {
		for _, extension := range extensionPaths {
			if filepath.Ext(extension) == filepath.Ext(vFile) {
				oldFile := vFile
				newFile := generateNewFolderName(kFile, extension)
				if !checkForExistingFile(newFile) {
					err := os.Rename(oldFile, newFile)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}
