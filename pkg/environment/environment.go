package environment

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type EnvironmentSettings struct {
	ConnectionString string
}

func GetEnvironmentVariables(logger log.Logger) (EnvironmentSettings, bool) {

	env := EnvironmentSettings{}

	connectionString, ok := os.LookupEnv("DUNGEON_FETCHER_CONNECTION_STRING")
	if !ok {
		level.Error(logger).Log("missing_environment_variable", "DUNGEON_FETCHER_CONNECTION_STRING")
		return env, ok
	}

	env.ConnectionString = connectionString

	return env, true
}
