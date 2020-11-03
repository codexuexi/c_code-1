package c_code

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

////mongodb 结果集数据量
//func MongodbResult(e error) error {
//	if e != nil {
//		if e.Error() == "mongo: no documents in result" {
//			return nil
//		}
//	}
//	return e
//}

//con := fmt.Sprintf("mongodb://%s:%s", _config.DbInfo.IP, _config.DbInfo.Port)
//connect_options := options.Client().ApplyURI(con)
//if _config.Auth {
//connect_options = options.Client().ApplyURI(con).SetAuth(options.Credential{
//AuthSource: "admin",
//Username:   "codexuexiAdminRoot",
//Password:   "s2CKFngFPq4NuA6D",
//})
//}

func MongodbConnect(connect_options *options.ClientOptions) (db *mongo.Client, err error) {
	db, err = mongo.NewClient(connect_options)
	if err != nil {
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(ctx)
	if err != nil {
		return
	}
	err = db.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	return
}
