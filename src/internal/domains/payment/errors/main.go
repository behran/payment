package errors

//PaymentError ...
type PaymentError struct {
	Code   int
	Title  string
	Detail string
}

//StatusCode ...
func (r PaymentError) StatusCode() int { return r.Code }

//Error ...
func (r PaymentError) Error() string { return r.Title }

//Details ...
func (r PaymentError) Details() string { return r.Detail }
