package XmlRkeeper

import (
	"github.com/beevik/etree"
)

// `<?xml version="1.0" encoding="utf-8" standalone="yes" ?>
// <RK7Query>
// <RK7CMD CMD="GetRefDataFiltered" RefName="MenuItems" IgnoreEnums="1" WithChildItems="1" WithMacroProp="1"
// 		PropMask="items.(Code,Name,ModiScheme,LargeImagePath,Status,CategPath,HighLevelGroup1,HighLevelGroup2,HighLevelGroup3,HighLevelGroup4,ComboScheme,PriceTypes*)"/>
// <Station id="15033"/>
// </RK7Query>`
func RequestGetRefData(refname string, propmask string) ([]byte, error) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	rk7query := doc.CreateElement("RK7Query")

	rk7cmd2 := rk7query.CreateElement("RK7Command2")
	rk7cmd2.CreateAttr("CMD", "GetRefDataFiltered")
	rk7cmd2.CreateAttr("RefName", "MenuItems")
	rk7cmd2.CreateAttr("IgnoreEnums", "1")
	rk7cmd2.CreateAttr("WithChildItems", "1")
	rk7cmd2.CreateAttr("WithMacroProp", "1")
	rk7cmd2.CreateAttr("PropMask", `items.(Code,Name,Status,CategPath,HighLevelGroup1,HighLevelGroup2,
		HighLevelGroup3,HighLevelGroup4,ComboScheme,PriceTypes*,genhideDish,genphotolink)`)
	doc.Indent(2)
	data, err := doc.WriteToBytes()
	if err != nil {
		return nil, err
	}
	return data, nil
}
