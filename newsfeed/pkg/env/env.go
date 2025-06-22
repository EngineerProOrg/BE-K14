package env

import "os"

type EnvType string

const (
	EnvTypeLocal EnvType = "local"
	EnvTypeProd  EnvType = "prod"

	envKey = "ENV"
)

// GetEnv return the current environment: EnvTypeLocal, EnvTypeProd
func GetEnv() EnvType {
	envType := os.Getenv(envKey)
	if envType == string(EnvTypeProd) {
		return EnvTypeProd
	}
	return EnvTypeLocal
}
