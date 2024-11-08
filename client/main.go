package main
import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("客户端启动>> net.Dia>>>>")
	conn,err :=l("tcp","127.0.0.1:8888")
	if(err!= nil){
		fmt.Println("连接失败: err",err)
		return
	}	
	fmt.Println("连接成功, conn:",conn);

	reader := bufio.NewReader(os.stdin)//读取控制台输入
	str,err := reader.ReadString('\n')//读取一行输入
	if err != nill{
		fmt.Println("读取控制台输入失败: err",err)
	}
	n,err := conn.Write([]byte(str))
}
