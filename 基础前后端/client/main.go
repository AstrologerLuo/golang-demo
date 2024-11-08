package main
import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	fmt.Println("客户端启动>> net.Dia>>>>")
	conn,err :=net.Dial("tcp","127.0.0.1:8888")
	if err != nil{
		fmt.Println("连接失败: err",err)
		return
	}	
	fmt.Println("连接成功, conn:",conn);

	reader := bufio.NewReader(os.Stdin)//读取控制台输入

	str,err := reader.ReadString('\n')//读取一行输入
	if err != nil{
		fmt.Println("读取控制台输入失败: err",err)
	}
	n,err := conn.Write([]byte(str))//将读取的str发送给服务端
	if err !=  nil{
		fmt .Println("发送数据失败: err",err)
	}

	fmt.Printf("已发送 %d 字节数据\n",n)
}
