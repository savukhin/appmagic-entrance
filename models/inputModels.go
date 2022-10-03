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

type Transaction struct {
	Time           JsonTime `json:"time"`
	GasPrice       float64  `json:"gasPrice"`
	GasValue       float64  `json:"gasValue"`
	Average        float64  `json:"average"`
	MaxGasPrice    float64  `json:"maxGasPrice"`
	MedianGasPrice float64  `json:"medianGasPrice"`
}
