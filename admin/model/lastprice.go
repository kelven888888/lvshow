package model

type LastPrice struct {
	Tickers   []Tickers `json:"tickers"`
	Status    string    `json:"status"`
	RequestID string    `json:"request_id"`
	Count     int       `json:"count"`
}
type Day struct {
	O  int `json:"o"`
	H  int `json:"h"`
	L  int `json:"l"`
	C  int `json:"c"`
	V  int `json:"v"`
	Vw int `json:"vw"`
}

type LastTrade struct {
	C  []int   `json:"c"`
	I  string  `json:"i"`
	P  float64 `json:"p"`
	S  int     `json:"s"`
	T  int64   `json:"t"`
	X  int     `json:"x"`
	Ds string  `json:"ds"`
}
type Min struct {
	Dv  string  `json:"dv"`
	Dav string  `json:"dav"`
	Av  int     `json:"av"`
	T   int64   `json:"t"`
	N   int     `json:"n"`
	O   float64 `json:"o"`
	H   float64 `json:"h"`
	L   float64 `json:"l"`
	C   float64 `json:"c"`
	V   int     `json:"v"`
	Vw  float64 `json:"vw"`
}
type PrevDay struct {
	O  float64 `json:"o"`
	H  float64 `json:"h"`
	L  int     `json:"l"`
	C  float64 `json:"c"`
	V  float64 `json:"v"`
	Vw float64 `json:"vw"`
}
type Tickers struct {
	Ticker           string    `json:"ticker"`
	TodaysChangePerc float64   `json:"todaysChangePerc"`
	TodaysChange     float64   `json:"todaysChange"`
	Updated          int64     `json:"updated"`
	Day              Day       `json:"day"`
	LastTrade        LastTrade `json:"lastTrade"`
	Min              Min       `json:"min"`
	PrevDay          PrevDay   `json:"prevDay"`
}
