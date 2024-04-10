package tcpserver

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	// wait for the quit channel
	go s.acceptLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			// should check for end of file stuff etc
			fmt.Println("read error", err)
			continue
		}
		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
		conn.Write([]byte("Thank you for your message"))
	}
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}
		go s.readLoop(conn)
	}
}

func TCPServer() {
	server := NewServer(":3000")

	go func() {
		for msg := range server.msgch {
			if string(msg.payload) == "close" {
				server.quitch <- struct{}{}
			}
			fmt.Println("recieved message from connections:", string(msg.payload), msg.payload)
		}
	}()

	log.Fatal(server.Start())
}
