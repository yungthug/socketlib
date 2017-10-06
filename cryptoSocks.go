package main

import (
	"bufio"
	"strconv"
	"log"
	"net"
	"fmt"
)
//logs information to server console. if error, its logged and program exits.
func print(data interface{}) {
	verbose := true
	if verbose == true && data != nil {
		if x,y := data.(error);y{
			log.Fatal(x)
		} else {
			log.Println(data)
		}
	}
}
//creates udp or tcp socket based on first argument.
func Create_server(Sock_type string,ip string,port int)(interface{}){
	var serv interface{}
	socket := ip+":"+strconv.Itoa(port)
	if Sock_type == "tcp"{
		s, err := net.Listen(Sock_type,socket)
		print(err)
		serv = s
	} else if Sock_type == "udp"{
		ServerAddr,err := net.ResolveUDPAddr(Sock_type,socket)
		print(err)
		s,err := net.ListenUDP(Sock_type,ServerAddr)
		print(err)
		serv = s
	} else {
		log.Fatal("unsupported socket type! must be udp or tcp")
	}
	print(Sock_type+" server sucessfully created on port: "+strconv.Itoa(port)+", at "+ip)
	return serv
}
//waits for active connection to server then returns the active connection
func getConn(protocol interface{}) interface{} {
	var ret interface{}
	print("waiting for connection...")
	if server,Eval := protocol.(*net.TCPListener); Eval {
		conn,_ := server.Accept()
		print("Connection established...")
		ret = conn
	} else if server,Eval := protocol.(*net.UDPConn); Eval {
		ret = server
	} else {
		log.Fatal("error resolving protocol")
	}
	return ret
}
//waits for string of data from connection then returns string for handling.
func recv(protocol interface{}) string {
	var ret string
	print("listening for data...")
	if client,Eval := protocol.(*net.TCPConn); Eval {
		for {
			connbuf := bufio.NewReader(client)
			str, _ := connbuf.ReadString('\n')
			if str != ""{
				print("client said: "+str)
				ret = str
				break
			}
		}		
	} else if client,Eval := protocol.(*net.UDPConn); Eval {
		for {
			connbuf := bufio.NewReader(client)
			str, _ := connbuf.ReadString('\n')
			if str != ""{
				print("client said: "+str)
				ret = str
				break
			}
		}
	} else {
		log.Fatal("not a valid socket to recieve from")
	}
	return ret
}

func main() {
	//creating a tcp/udp server couldn't be easier. just change first arg to "udp" or "tcp"
	s := Create_server("udp","localhost",123)
	//getconn() works with both sock types. 
	client := getConn(s)
	for{
		//recv() waits for data to be recieved from client and returns as string to be handled
		x := recv(client)
		fmt.Println(x)
	}
}
/* this is a working demo of my socket lib for you to mess around with, however it is
not complete yet so there may be a few bugs... as far as i know, the 3 main functions
create_server, getconn, and recv work as planned without any errors.
*/