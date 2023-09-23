package handler

import (
	"flare/internal/models"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	createFile byte = 76
	writeFile  byte = 42
	closeFile  byte = 99
)

type ConnectionHandler struct {
	openedFiles map[uint64]*os.File
}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		openedFiles: make(map[uint64]*os.File),
	}
}

func (c *ConnectionHandler) Handle(conn net.Conn) {
	for {
		buffer := make([]byte, 17)
		_, err := conn.Read(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		log.Print(buffer)

		action := buffer[0]
		packet := models.PacketInfo{
			Id:            getUInt64(buffer[1:9]),
			ContentLength: getUInt64(buffer[9:17]),
		}

		log.Printf("Id: %v, Content-Length: %d", packet.Id, packet.ContentLength)

		err = c.handle(action, packet, conn)
		if err != nil {
			c.closeOpenedFile(packet.Id)
		}
	}
}

func (c *ConnectionHandler) closeOpenedFile(id uint64) {
	file, ok := c.openedFiles[id]
	if !ok {
		file.Close()
		delete(c.openedFiles, id)
	}
}

func (c *ConnectionHandler) handle(action byte, packetInfo models.PacketInfo, conn net.Conn) error {

	logAction(action)

	switch action {
	case createFile:
		file, err := handleCreateFile(conn, packetInfo)
		if err != nil {
			return err
		}

		c.openedFiles[packetInfo.Id] = file
	case writeFile:
		file, ok := c.openedFiles[packetInfo.Id]
		if !ok {
			return nil
		}

		return handleWriteFile(conn, packetInfo, file)

	case closeFile:
		file, ok := c.openedFiles[packetInfo.Id]

		if !ok {
			return nil
		}

		err := file.Close()

		if err != nil {
			return err
		}

		delete(c.openedFiles, packetInfo.Id)
		log.Print("File Closed")
	}

	return nil
}

func logAction(action byte) {
	switch action {
	case createFile:
		log.Println("Command: Create File")
	case writeFile:
		log.Println("Command: Write File")
	case closeFile:
		log.Println("Command: Close File")
	default:
		log.Println("Command: Unknown")
	}
}

func getUInt64(bytes []byte) uint64 {
	var number uint64 = 0

	for i := 0; i < 8; i++ {
		number = number << 8
		number = number | uint64(bytes[i])
	}

	return number
}
