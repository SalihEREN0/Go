package main

import (
	"bufio"
	"net"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

type Client struct {
	name             string
	conn             *net.Conn
	chatRoom         *ChatRoom
	incomingMessages chan *Message
	outgoingMessages chan *Message
	reader           *bufio.Reader
	writer           *bufio.Writer
}

type ChatRoom struct {
	name     string
	clients  []Client
	messages []Message
}

type Message struct {
	time   time.Time
	client *Client
	text   string
}

type Lobby struct {
	clients          []Client
	chatRooms        map[string]*ChatRoom
	incomingMessages chan *Message
	join             chan *Client
	leave            chan *Client
	deleteRoom       chan *ChatRoom
}

func (l *Lobby) Listen() {
	go func() {
		for {
			select {
			case client := <-l.join:
				l.clients = append(l.clients, *client)
			case client := <-l.leave:
				for i, c := range l.clients {
					if c == *client {
						l.clients = append(l.clients[:i], l.clients[i+1:]...)
						break
					}
				}
			case message := <-l.incomingMessages:
				for _, client := range l.clients {
					client.outgoingMessages <- message
				}
			case chatRoom := <-l.deleteRoom:
				delete(l.chatRooms, chatRoom.name)
			}
		}
	}()
}

func NewLobby() *Lobby {
	lobby := &Lobby{
		clients:          make([]Client, 0),
		chatRooms:        make(map[string]*ChatRoom),
		incomingMessages: make(chan *Message),
		join:             make(chan *Client),
		leave:            make(chan *Client),
		deleteRoom:       make(chan *ChatRoom),
	}
	lobby.Listen() //Burayı goroutine yapmaya gerek var mı??
	return lobby
}

func (l *Lobby) Join(client *Client) {
}

func (l *Lobby) Leave(client *Client) {
}

func (l *Lobby) CreateChatRoom(c *Client, name string) {
	if l.chatRooms[name] != nil {
		return
	}

}

func (l *Lobby) DeleteChatRoom(name string) {
}

func main() {
}
