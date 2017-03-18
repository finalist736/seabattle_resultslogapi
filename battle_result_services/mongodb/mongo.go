package mongodb

import (
	"fmt"

	"github.com/finalist736/seabattle_resultslogapi"
	"github.com/finalist736/seabattle_resultslogapi/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
}

func NewService() seabattle_resultslogapi.BattleResultService {
	return new(Service)
}

func (*Service) BattleID(id int64) *seabattle_resultslogapi.BattleResult {
	conf := config.GetConfiguration()
	connectString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		conf.Mongo.User,
		conf.Mongo.Pass,
		conf.Mongo.Host,
		conf.Mongo.Port,
		conf.Mongo.Name)

	session, err := mgo.Dial(connectString)
	if err != nil {
		fmt.Printf("mongo connection error: %s\n", err.Error())
		return nil
	}
	defer session.Close()

	c := session.DB(conf.Mongo.Name).C("games")

	var result *seabattle_resultslogapi.BattleResult
	result = new(seabattle_resultslogapi.BattleResult)

	err = c.Find(bson.M{"battle": id}).One(result)
	if err != nil {
		fmt.Printf("mongo find error: %s\n", err.Error())
		return nil
	}
	return result
}
