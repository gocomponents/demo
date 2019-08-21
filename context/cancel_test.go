package context

import (
	"fmt"
	"net"
	"testing"
)

func init()  {
	Cancel()
}
func TestCancel(t *testing.T) {
	fmt.Println("start client ......")
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("dial failed...", err)
		return
	}
	defer conn.Close()
	for {
		data := "jin"
		if data == "quit" { //输入quit退出
			return
		}
		_, err := conn.Write([]byte(data)) //发送数据
		if err != nil {
			fmt.Println("send data error:", err)
			return
		}
		buf := make([]byte, 512)
		n, err := conn.Read(buf) //读取服务端端数据
		fmt.Println("from server:", string(buf[:n]))
	}
}
