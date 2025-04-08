package nonna

import (
	"time"

	"github.com/bonavadeur/nonna/pkg/bonalib"
)

func (q *ExtraQueue) SortAlgorithm(p *Packet) {
	bonalib.Info("SortAlgorithm", p)
	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "Incoming-N-Moment",
		Value: time.Now().Format("15:04:05.000000"),
	})

	// Copy lại các header cũ với tên mới
	copyFields := map[string]string{
		"Incoming-J-Moment-Responsed":  "Incoming-J-Moment",
		"Outcoming-J-Moment-Responsed": "Outcoming-J-Moment",
		"Queue-J-Length-Responsed":     "Queue-J-Length",
		"Lb-Momment-Responsed":         "Lb-Momment",
	}

	for newField, oldField := range copyFields {
		if val, ok := p.GetHeader(oldField); ok {
			bonalib.Log("SortAlgorithm", "copy header ", oldField, "to", newField, val, ok)

			p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
				Field: newField,
				Value: val,
			})
		}
	}
	time.Sleep(3 * time.Second)
	q.Queue = append([]*Packet{p}, q.Queue...)
	// var val string
	// val, _ = p.GetHeader("Incoming-J-Moment-Responsed")
	// bonalib.Succ("SortAlgorithm", "Incoming-J-Moment-Responsed", val)

	// val, _ = p.GetHeader("Incoming-J-Moment")
	// bonalib.Log("SortAlgorithm", "Incoming-J-Moment", val)

	// val, _ = p.GetHeader("Outcoming-J-Moment-Responsed")
	// bonalib.Succ("SortAlgorithm", "Outcoming-J-Moment-Responsed", val)

	// val, _ = p.GetHeader("Outcoming-J-Moment")
	// bonalib.Log("SortAlgorithm", "Outcoming-J-Moment", val)

}
