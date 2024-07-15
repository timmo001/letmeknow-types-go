package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
	UserID     *string
}

func (c *Client) Display() string {
	if c.UserID == nil {
		return fmt.Sprintf("Client(Connection=%s, UserID=nil)", c.Connection.RemoteAddr())
	}
	return fmt.Sprintf("Client(Connection=%s, UserID=%s)", c.Connection.RemoteAddr(), *c.UserID)
}

// TODO: Make sure registration contains password/token
type RequestRegister struct {
	Type   string `json:"type"` // "register"
	UserID string `json:"userID"`
}

type RequestNotification struct {
	Type    string       `json:"type"` // "notification"
	Data    Notification `json:"data"`
	Targets []string     `json:"targets"`
}

type ResponseError struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Error   *string `json:"error"`
}

type ResponseSuccess struct {
	Type      string `json:"type"`
	Succeeded bool   `json:"succeeded"`
	Message   string `json:"message"`
}
