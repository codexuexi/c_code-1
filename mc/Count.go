package mc

func (q *MQ) Count() (count int, err error) {

	count64, err := q.table.CountDocuments(TimeOut(), q.where)
	if err != nil {
		return
	}
	count = int(count64)
	return
}
