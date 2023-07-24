package XmlRkeeper

import (
	"fmt"

	"github.com/beevik/etree"
	"golang.org/x/net/html/charset"
)

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

func ResponseGetSystemInfo(xmlData []byte) (*SysInfoResponse, error) {
	resp := new(SysInfoResponse)

	doc := etree.NewDocument()
	doc.ReadSettings.CharsetReader = charset.NewReaderLabel
	if err := doc.ReadFromBytes(xmlData); err != nil {
		fmt.Println("Read Xml from bytes error:", err)
	}
	root := doc.SelectElement("RK7QueryResult")
	serverVersion := root.SelectAttrValue("ServerVersion", "")
	resp.ServerVersion = serverVersion

	reqStatus := root.SelectAttrValue("Status", "")
	resp.Status = reqStatus
	arrivalTime := root.SelectAttrValue("ArrivalDateTime", "")
	resp.ArrivalTime = arrivalTime

	commandResult := root.SelectElement("CommandResult")
	cmdText := commandResult.SelectAttrValue("CMD", "")
	resp.CMD = cmdText
	errorText := commandResult.SelectAttrValue("ErrorText", "")
	resp.ErrorText = errorText

	SystemInfo := commandResult.SelectElement("SystemInfo")
	netnameCashServer := SystemInfo.SelectAttrValue("NetName", "")
	resp.Data.NetName = netnameCashServer
	shiftDate := SystemInfo.SelectAttrValue("ShiftDate", "")
	resp.Data.ShiftDate = shiftDate
	restFullCode := SystemInfo.SelectAttrValue("RestCode", "")
	resp.Data.RestFullCode = restFullCode

	CashGroup := SystemInfo.SelectElement("CashGroup")
	CashGroupId := CashGroup.SelectAttrValue("id", "")
	resp.Data.CashGroupId = CashGroupId

	Restaurant := SystemInfo.SelectElement("Restaurant")
	RestId := Restaurant.SelectAttrValue("id", "")
	resp.Data.RestId = RestId
	RestCode := Restaurant.SelectAttrValue("code", "")
	resp.Data.RestCode = RestCode
	RestName := Restaurant.SelectAttrValue("name", "")
	resp.Data.RestName = RestName
	return resp, nil
}
