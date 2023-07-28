package models

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
