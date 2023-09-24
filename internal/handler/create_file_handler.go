package handler

import (
	"flare/internal/durable"
	"flare/internal/models"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
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

	filePath := filepath.Join(durable.Config.SaveFolderPath, fileName)
	dirPath := filepath.Dir(filePath)

	log.Printf("File Path, %v", filePath)
	log.Printf("Dir Path, %v", dirPath)

	err = os.MkdirAll(dirPath, os.ModeDir)

	if err != nil {
		return nil, err
	}

	file, err := os.Create(filePath)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Printf("File %v created", fileName)

	return file, nil
}
