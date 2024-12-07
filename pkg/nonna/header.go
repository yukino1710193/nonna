package nonna

func (q *ExtraQueue) HeaderModifierAlgorithm(p *Packet) {
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "Queue-N-Field-2",
		Value: "biết khóc",
	})
}
