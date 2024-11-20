package chat

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

var (
	// 连接,用于关闭连接(用户关闭页面断开连接)
	connMap map[*websocket.Conn]int64 = make(map[*websocket.Conn]int64)
	userMap map[int64]*websocket.Conn = make(map[int64]*websocket.Conn)
)

// GetConnByUserId 获取长链
func GetConnByUserId(uid int64) (*websocket.Conn, error) {
	conn, exist := userMap[uid]
	if !exist {
		log.Println("getConnByUserId:", uid, "不存在")
		return nil, errors.New(fmt.Sprintf("getConnByUserId:%d不存在", uid))
	}
	return conn, nil
}

func PutUserToOnline(uid int64, conn *websocket.Conn) {
	if conn == nil {
		return
	}
	if _, exist := connMap[conn]; !exist {
		connMap[conn] = uid
	}
	if _, exist := userMap[uid]; !exist {
		userMap[uid] = conn
	}
	log.Println(connMap)
	log.Println(userMap)
}

func DeleteUserFromOnline(conn *websocket.Conn) (uid int64) {
	if conn == nil {
		return
	}
	userId, exist := connMap[conn]
	if !exist {
		return
	}
	err := conn.Close()
	if err != nil {
		return
	}
	delete(connMap, conn)
	delete(userMap, userId)
	return userId
}

// NoticeAllUser 通知所有用户
func NoticeAllUser(message string) {
	for _, conn := range userMap {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			return
		}
	}
}
