package seabattle_resultslogapi

type BattleResult struct {
	BattleID  int64                 `bson:"battle"`
	StartTime int64                 `bson:"start"`
	EndTime   int64                 `bson:"end"`
	Winner    int64                 `bson:"winner"`
	Sides     [2]*BattleResultSides `bson:"sides"`
	Turns     []*BattleResultTurn   `bson:"turns"`
}

type BattleResultSides struct {
	ID   int64     `bson:"id"`
	Name string    `bson:"name"`
	Sea  *[100]int `bson:"sea"`
}

type BattleResultTurn struct {
	ID     int64  `bson:"id"`
	Shot   [2]int `bson:"shot"`
	Result int    `bson:"result"`
}

type BattleResultService interface {
	BattleID(int64) *BattleResult
}
