package main
import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error listening: err", err)
		return
	}

	for{
		conn,err2 := listen.Accept()
		if err2 != nil{
			fmt.Println("Error accepting: err2", err2)
		}else{
			fmt.Printf("等待连接成功 , con=%v, 接受到客户端信息：%v \n",conn,conn.RemoteAddr().String())
		}
	}
}