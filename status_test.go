package clevermining

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPublic(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			   "scrypt":{
			      "reference":"LTC",
			      "profitability":{
				 "live":{
				    "btc_mh":"0.0000225",
				    "vs_reference":"128.50",
				    "lastupdate":"2017-02-21 21:20:02"
				 },
				 "last_hour":{
				    "btc_mh":"0.0000190",
				    "vs_reference":"108.71"
				 },
				 "24_hours":{
				    "btc_mh":"0.0000236",
				    "vs_reference":"132.71"
				 },
				 "history_daily":[
				    {
				       "time":"2017-02-21",
				       "btc_mh":"0.0000239",
				       "vs_reference":"135.29"
				    },
				    {
				       "time":"2017-02-20",
				       "btc_mh":"0.0000225",
				       "vs_reference":"119.41"
				    }
				 ],
				 "history_hourly":[
				    {
				       "time":"2017-02-21 21:00:00",
				       "btc_mh":"0.0000190",
				       "vs_reference":"108.71"
				    },
				    {
				       "time":"2017-02-21 20:00:00",
				       "btc_mh":"0.0000182",
				       "vs_reference":"104.00"
				    }
				 ]
			      }
			   },
			   "x11":{
			      "reference":"DASH",
			      "profitability":{
				 "live":{
				    "btc_mh":"0.0000455",
				    "vs_reference":"372.44",
				    "lastupdate":"2017-02-21 21:20:05"
				 },
				 "last_hour":{
				    "btc_mh":"0.0000052",
				    "vs_reference":"44.13"
				 },
				 "24_hours":{
				    "btc_mh":"0.0000082",
				    "vs_reference":"100.30"
				 },
				 "history_daily":[
				    {
				       "time":"2017-02-21",
				       "btc_mh":"0.0000085",
				       "vs_reference":"98.14"
				    }
				 ],
				 "history_hourly":[
				    {
				       "time":"2017-02-21 21:00:00",
				       "btc_mh":"0.0000052",
				       "vs_reference":"44.13"
				    }
				 ]
			      }
			   },
			   "exchange_rates":{
			      "USD_BTC":"1116.57",
			      "EUR_BTC":"1054.66000"
			   }
			}`

	expectedItem := PublicStat{
		Scrypt: PoolStats{
			Reference: "LTC",
			Profitability:PoolProfitability{
				Live:ProfitabilityStats{
					BtcMh:0.0000225,
					VsReference:128.50,
					LastUpdate:"2017-02-21 21:20:02",
				},
				LastHour:ProfitabilityStats{
					BtcMh:0.0000190,
					VsReference:108.71,
				},
				LastDay:ProfitabilityStats{
					BtcMh:0.0000236,
					VsReference:132.71,
				},
				HistoryDaily:[]HistoryStat{
					HistoryStat{
						Time:"2017-02-21",
						BtcMh:0.0000239,
						VsReference:135.29,
					},
					HistoryStat{
						Time:"2017-02-20",
						BtcMh:0.0000225,
						VsReference:119.41,
					},
				},
				HistoryHourly:[]HistoryStat{
					HistoryStat{
						Time:"2017-02-21 21:00:00",
						BtcMh:0.0000190,
						VsReference:108.71,
					},
					HistoryStat{
						Time:"2017-02-21 20:00:00",
						BtcMh:0.0000182,
						VsReference:104.00,
					},
				},
			},
		},
		X11: PoolStats{
			Reference: "DASH",
			Profitability:PoolProfitability{
				Live:ProfitabilityStats{
					BtcMh:0.0000455,
					VsReference:372.44,
					LastUpdate:"2017-02-21 21:20:05",
				},
				LastHour:ProfitabilityStats{
					BtcMh:0.0000052,
					VsReference:44.13,
				},
				LastDay:ProfitabilityStats{
					BtcMh:0.0000082,
					VsReference:100.30,
				},
				HistoryDaily:[]HistoryStat{
					HistoryStat{
						Time:"2017-02-21",
						BtcMh:0.0000085,
						VsReference:98.14,
					},
				},
				HistoryHourly:[]HistoryStat{
					HistoryStat{
						Time:"2017-02-21 21:00:00",
						BtcMh:0.0000052,
						VsReference:44.13,
					},
				},
			},
		},
		ExchangeRates:ExchangeRates{
			UsdBtc:1116.57,
			EurBtc:1054.66000,
		},
	}

	mux.HandleFunc("/api/v1/public", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	cmClient := NewCleverMiningClient(httpClient, "http://dummy.com/", "")
	status, err := cmClient.Public()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, status)
}