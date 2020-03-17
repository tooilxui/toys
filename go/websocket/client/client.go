package main

import (
	"os"
	"os/signal"

	// "encoding/base64"
	"flag"
	"fmt"

	// "net/http"

	"github.com/gorilla/websocket"

	// "io/ioutil"
	"log"
	// "net/http"
	"net/url"
	"time"
)

// var addr = flag.String("addr", "192.1.15.217:9000", "http service address")
var addr = flag.String("addr", "192.168.0.108:9000", "http service address")

func main() {

	test()
	// http.HandleFunc("/writecookie", WriteCookieServer)
	// http.ListenAndServe(":80", nil)

	// ParamURL := url.QueryEscape("rtsp://192.169.1.60/user=admin&password=&channel=1&stream=0.sdp")
	//ParamURL := url.QueryEscape("rtsp://192.1.15.218/user=admin&password=&channel=1&stream=0.sdp")
	//rtsp%3A%2F%2F192.1.15.218%2Fuser%3Dadmin%26password%3D%26channel%3D1%26stream%3D0.sdp

	// log.Println(ParamURL)

	// DecodedUrl, err := url.QueryUnescape(ParamURL)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(DecodedUrl)

	//108
	//face@aptg.com.tw
	//123456

}

func test() {

	f, err := os.OpenFile("out.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	// log.Println("This is a test log entry")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// cookie := http.Cookie{Name: "session", Value: "9d881751-812a-4ea3-a28f-5f5fa9e3fa9f", Path: "/vedio", MaxAge:86400, Secure: false}
	// http.SetCookie(w, &cookie)
	// w.Write([]byte("<b>设置cookie成功。</b>\n"))

	// ParamURL := url.QueryEscape("rtsp://192.1.15.218/user=admin&password=&channel=1&stream=0.sdp")
	ParamURL := url.QueryEscape("rtsp://192.168.0.11/user=admin&password=&channel=1&stream=0.sdp")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/video", ForceQuery: true, RawQuery: "url=" + ParamURL}

	x := u.String()
	log.Println(x)

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Print(err)
		log.Fatal("dial:", err)
	}

	defer c.Close()

	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
		log.Printf("error: %v", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
