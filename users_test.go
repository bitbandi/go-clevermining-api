package clevermining

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
			   "general":{
			      "currency":"BTC",
			      "balance":0.00616682
			   },
			   "scrypt":{
			      "live":{
				 "hashrate_valid":619741,
				 "hashrate_invalid":0
			      },
			      "history":[
				 {
				    "time":"2017-02-07 23:00",
				    "hashrate_valid":591951,
				    "hashrate_invalid":1541,
				    "btc_mh":"0.00003827",
				    "balance":0.00094328
				 },
				 {
				    "time":"2017-02-08 00:00",
				    "hashrate_valid":604672,
				    "hashrate_invalid":187,
				    "btc_mh":"0.00001763",
				    "balance":0.00139744
				 }
			      ]
			   },
			   "x11":{
			      "live":{
				 "hashrate_valid":0,
				 "hashrate_invalid":0
			      },
			      "history":[
				 {
				    "time":"2017-02-07 23:00",
				    "hashrate_valid":null,
				    "hashrate_invalid":null,
				    "btc_mh":"0.00000634",
				    "balance":0.00094328
				 }
			      ]
			   },
			   "payouts":[
			      {
				 "time":"2017-02-21 13:01:58",
				 "amount":0.01311225,
				 "txid":"fb9a97d65e1c9f9d028c569e04569aaec9ee8ae41d1692cb4934a50ee9a7563a"
			      }
			   ],
			   "error":""
			}`

	expectedItem := UserStat{
		General:GeneralStats{
			Currency:"BTC",
			Balance:0.00616682,
		},
		Scrypt:UserPoolStat{
			Live:LiveStat{
				HashrateValid:619741,
				HashrateInvalid:0,
			},
			History:[]HistoryItem{
				HistoryItem{
					Time:"2017-02-07 23:00",
					HashrateValid:591951,
					HashrateInvalid:1541,
					BtcMh:0.00003827,
					Balance:0.00094328,
				},
				HistoryItem{
					Time:"2017-02-08 00:00",
					HashrateValid:604672,
					HashrateInvalid:187,
					BtcMh:0.00001763,
					Balance:0.00139744,
				},
			},
		},
		X11:UserPoolStat{
			Live:LiveStat{
				HashrateValid:0,
				HashrateInvalid:0,
			},
			History:[]HistoryItem{
				HistoryItem{
					Time:"2017-02-07 23:00",
					HashrateValid:0,
					HashrateInvalid:0,
					BtcMh:0.00000634,
					Balance:0.00094328,
				},
			},
		},
		Payouts:[]Payout{
			Payout{
				Time:"2017-02-21 13:01:58",
				Amount:0.01311225,
				TxId:"fb9a97d65e1c9f9d028c569e04569aaec9ee8ae41d1692cb4934a50ee9a7563a",
			},
		},
	}

	mux.HandleFunc("/api/v1/users/1234", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	cmClient := NewCleverMiningClient(httpClient, "http://dummy.com/", "")
	status, err := cmClient.UserStat("1234")

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, status)
}