package hubx

import (
	"encoding/json"
)

type Handler func(clientHubMessage *ClientHubMessage)
type Filter func(clientHubMessage *ClientHubMessage, next func())

//IHub like chat room
type IHub interface {
	//Get hub's id
	Id() interface{}
	//before new client join
	BeforeJoin(callback func(client IClient) error)
	AfterJoin(callback func(client IClient))
	//On client leave
	BeforeLeave(callback func(client IClient))
	AfterLeave(callback func(client IClient))
	//Add filter
	Use(filter Filter)
	// Attach an event handler function
	On(subject string, handler Handler)
	//Dettach an event handler function
	Off(subject string, handler Handler)
	//Send message to all clients
	SendAll(subject string, message interface{}, sender IClient)
	Close()
	// SetSelf(self IHub)
	Run()
	UnregisterChan() chan IClient
	RegisterChan() chan IClient
	MessageChan() chan *ClientHubMessage
	Clients() map[IClient]bool
}
type IClient interface {
	Close()
	// Send() chan []byte
	SendChan() chan []byte
	Send(subject string, msg interface{})
	Hub() IHub
	WritePump()
	ReadPump()
	Get(key interface{}) interface{}
	Set(key, value interface{})
	NewClientHubMessage(data []byte) (*ClientHubMessage, error)
	GetClient() *Client
}

type IFilters interface {
	Do(fn Handler)
}

//find callbacks in hub
type IRoute interface {
	Route(subject string) []Handler
	// Attach an event handler function
	On(subject string, handler Handler)
	//Dettach an event handler function
	Off(subject string, handler Handler)
}

//send Message should have this format
type HubMessageOut struct {
	Subject string
	Data    interface{}
}
type HubMessageIn struct {
	Subject string
	Data    *json.RawMessage
}

type ClientHubMessage struct {
	*HubMessageIn
	Client IClient
}

func (m *HubMessageIn) Decode(obj interface{}) error {
	if nil != m.Data {
		bs, err := m.Data.MarshalJSON()
		err = json.Unmarshal(bs, &obj)
		return err
	}
	return nil
}

func (m HubMessageOut) Encode() ([]byte, error) {
	return json.Marshal(m)
}
