package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
)

type SlugError struct {
	slug      string
	error     string
	errorType ErrorType
}

func (s SlugError) Error() string {
	return s.error
}

func (s SlugError) Slug() string {
	return s.slug
}

func (s SlugError) ErrorType() ErrorType {
	return s.errorType
}

func NewIncorrectInputError(slug, err string) SlugError {
	return SlugError{
		slug:      slug,
		error:     err,
		errorType: ErrorTypeIncorrectInput,
	}
}

func NewUnknownErr(slug, err string) SlugError {
	return SlugError{
		slug:      slug,
		error:     err,
		errorType: ErrorTypeUnknown,
	}
}
