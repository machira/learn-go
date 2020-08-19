package pointers_and_errors

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}
		got, _ := wallet.Withdraw(5)

		want := Bitcoin(5)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		_, err := wallet.Withdraw(20)

		assertBalance(t, wallet, Bitcoin(10))

		if err != insuffientFunds {
			t.Error("wanted an error but didn't get one")
		}
	})
}
