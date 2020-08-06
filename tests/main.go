package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/tebeka/selenium"
)

func main() {
	port, err := pickUnusedPort()

	var opts []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService("chromedriver",
		port, opts...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:"+strconv.Itoa(port)+"/wd/hub")
	if err != nil {
		panic(err)
	}

	wd.Refresh()

	wd.Get("http://localhost:8000/")
	defer service.Stop()
}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}
