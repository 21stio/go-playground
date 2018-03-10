package main

import (
	"log"
	"time"
	"os/exec"
	"os"
	"strconv"
	"syscall"
	"strings"
	"github.com/cbednarski/hostess"
	"net"
)

func main() {
	err := host()
	if err != nil {
		log.Fatal(err)
	}
}

func loopKill() (err error) {
	for true {
		err := run()
		if err != nil {
			log.Printf("%+v", err)
		}
		time.Sleep(500 * time.Millisecond)
		println("wait")
	}

	return
}

func host() (err error) {
	//hosts := hostess.NewHostlist()

	file := hostess.NewHostfile()
	err = file.Read()
	if err != nil {
		return
	}

	errs := file.Parse()
	for _, err = range errs {
		if err != nil {
			return
		}
	}

	print(string(file.GetData()))

	print(string(file.Format()))

	print("------------")

	print(string(file.Hosts.Format()))

	file.Hosts.Add(&hostess.Hostname{"www.google.nl", net.IPv4(byte(127), byte(0), byte(0), byte(1)), true, false})
	if err != nil {
		return
	}

	err = file.Save()
	if err != nil {
		return
	}

	//dumb, err := hosts.Dump()
	//hosts.Add(hostess.MustHostname("google.com", "127.0.0.1", false))
	//hosts.Add(hostess.MustHostname("google.com", "::1", true))
	//
	//fmt.Printf("%s\n", hosts.Format())

	return
}

func run() (err error) {
	pid, err := find("/Applications/DataGrip.app/Contents/MacOS/datagrip")
	if err != nil {
		return
	}

	pro, err := os.FindProcess(pid)
	if err != nil {
		return
	}

	err = pro.Signal(syscall.SIGTERM)
	if err != nil {
		return
	}

	return
}

func find(executable string) (pid int, err error) {
	outBytes, err := exec.Command("pgrep", "-f", executable).Output()
	if err != nil {
		return
	}

	out := strings.Replace(string(outBytes), "\n", "", -1)

	pid64bit, err := strconv.ParseInt(out, 10, 0)
	if err != nil {
		return
	}

	pid = int(pid64bit)

	return
}
