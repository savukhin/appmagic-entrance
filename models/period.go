package models

type StatElement struct {
	Average  float64 `json:"average"`
	Count    float64 `json:"count"`
	Title    string  `json:"title"`
	FullCost float64 `json:"full cost"`
}

func (stat *StatElement) AddStat(transaction *Transaction) {
	stat.Count += 1
	stat.FullCost += transaction.GasValue * transaction.GasPrice
	stat.Average = stat.FullCost / stat.Count
}

type iPeriod interface {
	ExtractKey(transaction *Transaction) string
	AddStat(transaction *Transaction)
}

type AbstractPeriod struct {
	iPeriod
	Stat map[string]*StatElement `json:"statistics"`
}

func createPeriod() *AbstractPeriod {
	return &AbstractPeriod{
		Stat: make(map[string]*StatElement),
	}
}

func (period *AbstractPeriod) AddStat(transaction *Transaction) {
	key := period.ExtractKey(transaction)

	if period.Stat[key] == nil {
		period.Stat[key] = &StatElement{Title: key}
	}
	period.Stat[key].AddStat(transaction)
}
