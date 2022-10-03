package analytics

import (
	"appmagic-entrance/models"
)

func Process(ethereum *models.Ethereum) (*models.Statistics, error) {
	result := models.CreateStatistics()

	for _, transaction := range ethereum.Transactions.Transaction {
		result.AddStat(&transaction)
	}

	return result, nil
}
