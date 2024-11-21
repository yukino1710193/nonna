package nonna

func (q *ExtraQueue) SortAlgorithm(p *Packet) {
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "naniField",
		Value: "naniValue",
	})

	q.Queue = append([]*Packet{p}, q.Queue...)
}
