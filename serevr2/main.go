package main

import (
	"bufio"
	"fmt"

	"log"
	"net"
	
)


	func main(){
		li,err:=net.Listen("tcp",":8081")
		if err!=nil{
		  log.Panic(err)
		}
		defer li.Close()
	
		for{
		  conn,err:=li.Accept()
		  if err!=	nil{
			log.Panic(err)
		  }
		  go handler(conn)
		  
		}
	  }

	  func handler(conn net.Conn){
       scanner:=bufio.NewScanner(conn)
	   for scanner.Scan(){
		   ln:=scanner.Text()
		   fmt.Println(ln)
		   fmt.Fprintf(conn,"hey: %s\n",ln)
		
	   }
	   defer conn.Close()
	   fmt.Println("code got here")
	  }
