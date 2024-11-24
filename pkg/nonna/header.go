package nonna

func (q *ExtraQueue) HeaderModifier(p *Packet) {
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "naniField-Pop",
		Value: "naniValue-Pop",
	})
}
