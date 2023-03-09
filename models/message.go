package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

// 消息
type Message struct {
	gorm.Model
	FormId   int64  //发送者
	TargetId int64  // 接受者
	Type     int    //发送类型 1私聊 2 群聊 3广播
	Media    int    // 消息类型 1文字  2 表情包 3 图片  4 音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 需求：发送者ID，接受者ID，消息类型，发送的内容，发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	//1.获取参数， 校验 token 合法性
	//token := query.Get("token")
	query := request.URL.Query()
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//msgType := query.Get("Type")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	isvalida := true //checkToken() 待......
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//3.用户关系
	//4.userid 跟node绑定 并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	//5.完成发送的逻辑
	go sendProc(node)
	//6.完成接收的逻辑
	go recvProc(node)
	sendMsg(userId, []byte("欢迎进入聊天系统"))
}
func sendProc(node *Node) {
	fmt.Println("sendProc函数")
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// 后端接收消息
func recvProc(node *Node) {
	fmt.Println("recvProc函数")
	for {
		_, data, err := node.Conn.ReadMessage()

		if err != nil {
			fmt.Println("recvProc err:", err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws]<<<<<<", data)
	}

}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}
func init() {
	go updSendProc()
	go updRecvProc()
}

// 完成upd数据发送协程
func updSendProc() {
	fmt.Println("updSendProc函数")
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 10),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// 完成udp数据接收协程
func updRecvProc() {
	fmt.Println("updRecvProc函数")
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	//私信
	case 1:
		sendMsg(msg.TargetId, data)
		//case 2:
		//   sendGroupMsg()
		//case 3:
		//   sendAllMsg()
		//case 4:
		//
	}
}
func sendMsg(userId int64, msg []byte) {
	fmt.Println("sendMsg函数")
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
