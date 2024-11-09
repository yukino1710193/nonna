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
	queueLock       sync.Mutex
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
		queueLock:       sync.Mutex{},
	}
	newExtraQueue.pushBridge = hashi.NewHashi(
		"PushBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/push-bridge",
		bonalib.Cm2Int("nonna-threads"),
		reflect.TypeOf(PushRequest{}),
		reflect.TypeOf(PushResponse{}),
		newExtraQueue.PushResponseAdapter,
	)
	newExtraQueue.popBridge = hashi.NewHashi(
		"PopBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/pop-bridge",
		bonalib.Cm2Int("nonna-threads"),
		reflect.TypeOf(PopRequest{}),
		reflect.TypeOf(PopResponse{}),
		newExtraQueue.PopResponseAdapter,
	)

	return newExtraQueue
}

func (q *ExtraQueue) PushResponseAdapter(params ...interface{}) (interface{}, error) {
	pushRequest := params[0].(*PushRequest)
	q.Push(pushRequest)
	return &PushResponse{SymbolizeResponse: Status_Success}, nil
}

func (q *ExtraQueue) PopResponseAdapter(params ...interface{}) (interface{}, error) {
	// _ = params[0].(*PopRequest)
	popPacket := q.Pop()
	return packet2PopResponse(popPacket), nil
}

func (q *ExtraQueue) Push(pushPacket *PushRequest) {
	q.queueLock.Lock()
	q.sort(pushRequest2Packet(pushPacket))
	q.queueLock.Unlock()
	q.Next <- true
}

func (q *ExtraQueue) Pop() *Packet {
	q.popLock.Lock()
	defer q.popLock.Unlock()

	<-q.Next // hangout until len of Queue > 0
	q.queueLock.Lock()
	popPacket := q.Queue[len(q.Queue)-1]
	q.Queue = q.Queue[:len(q.Queue)-1]
	q.queueLock.Unlock()

	return popPacket
}

// custom this
func (q *ExtraQueue) sort(p *Packet) {
	q.sortLock.Lock()
	defer q.sortLock.Unlock()

	// example of adding header
	p.Headers = append(p.Headers, &PushRequest_HeaderSchema{
		Field: "naniField",
		Value: "naniValue",
	})

	// for i, header := range p.Headers {
	// 	bonalib.Info("Header", i, header.Field, header.Value)
	// }

	// position := -1
	// position = len(q.Queue) - 1
	// q.Queue = append(q.Queue[:position+1], append([]*Packet{p}, q.Queue[position+1:]...)...)
	q.Queue = append([]*Packet{p}, q.Queue...)
}
