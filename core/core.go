package core

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func SSHConnect(user, password, host string, port int) (*ssh.Session, error) {

	//get auth method
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig := &ssh.ClientConfig{
		User: user,
		Auth: auth,
		//Timeout: 30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	//connect to ssh
	addr := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return nil, err
	}

	defer client.Close()
	//create session
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func RunSSH() {

	var stdOut, stdErr bytes.Buffer

	session, err := SSHConnect("tony", "LIUchong1987!", "172.23.238.96", 22)
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	if session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal(stdErr.String())
	}
	// ret, err := strconv.Atoi(str.Replace(stdOut.String(), "\n", "", -1))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%d,%s\n", ret, stdErr.String())
	fmt.Print(stdOut.String())
}
