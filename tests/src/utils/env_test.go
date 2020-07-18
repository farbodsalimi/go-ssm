package utils

import (
	"go-ssm/pkg/utils"
	"os"
	"testing"
)

func TestGetEnvStr(t *testing.T) {
	mockValue := "mock_value"
	os.Setenv("mock_key", mockValue)
	value, _ := utils.GetEnvStr(utils.GetEnvArgs{Key: "mock_key"})

	if value != mockValue {
		t.Errorf("GetEnvStr() FAILED, expected %v but got value %v", mockValue, value)
	}

	os.Unsetenv("mock_key")
}

func TestGetEnvStrError(t *testing.T) {
	_, err := utils.GetEnvStr(utils.GetEnvArgs{Key: "mock_key"})

	if err == nil {
		t.Errorf("GetEnvStr() FAILED, expected %v but got value %v", utils.ErrEnvVarEmpty("mock_key"), err)
	}
}
