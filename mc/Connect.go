package mc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var mgo_now *mongo.Client

var db_name = ""

func SetDB(client *mongo.Client) {
	mgo_now = client
}
func SetDBName(string2 string) {
	db_name = string2
}
func Coll(table_name string) *mongo.Collection {
	return mgo_now.Database(db_name).Collection(table_name)
}

func Client() *mongo.Database {
	return mgo_now.Database(db_name)
}

//func Command()  {
//	return mgo_now.Database(db_name).RunCommand
//}

func TimeOut() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

//
//
//import (
//"context"
//"fmt"
//"go.mongodb.org/mongo-driver/mongo"
//"go.mongodb.org/mongo-driver/mongo/options"
//"go.mongodb.org/mongo-driver/mongo/readpref"
//"log"
//"time"
//)
//
//var mgo_now *mongo.Client
//var _config = config.GetConfig()
//
//func Connect(d_info config.ConfigDB) {
//
//	con := fmt.Sprintf("mongodb://%s:%s", d_info.IP, d_info.Port)
//	connect_options := options.Client().ApplyURI(con)
//	if d_info.Username != "" && d_info.Password != "" {
//		connect_options = options.Client().ApplyURI(con).SetAuth(options.Credential{
//			AuthSource: d_info.AuthSource,
//			Username:   d_info.Username,
//			Password:   d_info.Password,
//		})
//	}
//
//	old_client, err := mongo.NewClient(connect_options)
//	if err != nil {
//		log.Fatal(1, err)
//	}
//	mgo_now = old_client
//	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//	err = mgo_now.Connect(ctx)
//	if err != nil {
//		log.Fatal(2, " Now 数据库", err)
//	}
//	err = mgo_now.Ping(ctx, readpref.Primary())
//	if err != nil {
//		log.Fatal(3, " Now", err)
//	}
//}
//
//func Coll(table_name string) *mongo.Collection {
//	return mgo_now.Database("ask").Collection(table_name)
//}
//
//func TimeOut() context.Context {
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	return ctx
//}
//
