package db_sqlserver

func checkerr() {
	if err := recover(); err != nil {
		errorinfo(err)
	}
}
