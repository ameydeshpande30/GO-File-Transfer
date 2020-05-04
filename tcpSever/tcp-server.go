package tcpSever

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

// only needed below f	or sample processing

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func sendFile(server net.Conn, srcFile string) {
	// accept connection
	screen.Clear()
	// open file to send
	// dir, _ := os.Getwd()
	// print(dir)
	start := time.Now()
	fmt.Println("Sending File", srcFile)
	// defer elapsed(srcFile)
	src := strings.Trim(srcFile, "\n")
	fi, err := os.Open("send" + string(os.PathSeparator) + src)
	checkError(err)
	defer fi.Close()

	// send file to client
	_, err = io.Copy(server, fi)

	checkError(err)
	fmt.Printf("%s took %v\n", srcFile, time.Since(start))

}

// StratServer Fuction to start server
func StartServer(ip string, ss string) {
	ln, _ := net.Listen("tcp", ip)
	conn, _ := ln.Accept()
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Received:", string(message))
	conn.Write([]byte(ss + "\n"))
	message1, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message1)
	sendFile(conn, message1)
	conn.Close()
	ln.Close()
}
