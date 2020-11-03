package mc

func (q *MQ) Insert(data interface{}) (err error) {
	_, err = q.table.InsertOne(TimeOut(), data)
	return
}
