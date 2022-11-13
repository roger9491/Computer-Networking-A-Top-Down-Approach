package main

/*

原作業文檔及python模板

https://gaia.cs.umass.edu/kurose_ross/programming/Python_code_only/WebServer_programming_lab_only.pdf

實作一個基於tcp協議的web服務器，使用瀏覽器輸入
http://127.0.0.1:6789/HelloWorld.html
後綴為請求文件名稱
從伺服器文件系統找到文件並回傳
找不到文件響應404報文
*/
import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

var (
	ipPort string = "127.0.0.1:6789" // ip port
)

func process(conn net.Conn) {
	// 關閉連接
	defer conn.Close()

	// 儲存回傳訊息的資料結構
	msg := make([]byte, 1024)
	_, err := conn.Read(msg)
	if err != nil {
		fmt.Println(err)
	}
	msgStr := string(msg)
	fmt.Println(msgStr)
	strArr := strings.Split(msgStr, "\n")
	fmt.Println(strArr[0])

	requestLine := strArr[0]
	fmt.Println("open file ", requestLine)

	requestLineArr := strings.Split(requestLine, " ")
	openFileName := requestLineArr[1][1:]
	fmt.Println("open file ", openFileName)
	bytes, err := ioutil.ReadFile(openFileName)
	if err != nil {
		fmt.Println(err)
		head := []byte("HTTP/1.1 404 notfundasdasd")
		conn.Write(head)
		return
	}

	// 回傳http報文
	// header := []byte("HTTP/1.1 200 OK\n")
	// conn.Write(header)
	// 回傳文件
	conn.Write(bytes)
	conn.Close()
}

func main() {
	// 建立tcp服務
	listen, err := net.Listen("tcp", ipPort)
	if err != nil {
		log.Println("listen failed, err:", err)
	}

	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}
}
