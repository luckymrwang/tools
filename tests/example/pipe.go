package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"os"
	"time"
)

func SendCommand(in io.WriteCloser, cmd string) error {
	if _, err := in.Write([]byte(cmd + "\n")); err != nil {
		return err
	}

	return nil
}

func main() {
	// Prompt for Username
	fmt.Print("Username: ")
	r := bufio.NewReader(os.Stdin)
	username, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Prompt for password
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}

	username = "root"
	password = []byte("200800")
	// Setup configuration for SSH client
	config := &ssh.ClientConfig{
		Timeout: time.Second * 5,
		User:    username,
		Auth: []ssh.AuthMethod{
			ssh.Password(string(password)),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the client
	client, err := ssh.Dial("tcp", "10.211.55.6:22", config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Create a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Setup StdinPipe to send commands
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()

	// Route session Stdout/Stderr to system Stdout/Stderr
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// Start a shell
	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}

	// Send username
	if _, err := stdin.Write([]byte(username)); err != nil {
		log.Fatal(err)
	}
	// Send password
	SendCommand(stdin, string(password))

	// Run configuration commands
	SendCommand(stdin, "cd /root")
	SendCommand(stdin, "ls -h")
	SendCommand(stdin, "cd security-demos")
	SendCommand(stdin, "ls -h")
	SendCommand(stdin, "exit")
	SendCommand(stdin, "N")

	session.Wait()
}
