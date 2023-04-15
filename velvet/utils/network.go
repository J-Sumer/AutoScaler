package utils

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
)

func getPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func GetFreePort() (int, error) {
	port, err := getPort()
	if err != nil {
		return 0, err
	}

	fmt.Println("Port ", port)

	// To expose the port to outside world
	// sudo iptables -I INPUT -p tcp -s 0.0.0.0/0 --dport 3306 -j ACCEPT
	cmd := exec.Command("sudo", "iptables", "-I", "INPUT", "-p", "tcp", "-s", "0.0.0.0/0", "--dport", strconv.Itoa(port), "-j", "ACCEPT")
	eRun := cmd.Run()
	if eRun != nil {
		log.Fatal(err)
	}

	return port, nil
}