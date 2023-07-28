package Rk7CMD

import (
	"fmt"

	"github.com/beevik/etree"
	"golang.org/x/net/html/charset"
)

//Структура для ответа
type SysInfoResponse struct {
	ServerVersion string
	Status        string
	ArrivalTime   string
	CMD           string
	ErrorText     string
	Data          struct {
		NetName      string
		ShiftDate    string
		RestFullCode string
		CashGroupId  string
		RestId       string
		RestCode     string
		RestName     string
	}
}

// `<?xml version="1.0" encoding="utf-8"?>
// 	 <RK7Query>
// 		   <RK7Command2 CMD="GetSystemInfo"/>
// 	 </RK7Query>`
func RequestGetSystemInfo() ([]byte, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	rk7query := doc.CreateElement("RK7Query")

	rk7cmd2 := rk7query.CreateElement("RK7Command2")
	rk7cmd2.CreateAttr("CMD", "GetSystemInfo")

	doc.Indent(2)
	data, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ResponseGetSystemInfo(xmlData []byte) (*SysInfoResponse, error) {
	resp := new(SysInfoResponse)

	doc := etree.NewDocument()
	doc.ReadSettings.CharsetReader = charset.NewReaderLabel
	if err := doc.ReadFromBytes(xmlData); err != nil {
		fmt.Println("Read Xml from bytes error:", err)
	}
	root := doc.SelectElement("RK7QueryResult")

	resp.ServerVersion = root.SelectAttrValue("ServerVersion", "")
	resp.Status = root.SelectAttrValue("Status", "")
	resp.ArrivalTime = root.SelectAttrValue("ArrivalDateTime", "")

	commandResult := root.SelectElement("CommandResult")
	resp.CMD = commandResult.SelectAttrValue("CMD", "")
	resp.ErrorText = commandResult.SelectAttrValue("ErrorText", "")

	SystemInfo := commandResult.SelectElement("SystemInfo")
	resp.Data.NetName = SystemInfo.SelectAttrValue("NetName", "")
	resp.Data.ShiftDate = SystemInfo.SelectAttrValue("ShiftDate", "")
	resp.Data.RestFullCode = SystemInfo.SelectAttrValue("RestCode", "")

	CashGroup := SystemInfo.SelectElement("CashGroup")
	resp.Data.CashGroupId = CashGroup.SelectAttrValue("id", "")

	Restaurant := SystemInfo.SelectElement("Restaurant")
	resp.Data.RestId = Restaurant.SelectAttrValue("id", "")
	resp.Data.RestCode = Restaurant.SelectAttrValue("code", "")
	resp.Data.RestName = Restaurant.SelectAttrValue("name", "")
	return resp, nil
}
