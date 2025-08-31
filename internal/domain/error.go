package domain

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
