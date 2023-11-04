package functions

func Map[T any, U any](values []T, fn func(T) U) []U {
	result := make([]U, 0);
	for _, v := range values {
		result = append(result, fn(v));
	}
	return result;
}