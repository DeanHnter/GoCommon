package commonos

import (
	"os"
	"strings"
)

type EnvVars map[string]string

func (e EnvVars) Exists(key string) (string, bool) {
	val, ok := e[key]
	return val, ok
}

func ReadEnvironmentVariables() EnvVars {
	envVars := make(EnvVars)
	for _, env := range os.Environ() {
		splitEnv := strings.SplitN(env, "=", 2)
		envVars[splitEnv[0]] = splitEnv[1]
	}
	return envVars
}
