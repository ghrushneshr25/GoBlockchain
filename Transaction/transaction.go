package Transaction

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("Sender %s\n", t.senderBlockchainAddress)
	fmt.Printf("Recipient %s\n", t.recipientBlockchainAddress)
	fmt.Printf("Value %.2f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"SenderBlockchainAddress"`
		Recipient string  `json:"RecipientBlockchainAddress"`
		Value     float32 `json:"Value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func (t *Transaction) GetTransactionDetails() (string, string, float32) {
	return t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}
