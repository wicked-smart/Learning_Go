package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)

	if err != nil{
		panic(err)
	}
	
	p := fmt.Println

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	p(u.User.Username())

	pass, _ := u.User.Password()

	p(pass)


	p(u.Host)
	host, port, _ :=  net.SplitHostPort(u.Host)

	p(host)
	p(port)


}
