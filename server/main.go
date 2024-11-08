package main
import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()
	for{
		buf := make([]byte ,1024)

		n,err := conn.Read(buf)
		if err != nil{
			return 
		}
		fmt.Println(string(buf[0:n]))
	}
}

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

		go process(conn)//不同客户端的请求，链接的conn不一样，所以需要用go来处理请求
	}
}