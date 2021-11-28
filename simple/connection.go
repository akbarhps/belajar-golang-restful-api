package simple

import "fmt"

type Connection struct {
	*File
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{file}
	return connection, func() {
		connection.Close()
	}
}

func (c *Connection) Close() {
	fmt.Println("Closing connection", c.File.Name)
}
