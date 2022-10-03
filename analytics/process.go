package analytics

import (
	"appmagic-entrance/models"
	"sync"
)

func Process(ethereum *models.Ethereum) (*models.Statistics, error) {
	result := models.CreateStatistics()

	var wg sync.WaitGroup
	workers := 1

	transactionChannel := make(chan *models.Transaction, workers)

	for worker := 0; worker < workers; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for transaction := range transactionChannel {
				result.AddStat(transaction)
			}

		}()
	}

	for _, transaction := range ethereum.Transactions.Transaction {
		transactionChannel <- &transaction
	}
	close(transactionChannel)

	wg.Wait()

	return result, nil
}
