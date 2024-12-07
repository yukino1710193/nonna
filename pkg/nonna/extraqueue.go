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
	headerModBridge *hashi.Hashi
	MsgIDLock       sync.Mutex
	MsgIDCount      uint32
	sortLock        sync.Mutex
	queueLock       sync.Mutex
}

func NewExtraQueue() *ExtraQueue {
	containerConcurrency := bonalib.Str2Int(os.Getenv("CONTAINER_CONCURRENCY"))
	nonnaThreads := bonalib.Cm2Int("nonna-threads")

	newExtraQueue := &ExtraQueue{
		Queue:           make([]*Packet, 0),
		Next:            make(chan bool, nonnaThreads),
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
		nonnaThreads,
		reflect.TypeOf(PushRequest{}),
		reflect.TypeOf(PushResponse{}),
		newExtraQueue.PushResponseAdapter,
	)
	newExtraQueue.popBridge = hashi.NewHashi(
		"PopBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/pop-bridge",
		nonnaThreads,
		reflect.TypeOf(PopRequest{}),
		reflect.TypeOf(PopResponse{}),
		newExtraQueue.PopResponseAdapter,
	)
	newExtraQueue.headerModBridge = hashi.NewHashi(
		"HeaderModBridge",
		hashi.HASHI_TYPE_SERVER,
		BASE_PATH+"/header-mod-bridge",
		nonnaThreads,
		reflect.TypeOf(HeaderModRequest{}),
		reflect.TypeOf(HeaderModResponse{}),
		newExtraQueue.HeaderModResponseAdapter,
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

func (q *ExtraQueue) HeaderModResponseAdapter(params ...interface{}) (interface{}, error) {
	headerModRequest := params[0].(*HeaderModRequest)
	modifiedPacket := q.HeaderModifier(headerModRequest2Packet(headerModRequest))
	return packet2HeaderModResponse(modifiedPacket), nil
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
	q.HeaderModifier(popPacket)
	q.Queue = q.Queue[:len(q.Queue)-1]
	q.queueLock.Unlock()

	return popPacket
}

func (q *ExtraQueue) HeaderModifier(p *Packet) *Packet {
	q.HeaderModifierAlgorithm(p)
	return p
}

// custom this
func (q *ExtraQueue) sort(p *Packet) {
	q.sortLock.Lock()
	defer q.sortLock.Unlock()

	q.SortAlgorithm(p)
}
