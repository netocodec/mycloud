package mfs

import (
	"testing"

	"../config"
)

const userIDTest int = 1

func TestInit(test *testing.T) {
	if !checkDir(rootMsfDir) {
		test.Errorf("Cannot create main root folder %s%s", config.GetRootLocation(), rootMsfDir)
	}
}

func TestUserLoadFolder(test *testing.T) {
	if !PrepareUserCloud(userIDTest) {
		test.Errorf("User with ID %d cannot be loaded on MSF Module!", userIDTest)
	}
}
