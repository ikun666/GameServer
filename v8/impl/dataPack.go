package impl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ikun666/v8/iface"
)

type DataPack struct {
}

func (d *DataPack) GetHeadLen() uint32 {
	//len 4 + id 4
	return 8
}
func (d *DataPack) Pack(message iface.IMessage) ([]byte, error) {
	// fmt.Println("Pack Pack---------------")
	buf := bytes.NewBuffer([]byte{})
	//write len
	// err := buf.WriteByte(byte(message.GetMsgLen()))
	// fmt.Println(message.GetMsgLen())
	err := binary.Write(buf, binary.LittleEndian, message.GetMsgLen())
	if err != nil {
		fmt.Printf("write len err:%v\n", err)
		return nil, err
	}

	// fmt.Println("pack buf0:", buf.Len())
	//write id
	// err = buf.WriteByte(byte(message.GetMsgID()))
	err = binary.Write(buf, binary.LittleEndian, message.GetMsgID())
	if err != nil {
		fmt.Printf("write id err:%v\n", err)
		return nil, err
	}
	// fmt.Println(message.GetMsgID())
	// fmt.Println("pack buf1:", buf.Len())
	// _, err = buf.Write(message.GetMsgData())
	// fmt.Println(string(message.GetMsgData()))
	err = binary.Write(buf, binary.LittleEndian, message.GetMsgData())
	if err != nil {
		fmt.Printf("write data err:%v\n", err)
		return nil, err
	}
	// fmt.Println("pack buf2:", buf.Len(), buf.Bytes())
	return buf.Bytes(), nil
}

// func (d *DataPack) UnPack(conn net.Conn) (iface.IMessage, error) {
// 	// fmt.Println("UnPack UnPack---------------")

// 	var msgLen, msgID uint32 = 0, 0
// 	err := binary.Read(conn, binary.LittleEndian, &msgLen)
// 	if err != nil {
// 		fmt.Printf("read len err:%v\n", err)
// 		return nil, err
// 	}
// 	fmt.Printf("len:%v\n", msgLen)
// 	if msgLen > uint32(conf.GConfig.MaxPackageSize) {
// 		fmt.Printf("read len %v over MaxPackageSize %v\n", msgLen, conf.GConfig.MaxPackageSize)
// 		return nil, fmt.Errorf("read len over MaxPackageSize")
// 	}
// 	err = binary.Read(conn, binary.LittleEndian, &msgID)
// 	if err != nil {
// 		fmt.Printf("read len err:%v\n", err)
// 		return nil, err
// 	}
// 	fmt.Printf("id:=%v\n", msgID)
// 	msgData := make([]byte, msgLen)
// 	// _, err = io.ReadFull(conn, msgData)
// 	err = binary.Read(conn, binary.LittleEndian, msgData)
// 	if err != nil {
// 		fmt.Printf("read data err:%v\n", err)
// 		return nil, err
// 	}
// 	// fmt.Printf("%v\n", msgData)

// 	msg := &Message{
// 		Len:  msgLen,
// 		ID:   msgID,
// 		Data: msgData,
// 	}
// 	// fmt.Println(msg)
// 	return msg, nil
// }

func (d *DataPack) UnPack(conn io.Reader) (iface.IMessage, error) {
	// fmt.Println("UnPack UnPack---------------")

	head := make([]uint32, 2)
	// fmt.Println("head:", head, "conn:", conn)
	err := binary.Read(conn, binary.LittleEndian, head)
	if err != nil {
		fmt.Printf("read head err:%v\n", err)
		return nil, err
	}
	// if head[0] > uint32(conf.GConfig.MaxPackageSize) {
	// 	fmt.Printf("read len %v over MaxPackageSize %v\n", head[0], conf.GConfig.MaxPackageSize)
	// 	return nil, fmt.Errorf("read len over MaxPackageSize")
	// }
	if head[0] > 2048 {
		fmt.Printf("read len %v over MaxPackageSize %v\n", head[0], 2048)
		return nil, fmt.Errorf("read len over MaxPackageSize")
	}
	// fmt.Println("head:", head, "conn:", conn)
	body := make([]byte, head[0])
	// _, err = io.ReadFull(conn, msgData)
	err = binary.Read(conn, binary.LittleEndian, body)
	if err != nil {
		fmt.Printf("read body err:%v\n", err)
		return nil, err
	}
	// fmt.Printf("%v\n", msgData)

	msg := &Message{
		Len:  head[0],
		ID:   head[1],
		Data: body,
	}
	// fmt.Println(msg)
	return msg, nil
}
