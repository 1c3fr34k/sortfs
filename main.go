package main

import (
	"fmt"
	"time"

	"github.com/1c3fr34k/sortFS/sortfs"
)

func main() {
	start := time.Now()
	// folders, files := sortfs.GetFolderAndFilePaths("C:\\Users\\manue\\Downloads")
	rootpath := sortfs.New("C:\\Users\\manue\\Desktop\\DEV\\Go\\sortFS\\Test\\Test")
	rootpath.Sort()

	elapsed := time.Since(start)

	fmt.Printf("\nElapsed time: %s", elapsed)
}
