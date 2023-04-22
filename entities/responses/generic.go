package responses

type GenericResponse[T interface{}] struct {
	Message string
	Data    T
}
