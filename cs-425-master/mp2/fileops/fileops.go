package fileops

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"../config"
	"../logger"
)

// Return array of file names under the "sdfs/" dir
func ListFiles() []string {
	files, _ := ioutil.ReadDir(config.SDFS_DIR)
	filesList := make([]string, len(files))
	for idx, file := range files {
		filesList[idx] = file.Name()
	}
	return filesList
}

func DeleteFile(filename string) bool {
	newFilename := strings.ReplaceAll(filename, "/", "_")
	err := os.Remove(config.SDFS_DIR + newFilename)

	if err != nil {
		return false
	}

	logger.PrintInfo("Successfully deleted local file:", config.SDFS_DIR+newFilename)
	return true
}

// Fetches byte array and stores to local dir
func PutFile(filename string, data []byte, isSdfsDir, shouldOverwrite bool) bool {
	newFilename := strings.ReplaceAll(filename, "/", "_")

	file := ""

	if isSdfsDir {
		file = config.SDFS_DIR + newFilename
	} else {
		file = filename
	}

	if shouldOverwrite {
		err := ioutil.WriteFile(file, data, config.FILE_PERM)

		if err != nil {
			return false
		}
	} else {
		f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, config.FILE_PERM)
		if err != nil {
			return false
		}

		defer f.Close()

		if _, err = f.Write(data); err != nil {
			return false
		}
	}

	logger.PrintInfo("Successfully put local file:", file)
	return true
}

// Goroutine for reading a file n bytes at a time, using channels to block and
// transmit the read bytes over to the caller of the thread
func ReadFileThread(filename string, isSdfsDir bool, n int, c chan []byte) {
	newFilename := strings.ReplaceAll(filename, "/", "_")

	file := ""

	if isSdfsDir {
		file = config.SDFS_DIR + newFilename
	} else {
		file = filename
	}

	fd, err := os.Open(file)
	if err != nil {
		logger.PrintError("Error opening file:", err)
		c <- nil
		return
	}

	defer fd.Close()

	for {
		dataArr := make([]byte, n)
		numBytes, err := fd.Read(dataArr)
		dataArr = dataArr[:numBytes]

		if numBytes == 0 || len(dataArr) == 0 || err != nil {
			if err != io.EOF {
				logger.PrintError("Error reading file:", err)
			}

			c <- nil
			break
		}

		c <- dataArr
	}
}

// Initiate the SDS directory by erasing all existing files and creating the directory from scratch
func InitSdfs() {
	os.RemoveAll(config.SDFS_DIR)
	os.MkdirAll(config.SDFS_DIR, config.FILE_PERM)
}
