package errs

var (
	ErrorAgeMinimum       = "User does not meet minimum age requirement"
	ErrorEmailFormat      = "User email must be properly formatted"
	ErrorEmailRequired    = "User email required"
	ErrorNameRequired     = "user first/last names required"
	ErrorNameUnique       = "User with the same first and last name already exists"
	ErrorBadRequest       = "Bad request"
	ErrorUserNotFound     = "User not found"
	ErrorValidationFailed = "User did not pass validation"
	ErrorUnexpected       = "Unexpected error"
)

type BadRequestError struct {
	Message string
}

func NewBadRequestError() *BadRequestError {
	return &BadRequestError{Message: ErrorBadRequest}
}

func (e BadRequestError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string
}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{Message: ErrorUserNotFound}
}

func (e NotFoundError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
	Details []string
}

func NewValidationError(details []string) *ValidationError {
	return &ValidationError{Message: ErrorValidationFailed, Details: details}
}

func (e ValidationError) Error() string {
	return e.Message
}

type UnexpectedError struct {
	Message string
}

func NewUnexpectedError() *UnexpectedError {
	return &UnexpectedError{Message: ErrorUnexpected}
}

func (e UnexpectedError) Error() string {
	return e.Message
}
