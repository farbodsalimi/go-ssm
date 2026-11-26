package utils

import (
	"fmt"
	"os"
)

// GetEnvArgs struct
type GetEnvArgs struct {
	Key          string
	DefaultValue string
}

// ErrEnvVarEmpty returns custom error for empty or undefined environment variable
func ErrEnvVarEmpty(key string) error {
	return fmt.Errorf("getenv: %s environment variable empty", key)
}

// GetEnvStr returns environment variable in string
func GetEnvStr(params GetEnvArgs) (string, error) {
	v, exists := os.LookupEnv(params.Key)
	if exists {
		return v, nil
	} else if !exists && params.DefaultValue != "" {
		return params.DefaultValue, nil
	}
	return v, ErrEnvVarEmpty(params.Key)
}
