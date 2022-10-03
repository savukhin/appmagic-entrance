package models

import (
	"strings"
	"time"
)

type Ethereum struct {
	Transactions Transactions `json:"ethereum"`
}

type Transactions struct {
	Transaction []Transaction `json:"transactions"`
}

type JsonTime struct {
	time time.Time
}

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("06-01-02 15:04", s)
	if err != nil {
		return err
	}
	*j = JsonTime{t}
	return nil
}

// type JsonTime time.Time

// func (time1 *JsonTime) Before(time2 JsonTime) bool {
// 	return time1.time.Before(time2.time)
// }

type Transaction struct {
	Time JsonTime `json:"time"`
	// Time           string  `json:"time"`
	// Time           time.Time `json:"time"`
	GasPrice       float64 `json:"gasPrice"`
	GasValue       float64 `json:"gasValue"`
	Average        float64 `json:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice"`
}

// func (a Transactions) Len() int { return len(a.Transaction) }

// func (a Transactions) Less(i, j int) bool {
// 	return a.Transaction[i].Time.Before(a.Transaction[j].Time)
// }
// func (a Transactions) Swap(i, j int) {
// 	a.Transaction[i], a.Transaction[j] = a.Transaction[j], a.Transaction[i]
// }
