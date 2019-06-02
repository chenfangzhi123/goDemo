package main

import (
	"io/ioutil"
	//"fmt"
	"net"
	//"os"
	//"bufio"
	"log"
)

const (
	ip   = ""
	port = 8899
)

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port})
	if err != nil {
		log.Println("监听端口失败:", err.Error())
		return
	}
	log.Println("已初始化连接，等待客户端连接...")
	Server(listen)
}

func Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("接受客户端连接异常:", err.Error())
			continue
		}
		log.Println("客户端连接来自:", conn.RemoteAddr().String())
		defer conn.Close()
		go func() {
			data := make([]byte, 8192)
			for {
				i, err := conn.Read(data)
				//log.Println("客户端发来数据:\n", string(data[0:i]))
				// log.Println("客户端发来数据:\n", data)
				log.Println("客户端发来数据:", i)
				if err != nil {
					log.Println("读取客户端数据错误:", err.Error())
					break
				}
				go conn.Write(Send(data))
			}
		}()
	}
}

func Send(data []byte) (buf []byte) {
	pTCPConn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		//log.Fprintf(os.Stderr, "Error: %s", err.Error())
		log.Printf("Error: %s", err.Error())
		return
	}

	n, errWrite := pTCPConn.Write(data)
	if errWrite != nil {
		//log.Fprintf(os.Stderr, "Error: %s", errWrite.Error())
		log.Printf("Error: %s", errWrite.Error())
		return
	}
	defer pTCPConn.Close()

	//log.Fprintf(os.Stdout, "writed: %d\n", n)
	log.Printf("writed: %d\n", n)

	buf, errRead := ioutil.ReadAll(pTCPConn)
	//log.Println("服务端发来数据:\n", string(buf))
	log.Println("服务端发来数据:", len(buf))
	if errRead != nil {
		//log.Fprintf(os.Stderr, "Error: %s", errRead.Error())
		log.Printf("Error: %s", errRead.Error())
		return
	}
	//r := len(buf)
	//fmt.Fprintf(os.Stdout, string(buf[:r]))
	//fmt.Fprintf(os.Stdout, "readed: %v\n", buf)
	return
}
