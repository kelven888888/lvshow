package model

import "time"

type Options struct {
	Id                int       `comment:"ID" types:"" text:"" json:"id" form:"id" range:"" edit:"0" `
	CreateTime        time.Time `comment:"" types:"" text:"" time_format:"2006-01-02" json:"create_time" form:"create_time" range:"" edit:"0"`
	UpdateTime        time.Time `comment:"" types:"" text:"" json:"update_time" form:"update_time" range:""`
	Remarks           string    `comment:"" types:"" text:"" json:"remarks" form:"remarks" range:""`
	Ticket            string    `comment:"股票代码" types:"" text:"" json:"ticket" form:"ticket" range:""`
	Cfi               string    `comment:"" types:"" text:"" json:"cfi" form:"cfi" range:""`
	ContractType      string    `comment:"合约类型" types:"" text:"" json:"contract_type" form:"contract_type" range:""`
	ExerciseStyle     string    `comment:"行权方式" types:"" text:"" json:"exercise_style" form:"exercise_style" range:""`
	ExpirationDate    time.Time `comment:"合同到期日期" time_format:"2006-01-02"  types:"" text:"" json:"expiration_date" form:"expiration_date" range:""`
	PrimaryExchange   string    `comment:"" types:"" text:"" json:"primary_exchange" form:"primary_exchange" range:""`
	SharesPerContract int       `comment:"" types:"" text:"" json:"shares_per_contract" form:"shares_per_contract" range:""`
	StrikePrice       float64   `comment:"行使价" types:"" text:"" json:"strike_price" form:"strike_price" range:""`
	UnderlyingTicker  string    `comment:"底层股票代码" types:"" text:"" json:"underlying_ticker" form:"underlying_ticker" range:""`
	Status            *int      `comment:"状态" types:"radio" text:"停用,启用" json:"status" form:"status" range:"0,1"`
	Delta             float64   `comment:"" types:"" text:"" json:"delta" form:"delta" range:""`
	Gamma             float64   `comment:"" types:"" text:"" json:"gamma" form:"gamma" range:""`
	Theta             float64   `comment:"" types:"" text:"" json:"theta" form:"theta" range:""`
	Vega              float64   `comment:"" types:"" text:"" json:"vega" form:"vega" range:""`
	LastPrice         float64   `comment:"最新价" types:"" text:"" json:"last_price" form:"last_price" range:""`
	AfterChangeVal    float64   `comment:"" types:"" text:"" json:"after_change_val" form:"after_change_val" range:""`
	AfterLowPrice     float64   `comment:"" types:"" text:"" json:"after_low_price" form:"after_low_price" range:""`
	AfterHighPrice    float64   `comment:"" types:"" text:"" json:"after_high_price" form:"after_high_price" range:""`
	AfterPrice        float64   `comment:"" types:"" text:"" json:"after_price" form:"after_price" range:""`
	CChangeRate       float64   `comment:"" types:"" text:"" json:"c_change_rate" form:"c_change_rate" range:""`
	CChangeVal        float64   `comment:"" types:"" text:"" json:"c_change_val" form:"c_change_val" range:""`
	CPrice            float64   `comment:"" types:"" text:"" json:"c_price" form:"c_price" range:""`
	PreChangeRate     float64   `comment:"" types:"" text:"" json:"pre_change_rate" form:"pre_change_rate" range:""`
	PreChangeVal      float64   `comment:"" types:"" text:"" json:"pre_change_val" form:"pre_change_val" range:""`
	PreLowPrice       float64   `comment:"" types:"" text:"" json:"pre_low_price" form:"pre_low_price" range:""`
	PreHighPrice      float64   `comment:"" types:"" text:"" json:"pre_high_price" form:"pre_high_price" range:""`
	PrePrice          float64   `comment:"" types:"" text:"" json:"pre_price" form:"pre_price" range:""`
	Volume            float64   `comment:"" types:"" text:"" json:"volume" form:"volume" range:""`
	PrevClosePrice    float64   `comment:"" types:"" text:"" json:"prev_close_price" form:"prev_close_price" range:""`
	LowPrice          float64   `comment:"" types:"" text:"" json:"low_price" form:"low_price" range:""`
	HighPrice         float64   `comment:"" types:"" text:"" json:"high_price" form:"high_price" range:""`
	OpenPrice         float64   `comment:"" types:"" text:"" json:"open_price" form:"open_price" range:""`
	ImpliedVolatility float64   `comment:"" types:"" text:"" json:"implied_volatility" form:"implied_volatility" range:""`
	OpenInterest      float64   `comment:"" types:"" text:"" json:"open_interest" form:"open_interest" range:""`
	BreakEvenPrice    float64   `comment:"" types:"" text:"" json:"break_even_price" form:"break_even_price" range:""`
	Fmv               float64   `comment:"" types:"" text:"" json:"fmv" form:"fmv" range:""`
	ChangeToBreakEven float64   `comment:"" types:"" text:"" json:"change_to_break_even" form:"change_to_break_even" range:""`
	Price             float64   `comment:"股票价格" types:"" text:"" json:"price" form:"price" range:""`
	Value             float64   `comment:"" types:"" text:"" json:"value" form:"value" range:""`
	Name              string    `comment:"" types:"" text:"" json:"name" form:"name" range:""`
	Ask               float64   `comment:"" types:"" text:"" json:"ask" form:"ask" range:""`
	AskSize           float64   `comment:"" types:"" text:"" json:"ask_size" form:"ask_size" range:""`
	Bid               float64   `comment:"" types:"" text:"" json:"bid" form:"bid" range:""`
	BidSize           float64   `comment:"" types:"" text:"" json:"bid_size" form:"bid_size" range:""`
	Midpoint          float64   `comment:"" types:"" text:"" json:"midpoint" form:"midpoint" range:""`
	MarkPrice         float64   `comment:"" types:"" text:"" json:"mark_price" form:"mark_price" range:""`
	TimeValue         float64   `comment:"" types:"" text:"" json:"time_value" form:"time_value" range:""`
	InnerValue        float64   `comment:"" types:"" text:"" json:"inner_value" form:"inner_value" range:""`
	LeverRate         float64   `comment:"" types:"" text:"" json:"lever_rate" form:"lever_rate" range:""`
	PremiumRate       float64   `comment:"" types:"" text:"" json:"premium_rate" form:"premium_rate" range:""`
	BuyProfitRate     float64   `comment:"" types:"" text:"" json:"buy_profit_rate" form:"buy_profit_rate" range:""`
}

