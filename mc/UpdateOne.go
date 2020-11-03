package mc

import (
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (q *MQ) UpdateOne(data interface{}) (err error) {
	_, err = q.table.UpdateOne(TimeOut(), q.where, bson.M{"$set": data})
	return
}
func (q *MQ) UpdateAll(data interface{}) (err error) {
	_, err = q.table.UpdateMany(TimeOut(), q.where, bson.M{"$set": data})
	return
}
func (q *MQ) UpdateOneIsEmptyNewInsert(data interface{}) (err error) {
	_, err = q.table.UpdateOne(TimeOut(), q.where, bson.M{"$set": data}, options.Update().SetUpsert(true))
	return
}

//字段自增
func (q *MQ) FieldAddOrDel(fileld_name string, add_or_del int) (err error) {
	_, err = q.table.UpdateOne(TimeOut(), q.where, bson.M{"$inc": bson.M{fileld_name: add_or_del}}, options.Update().SetUpsert(true))
	return
}
