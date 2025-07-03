package responsebuilder

type ResponseJson[T any] struct {
	Status             int                 `json:"status"`
	Message            string              `json:"message"`
	Data               *T                  `json:"data,omitempty"`
	ResponsePagination *ResponsePagination `json:"pagination,omitempty"`
	Errors             *T                  `json:"errors,omitempty"`
}

func NewResponseJson[T any]() *ResponseJson[T] {
	return &ResponseJson[T]{}
}

func (rj *ResponseJson[T]) SetStatus(status int) *ResponseJson[T] {
	rj.Status = status

	return rj
}

func (rj *ResponseJson[T]) SetMessage(message string) *ResponseJson[T] {
	rj.Message = message

	return rj
}

func (rj *ResponseJson[T]) SetData(data *T) *ResponseJson[T] {
	if rj.Errors == nil {
		rj.Data = data
	}

	return rj
}

func (rj *ResponseJson[T]) SetPagination(totalRows int, page int, size int) *ResponseJson[T] {
	if rj.Errors == nil {
		rj.ResponsePagination = NewResponsePagination(totalRows, page, size)
	}

	return rj
}

func (rj *ResponseJson[T]) SetErrors(errors *T) *ResponseJson[T] {
	if rj.Data == nil || rj.ResponsePagination == nil {
		rj.Errors = errors
	}

	return rj
}
