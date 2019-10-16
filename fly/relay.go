package fly

import (
	"fmt"
	"github.com/xtaci/kcp-go"
	"net"
)

func TCPToUDP(session *kcp.UDPSession, conn net.Conn) {
	buff := make([]byte, 4096)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			logger.Println(err)
			break
		}
		//logs.Printf("TCPToUDP read %d byte\n", n)
		n, err = session.Write(buff[:n])
		if err != nil {
			logger.Println(err)
			break
		}
	}
}

func UDPToTCP(conn net.Conn, session *kcp.UDPSession) {
	buff := make([]byte, 4096)
	for {
		n, err := session.Read(buff)
		if err != nil {
			logger.Println(err)
			break
		}
		//logs.Printf("UDPToTCP read %d byte\n", n)
		n, err = conn.Write(buff[:n])
		if err != nil {
			logger.Println(err)
			break
		}
	}
}

func RelayTraffic(dst, src net.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := src.Read(buff)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Read", n)
		if n > 0 {
			if n, err = dst.Write(buff[:n]); err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}
