package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
    service := "127.0.0.1:12000"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)
    conn, err := net.ListenUDP("udp", udpAddr)
    checkError(err)
    fmt.Println("server is running ")
    for {
        handleClient(conn)
    }
}

func handleClient(conn *net.UDPConn) {
    fmt.Println(1)
    var buf [1024]byte
    _, addr, err := conn.ReadFromUDP(buf[0:])
    fmt.Println(2)
    if err != nil {
        return
    }
    fmt.Println(3)
	message := string(buf[0:])
	message = strings.ToUpper(message)
    // 產生亂數
    rand.Seed(time.Now().UnixNano())
    randNumber := rand.Intn(10)
    if randNumber < 3 {
        return
    }
    // time.Sleep(10*time.Second)
    fmt.Println("response")
	conn.WriteToUDP([]byte(message), addr)
}


func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
        os.Exit(1)
    }
}