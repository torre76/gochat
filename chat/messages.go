package chat

import (
	"net"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// broadcastMessage send a message to other client except for the source.
//
// If the source is nil, the broadcast is sent to everyone
func broadcastMessage(msg string, source net.Conn) error {
	var targetName string

	if source != nil {
		name, ok := connections[source]

		if !ok {
			log.Errorf("Unable to find connection [%v] in connection pool", source)
			return errors.Errorf("connection [%v] not found", source)
		}

		targetName = name
	}

	counter := 0
	for conn, name := range connections {
		if name != targetName {
			if _, err := conn.Write([]byte(msg)); err != nil {
				log.Errorf("Unable to send message to [%v]", name)
				return errors.Wrapf(err, "unable to send message to [%v]", name)
			}

			counter++
		}
	}

	if targetName != "" {
		log.Infof("Sent message [%v] from [%v] to [%d] client(s).", msg, targetName, counter)
	} else {
		log.Infof("Broadcasted message [%v] to [%d] client(s).", msg, counter)
	}

	return nil
}

// broadcast send message to all clients connected
func broadcast(msg string) error {
	return broadcastMessage(msg, nil)
}

// send sends a message to a specific client
func send(msg string, target net.Conn) error {
	name, ok := connections[target]

	if !ok {
		log.Errorf("Unable to find connection [%v] in connection pool", target)
		return errors.Errorf("connection [%v] not found", target)
	}

	if _, err := target.Write([]byte(msg)); err != nil {
		log.Errorf("Unable to send message to [%v]", name)
		return errors.Wrapf(err, "unable to send message to [%v]", name)
	}

	log.Infof("Sent message [%v] to [%v].", msg, name)

	return nil

}
