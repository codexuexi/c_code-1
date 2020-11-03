package mc

import (
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MQ struct {
	table      *mongo.Collection
	where      bson.M
	order      bson.M
	projection bson.M
	limit      int64
	skip       int64
	page       int64
}

func Table(table_name string) *MQ {
	q := &MQ{}
	q.table = Coll(table_name)
	return q
}

func (q *MQ) Projection(bbj bson.M) *MQ {
	q.projection = bbj
	return q
}

func (q *MQ) Where(bbj bson.M) *MQ {
	q.where = bbj
	return q
}
func (q *MQ) Order(bbj bson.M) *MQ {
	q.order = bbj
	return q
}
func (q *MQ) Limit(bbj int64) *MQ {
	q.limit = bbj
	return q
}
func (q *MQ) Skip(bbj int64) *MQ {
	q.skip = bbj
	return q
}
func (q *MQ) Page(bbj int64) *MQ {
	q.page = bbj
	return q
}
