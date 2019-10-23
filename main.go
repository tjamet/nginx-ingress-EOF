package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var host = fmt.Sprintf("http://http-echo.example.com/%d", time.Now().Unix())

func wrapDialer(f func(ctx context.Context, network, addr string) (net.Conn, error)) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		c, err := f(ctx, network, "localhost:8080")
		if err != nil {
			return c, err
		}
		fmt.Println("connecting to proxy, original address:", addr)
		_, err = c.Write([]byte("PROXY TCP4 192.168.1.3 127.0.0.1 384 8080\r\n"))
		return c, err
		// return f(ctx, network, addr)
	}
}

type HookReader struct {
	reader io.Reader
	f      func()
}

func kubectl(args ...string) {
	c := exec.Command("kubectl", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

func updateIng() {
	kubectl("annotate", "--overwrite=true", "ing/http-echo", fmt.Sprintf("nginx.ingress.kubernetes.io/configuration-snippet=# %s", time.Now().Format(time.RFC3339)))
}

func (a *HookReader) Read(p []byte) (n int, err error) {

	if a.f != nil {
		a.f()
	}
	return a.reader.Read(p)
}

func test(client *http.Client, reader io.Reader, writer io.Writer, it int) {

	resp, err := client.Post(host+fmt.Sprintf("?count=%d", it), "application/json", reader)
	if err != nil {
		fmt.Println(err)
		fmt.Println("it looks like we reproduced the problem, exiting")
		os.Exit(1)
	} else {
		defer resp.Body.Close()
		ioutil.ReadAll(resp.Body)
		//fmt.Println("ok")
	}
}

type testWriter struct {
}

func (testWriter) Write(p []byte) (n int, err error) {
	time.Sleep(50 * time.Millisecond)
	return len(p), nil
}

type nopWriter struct {
}

func (nopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func loopLongRead(client *http.Client) {
	count := 0
	for {
		test(client, &HookReader{strings.NewReader(`{"hello":"world"}`), func() { time.Sleep(50 * time.Millisecond) }}, nopWriter{}, count)
		time.Sleep(50 * time.Millisecond)
		count++
	}
}

func loopLongWrite(client *http.Client) {
	count := 0
	for {
		test(client, &HookReader{strings.NewReader(`{"hello":"world"}`), func() {}}, testWriter{}, count)
		time.Sleep(50 * time.Millisecond)
		count++
	}
}

func main() {
	transport := http.DefaultTransport.(*http.Transport)
	d := &net.Dialer{
		Timeout:   1 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	transport.DialContext = wrapDialer(d.DialContext)
	transport.MaxIdleConns = 300
	transport.MaxIdleConnsPerHost = 200
	client := http.Client{
		Transport: transport,
	}
	go loopLongRead(&client)
	//go loopLongWrite(&client)
	time.Sleep(1 * time.Second)
	updateIng()
	time.Sleep(3 * time.Second)
	fmt.Println("It looks like we didn't reproduce the problem")
}
