package chat

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 协议升级
var upgrade = websocket.Upgrader{
	//允许跨域访问
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func PeerToPeerChat(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("userId is nil")
		return
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		log.Println("uid parse int err:", err)
	}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade fail:", err)
		return
	}
	defer CloseConn(conn)
	//推送用户进入在线列表
	PutUserToOnline(uid, conn)
	//用户进线通知
	NoticeAllUser(fmt.Sprintf("用户:%d 进线通知", uid))
	//监听客户端消息
	for {
		msgType, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("read message error:", err)
			break
		}
		// 文本类型消息处理
		if msgType == websocket.TextMessage {
			log.Println(string(bytes))
			var msg Message
			err := json.Unmarshal(bytes, &msg)
			if err != nil {
				log.Println("text message parse error:", err)
				continue
			}
			// 消息转发和发送
			msg.Session = uuid.NewString()
			msg.Timestamp = time.Now().UnixMilli()
			res := MessageTransmit(msg)
			if !res {
				log.Println("消息:", msg.Body, "发送失败")
			}
		}

	}

}

func CloseConn(conn *websocket.Conn) {
	uid := DeleteUserFromOnline(conn)
	NoticeAllUser(fmt.Sprintf("用户:%d,退线通知", uid))
}

// MessageTransmit 消息转发
func MessageTransmit(msg Message) (res bool) {
	if msg.From == 0 || msg.To == 0 || msg.Body == "" {
		return false
	}
	conn, err := GetConnByUserId(msg.To)
	if err != nil {
		return false
	}
	return SendMessage(conn, msg)
}

// SendMessage 发送消息
func SendMessage(conn *websocket.Conn, msg Message) (res bool) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Println("SendMessage json parse error:", err)
		return false
	}
	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		log.Println("SendMessage write msg error:", err)
		return false
	}
	return true
}
