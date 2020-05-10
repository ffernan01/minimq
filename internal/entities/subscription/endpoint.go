package subscription

// Address is the endpoint address
type Address string

// Endpoint for calling the subscriptor
type Endpoint struct {
	Protocol Protocol
	Address  Address
}
