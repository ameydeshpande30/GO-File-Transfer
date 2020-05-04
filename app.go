package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	client "./tcpClient"
	server "./tcpSever"
)

func sendFile(ss string) {
	fmt.Println("Send File")
	server.StartServer(":8081", ss)
}

func recvFile(ip string) {
	fmt.Println("Recv File")
	client.RecvFile(ip)
}

func readFiles() string {
	files, _ := ioutil.ReadDir("send")
	var ll []string
	for _, f := range files {
		// fmt.Println(f.Name())
		ll = append(ll, f.Name())
	}
	rr, _ := json.Marshal(ll)
	ss := string(rr)
	return ss
}

func main() {
	var a int = 5
	_ = os.Mkdir("send", os.ModePerm)
	_ = os.Mkdir("recv", os.ModePerm)
	ss := readFiles()
	for a > 0 {
		var input int
		fmt.Print("1) Send File\n2) Recv File\n3) Quit\n:")
		fmt.Scanln(&input)
		if input == 1 {
			sendFile(ss)
		}
		if input == 2 {
			var ip string
			fmt.Println("Give IP of TCP Server")
			fmt.Scanf("%s", &ip)
			recvFile(ip)
		}
		if input == 3 {
			a = 0
		}
	}

}
