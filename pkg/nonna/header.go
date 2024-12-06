package nonna

func (q *ExtraQueue) HeaderModifierAlgorithm(p *Packet) {
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "NonnaField-2",
		Value: "NonnaValue-2",
	})
}
