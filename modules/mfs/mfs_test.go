package mfs

import (
	"log"
	"testing"

	"../config"
)

const userIDTest int = 1

var newContentTest ContentInformation = ContentInformation{
	ContentName:     "teste.txt",
	ContentData:     "This is a test!",
	ContentFullRoot: "",
	Type:            File,
}

func TestInit(test *testing.T) {
	if !checkDir(rootMsfDir) {
		test.Errorf("Cannot create main root folder %s%s", config.GetRootLocation(), rootMsfDir)
	}
}

func TestUserLoadFolder(test *testing.T) {
	if PrepareUserCloud(userIDTest) {
		var newContent ContentInformation = newContentTest

		if CreateContentOnUserCloud(userIDTest, newContent) {
			if getSuccess, userContent := GetContentOnUserCloud(userIDTest, newContentTest); !getSuccess {
				test.Errorf("User with ID %d cannot be loaded on MSF Module!", userIDTest)
			} else {
				log.Println("User content!")
				log.Println(userContent)
			}
		} else {
			test.Errorf("User with ID %d cannot be loaded on MSF Module!", userIDTest)
		}
	} else {
		test.Errorf("User with ID %d cannot be loaded on MSF Module!", userIDTest)
	}
}

func TestRemoveFile(test *testing.T) {
	if !RemoveContentOnUserCloud(userIDTest, newContentTest) {
		test.Errorf("User with ID %d cannot delete this file %s on directory %s", userIDTest, newContentTest.ContentName, newContentTest.ContentFullRoot)
	}
}
