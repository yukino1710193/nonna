package nonna

type Packet struct {
	ID       uint32
	SourceIP string
	Domain   string
	URI      string
	Method   string
	Headers  []*PushRequest_HeaderSchema
}

func (p *Packet) GetHeader(key string) (string, bool) {
	for _, h := range p.Headers {
		if h.Field == key {
			return h.Value, true
		}
	}
	return "", false
}