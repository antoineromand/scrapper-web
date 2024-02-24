package errors

type InsertError struct {
    Error error
    Success bool
}