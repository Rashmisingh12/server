package main
import(
	"net"
	"fmt"
	"log"
	
	"io"
)

func main(){
    li,err:=net.Listen("tcp",":8080")
    if err!=nil{
      log.Panic(err)
    }
    defer li.Close()

    for{
      conn,err:=li.Accept()
      if err!=	nil{
        log.Println(err)
      }
      io.WriteString(conn,"\nHello from this side\n")
      fmt.Fprintln(conn,"how is your  days?")
      fmt.Fprintf(conn,"%v","well,I hope")

      conn.Close()
    }
  }