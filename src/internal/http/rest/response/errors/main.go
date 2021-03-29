package errors

//ResponseError ...
type ResponseError struct {
	Code   int
	Title  string
	Detail string
}

//StatusCode ...
func (r ResponseError) StatusCode() int { return r.Code }

//Error ...
func (r ResponseError) Error() string { return r.Title }

//Details ...
func (r ResponseError) Details() string { return r.Detail }
