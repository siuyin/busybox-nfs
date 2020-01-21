package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/siuyin/dflt"
)

func main() {
	fmt.Println("hello-nfs")
	for {
		updateDatFile()
		time.Sleep(5 * time.Second)
	}
}
func updateDatFile() {
	datFile := dflt.EnvString("DAT_FILE", "/data/junk.txt")
	f, err := os.Create(datFile)
	if err != nil {
		log.Printf("could not open data file: %s: %v\n", datFile, err)
		return
	}
	defer f.Close()
	fmt.Fprintf(f, "%s\n", time.Now().Format("2006-01-02 15:04:05.000 MST"))
}
