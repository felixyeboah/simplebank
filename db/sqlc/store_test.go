package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore(t *testing.T) {
	store := NewStore(testDB)

	// Create an account
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// run n concurrent transfer transactions
	n := 10
	amount := int64(10)

	// Channel to receive the errors
	errs := make(chan error)
	// Channel to receive the results
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		fromAccountId := account1.ID
		toAccountId := account2.ID

		if i%2 == 0 {
			fromAccountId = account2.ID
			toAccountId = account1.ID
		}

		go func() {
			// Run the transfer transaction
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountId,
				ToAccountID:   toAccountId,
				Amount:        amount,
			})

			errs <- err
		}()

		// Check the transfer result
		for i := 0; i < n; i++ {
			err := <-errs
			require.NoError(t, err)

			result := <-results
			require.NotEmpty(t, result)
		}

		// Check the final account balances
		updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
		require.NoError(t, err)

		updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
		require.NoError(t, err)

		fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
		require.Equal(t, account1.Balance, updatedAccount1.Balance)
		require.Equal(t, account2.Balance, updatedAccount2.Balance)
	}
}
