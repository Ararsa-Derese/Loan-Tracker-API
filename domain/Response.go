package domain

type Response struct {
	Err        error
	Message    string
	Data 	 interface{}
}
