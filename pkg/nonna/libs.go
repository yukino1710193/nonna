package nonna

func pushRequest2Packet(pushRequest *PushRequest) *Packet {
	retPacket := &Packet{
		ID:       pushRequest.ID,
		SourceIP: pushRequest.SourceIP,
		Domain:   pushRequest.Domain,
		URI:      pushRequest.URI,
		Method:   pushRequest.Method,
		Headers:  make([]*PushRequest_HeaderSchema, 0),
	}

	for _, header := range pushRequest.Headers {
		retPacket.Headers = append(retPacket.Headers, &PushRequest_HeaderSchema{
			Field: header.Field,
			Value: header.Value,
		})
	}

	return retPacket
}

func packet2PopResponse(packet *Packet) *PopResponse {
	retPopResponse := &PopResponse{
		ID:       packet.ID,
		SourceIP: packet.SourceIP,
		Domain:   packet.Domain,
		URI:      packet.URI,
		Method:   packet.Method,
		Headers:  make([]*PopResponse_HeaderSchema, 0),
	}

	for _, header := range packet.Headers {
		retPopResponse.Headers = append(retPopResponse.Headers, &PopResponse_HeaderSchema{
			Field: header.Field,
			Value: header.Value,
		})
	}

	return retPopResponse
}

func headerModRequest2Packet(headerModRequest *HeaderModRequest) *Packet {
	retPacket := &Packet{
		ID:       headerModRequest.ID,
		SourceIP: headerModRequest.SourceIP,
		Domain:   headerModRequest.Domain,
		URI:      headerModRequest.URI,
		Method:   headerModRequest.Method,
		Headers:  make([]*PushRequest_HeaderSchema, 0),
	}

	for _, header := range headerModRequest.Headers {
		retPacket.Headers = append(retPacket.Headers, &PushRequest_HeaderSchema{
			Field: header.Field,
			Value: header.Value,
		})
	}

	return retPacket
}

func packet2HeaderModResponse(packet *Packet) *HeaderModResponse {
	retHeaderModResponse := &HeaderModResponse{
		ID:       packet.ID,
		SourceIP: packet.SourceIP,
		Domain:   packet.Domain,
		URI:      packet.URI,
		Method:   packet.Method,
		Headers:  make([]*HeaderModResponse_HeaderSchema, 0),
	}

	for _, header := range packet.Headers {
		retHeaderModResponse.Headers = append(retHeaderModResponse.Headers, &HeaderModResponse_HeaderSchema{
			Field: header.Field,
			Value: header.Value,
		})
	}

	return retHeaderModResponse
}
