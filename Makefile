gen-messages:
	protoc -I=. --go_out=. pkg/nonna/messages.proto
