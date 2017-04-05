package mongodb

import (
	"fmt"

	"github.com/finalist736/seabattle_resultslogapi/tools/logger"

	"github.com/finalist736/seabattle_resultslogapi"
	"github.com/finalist736/seabattle_resultslogapi/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type Service struct {
}

func NewService() seabattle_resultslogapi.BattleResultService {
	return new(Service)
}

func (*Service) BattleID(id int64) *seabattle_resultslogapi.BattleResult {
	var err error
	conf := config.GetConfiguration()
	if session == nil {
		connectString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
			conf.Mongo.User,
			conf.Mongo.Pass,
			conf.Mongo.Host,
			conf.Mongo.Port,
			conf.Mongo.Name)
		logger.StdOut().Debugf("mongo connection string: %v", connectString)
		session, err = mgo.Dial(connectString)
		if err != nil {
			logger.StdErr().Errorf("mongo connection error: %s\n", err.Error())
			return nil
		}
		//defer session.Close()
	} else {
		session.Refresh()
	}

	c := session.DB(conf.Mongo.Name).C("games")

	var result seabattle_resultslogapi.BattleResult

	err = c.Find(bson.M{"battle": id}).One(&result)
	if err != nil {
		logger.StdErr().Errorf("mongo find error: %s\n", err.Error())
		return nil
	}
	return &result
}
