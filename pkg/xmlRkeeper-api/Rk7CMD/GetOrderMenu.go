package Rk7CMD

import (
	"fmt"

	"github.com/beevik/etree"
	"golang.org/x/net/html/charset"
)

//Структура для ответа
type GetOrderMenuResponse struct {
	ServerVersion string
	Status        string
	ArrivalTime   string
	CMD           string
	ErrorText     string
	Data          struct {
		Price      PriceScale
		TradeGroup TradeGroup
		Dishes     []Dish
		Modifiers  []Modificator
		OrderTypes []OrderType
	}
}

type PriceScale struct {
	Id   string
	Code string
	Name string
}

type TradeGroup struct {
	Id   string
	Code string
	Name string
}

type Dish struct {
	Ident string
	Price string
}

type Modificator struct {
	Ident string
	Name  string
	Price string
}

type OrderType struct {
	Ident string
}

// <?xml version="1.0" encoding="UTF-8"?>
// <RK7Query >
// 	<RK7Command2 CMD="GetOrderMenu" >
// 		<Station code="68"/>
// 	</RK7Command2>
// </RK7Query>
func RequestGetOrderMenu() ([]byte, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	rk7query := doc.CreateElement("RK7Query")

	rk7cmd2 := rk7query.CreateElement("RK7Command2")
	rk7cmd2.CreateAttr("CMD", "GetOrderMenu")
	station := rk7cmd2.CreateElement("Station")
	station.CreateAttr("code", "68")
	data, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseGetOrderMenu(xmlData []byte) (*GetOrderMenuResponse, error) {
	resp := new(GetOrderMenuResponse)

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

	price := commandResult.SelectElement("PriceScale")
	resp.Data.Price.Id = price.SelectAttrValue("id", "")
	resp.Data.Price.Code = price.SelectAttrValue("code", "")
	resp.Data.Price.Name = price.SelectAttrValue("name", "")

	tradegroup := commandResult.SelectElement("TradeGroup")
	resp.Data.TradeGroup.Id = tradegroup.SelectAttrValue("id", "")
	resp.Data.TradeGroup.Code = tradegroup.SelectAttrValue("code", "")
	resp.Data.TradeGroup.Name = tradegroup.SelectAttrValue("name", "")

	Dishes := commandResult.SelectElement("Dishes")
	for _, elem := range Dishes.SelectElements("Item") {
		item := new(Dish)
		item.Ident = elem.SelectAttrValue("Ident", "")
		item.Price = elem.SelectAttrValue("Price", "")
		resp.Data.Dishes = append(resp.Data.Dishes, *item)
	}
	return resp, nil
}
