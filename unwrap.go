package utils

/*
Unwrap a tuple containing the desired output and an error.
Panic if err is not nil.
*/
func Unwrap[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}

	return a
}
