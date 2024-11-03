package nonna

import (
	"os"
	"reflect"
	"sync"

	"github.com/bonavadeur/nonna/pkg/bonalib"
	"github.com/bonavadeur/nonna/pkg/hashi"
)

type ExtraQueue struct {
	Queue           []*Packet
	Next            chan bool
	NextQueueSize   int
	NextQueueLength int
	popLock         sync.Mutex
	pushBridge      *hashi.Hashi
	popBridge       *hashi.Hashi
	MsgIDLock       sync.Mutex
	MsgIDCount      uint32
	sortLock        sync.Mutex
}

func NewExtraQueue() *ExtraQueue {
	containerConcurrency := bonalib.Str2Int(os.Getenv("CONTAINER_CONCURRENCY"))

	newExtraQueue := &ExtraQueue{
		Queue:           make([]*Packet, 0),
		Next:            make(chan bool, containerConcurrency),
		NextQueueSize:   containerConcurrency,
		NextQueueLength: 0,
		popLock:         sync.Mutex{},
		MsgIDLock:       sync.Mutex{},
		MsgIDCount:      0,
		sortLock:        sync.Mutex{},
	}
	newExtraQueue.pushBridge = hashi.NewHashi(
		"PushBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/push-bridge",
		10,
		reflect.TypeOf(PushRequest{}),
		reflect.TypeOf(PushResponse{}),
		newExtraQueue.PushResponseAdapter,
	)
	newExtraQueue.popBridge = hashi.NewHashi(
		"PopBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/pop-bridge",
		10,
		reflect.TypeOf(PopRequest{}),
		reflect.TypeOf(PopResponse{}),
		newExtraQueue.PopResponseAdapter,
	)

	// go newExtraQueue.controller()

	return newExtraQueue
}

// func (q *ExtraQueue) controller() {
// 	for {
// 		<-q.Next
// 		if len(q.Queue) > 0 {
// 			q.Pop()
// 		}
// 	}
// }

func (q *ExtraQueue) PushResponseAdapter(params ...interface{}) (interface{}, error) {
	pushRequest := params[0].(*PushRequest)
	q.Push(pushRequest)
	return &PushResponse{SymbolizeResponse: Status_Success}, nil
}

func (q *ExtraQueue) PopResponseAdapter(params ...interface{}) (interface{}, error) {
	bonalib.Info("PopResponseAdapter")
	_ = params[0].(*PopRequest)
	bonalib.Info("PopResponseAdapter 2")
	popPacket := q.Pop()
	return packet2PopResponse(popPacket), nil
}

func (q *ExtraQueue) Push(pushPacket *PushRequest) {
	q.sort(pushRequest2Packet(pushPacket))
}

func (q *ExtraQueue) Pop() *Packet {
	bonalib.Info("Pop")
	q.popLock.Lock()
	defer q.popLock.Unlock()

	popPacket := q.Queue[len(q.Queue)-1]
	q.Queue = q.Queue[:len(q.Queue)-1]
	return popPacket
}

// custom this
func (q *ExtraQueue) sort(p *Packet) {
	q.sortLock.Lock()
	defer q.sortLock.Unlock()

	position := -1
	// for i, packet := range q.Queue {
	// 	if p.Priority > packet.Priority {
	// 		position = i
	// 		continue
	// 	}
	// 	if p.Priority == packet.Priority {
	// 		position = i - 1
	// 		continue
	// 	}
	// 	if packet.Priority > p.Priority {
	// 		break
	// 	}
	// }
	position = len(q.Queue) - 1
	q.Queue = append(q.Queue[:position+1], append([]*Packet{p}, q.Queue[position+1:]...)...)
}
