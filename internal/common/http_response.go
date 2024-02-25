package common

type HttpResponse struct {
    Error error
    Success bool
    Code int
    Message string
    Data interface{}
}
