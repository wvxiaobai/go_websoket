package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin:func(r *http.Request) bool {
			return true;
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request){
	var (
		wsConn *websocket
		err error
		conn *impl.Connection
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil{
		return 
	}

	if conn, err = impl.InitConection(wsConn); err != nil {
		goto ERR
	}

	

	for {
		if data, err = conn.ReadMessage(); err != nill{
			goto ERR
		}

		if err = conn.WriteMessage(); err != nil{
			goto ERR
		}

		go func(){
			var (
				err error
			)
	
			if err = conn.WriteMessage([]byte("11111")); err != nil {
				return
			}
	
			time.Sleep(1*time.Second)
		}()
	}

	ERR:
		conn.Close()
}

func main(){
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:7777",nil)
}
