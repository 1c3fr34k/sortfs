package sortfs

import (
	"testing"
)

// TODO: Need more tests

func TestGetFolderAndFilePathsLenFiles(t *testing.T) {
	files, _ := getFolderAndFilePaths("C:\\Users\\manue\\Desktop\\DEV\\Go\\sortFS\\Test\\Unittest")
	if len(files) != 1 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}

func TestGetFolderAndFilePathsLenFolders(t *testing.T) {
	_, folders := getFolderAndFilePaths("C:\\Users\\manue\\Desktop\\DEV\\Go\\sortFS\\testfolder")
	if len(folders) != 2 {
		t.Errorf("Expected 1 folder, got %d", len(folders))
	}
}

func BenchmarkGetFolderAndFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFolderAndFilePaths("C:\\Users\\manue\\Desktop\\DEV\\Go\\sortFS\\testfolder")
	}
}
