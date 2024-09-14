package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	Name          string
	Type          string
	Size          int64
	CreationTime  time.Time
	FileExtension string
}

func List(directory string) ([]FileInfo, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error: Unable to read directory %s\n", directory)
		return nil, err
	}

	var fileList []FileInfo
	for _, file := range files {
		info, _ := file.Info()
		fileType := "FILE"
		if file.IsDir() {
			fileType = "DIR"
		}

		fileExt := getFileExtension(file.Name())
		creationTime := info.ModTime()
		fileList = append(fileList, FileInfo{
			Name:          file.Name(),
			Type:          fileType,
			Size:          info.Size(),
			CreationTime:  creationTime,
			FileExtension: fileExt,
		})
	}
	return fileList, nil
}

func getFileExtension(filename string) string {
	if dot := strings.LastIndex(filename, "."); dot != -1 {
		return filename[dot:]
	}
	return ""
}
