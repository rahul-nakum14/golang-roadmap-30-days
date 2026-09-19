package main

import (
	"errors"
	"fmt"
)

type notification interface {
	sendNotification() string
}
type email struct {
	sender   string
	receiver string
	subject  string
}
type sms struct {
	sender   string
	receiver string
	message  string
}
type whatsapp struct {
	sender   string
	receiver string
	message  string
}

func (e email) sendNotification() string {
	return fmt.Sprintf("Email sent from %s to %s with subject: %s", e.sender, e.receiver, e.subject)
}
func (s sms) sendNotification() string {
	return fmt.Sprintf("SMS sent from %s to %s with message: %s", s.sender, s.receiver, s.message)
}
func (w whatsapp) sendNotification() string {
	return fmt.Sprintf("WhatsApp message sent from %s to %s with message: %s", w.sender, w.receiver, w.message)
}

func sendNotification(n notification) {
    fmt.Println(n.sendNotification())
}
func main() {

	// var n notification
	// n = email{sender: "Rahul", receiver: "John", subject: "Hello"}
	// fmt.Println(n.sendNotification())

	// n = sms{sender: "Rahul", receiver: "John", message: "Hello"}
	// fmt.Println(n.sendNotification())

	// n = whatsapp{sender: "Rahul", receiver: "John", message: "Hello"}
	// fmt.Println(n.sendNotification())

	// sendNotification(email{sender: "Rahul", receiver: "John", subject: "Hello"})

	var input int
	fmt.Print("Enter your age: ")
	fmt.Scan(&input)
	err := checkAge(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You are an adult.")
	}

	err = withdraw(1000, 1500)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Withdrawal successful.")
	}

}

func checkAge(age int) error {
    if age < 18 {
        return errors.New("you are a minor")
    }

    return nil
}
func withdraw(balance, amount float64) error {
	if amount > balance {
		return fmt.Errorf("insufficient funds available balance: %.2f", balance)
	}
	return nil
}