func (Options) TableName() string {
	return "options"
}

type OptionsBase struct {
	Id                int
	CreateTime        time.Time
	UpdateTime        time.Time
	Remarks           string
	Ticket            string    `comment:"期权合约的股票代码。"`
	Cfi               string    `comment:"合同的 6 个字母 CFI 代码（"`
	ContractType      string    `comment:"合约类型。可以是“看跌”、“看涨”，或者在极少数情况下为“其他”。"`
	ExerciseStyle     string    `comment:"此合约的行权方式。有关行权方式的更多详情"`
	ExpirationDate    time.Time `comment:"合同的到期日期（YYYY-MM-DD 格式"`
	PrimaryExchange   string    `comment:"该合约上市的主要交易所的 MIC 代码。"`
	SharesPerContract int       `comment:"该合约的每份合约的股数。"`
	StrikePrice       float64   `comment:"期权合约的执行价格行使价"`
	UnderlyingTicker  string    `comment:"底层股票代码"`
	Status            int8      `comment:"0 停用 1启用 "`
	Delta             float64
	Gamma             float64
	Theta             float64
	Vega              float64
	LastPrice         float64
	AfterChangeVal    float64
	AfterLowPrice     float64
	AfterHighPrice    float64
	AfterPrice        float64
	CChangeRate       float64
	CChangeVal        float64
	CPrice            float64
	PreChangeRate     float64
	PreChangeVal      float64
	PreLowPrice       float64
	PreHighPrice      float64
	PrePrice          float64
	Volume            float64
	PrevClosePrice    float64
	LowPrice          float64
	HighPrice         float64
	OpenPrice         float64
	ImpliedVolatility float64
	OpenInterest      float64
	BreakEvenPrice    float64
	Fmv               float64
	ChangeToBreakEven float64
	Price             float64
	Value             float64
	Name              string
	Ask               float64
	AskSize           float64
	Bid               float64
	BidSize           float64
	Midpoint          float64
	MarkPrice         float64
}

type Optionstask struct {
	TdAcc            string
	Security         string
	AbAmount         int
	LockAmount       int
	UnderlyingTicker string `comment:"底层股票代码"`
	AgPrice          float64
	ExpirationDate   time.Time `comment:"合同的到期日期（YYYY-MM-DD 格式"`
	StrikePrice      float64   `comment:"期权合约的执行价格行使价"`
	Price            float64   `comment:"底层股票价格。"`
	ExerciseStyle    string    `comment:"此合约的行权方式。有关行权方式的更多详情"`
	ContractType     string    `comment:"合约类型。可以是“看跌”、“看涨”，或者在极少数情况下为“其他”。"`
	LastPrice        float64
	Exercise_Status  *int
}
