package handler

import (
	"flare/internal/models"
	"io"
	"log"
	"net"
	"os"
)

func handleCreateFile(conn net.Conn, packetInfo models.PacketInfo) (*os.File, error) {
	fileNameLength := packetInfo.ContentLength

	log.Printf("File length %v", fileNameLength)

	buffer := make([]byte, fileNameLength)

	_, err := io.ReadFull(conn, buffer)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	fileName := string(buffer[:])

	log.Printf("File Name, %v", fileName)

	file, err := os.Create(fileName)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Printf("File %v created", fileName)

	return file, nil

}
