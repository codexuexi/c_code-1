package mc

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createOpsOne(q *MQ) []*options.FindOneOptions {
	ops := []*options.FindOneOptions{}
	if len(q.order) > 0 {
		ops = append(ops, options.FindOne().SetSort(q.order))
	}
	return ops
}

func (q *MQ) FindOne(find_one interface{}) (err error) {
	d9999 := map[string]interface{}{}
	ops := createOpsOne(q)
	if len(ops) > 0 {
		err = q.table.FindOne(TimeOut(), q.where, ops...).Decode(&d9999)
	} else {
		err = q.table.FindOne(TimeOut(), q.where).Decode(&d9999)
	}
	temporaryBytes, err := json.Marshal(d9999)
	if err != nil {
		return
	}
	//fmt.Println(marshal, err)
	//temporaryBytes, _ := bson.MarshalJSON(d9999)
	err = json.Unmarshal(temporaryBytes, &find_one)
	return
}
