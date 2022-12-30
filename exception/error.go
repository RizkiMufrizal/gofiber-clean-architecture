package exception

func PanicLogging(err interface{}) {
	if err != nil {
		panic(err)
	}
}
