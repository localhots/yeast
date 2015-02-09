package unit

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

type (
	Unit struct {
		Name string
	}
)

func New(name string) *Unit {
	return &Unit{
		Name: name,
	}
}

func (u *Unit) Send(data []byte) (resp []byte, err error) {
	conn, err := net.DialUnix("unix", nil, &net.UnixAddr{u.socketPath(), "unix"})
	if err != nil {
		fmt.Println("Failed opening socket:", err.Error())
		return
	}
	defer conn.Close()

	if _, err = conn.Write(data); err != nil {
		fmt.Println("Failed to write data to socket:", err.Error())
		return
	}
	if err = conn.CloseWrite(); err != nil {
		fmt.Println("Failed to close socket for reading:", err.Error())
		return
	}

	var respBuf bytes.Buffer
	if _, err = respBuf.ReadFrom(conn); err != nil {
		fmt.Println("Failed read data from socket:", err.Error())
		return
	}
	resp = respBuf.Bytes()

	return
}

func (u *Unit) socketPath() string {
	return strings.Join([]string{"/tmp/unit_", u.Name, ".sock"}, "")
}
