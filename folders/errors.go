package folders

import "fmt"

/*
	Error Package for Folders. Built off Golang's existing `error` interface
	Contains an 'enum' of Error types, including the Error code and human readable Error string.
	Easily extensible if new errors are introduced.
*/

type FetchFolderErrorCode int
type FetchFolderError struct {
	ErrCode FetchFolderErrorCode
	Message string
}

const (
	ErrInvalidRequest FetchFolderErrorCode = iota
	ErrInvalidUUID    FetchFolderErrorCode = iota
)

var errorMessages = map[FetchFolderErrorCode]string{
	ErrInvalidRequest: "invalid request, request cannot be nil",
	ErrInvalidUUID:    "invalid UUID, uuid cannot be found",
}

func (e *FetchFolderError) Error() string {
	return fmt.Sprintf("error code %d: %s", e.ErrCode, e.Message)
}

func NewFetchFolderError(code FetchFolderErrorCode) error {
	message, ok := errorMessages[code]
	if !ok {
		message = "Unknown error"
	}
	return &FetchFolderError{ErrCode: code, Message: message}
}
