package payment

import (
	"github.com/kukgini/cckit2/examples/payment/schema"
	m "github.com/kukgini/cckit2/state/mapping"
)

var (
	// State mappings
	StateMappings = m.StateMappings{}.Add(
		&schema.Payment{},             // state entry value will contain marshaled protobuf schema.Payment
		m.PKeyAttr(`Type`, `Id`),      // state entry key will be composite key <'Payment',{Type}, {Id}>
		m.List(&schema.PaymentList{})) // state.list() method will return marshaled protobuf schema.PaymentList
	// same same
	//Add(&schema.Payment{}, m.PKeySchema(&schema.PaymentId{}))

	// Event mappings
	EventMappings = m.EventMappings{}.
			Add(&schema.PaymentEvent{})
)
