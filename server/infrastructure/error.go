package infrastructure

func CheckError(err interface{}) {
	if err != nil {
		panic(err)
	}
}
