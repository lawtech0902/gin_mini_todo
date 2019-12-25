package e

import "fmt"

// frontend usage
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// backend usage
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// New custom err
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

// DecodeErr decode custom error
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	
	switch eType := err.(type) {
	case *Err:
		return eType.Code, eType.Message
	case *Errno:
		return eType.Code, eType.Message
	default:
	}
	
	return InternalServerError.Code, err.Error()
}
