package pubSub

type subscriberFn func(message interface{})

type cancelSub func()

type PubSub struct {
	topics  map[string]map[int]subscriberFn
	counter int
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics:  make(map[string]map[int]subscriberFn),
		counter: 0,
	}
}

func (ps *PubSub) Subscribe(topic string, subFn subscriberFn) cancelSub {
	if ps.topics[topic] == nil {
		ps.topics[topic] = make(map[int]subscriberFn)
	}

	currCounter := ps.counter
	ps.topics[topic][currCounter] = subFn
	ps.counter++

	return func() {
		delete(ps.topics[topic], currCounter)
	}
}

func (ps *PubSub) Publish(topic string, message interface{}) {
	subscribers, ok := ps.topics[topic]
	if !ok {
		return
	}

	for _, subFn := range subscribers {
		subFn(message)
	}
}
