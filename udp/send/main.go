package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"main/layers"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp4", "10.0.3.101:4502")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("nodeID:")
	scanner.Scan()
	strNodeID := scanner.Text()
	fmt.Printf("commonKey:")
	scanner.Scan()
	strCommonKey := scanner.Text()
	fmt.Printf("pathID:")
	scanner.Scan()
	strPathID := scanner.Text()

	nodeID, err := decodeBase64(strNodeID)
	if err != nil {
		fmt.Println("decode failed nodeID", err)
	}

	commonKey, err := decodeBase64(strCommonKey)
	if err != nil {
		fmt.Println("decode failed nodeID", err)
	}

	pathID, err := decodeBase64(strPathID)
	if err != nil {
		fmt.Println("decode failed nodeID", err)
	}

	noticonf := layers.NotificationConfirmation{}
	noticonf.GenerateNotificationConfirmation(pathID, nodeID)

	buffer, err := noticonf.Marshal(commonKey)
	if err != nil {
		fmt.Println("decode failed nodeID", err)
	}

	// buffer := []byte{
	// 	// Base Header
	// 	0x00, 0x00, 0x00, 0x00, // TransactionID
	// 	0x01, 0x00, 0x09, 0x00, // Version,Flag,Type,Count
	// 	0x00, 0x00, 0x00, 0x00, // SequenceNumber
	// 	0x00, 0x60, 0x00, 0x00, // MessageLength, Next Opt, Reserved
	// 	0x66, 0xc3, 0xb3, 0xb7, // ID
	// 	0x5f, 0xab, 0x5e, 0x60,
	// 	0xa8, 0x12, 0xec, 0xd5,
	// 	0x89, 0x57, 0x59, 0xbf,
	// 	// Encrypted Path ID
	// 	0x1f, 0x6b, 0x5b, 0x4e,
	// 	0x35, 0x1e, 0x66, 0x14,
	// 	0x1d, 0x58, 0x39, 0x64,
	// 	0x14, 0xa6, 0xd8, 0xca,
	// 	0xc4, 0x62, 0x69, 0x8b,
	// 	0xd7, 0x7f, 0x52, 0xbe,
	// 	0x14, 0x60, 0x13, 0x5b,
	// 	0xc5, 0x8c, 0x82, 0x6a,
	// 	0x54, 0xe6, 0x0f, 0x79,
	// 	0x4b, 0xcf, 0xe8, 0xfb,
	// 	0xb7, 0x20, 0x07, 0x5f,
	// 	0x24, 0xb8, 0x15, 0x2a,
	// 	// HMAC

	// 	0x22, 0x95, 0x3c, 0xc4,
	// 	0x34, 0x04, 0x4a, 0xa1,
	// 	0x6a, 0x1c, 0x44, 0xdf,
	// 	0x7d, 0x01, 0x7c, 0x0c,
	// }

	fmt.Println("Sending to server")
	_, err = conn.Write(buffer)
	if err != nil {
		panic(err)
	}
}

func decodeBase64(data string) ([]byte, error) {
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return dec, nil
}
