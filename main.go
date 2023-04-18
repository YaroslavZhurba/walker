package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	dir    = "./testdir"
	output = "output.txt"
)

func getFileInfos(dirName string) ([]os.FileInfo, error) {
	f, err := os.Open(dirName)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func walkAndWrite(f *os.File) error {
	w := bufio.NewWriter(f)
	fileInfos, err := getFileInfos(dir)
	if err != nil {
		return fmt.Errorf("error reading %s error: %w", dir, err)
	}

	stackDir := NewStackDir()
	stackDir.Push(dir)

	for !stackDir.IsEmpty() {
		dirName := stackDir.Pop()
		fileInfos, err = getFileInfos(dirName)
		if err != nil {
			return fmt.Errorf("error reading %s error: %w", dirName, err)
		}
		for _, fileInfo := range fileInfos {
			if !fileInfo.IsDir() {
				//fmt.Println(fileInfo.Name())
				_, err = w.WriteString(fileInfo.Name() + "\n")
				if err != nil {
					return fmt.Errorf("error writing %s error: %w", output, err)
				}
				continue
			}
			stackDir.Push(dirName + "/" + fileInfo.Name())
		}
	}

	return w.Flush()
}

func main() {
	f, err := os.Create(output)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = walkAndWrite(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finished")
}
