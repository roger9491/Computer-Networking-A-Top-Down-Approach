package main

/*

原作業文檔及python模板

https://gaia.cs.umass.edu/kurose_ross/programming/Python_code_only/UDP_Pinger_programming_lab_only.pdf


實作一個基於udp協議的web客戶端
功能:
	ping 10 次 server端 ，並計算RTT時間，1秒為超時時間
server端程式碼為教材提供


*/

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
    
    udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:12000")
    checkError(err)
    conn, err := net.DialUDP("udp", nil, udpAddr)
    checkError(err)

	// Ping 10 次
	for i := 1; i <= 10; i++ {
		start := time.Now().UnixNano()
		_, err = conn.Write([]byte("aksdlkal"))
		checkError(err)

		// 設置超時時間
		err = conn.SetReadDeadline(time.Now().Add(1*time.Second))
		if err != nil {
			log.Println("SetReadDeadline failed: ", err)
		}
		var buf [1024]byte
		_, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Printf("Sequence %d: Request timed out\n", i)
			continue
		}

		fmt.Printf("Sequence %d: Reply from 127.0.0.1    RTT = %d ns\n", i, time.Now().UnixNano() - start)

	}
    os.Exit(0)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
        os.Exit(1)
    }
}
