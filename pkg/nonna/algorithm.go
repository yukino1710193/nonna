package nonna

import "github.com/bonavadeur/nonna/pkg/bonalib"

func (q *ExtraQueue) SortAlgorithm(p *Packet) {
	bonalib.Info("SortAlgorithm", p)
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "NonnaField-1",
		Value: "NonnaValue-1",
	})

	q.Queue = append([]*Packet{p}, q.Queue...)
}
