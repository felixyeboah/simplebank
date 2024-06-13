package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestStore is a test function to simulate concurrent transfer transactions between two accounts.
// It creates two random accounts, runs n concurrent transfer transactions, and checks the final account balances.
func TestStore(t *testing.T) {
	store := NewStore(testDB)

	// Create two random accounts
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// run n concurrent transfer transactions
	n := 10
	amount := int64(10)

	// Channel to receive the errors
	errs := make(chan error)

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
			fmt.Println(">> transfer done:", i)
		}

		// Check the final account balances
		fmt.Println(">> tx:", account1.Balance, account2.Balance)
		updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
		require.NoError(t, err)

		updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
		require.NoError(t, err)

		fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
		require.Equal(t, account1.Balance, updatedAccount1.Balance)
		require.Equal(t, account2.Balance, updatedAccount2.Balance)
	}
}
