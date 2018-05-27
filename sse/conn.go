package sse

import (
	"encoding/json"
)

// Conn describes an SSE connection.
type Conn struct {
	shutdown chan bool
	messages chan message
	isOpen   bool
}

// message describes an SSE message.
type message struct {
	id      string
	name    string
	message []byte
}

// Write sends a byte-slice to the connected client, returns an errorif the
// connection is already closed.
func (c *Conn) Write(msg []byte) error {
	return c.WriteEvent("", msg)
}

// WriteEvent sends a byte-slice to the connected client and triggers the
// specified event with the data, returns an error if the connection is already
// closed.
func (c *Conn) WriteEvent(name string, msg []byte) error {
	return c.WriteEventWithID("", name, msg)
}

// WriteEventWithID sends a byte-slice to the connected client and triggers the
// specified event with the data, returns an error if the connection is already
// closed.
func (c *Conn) WriteEventWithID(id, name string, msg []byte) error {
	if !c.isOpen {
		return errConnectionClosed
	}

	c.messages <- message{
		id:      id,
		name:    name,
		message: msg,
	}

	return nil
}

// WriteString sends a string to the connected client, returns an error if the
// connection is already closed.
func (c *Conn) WriteString(msg string) error {
	return c.WriteEvent("", []byte(msg))
}

// WriteStringEvent sends a string to the connected client, targeting the
// specified event, returns an error if the connection is already closed.
func (c *Conn) WriteStringEvent(name, msg string) error {
	return c.WriteEvent(name, []byte(msg))
}

// WriteJSON sends a json-encoded struct to the connected client, returns an
// error if the connection is already closed or if the encoding failed.
func (c *Conn) WriteJSON(value interface{}) error {
	return c.WriteJSONEvent("", value)
}

// WriteJSONEvent sends a json-encoded struct to the connected client, targeting
// the specified event, returns an error if the connection is already closed or
// if the encoding failed.
func (c *Conn) WriteJSONEvent(name string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.WriteEvent(name, b)
}

// IsOpen returns whether the connection is still opened.
func (c *Conn) IsOpen() bool {
	return c.isOpen
}

// Close forces the connection to close. The Conn object should not be used
// anymore.
func (c *Conn) Close() {
	c.shutdown <- true
}
