package clevermining

import (
	"encoding/json"
)

type HistoryItem struct {
	Time            string `json:"time"`
	HashrateValid   uint32 `json:"hashrate_valid"`
	HashrateInvalid uint32 `json:"hashrate_invalid"`
	BtcMh           float64 `json:"btc_mh,string"`
	Balance         float64 `json:"balance"`
}

func (item *HistoryItem) UnmarshalJSON(data []byte) error {
	var err error
	type Alias HistoryItem
	aux := &struct {
		HashrateValid   json.Number `json:"hashrate_valid"`
		HashrateInvalid json.Number `json:"hashrate_invalid"`
		*Alias
	}{
		Alias: (*Alias)(item),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.HashrateValid.String() != "" {
		val, err := aux.HashrateValid.Int64()
		if err != nil {
			return err
		}
		item.HashrateValid = uint32(val)
	} else {
		item.HashrateValid = 0
	}
	if aux.HashrateInvalid.String() != "" {
		val, err := aux.HashrateInvalid.Int64()
		if err != nil {
			return err
		}
		item.HashrateInvalid = uint32(val)
	} else {
		item.HashrateInvalid = 0
	}
	return nil
}

type LiveStat struct {
	HashrateValid   uint32 `json:"hashrate_valid"`
	HashrateInvalid uint32 `json:"hashrate_invalid"`
}

type UserPoolStat struct {
	Live    LiveStat `json:"live"`
	History []HistoryItem `json:"history,omitempty"`
}
type GeneralStats struct {
	Currency string `json:"currency"`
	Balance  float64 `json:"balance"`
}

type Payout struct {
	Time   string `json:"time"`
	Amount float64 `json:"amount"`
	TxId   string `json:"txid"`
}

type UserStat struct {
	General GeneralStats `json:"general"`
	Scrypt  UserPoolStat `json:"scrypt,omitempty"`
	X11     UserPoolStat `json:"x11,omitempty"`
	Payouts []Payout `json:"payouts,omitempty"`
}

func (client *CleverMiningClient) UserStat(address string) (UserStat, error) {
	stats := UserStat{}
	_, err := client.sling.New().Get("users/" + address).ReceiveSuccess(&stats)
	if err != nil {
		return stats, err
	}

	return stats, nil
}
