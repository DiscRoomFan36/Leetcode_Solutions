package main

type Bank struct {
    balance []int64;
}


func Constructor(balance []int64) Bank {
    result := Bank{
        balance: balance,
    };
    return result;
}

// remember to check for nil returns from this function.
func (this *Bank) maybe_get_account(account int) *int64 {
    if 1 <= account && account <= len(this.balance) {
        return &this.balance[account-1];
    } else {
        return nil;
    }
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    this_account1 := this.maybe_get_account(account1);
    if this_account1 == nil { return false; }
    this_account2 := this.maybe_get_account(account2);
    if this_account2 == nil { return false; }

    // check if it has enough money
    if *this_account1 < money { return false; }

    *this_account1 -= money;
    *this_account2 += money;
    return true;    
}


func (this *Bank) Deposit(account int, money int64) bool {
    this_account := this.maybe_get_account(account);
    if this_account == nil { return false; }
    *this_account += money;
    return true;
}


func (this *Bank) Withdraw(account int, money int64) bool {
    this_account := this.maybe_get_account(account);
    if this_account == nil { return false; }

    // check if it has enough money
    if *this_account < money { return false; }

    *this_account -= money;
    return true;
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */