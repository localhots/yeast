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

func (u *Unit) Call(data []byte) (resp []byte, err error) {
	var (
		addr = &net.UnixAddr{u.socketPath(), "unix"}
		conn *net.UnixConn
		buf  bytes.Buffer
	)

	if conn, err = net.DialUnix("unix", nil, addr); err != nil {
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

	if _, err = buf.ReadFrom(conn); err != nil {
		fmt.Println("Failed read data from socket:", err.Error())
		return
	}

	return buf.Bytes(), nil
}

func (u *Unit) Units() []string {
	return []string{u.Name}
}

func (u *Unit) socketPath() string {
	return strings.Join([]string{"/tmp/unit_", u.Name, ".sock"}, "")
}
