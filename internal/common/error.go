package common

type ErrorKind int

const (
	ErrNotFound   ErrorKind = iota // 404
	ErrInvalid                     // 400
	ErrPermission                  // 403
	ErrConflict                    // 409
	ErrInternal                    // 500
)

func (k ErrorKind) String() string {
	switch k {

	case ErrInvalid: // 400
		return "Invalid"
	case ErrPermission: // 403
		return "Permission"
	case ErrNotFound: // 404
		return "NotFound"
	case ErrConflict: // 409
		return "Conflict"
	case ErrInternal: // 500
		return "Internal"
	default:
		return "Unknown"
	}
}

func (k ErrorKind) StatusCode() int {
	switch k {

	case ErrInvalid: // 400
		return 400
	case ErrPermission: // 403
		return 403
	case ErrNotFound: // 404
		return 404
	case ErrConflict: // 409
		return 409
	case ErrInternal: // 500
		return 500
	default:
		return 500
	}
}

type AppError struct {
	Kind    ErrorKind
	Message string
}

func NewAppError(kind ErrorKind, msg string) *AppError {
	return &AppError{
		Kind:    kind,
		Message: msg,
	}
}

// これがあるから error インターフェースを満たす
func (e *AppError) Error() string {
	return e.Message
}
