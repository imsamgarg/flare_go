package handler

import (
	"flare/internal/models"
	"io"
	"net"
	"os"
)

func handleWriteFile(conn net.Conn, packetInfo models.PacketInfo, file *os.File) error {
	buffer := make([]byte, packetInfo.ContentLength)

	_, err := io.ReadFull(conn, buffer)

	if err != nil {
		return err
	}

	_, err = file.Write(buffer)

	if err != nil {
		return err
	}

	return nil
}
