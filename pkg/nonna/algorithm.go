package nonna

import (
	"time"

	"github.com/bonavadeur/nonna/pkg/bonalib"
)

func (q *ExtraQueue) SortAlgorithm(p *Packet) {
	bonalib.Info("SortAlgorithm", p)
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "Queue-N-Field-1",
		Value: "tim",
	})
	time.Sleep(5 * time.Second)
	q.Queue = append([]*Packet{p}, q.Queue...)
}
