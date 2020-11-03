package mc

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createOpsFind(q *MQ) []*options.FindOptions {
	ops := []*options.FindOptions{}
	if len(q.order) > 0 {
		ops = append(ops, options.Find().SetSort(q.order))
	}
	if len(q.projection) > 0 {
		ops = append(ops, options.Find().SetProjection(q.projection))
	}
	if q.limit > 0 {
		ops = append(ops, options.Find().SetLimit(q.limit))
	}

	if q.page > 0 {
		q.skip = (q.page - 1) * q.limit
	}

	if q.skip > 0 {
		ops = append(ops, options.Find().SetSkip(q.skip))
	}
	return ops
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func (q *MQ) Find(data interface{}) (err error) {

	var d2 []bson.M

	var find *mongo.Cursor

	ops := createOpsFind(q)
	if len(ops) > 0 {
		find, err = q.table.Find(TimeOut(), q.where, ops...)
	} else {
		find, err = q.table.Find(TimeOut(), q.where)
	}
	if err != nil {
		return
	}
	err = find.All(TimeOut(), &d2)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(d2)

	s_html := string(marshal)
	if err != nil {
		fmt.Println(s_html)
		return
	}

	err = json.Unmarshal(marshal, &data)
	return
}
