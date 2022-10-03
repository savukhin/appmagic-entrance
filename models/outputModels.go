package models

type Monthly struct {
	*AbstractPeriod `json:"period"`
}

func (*Monthly) ExtractKey(transaction Transaction) string {
	return transaction.Time.time.Month().String()
}

type Daily struct {
	*AbstractPeriod `json:"period"`
}

func (*Daily) ExtractKey(transaction Transaction) string {
	return transaction.Time.time.Format("Monday")
}

type Hourly struct {
	*AbstractPeriod `json:"period"`
}

func (*Hourly) ExtractKey(transaction Transaction) string {
	return transaction.Time.time.Format("15")
}

type Absolute struct {
	*AbstractPeriod `json:"period"`
}

func (*Absolute) ExtractKey(transaction Transaction) string {
	return "absolute"
}

type Statistics struct {
	Monthly  `json:"monthly"`
	Daily    `json:"daily"`
	Hourly   `json:"hourly"`
	Absolute `json:"absolute"`
}

func CreateStatistics() *Statistics {
	monthlyPeriod := createPeriod()
	dailyPeriod := createPeriod()
	hourlyPeriod := createPeriod()
	absolutePeriod := createPeriod()

	result := &Statistics{
		Monthly:  Monthly{AbstractPeriod: monthlyPeriod},
		Daily:    Daily{AbstractPeriod: dailyPeriod},
		Hourly:   Hourly{AbstractPeriod: hourlyPeriod},
		Absolute: Absolute{AbstractPeriod: absolutePeriod},
	}
	monthlyPeriod.iPeriod = &result.Monthly
	dailyPeriod.iPeriod = &result.Daily
	hourlyPeriod.iPeriod = &result.Hourly
	absolutePeriod.iPeriod = &result.Absolute

	return result
}

func (stat *Statistics) AddStat(transaction Transaction) {
	stat.Monthly.AddStat(transaction)
	stat.Daily.AddStat(transaction)
	stat.Hourly.AddStat(transaction)
	stat.Absolute.AddStat(transaction)
}
