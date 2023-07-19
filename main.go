package main

import (
	"fmt"
	"time"

	"github.com/1c3fr34k/sortFS/sortfs"
)

func main() {
	start := time.Now()

	rootpath := sortfs.New("C:\\Users\\manue\\Downloads\\Test\\Test")
	rootpath.Sort()

	elapsed := time.Since(start)

	fmt.Printf("\nElapsed time: %s", elapsed)
}
