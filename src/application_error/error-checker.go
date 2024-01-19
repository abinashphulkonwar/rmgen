package applicationerror

func ErrorChecker(err *error) {
	if *err != nil {
		panic(*err)
	}
}
