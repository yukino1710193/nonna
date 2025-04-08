package nonna

import (
	"time"

	"github.com/bonavadeur/nonna/pkg/bonalib"
	_ "github.com/bonavadeur/nonna/pkg/bonalib"
)

func (q *ExtraQueue) HeaderModifierAlgorithm(p *Packet) {

	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "Outcoming-N-Moment",
		Value: time.Now().Format("15:04:05.000000"),
	})
	bonalib.Log("HeaderModifierAlgorithm", "Outcoming-N-Moment", p.Headers[len(p.Headers)-1].Value)
}
