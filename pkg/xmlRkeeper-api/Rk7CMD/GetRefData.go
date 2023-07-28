package Rk7CMD

import (
	"fmt"

	"github.com/beevik/etree"
	"golang.org/x/net/html/charset"
)

type RequestGetRefData struct {
	CMD            string
	RefName        string
	IgnoreEnums    string
	WithChildItems string
	WithMacroProp  string
	PropMask       string
	PropFilter     string
}

type ResponseRefData struct {
	ServerVersion string
	Status        string
	ArrivalTime   string
	CMD           string
	ErrorText     string
	Dataversion   string
	Data          struct {
		Items []Item
	}
}

type Item struct {
	Code       string
	Ident      string
	Name       string
	MainParent string
	Parent     string
}

// `<?xml version="1.0" encoding="utf-8" standalone="yes" ?>
// <RK7Query>
// <RK7CMD CMD="GetRefDataFiltered" RefName="MenuItems" IgnoreEnums="1" WithChildItems="1" WithMacroProp="1"
// 		PropMask="items.(Code,Name,ModiScheme,LargeImagePath,Status,CategPath,HighLevelGroup1,HighLevelGroup2,HighLevelGroup3,HighLevelGroup4,ComboScheme,PriceTypes*)"/>
// <Station id="15033"/>
// </RK7Query>`
func CmdGetRefData(reqData *RequestGetRefData) ([]byte, error) {
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

func ParseGetRefData(xmlData []byte) (*ResponseRefData, error) {
	resp := new(ResponseRefData)

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
	RkReference := commandResult.SelectElement("RK7Reference")
	resp.Dataversion = RkReference.SelectAttrValue("DataVersion", "")
	Items := RkReference.SelectElement("Items")
	for _, elem := range Items.SelectElements("Item") {
		fmt.Println()
		item := new(Item)
		item.Code = elem.SelectAttrValue("Code", "")
		item.Name = elem.SelectAttrValue("Name", "")
		item.Ident = elem.SelectAttrValue("Ident", "")
		item.MainParent = elem.SelectAttrValue("MainParent", "")
		item.Parent = elem.SelectAttrValue("Parent", "")
		resp.Data.Items = append(resp.Data.Items, *item)
	}
	return resp, nil
}
