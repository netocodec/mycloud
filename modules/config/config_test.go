package config

import "testing"

const defaultConfigPort int = 8080

func TestDefaultConfig(test *testing.T) {
	defaultConfig := loadConfiguration()

	if defaultConfig.Server.Port != defaultConfigPort {
		test.Errorf("Server port not right! (Current config port: %d | Expected Port: %d", defaultConfig.Server.Port, defaultConfigPort)
	}
}
