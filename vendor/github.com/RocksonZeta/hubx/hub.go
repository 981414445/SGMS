package hubx

import "fmt"

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[IClient]bool
	// Inbound messages from the clients.
	message chan *ClientHubMessage
	// Register requests from the clients.
	register chan IClient
	// Unregister requests from clients.
	unregister chan IClient
	filters    []Filter
	handlers   map[string]Handler
	// hubs       *Hubs
	self IHub
	//identify hub
	id          interface{}
	beforeJoin  func(client IClient) error
	afterJoin   func(client IClient)
	beforeLeave func(client IClient)
	afterLeave  func(client IClient)
}

func NewHub(id interface{}) IHub {
	hub := &Hub{
		clients:    make(map[IClient]bool),
		message:    make(chan *ClientHubMessage),
		register:   make(chan IClient),
		unregister: make(chan IClient),
		handlers:   make(map[string]Handler),
		id:         id,
	}
	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("Hub.register")
			if nil != h.beforeJoin {
				err := h.beforeJoin(client)
				if err != nil {
					return
				}
			}
			h.clients[client] = true
			if nil != h.afterJoin {
				h.afterJoin(client)
			}
		case client := <-h.unregister:
			fmt.Println("Hub.unregister")
			if _, ok := h.clients[client]; ok {
				if nil != h.beforeLeave {
					h.beforeLeave(client)
				}
				delete(h.clients, client)
				close(client.SendChan())
				if nil != h.afterLeave {
					h.afterLeave(client)
				}
			}
		case message := <-h.message:
			h.OnMessage(message)
		}
	}
}

func (h *Hub) OnMessage(msg *ClientHubMessage) {
	subject := msg.HubMessageIn.Subject
	if handler, ok := h.handlers[subject]; ok {
		h.Filter(handler, msg)
	}
}

func (h *Hub) Filter(handler Handler, msg *ClientHubMessage) {
	if len(h.filters) <= 0 {
		handler(msg)
		return
	}
	pos := 0
	filter := h.filters[pos]
	var next func()
	next = func() {
		pos++
		if pos >= len(h.filters) {
			handler(msg)
			return
		}
		filter = h.filters[pos]
		filter(msg, next)
	}
	filter(msg, next)

}

func (h *Hub) BeforeJoin(callback func(client IClient) error) {
	h.beforeJoin = callback
}
func (h *Hub) AfterJoin(callback func(client IClient)) {
	h.afterJoin = callback
}
func (h *Hub) AfterLeave(callback func(client IClient)) {
	h.afterLeave = callback
}
func (h *Hub) BeforeLeave(callback func(client IClient)) {
	h.beforeLeave = callback
}

func (h *Hub) Use(filter Filter) {
	h.filters = append(h.filters, filter)
}

func (h *Hub) On(subject string, handler Handler) {
	h.handlers[subject] = handler
}

func (h *Hub) Off(subject string, handler Handler) {
	delete(h.handlers, subject)
}
func (h *Hub) Close() {
	// h.hubs.RemoveHub(h.Id())
	for client, _ := range h.clients {
		client.Close()
	}
}
func (h *Hub) Id() interface{} {
	return h.id
}
func (h *Hub) SendAll(subject string, message interface{}, sender IClient) {
	for client, _ := range h.clients {
		client.Send(subject, message)
	}
}

// func (h *Hub) SetSelf(self IHub) {
// 	h.self = self
// }
func (h *Hub) RegisterChan() chan IClient {
	return h.register
}
func (h *Hub) UnregisterChan() chan IClient {
	return h.unregister
}

// RegisterChan() chan *Client
func (h *Hub) MessageChan() chan *ClientHubMessage {
	return h.message
}
func (h *Hub) Clients() map[IClient]bool {
	return h.clients
}
