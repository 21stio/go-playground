package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"golang.org/x/net/proxy"
	"fmt"
	"os"
)

func main() {
	err := testSocks5()
	if err != nil {
		log.Print(err.Error())
	}
}

func testProxy() error {
	resp, err := http.Get("http://httpbin.org/anything")
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	println(string(body))

	return nil
}

func testSocks5() error {
	auth := proxy.Auth{}
	auth.User = "x6786021"
	auth.Password = "uxHQ5oYKQZ"
	proxyAddress := "proxy-nl.privateinternetaccess.com:1080"

	dialer, err := proxy.SOCKS5("tcp", proxyAddress, &auth, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)

		return err
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}

	httpTransport.Dial = dialer.Dial

	req, err := http.NewRequest("GET", "http://httpbin.org/anything", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)

		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)

		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)

		return err
	}
	fmt.Println(string(b))

	return nil
}
