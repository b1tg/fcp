package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

const help = `USAGE: 
Scene: computer A(192.168.1.2) send handbook.pdf to computer B(192.168.1.3)
On computer A: ./fcp s handbook.pdf
On computer B: ./fcp r 192.168.1.2:1234 handbook-copy.pdf
`

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("wrong arg")
		os.Exit(1)
	}
	cmd := args[1]

	if cmd == "s" { //fcp s a.txt
		//send
		filename := args[2]
		send(filename)
	} else if cmd == "r" { //fcp r 192.168.1.10:1234  b.txt
		//receive
		filename := args[2]
		address := args[3]
		receive(filename, address)
	} else {
		fmt.Println("wrong arg")
		os.Exit(1)
	}
}

func send(filename string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	realname := filepath.Base(filename)
	checkError(err)

	fmt.Printf("run 'fcp r %s:1234 %s' to download\n", getIP(), realname)
	fmt.Println("waiting for client.....")
	for {
		conn, err := listener.Accept()
		checkError(err)

		rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		defer conn.Close()
		f, err := os.Open(filename)
		checkError(err)

		bf := bufio.NewReader(f)
		_, err = bf.WriteTo(rw)
		checkError(err)

		err = rw.Flush() //!important
		checkError(err)
		fmt.Println("finish write!")
		os.Exit(0)
	}
}

func receive(address string, filename string) {

	conn, err := net.Dial("tcp", address)
	checkError(err)

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	f, err := os.Create(filename)
	checkError(err)
	defer f.Close()
	w := bufio.NewWriter(f)

	rw.WriteTo(w)
	w.Flush() //!import
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
