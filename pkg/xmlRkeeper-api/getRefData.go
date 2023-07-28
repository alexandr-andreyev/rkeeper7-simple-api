package XmlRkeeper

import (
	"fmt"
	"rkeeper7-simpleapi-service/pkg/xmlRkeeper-api/models"

	"github.com/beevik/etree"
	"golang.org/x/net/html/charset"
)

// `<?xml version="1.0" encoding="utf-8" standalone="yes" ?>
// <RK7Query>
// <RK7CMD CMD="GetRefDataFiltered" RefName="MenuItems" IgnoreEnums="1" WithChildItems="1" WithMacroProp="1"
// 		PropMask="items.(Code,Name,ModiScheme,LargeImagePath,Status,CategPath,HighLevelGroup1,HighLevelGroup2,HighLevelGroup3,HighLevelGroup4,ComboScheme,PriceTypes*)"/>
// <Station id="15033"/>
// </RK7Query>`
func RequestGetRefData(reqData *models.RequestGetRefData) ([]byte, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	rk7query := doc.CreateElement("RK7Query")

	rk7cmd2 := rk7query.CreateElement("RK7Command2")
	rk7cmd2.CreateAttr("CMD", reqData.CMD)
	rk7cmd2.CreateAttr("RefName", reqData.RefName)
	if reqData.IgnoreEnums != "" {
		rk7cmd2.CreateAttr("IgnoreEnums", reqData.IgnoreEnums)
	}
	if reqData.WithChildItems != "" {
		rk7cmd2.CreateAttr("WithChildItems", reqData.WithChildItems)
	}
	if reqData.WithMacroProp != "" {
		rk7cmd2.CreateAttr("WithMacroProp", reqData.WithMacroProp)
	}
	if reqData.PropMask != "" {
		rk7cmd2.CreateAttr("PropMask", reqData.PropMask)
	}

	data, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ResponseGetRefData(xmlData []byte) (*models.SysInfoResponse, error) {
	resp := new(models.SysInfoResponse)

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

	return resp, nil
}
