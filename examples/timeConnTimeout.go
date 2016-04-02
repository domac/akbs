package main

import (
	"github.com/phillihq/akbs/logger"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

type TimeoutConn struct {
	conn    net.Conn
	timeout time.Duration
}

func NewTimeoutConn(conn net.Conn, timeout time.Duration) *TimeoutConn {
	return &TimeoutConn{
		conn:    conn,
		timeout: timeout,
	}
}

func (c *TimeoutConn) Read(b []byte) (n int, err error) {
	c.SetReadDeadline(time.Now().Add(c.timeout))
	return c.conn.Read(b)
}

func (c *TimeoutConn) Write(b []byte) (n int, err error) {
	c.SetWriteDeadline(time.Now().Add(c.timeout))
	return c.conn.Write(b)
}

func (c *TimeoutConn) Close() error {
	return c.conn.Close()
}

func (c *TimeoutConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TimeoutConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TimeoutConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *TimeoutConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *TimeoutConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {

				logger.GetLogger().Infof("dia to %s://%s", netw, addr)
				conn, err := net.DialTimeout(netw, addr, time.Second*2)

				if err != nil {
					return nil, err
				}

				return NewTimeoutConn(conn, time.Second*2), nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	addr := StartTestServer()

	SendTestRequest(client, "1st", addr, "normal")
	SendTestRequest(client, "2st", addr, "normal")
	SendTestRequest(client, "3st", addr, "timeout")
	SendTestRequest(client, "4st", addr, "normal")

	time.Sleep(time.Second * 3)

	SendTestRequest(client, "5st", addr, "normal")
}

//开启测试服务器
func StartTestServer() string {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		logger.GetLogger().Infof("fail to listen - %s", err.Error())
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		http.HandleFunc("/normal", func(w http.ResponseWriter, req *http.Request) {
			time.Sleep(1000 * time.Millisecond)
			io.WriteString(w, "ok")
		})

		http.HandleFunc("/timeout", func(w http.ResponseWriter, req *http.Request) {
			time.Sleep(2500 * time.Millisecond)
			io.WriteString(w, "ok")
		})
		wg.Done()

		err = http.Serve(listener, nil)

		if err != nil {
			logger.GetLogger().Errorf("fail to start Http server - %s", err.Error())
		}
	}()
	wg.Wait()
	logger.GetLogger().Infof("start http server at http://%s/", listener.Addr())
	return listener.Addr().String()
}

//发送测试请求
func SendTestRequest(client *http.Client, id, addr, path string) {
	req, err := http.NewRequest("GET", "http://"+addr+"/"+path, nil)
	if err != nil {
		logger.GetLogger().Errorf("new request failed - %s", err)
	}
	req.Header.Add("Connection", "keep-alive")

	switch path {
	case "normal":
		if resp, err := client.Do(req); err != nil {
			logger.GetLogger().Fatalf("%s requst read failed - %s", id, err)
		} else {
			result, err2 := ioutil.ReadAll(resp.Body)
			if err2 != nil {
				logger.GetLogger().Fatalf("%s response read failed - %s", id, err2)
			}
			resp.Body.Close()
			logger.GetLogger().Infof("%s request - %s", id, result)
		}
	case "timeout":
		if _, err := client.Do(req); err == nil {
			logger.GetLogger().Infof("%s request not timeout", id)
		} else {
			logger.GetLogger().Errorf("%s request timeout - %s", id, err)
		}
	}
}
