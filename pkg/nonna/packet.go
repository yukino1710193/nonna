package nonna

type Packet struct {
	ID       uint32
	SourceIP string
	Domain   string
	URI      string
	Method   string
	Headers  []*PushRequest_HeaderSchema
}
