package network

import (
	"net"
	"fmt"
	"github.com/GreyHood-Studio/ai_server/utils"
	"bufio"
)

type TCPClient struct {
	// server info
	serverId	int
	serverIP	string
	serverPort	int

	// connect data
	commuicate	chan []byte
	buf			[]byte
	conn		net.Conn
	reader		*bufio.Reader
	writer		*bufio.Writer
}

func (TCPClient *TCPClient) write() {
	for {
		select {
			case data :=  <- TCPClient.commuicate:
			TCPClient.buf = append(data, '\n')
			//length, err := gameClient.eventConn.Write(data)
			//utils.CheckError(err, "write error")
			TCPClient.writer.Write(TCPClient.buf)
		}
	}
}

func (TCPClient *TCPClient) read() {
	for {
		data, err := TCPClient.reader.ReadBytes('\n')
		utils.NoDeadError(err, "ai server read error")
		fmt.Println(data)
	}
}

func (tcpClient *TCPClient) listen() {
	go tcpClient.write()
	go tcpClient.read()
}

func (tcpClient *TCPClient) Run() {
	serverAddress := fmt.Sprint(tcpClient.serverIP, ":", tcpClient.serverPort)
	conn, err := net.Dial("tcp", serverAddress)
	utils.CheckError(err, "connect error")

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	tcpClient.conn = conn
	tcpClient.reader = reader
	tcpClient.writer = writer

	tcpClient.listen()
}

func NewClient(serverId int, serverIP string, serverPort int) *TCPClient {
	client := &TCPClient{
		serverId: serverId,
		serverIP: serverIP,
		serverPort: serverPort,

		buf: make([]byte, 1024),
		commuicate: make(chan []byte),
	}
	return client
}