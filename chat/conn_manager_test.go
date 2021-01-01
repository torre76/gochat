package chat

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// mockedConn represents a Mock object for net.Conn
type mockedConn struct {
	InternalRef time.Time
}

func (mc mockedConn) Read(b []byte) (n int, err error) {
	return 0, nil
}

func (mc mockedConn) Write(b []byte) (n int, err error) {
	return 0, nil
}

func (mc mockedConn) Close() error {
	return nil
}

func (mc mockedConn) LocalAddr() net.Addr {
	return nil
}

func (mc mockedConn) RemoteAddr() net.Addr {
	return nil
}

func (mc mockedConn) SetDeadline(t time.Time) error {
	return nil
}

func (mc mockedConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (mc mockedConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestRegisterDeleteNewClient(t *testing.T) {
	fakeConnection := mockedConn{
		InternalRef: time.Now(),
	}

	name := registerNewClient(fakeConnection)

	assert.NotNil(t, name)
	assert.Len(t, connections, 1)

	_, err := deleteClient(fakeConnection)

	assert.NoError(t, err)

	_, err = deleteClient(fakeConnection)

	assert.Error(t, err)
}
