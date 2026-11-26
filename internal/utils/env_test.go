package utils

import (
	"os"
	"testing"
)

func TestGetEnvStr(t *testing.T) {
	mockValue := "mock_value"
	os.Setenv("mock_key", mockValue)
	value, _ := GetEnvStr(GetEnvArgs{Key: "mock_key"})

	if value != mockValue {
		t.Errorf("GetEnvStr() FAILED, expected %v but got value %v", mockValue, value)
	}

	os.Unsetenv("mock_key")
}

func TestGetEnvStrError(t *testing.T) {
	_, err := GetEnvStr(GetEnvArgs{Key: "mock_key"})

	if err == nil {
		t.Errorf(
			"GetEnvStr() FAILED, expected %v but got value %v",
			ErrEnvVarEmpty("mock_key"),
			err,
		)
	}
}
