package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ikun666/v10/impl"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	go func() {
		for {
			// head := make([]uint32, 2)
			// err := binary.Read(conn, binary.LittleEndian, head)
			// if err != nil {
			// 	fmt.Printf("read head err:%v\n", err)
			// }
			// // fmt.Println(head)
			// body := make([]byte, head[0])
			// binary.Read(conn, binary.LittleEndian, body)
			// if err != nil {
			// 	fmt.Printf("read body err:%v\n", err)
			// }
			// fmt.Println(string(body))

			//conf not init lead to read nil panic
			pack := impl.DataPack{}
			// fmt.Println("unpack start")
			msg, err := pack.UnPack(conn)

			if err != nil {
				fmt.Printf("unpack err:%v\n", err)
				return
			}
			fmt.Println("receive msg:", string(msg.GetMsgData()))

		}
	}()
	data := "ikun ctrl"
	for {
		msg := &impl.Message{
			Len:  uint32(len(data)),
			ID:   2,
			Data: []byte(data),
		}
		pack := impl.DataPack{}
		// fmt.Println("pack start")
		sendMsg, err := pack.Pack(msg)
		if err != nil {
			fmt.Printf("pack msg err:%v", err)
			return
		}
		// fmt.Println("pack end ", sendMsg)
		n, err := conn.Write(sendMsg)
		if err != nil || n <= 0 {
			fmt.Printf("write msg err:%v", err)
			return
		}
		fmt.Println("send msg:", data)
		time.Sleep(5 * time.Second)
	}

}
