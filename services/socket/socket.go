package socketservice

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type SocketConnection struct {
	id         string
	userId     string
	connection *websocket.Conn
}
type MessageContent struct{}

type BroadcastData struct {
	SenderId   string `json:"senderId`
	ChatRoomId string `json:"chatRoomId`
	Content    string `json:"content`
}

type SocketManager struct {
	// socketPool          map[string]*SocketConnection
	socketPool          map[string][]*SocketConnection
	register            chan *SocketConnection
	unregister          chan *SocketConnection
	broadcastToChatRoom chan *BroadcastData
}

var Manager = &SocketManager{
	socketPool:          make(map[string][]*SocketConnection),
	register:            make(chan *SocketConnection),
	unregister:          make(chan *SocketConnection),
	broadcastToChatRoom: make(chan *BroadcastData),
}

func Register(connection *websocket.Conn, userID string) *SocketConnection {
	// Create a SocketConnection and push it to the Socket Manager
	id, _ := uuid.NewV4()
	var socketConnection = &SocketConnection{
		id:         id.String(),
		userId:     userID,
		connection: connection,
	}
	Manager.register <- socketConnection
	return socketConnection
}

// Need to optimize the logic or change the structure of the socket pool
func getSocketConnectionFromSocketPool(userID string) []*SocketConnection {
	return Manager.socketPool[userID]
}

func BroadcastDataToChatRoom(data *BroadcastData) {
	var receiversIds = findMembersInChatRoom(data.ChatRoomId)
	for _, receiverId := range receiversIds {
		var socketConnections = getSocketConnectionFromSocketPool(receiverId)
		for _, socketConnection := range socketConnections {
			socketResponse, _ := json.Marshal(data)
			socketConnection.connection.WriteMessage(1, socketResponse)
		}
	}
}

func (socketConnection *SocketConnection) IncomingSocketMessageListener() {
	defer func() {
		Manager.unregister <- socketConnection
	}()
	for {
		_, message, err := socketConnection.connection.ReadMessage()
		if err != nil {
			Manager.unregister <- socketConnection
			break
		}

		// log.Print(bytes.NewBuffer(message)) // this works fine
		var socketMessage BroadcastData
		err = json.Unmarshal(message, &socketMessage) // handle error afterwards
		Manager.broadcastToChatRoom <- &socketMessage
	}
}

func findMembersInChatRoom(receiver string) []string {
	// hit the db and get the user / users id with the chat room id
	abc := []string{
		"sd",
	}
	return abc
}

func InitiateSocketWorker() {
	for {
		select {
		case conn := <-Manager.register:
			// add this connection to the manager's socket pool
			if _, ok := Manager.socketPool[conn.userId]; ok {
				Manager.socketPool[conn.userId] = append(Manager.socketPool[conn.userId], conn)
			}
			log.Print("Socket connection established :: ID :: ", conn.id, "::  for user id :: ", conn.userId)
		case conn := <-Manager.unregister:
			// remove the sockets form the socket pool
			var connections = Manager.socketPool[conn.userId]
			for i, connection := range connections {
				if conn.id == connection.id {
					connections = append(connections[:i], connections[i+1:]...)
					log.Print("Socket connection terminated :: ID :: ", conn.id, "::  for user id :: ", conn.userId)
					break
				}
			}
		case data := <-Manager.broadcastToChatRoom:
			// call the broadcast method
			log.Print(" Message received::: ", data)
			BroadcastDataToChatRoom(data)
		}
	}
}
