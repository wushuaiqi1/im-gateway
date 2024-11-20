package broadcast

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

// upgrade ws协议升级
var upgrade = websocket.Upgrader{
	//允许跨域访问
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var onlineUsers = make(map[*websocket.Conn]int64)

// AllUserApi 广播所有在线的用户
func AllUserApi(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		return
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return
	}
	//协议升级
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer closeConn(conn)
	//维持连接
	onlineUsers[conn] = uid
	//用户进线通知
	var msg string = fmt.Sprintf("用户:%d 进线通知", uid)
	noticeAllUser(msg)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func closeConn(conn *websocket.Conn) {
	uid, ok := onlineUsers[conn]
	if !ok {
		return
	}
	delete(onlineUsers, conn)
	err := conn.Close()
	if err != nil {
		return
	}
	noticeAllUser(fmt.Sprintf("用户:%d,退线通知", uid))
}

func noticeAllUser(message string) {
	for conn, _ := range onlineUsers {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			return
		}
	}
}
