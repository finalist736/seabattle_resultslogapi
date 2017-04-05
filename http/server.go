package http

import (
	"net/http"

	"github.com/finalist736/seabattle_resultslogapi/tools/logger"

	"github.com/finalist736/seabattle_resultslogapi"
	"github.com/finalist736/seabattle_resultslogapi/config"
	"github.com/finalist736/seabattle_resultslogapi/http/handlers/results"
	"github.com/gocraft/web"
)

func StartServer() error {

	router := web.New(seabattle_resultslogapi.HttpContext{})
	router.Middleware((*seabattle_resultslogapi.HttpContext).AuthMiddleWare)
	router.Get("/result/:battle", results.ResultsBattleID)

	logger.StdOut().Infof("Listening on port: %s", config.GetConfiguration().Port)

	return http.ListenAndServe(config.GetConfiguration().Port, router)
}
