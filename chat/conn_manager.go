// Package chat contains all function related to a chat server
package chat

import (
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/goombaio/namegenerator"
	log "github.com/sirupsen/logrus"
)

var (
	connections     = map[net.Conn]string{}
	connectionMutex = &sync.Mutex{}
)

// registerNewClient registers a new client into the client pool
func registerNewClient(conn net.Conn) string {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()

	// Generate a random name for this client
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	connections[conn] = nameGenerator.Generate()

	log.Infof("A new client has connected. We could refer to it as [%v].", connections[conn])
	log.Infof("Currently there is/are [%d] client(s) registered into this server.", len(connections))

	return connections[conn]
}

// deleteClient remove a client from the client pool
func deleteClient(conn net.Conn) (string, error) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()

	name, ok := connections[conn]

	if !ok {
		log.Errorf("Unable to remove connection [%v]", conn)
		return "", errors.Errorf("Unable to remove connection [%v]", conn)
	}

	delete(connections, conn)

	log.Infof("Our friend [%v] has left. I salute you!", name)
	log.Infof("Currently there is/are [%d] client(s) registered into this server.", len(connections))

	return name, nil
}
