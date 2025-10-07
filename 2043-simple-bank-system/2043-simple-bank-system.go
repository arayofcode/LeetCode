type Bank struct {
    total    int
    balances []int64
}


func Constructor(balance []int64) Bank {
    return Bank{
        total: len(balance),
        balances: append([]int64{0}, balance...),
    }
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if this.validAccount(account1) && this.validAccount(account2) && this.balances[account1] >= money {
        this.balances[account1] -= money
        this.balances[account2] += money
        return true
    }
    return false
}


func (this *Bank) Deposit(account int, money int64) bool {
    if this.validAccount(account) {
        this.balances[account] += money
        return true
    }
    return false
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if this.validAccount(account) && money <= this.balances[account] {
        this.balances[account] -= money
        return true
    }
    return false
}

func (this *Bank) validAccount(account int) bool {
    return account >= 1 && account <= this.total
}