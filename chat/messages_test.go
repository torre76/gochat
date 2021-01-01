package chat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBroadcastMessage(t *testing.T) {
	sourceConnection := mockedConn{
		InternalRef: time.Now(),
	}
	destinationConnection := mockedConn{
		InternalRef: time.Now(),
	}

	registerNewClient(sourceConnection)
	registerNewClient(destinationConnection)

	err := broadcastMessage("This is a test", sourceConnection)

	assert.NoError(t, err)

	err = broadcastMessage("This is a test", mockedConn{
		InternalRef: time.Now(),
	})

	assert.Error(t, err)

	err = broadcast("This is a test")

	assert.NoError(t, err)

	err = send("This is a test", sourceConnection)

	assert.NoError(t, err)

	err = send("This is a test", mockedConn{
		InternalRef: time.Now(),
	})

	assert.Error(t, err)

}
