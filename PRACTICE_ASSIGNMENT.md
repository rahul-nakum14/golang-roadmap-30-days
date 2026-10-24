# 🚀 Go Capstone Assignment: CLI Digital Banking & Wallet System

A comprehensive practice project covering core Go features in one cohesive application.

---

## 📋 Concepts Checklist

- [ ] **Variables, Types & Constants** (`iota` / custom transaction type constants)
- [ ] **Structs & Custom Types** (`Transaction`, `Account`, `Bank`)
- [ ] **Methods with Pointer Receivers** (mutating state: `Deposit`, `Withdraw`)
- [ ] **Methods with Value Receivers** (read-only: `GetStatement`)
- [ ] **Slices & Dynamic Arrays** (`[]Transaction` history)
- [ ] **Maps** (`map[string]*Account` for fast lookups)
- [ ] **Conditionals & Loops** (`if/else`, `switch`, `for range` filtering)
- [ ] **Multiple Return Values & Error Handling** (`fmt.Errorf`, custom errors)
- [ ] **Interfaces & Polymorphism** (`Notifier` interface with multiple implementations)
- [ ] *(Bonus)* **Concurrency** (Goroutines & Channels for async logging)

---

## 🏗️ Architecture & Requirements

### Step 1: Types, Structs & Constants

1. **Constants**:
   ```go
   type TransactionType string

   const (
       TxDeposit    TransactionType = "DEPOSIT"
       TxWithdrawal TransactionType = "WITHDRAWAL"
       TxTransfer   TransactionType = "TRANSFER"
   )
   ```

2. **`Transaction` Struct**:
   - `ID` (`string` or `int`)
   - `Type` (`TransactionType`)
   - `Amount` (`float64`)
   - `Timestamp` (`time.Time`)

3. **`Account` Struct**:
   - `AccountNumber` (`string`)
   - `HolderName` (`string`)
   - `Balance` (`float64`)
   - `Transactions` (`[]Transaction`)

4. **`Bank` Struct**:
   - `Accounts` (`map[string]*Account`)

---

### Step 2: Account Methods (Pointer vs Value Receivers)

Implement the following methods on `Account`:

1. **`Deposit(amount float64) error`** *(Pointer receiver `*Account`)*
   - Validate `amount > 0`. If invalid, return `errors.New("deposit amount must be greater than zero")`.
   - Add amount to `Balance`.
   - Append a new `Transaction` to `Transactions`.
   - Return `nil`.

2. **`Withdraw(amount float64) error`** *(Pointer receiver `*Account`)*
   - Validate `amount > 0`.
   - Check if `amount > Balance`. If true, return `fmt.Errorf("insufficient funds: balance is %.2f, attempted %.2f", a.Balance, amount)`.
   - Deduct amount from `Balance`.
   - Record the transaction.
   - Return `nil`.

3. **`GetStatement() string`** *(Value receiver `Account`)*
   - Return a formatted multi-line summary of the account holder, balance, and all recorded transactions.

4. **`FilterTransactions(txType TransactionType) []Transaction`**
   - Use a `for range` loop to return only transactions matching the given `txType`.

---

### Step 3: Bank Operations (Maps & Multi-Return)

Implement the following methods on `Bank`:

1. **`NewBank() *Bank`**
   - Constructor function that properly initializes the `Accounts` map.

2. **`CreateAccount(accNum string, name string, initialDeposit float64) (*Account, error)`**
   - Check if `accNum` already exists in `Accounts`. If so, return an error.
   - Validate initial deposit (`initialDeposit >= 0`).
   - Create account, save it into the map, and return the pointer.

3. **`GetAccount(accNum string) (*Account, error)`**
   - Look up the account by key.
   - If not found, return `nil, fmt.Errorf("account %s not found", accNum)`.

4. **`Transfer(fromAccNum, toAccNum string, amount float64) error`**
   - Retrieve both accounts using `GetAccount`.
   - Withdraw from the sender.
   - Deposit into the receiver.
   - Ensure the transaction is atomic (if withdrawal fails, do not proceed).

---

### Step 4: Interfaces & Polymorphism

Create a decoupled notification system:

```go
type Notifier interface {
    SendNotification(recipient string, message string) error
}
```

Implement two structs satisfying `Notifier`:

1. **`EmailNotifier`**:
   - Prints: `"[EMAIL] To: <recipient> | Message: <message>"`
2. **`SMSNotifier`**:
   - Prints: `"[SMS] To: <phone/recipient> | Message: <message>"`

Add a helper function:
```go
func Notify(n Notifier, recipient, message string) {
    if err := n.SendNotification(recipient, message); err != nil {
        fmt.Println("Notification failed:", err)
    }
}
```

Call this whenever money moves (deposit, withdrawal, transfer).

---

### Step 5: Main Orchestration (`main.go`)

In `func main()`:
1. Initialize a bank with `NewBank()`.
2. Create two accounts:
   - Alice: `$1000.00`
   - Bob: `$250.00`
3. Deposit `$500.00` into Alice's account.
4. Attempt to withdraw `$2000.00` from Alice's account (catch and print the error gracefully).
5. Transfer `$200.00` from Alice to Bob.
6. Dispatch notification messages using both `EmailNotifier` and `SMSNotifier`.
7. Print formatted statements for both accounts.
8. Filter and print only `"TRANSFER"` transactions for Alice.

---

### ⭐ Bonus Challenge: Concurrency with Channels & Goroutines

- Create a background logger channel:
  ```go
  logChan := make(chan string)
  ```
- Launch a worker goroutine:
  ```go
  go func() {
      for log := range logChan {
          fmt.Printf("[AUDIT LOG]: %s\n", log)
      }
  }()
  ```
- Send an audit event into `logChan` on every banking operation.
- Close the channel before exiting `main()`.
