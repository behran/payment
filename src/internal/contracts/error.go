package contracts

//IPaymentError ....
type IPaymentError interface {
	StatusCode() int
	Error() string
	Details() string
}
