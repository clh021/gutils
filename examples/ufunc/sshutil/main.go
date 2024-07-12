package main

import "log"

func main() {
	log.Println("---------------------- test ssh.go")
	testSsh()

	log.Println("---------------------- test scp.go")
	testScp()

	log.Println("---------------------- test ssh-client.go")
	testSshClient()

	log.Println("---------------------- test scp-client.go")
	testScpClient()
}
