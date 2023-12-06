package main

type PubSub struct {
	subscribers map[string][]chan interface{}
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (ps *PubSub) Publish(topic string, item interface{}) {
	for _, ch := range ps.subscribers[topic] {
		ch <- item
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan interface{} {
	ch := make(chan interface{})
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)

	return ch
}

func main() {
	ps := NewPubSub()
	ch := ps.Subscribe("t1")
	go ps.Publish("t1", "hello")

	item := <-ch
	println(item.(string))
}
