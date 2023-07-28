package models

type RequestGetRefData struct {
	CMD            string
	RefName        string
	IgnoreEnums    string
	WithChildItems string
	WithMacroProp  string
	PropMask       string
	PropFilter     string
}

type GetRefDataResponse struct {
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
