package clevermining

type ProfitabilityStats struct {
	BtcMh       float64 `json:"btc_mh,string"`
	VsReference float32 `json:"vs_reference,string"`
	LastUpdate  string  `json:"lastupdate,omitempty"`
}

type HistoryStat struct {
	BtcMh       float64 `json:"btc_mh,string"`
	VsReference float32 `json:"vs_reference,string"`
	Time        string  `json:"time"`
}

type PoolProfitability struct {
	Live          ProfitabilityStats `json:"live"`
	LastHour      ProfitabilityStats `json:"last_hour"`
	LastDay       ProfitabilityStats `json:"24_hours"`
	HistoryHourly []HistoryStat `json:"history_hourly"`
	HistoryDaily  []HistoryStat `json:"history_daily"`
}

type PoolStats struct {
	Reference     string `json:"reference"`
	Profitability PoolProfitability `json:"profitability"`
}

type ExchangeRates struct {
	UsdBtc float64 `json:"USD_BTC,string"`
	EurBtc float64 `json:"EUR_BTC,string"`
}

type PublicStat struct {
	Scrypt        PoolStats `json:"scrypt"`
	X11           PoolStats `json:"x11"`
	ExchangeRates ExchangeRates `json:"exchange_rates"`
}

func (client *CleverMiningClient) Public() (PublicStat, error) {
	public := PublicStat{}
	_, err := client.sling.New().Get("public").ReceiveSuccess(&public)
	if err != nil {
		return public, err
	}

	return public, err
}
