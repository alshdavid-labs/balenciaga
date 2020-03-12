package topic

import (
	"fmt"
	"parkour/node/platform/websockets"
)

type ConnectionMap map[string]*websockets.Websocket

type Container struct {
	topics map[string]ConnectionMap
}

func NewContainer() *Container {
	return &Container{
		topics: map[string]ConnectionMap{},
	}
}

func (c *Container) HasKey(topic string) bool {
	_, ok := c.topics[topic]
	return ok
}

func (c *Container) Get(topicName string) ConnectionMap {
	fmt.Println(topicName)
	ok := c.HasKey(topicName)
	if ok == false {
		panic("NO_TOPIC")
	}
	return c.topics[topicName]
}

func (c *Container) Add(topicName string, conn *websockets.Websocket) {
	if c.HasKey(topicName) == false {
		c.topics[topicName] = ConnectionMap{}
	}
	c.topics[topicName][conn.SocketID] = conn
}

func (c *Container) Remove(ID string) {
	for key, topic := range c.topics {
		delete(topic, ID)
		if len(topic) == 0 {
			delete(c.topics, key)
		}
	}
}
