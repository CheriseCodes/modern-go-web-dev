package repositories

import (
	"context"
	"database/sql"
)

func BeginTransaction(runnersRepository RunnersRepositoryInterface, resultsRepository ResultsRepositoryInterface) error {
	ctx := context.Background()
	transaction, err := resultsRepository.GetHandler().BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	runnersRepository.SetTransaction(transaction)
	resultsRepository.SetTransaction(transaction)
	return nil
}

func RollbackTransaction(runnersRepository RunnersRepositoryInterface, resultsRepository ResultsRepositoryInterface) error {
	transaction := runnersRepository.GetTransaction()
	runnersRepository.SetTransaction(nil)
	resultsRepository.SetTransaction(nil)
	return transaction.Rollback()
}

func CommitTransaction(runnersRepository RunnersRepositoryInterface, resultsRepository ResultsRepositoryInterface) error {
	transaction := runnersRepository.GetTransaction()
	runnersRepository.SetTransaction(nil)
	resultsRepository.SetTransaction(nil)
	return transaction.Commit()
}
