package seNet

import (
	"TCP-server-framework/seInterface"
	"TCP-server-framework/utils"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取长度
func (d *DataPack) GetHeadLen() uint32 {
	return 8
}

// 封包
func (d *DataPack) Pack(msg seInterface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	// 将datalen写入buff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMesLen()); err != nil {
		return nil, err
	}
	// 将dataid写入buff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMesId()); err != nil {
		return nil, err
	}
	// 将data数据写入buff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMesData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// 拆包,把包的head信息读出来，之后再根据head信息里的data的长度，再进行一次读
func (d *DataPack) UnPack(data []byte) (seInterface.IMessage, error) {
	dataBuff := bytes.NewReader(data)
	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("Too long msg")
	}
	return msg, nil
}
