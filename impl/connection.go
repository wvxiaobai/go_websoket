package impl

import "github.com/gorilla/websocket"

type Connection struct{
	wsConn * websocket.Conne
	inChan chan []byte
	outChan chan []byte
	closeChan chan[]bool
	isClosed bool 
}

func InitConection(wsConn *websocket.Conn)(conn *Connection, err error){
	conn = &Connection{
		wsConn: wsConn,
		inChan: make(chan []byte, 1000),
		outChan: make(chan []byte, 1000),
		closeChan: make(chan []byte 1),
		
	}

	go readLoop()
	go writeLoop()
	return
}

func (conn *Connection) ReadMessage()(data []byte, err error){
	select{
	case data= <- conn.inChan:
	case <- conn.closeChan:
		err = errors.New('connect is close!')
	}
	return 
}

func (conn *Connection) WriteMessage(data []byte)(err error){
	select{
	case 	conn.outChan <- data:
	case <- conn.closeChan:
		err =  errors.New('connect is close!')
	}

}

func (conn *Connection) Close(){
	conn.wsConn.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.unLock()
}

func (conn *Connection) readLoop(){
	var(
		data []byte
		err error
	)

	for{
		if _,data,err = conn.wsConn.ReadMessage(); err != nil{
			goto ERR
		}

		select{
		case conn.inChan <- data:
		case <- conn.closeChan:
			goto ERR
		}
	}

	ERR:
		conn.Close()
}


func (conn *Connection) writeLoop(){
	var{
		data []byte
		err error
	}

	for{
		data = <- conn.outChan
		select{
		case data = <- conn.outChan:
			case <- conn.closeChan:	
				goto ERR
		}
		if err =  conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nill {
			goto ERR
		}
	}

	ERR:
		conn.Close()
}