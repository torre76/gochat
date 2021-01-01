package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	incomingConnections = make(chan net.Conn)
	deadConnections     = make(chan net.Conn)
	messages            = make(chan messageWithSource)
)

// messageWithSource represents the message sent and the source of message, if any
type messageWithSource struct {
	Message string
	Source  net.Conn
}

// StartServer initialize a simple chat server
func StartServer(host string, port uint16) {

	log.Infof("Starting chat server at address [%s] - port [%d].", host, port)

	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Unable to start server. System reported [%v]", err.Error())
		os.Exit(1)
	}

	// Goroutine used to accept incoming connection
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Errorf("Unable to accept an incoming connection. System reported [%v]", err.Error())
			}
			incomingConnections <- conn
		}
	}()

	// Main server handling routine
	for {
		select {
		// Accept incoming connections.
		case conn := <-incomingConnections:
			name := registerNewClient(conn)

			broadcastMessage(fmt.Sprintf("SYSTEM: We salute [%s], our new beloved friend.\n", name), conn)

			// Goroutine used to read text from client
			go func(conn net.Conn, name string) {
				rd := bufio.NewReader(conn)
				for {
					m, err := rd.ReadString('\n')
					if err != nil {
						break
					}
					messages <- messageWithSource{
						Message: fmt.Sprintf("%s: %s", name, m),
						Source:  conn,
					}
				}

				// If here, the client has closed connection
				deadConnections <- conn
			}(conn, name)

			send(fmt.Sprintf("SYSTEM: Hi, welcome to this chat server. From now on you are recognized as [%s]. We salute you!\n", name), conn)

		case message := <-messages:
			broadcastMessage(message.Message, message.Source)
		case conn := <-deadConnections:
			if name, err := deleteClient(conn); err == nil {
				broadcast(fmt.Sprintf("SYSTEM: Hail [%s] that has left.\n", name))
			}

		}
	}

}
