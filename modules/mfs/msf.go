package mfs

import (
	"fmt"
	"log"
	"os"

	"../config"
)

type ContentType int8

type ContentInformation struct {
	ContentName,
	ContentFullRoot,
	ContentData string
	ContentSize int32
	Type        ContentType
}

const rootMsfDir string = "fcloud"
const (
	Directory ContentType = iota
	File      ContentType = iota
)

func init() {
	log.Println("Load MSF module...")
	if !checkDir(rootMsfDir) {
		log.Fatalf("Cannot create main root folder %s%s", config.GetRootLocation(), rootMsfDir)
	}
}

func PrepareUserCloud(userID int) bool {
	return checkDir(fmt.Sprintf("%s/%d", rootMsfDir, userID))
}

func CreateContentOnUserCloud(userID int, content ContentInformation) bool {
	fullDir := fmt.Sprintf("%s/%s/%s", getUserFullDir(userID, false), content.ContentFullRoot, content.ContentName)
	var result bool = false

	if content.Type == Directory {
		result = checkDir(fullDir)
	} else {
		result = checkFile(fullDir, content.ContentData)
	}

	return result
}

func checkFile(filename, content string) bool {
	var result bool = false

	if newFile, newFileErr := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755); newFileErr == nil {
		result = true
		newFile.WriteString(content)
		newFile.Close()
	}

	return result
}

func checkDir(dir string) bool {
	var result bool = false
	fullDir := fmt.Sprintf("%s%s", config.GetRootLocation(), dir)
	if _, err := os.Stat(fullDir); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(fullDir, 0755); err == nil {
				result = true
			}
		}
	} else {
		result = true
	}

	return result
}

func getUserFullDir(userID int, includeRoot bool) string {
	fullDir := fmt.Sprintf("%s%s/%d", config.GetRootLocation(), rootMsfDir, userID)

	if !includeRoot {
		fullDir = fmt.Sprintf("%s/%d", rootMsfDir, userID)
	}

	return fullDir
}
