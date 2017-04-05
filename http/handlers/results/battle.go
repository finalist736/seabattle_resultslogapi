package results

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/finalist736/seabattle_resultslogapi/battle_result_services/mongodb"
	"github.com/finalist736/seabattle_resultslogapi/tools/logger"

	"github.com/finalist736/seabattle_resultslogapi"
	"github.com/gocraft/web"
)

func ResultsBattleID(ctx *seabattle_resultslogapi.HttpContext, rw web.ResponseWriter, req *web.Request) {
	logger.StdOut().Debugln("ResultsBattleID called")
	battleIdString := req.PathParams["battle"]
	logger.StdOut().Debugf("battle id requested: %v", battleIdString)
	battleIdInt64, err := strconv.ParseInt(battleIdString, 10, 64)
	if err != nil {
		logger.StdErr().Infof("battle id parse error: %v", err)
		rw.Write([]byte("error"))
		return
	}
	logger.StdOut().Debugf("battle id after parsing: %v", battleIdInt64)
	if battleIdInt64 == 0 {
		logger.StdErr().Infoln("battle id zero")
		rw.Write([]byte("error"))
		return
	}
	mservice := mongodb.NewService()
	mbresult := mservice.BattleID(battleIdInt64)
	if mbresult == nil {
		logger.StdErr().Warnf("ResultsBattleID id not found: %v", battleIdInt64)
		rw.Write([]byte("not found"))
		return
	}

	baresult, err := json.Marshal(mbresult)
	if err != nil {
		logger.StdErr().Warnf("ResultsBattleID json marshal error: %v\ndata: %+v", err, mbresult)
		rw.Write([]byte("internal json error"))
		return
	}
	rw.Header().Set("Content-type", "application/json")
	fmt.Fprintf(rw, "%s", baresult)
}
