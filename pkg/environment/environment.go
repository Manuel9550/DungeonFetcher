package environment

import (
	"os"

	"github.com/Manuel9550/DungeonFetcher/pkg/dal"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type EnvironmentSettings struct {
	ConnectionString string
	DBType           string
}

func GetEnvironmentVariables(logger log.Logger) (EnvironmentSettings, bool) {

	env := EnvironmentSettings{}

	connectionString, ok := os.LookupEnv("DUNGEON_FETCHER_CONNECTION_STRING")
	if !ok {
		level.Error(logger).Log("missing_environment_variable", "DUNGEON_FETCHER_CONNECTION_STRING")
		return env, ok
	}

	DBTypeString, ok := os.LookupEnv("DUNGEON_FETCHER_DB_TYPE")
	if !ok {
		level.Error(logger).Log("missing_environment_variable", "DUNGEON_FETCHER_DB_TYPE")
		return env, ok
	}

	// Must make sure db is a supported type
	if _, ok := dal.GetDBTypeMap()[DBTypeString]; !ok {
		level.Error(logger).Log("improper_db_type", DBTypeString)
		return env, ok
	}

	env.ConnectionString = connectionString
	env.DBType = DBTypeString

	return env, true
}
