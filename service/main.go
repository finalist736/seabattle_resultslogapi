package main

import (
	"flag"

	"github.com/finalist736/seabattle_resultslogapi/tools/logger"

	"github.com/finalist736/seabattle_resultslogapi/config"
	"github.com/finalist736/seabattle_resultslogapi/http"
	"github.com/finalist736/seabattle_resultslogapi/tools/profiling"
)

var config_path *string = flag.String("config", "../config.json", "config file path")

func main() {
	flag.Parse()
	config.SetConfigFile(*config_path)
	logger.ReloadLogs()

	logger.StdOut().Infof("config file: %s", *config_path)

	logger.StdOut().Infof("profile: %+v", config.GetConfiguration())
	if config.GetConfiguration().ProfilingCPU {
		profiling.ProfileCPU()
		defer profiling.CloseCPU()
	}

	http.StartServer()
}
