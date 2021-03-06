package ipc

import (
	"encoding/json"
	"fmt"
)

// Request struct
type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

// Response struct
type Response struct {
	// response Code
	Code string
	// response Body
	Body string
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c

			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			resp := server.Handle(req.Method, req.Params)

			b, err := json.Marshal(resp)

			c <- string(b)
		}
	}(session)

	fmt.Println("A new session has bean created successfully")

	return session
}
