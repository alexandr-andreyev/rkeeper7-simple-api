package XmlRkeeper

import (
	"fmt"
	"rkeeper7-simpleapi-service/pkg/xmlRkeeper-api/Rk7CMD"
)

func (c *Client) GetRefData(refData Rk7CMD.RequestGetRefData) (any, error) {
	//Формируем тело для запроса XML
	cmd, err := Rk7CMD.CmdGetRefData(&refData)
	if err != nil {
		return nil, err
	}
	fmt.Println("CMD>", string(cmd))
	//Готовим запрос
	req, err := c.newRequest("GET", cmd)
	if err != nil {
		return nil, err
	}
	//Отправка запроса
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	//Парсинг результатов в структуру
	result, err := Rk7CMD.ParseGetRefData(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
