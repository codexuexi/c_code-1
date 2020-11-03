package mc

func (q *MQ) DelOne() {
	q.table.DeleteOne(TimeOut(), q.where)
}
func (q *MQ) DelAll() {
	q.table.DeleteMany(TimeOut(), q.where)
}
