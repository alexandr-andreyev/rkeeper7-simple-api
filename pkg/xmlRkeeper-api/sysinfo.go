package XmlRkeeper

import (
	"rkeeper7-simpleapi-service/pkg/xmlRkeeper-api/Rk7CMD"
)

func (c *Client) GetSystemInfo() (any, error) {
	//Формируем тело для запроса XML
	cmd, err := Rk7CMD.RequestGetSystemInfo()
	if err != nil {
		return nil, err
	}
	//Готовим запрос
	req, err := c.newRequest("POST", cmd)
	if err != nil {
		return nil, err
	}
	//Отправка запроса
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	//Парсинг результатов в структуру
	result, err := Rk7CMD.ResponseGetSystemInfo(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}
