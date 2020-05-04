package tcpClient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func recieveFile(server net.Conn, dstFile string) {
	// accept connection
	screen.Clear()
	// create new file
	start := time.Now()
	fo, err := os.Create("recv" + string(os.PathSeparator) + dstFile)
	checkError(err)
	defer fo.Close()
	fmt.Println("Reciving File", dstFile)

	// accept file from client & write to new file
	_, err = io.Copy(fo, server)
	checkError(err)
	fmt.Printf("%s took %v\n", dstFile, time.Since(start))

}

func RecvFile(ip string) {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Text to send: ")
	// text, _ := reader.ReadString('\n')
	// conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	conn, _ := net.Dial("tcp", ip)
	fmt.Fprint(conn, "OK\n")
	ll, _ := bufio.NewReader(conn).ReadString('\n')
	var ss []string
	var n int
	json.Unmarshal([]byte(ll), &ss)
	for i := 0; i < len(ss); i++ {
		fmt.Println(i+1, ss[i], "\n")
	}
	fmt.Println("Choose File : ")
	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace(text)
	// n, _ := strconv.ParseInt(text, 10, 64)
	// fmt.Scanf("%d\n", &n)
	stdin := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Fscan(stdin, &n)
		if err == nil {
			break
		}

		stdin.ReadString('\n')
		fmt.Println("Sorry, invalid input. Please enter an integer: ")
	}
	fmt.Fprintln(conn, ss[n-1])
	recieveFile(conn, ss[n-1])
}
