package db_sqlserver

func ExecNoQuery(strStmt string, args ...interface{}) error {

	sqlStmt, err := database.Prepare(strStmt)
	defer checkerr()
	defer sqlStmt.Close()
	if err != nil {
		errorinfo("ExecNoQuery,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
		return err
	}
	_, err = sqlStmt.Exec(args...)
	if err != nil {
		errorinfo("ExecNoQuery,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
		return err
	}
	//loginfo("ExecNoQuery，sql:%s,param:%s", strStmt, argsToString(args...))
	return nil
}

func ExecInsertGetLastId(strStmt string, args ...interface{}) (int64, error) {
	strStmt += ";select isnull(SCOPE_IDENTITY(),0)"
	sqlStmt, err := database.Prepare(strStmt)
	if err != nil {
		return 0, err
	} else {
		var seqId int64
		row := sqlStmt.QueryRow(args...)
		if err = row.Scan(&seqId); err != nil {
			return 0, err
		} else {
			return seqId, nil
		}
	}
}

func ExecQuery(dest interface{}, strStmt string, args ...interface{}) error {
	defer checkerr()
	//方法一
	if err := database.Select(dest, strStmt, args...); err != nil {
		errorinfo("ExecQuery,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
		return err
	}
	//loginfo("ExecQuery,sql:%s,param:%s", strStmt, argsToString(args...))
	return nil

}

func ExecCheckExists(strStmt string, args ...interface{}) (bool, error) {
	//defer checkerr()
	strSql := "select case when exists(" + strStmt + ") then 1 else 0 end as tt"

	var exists int

	stmt, err := database.Prepare(strSql)
	defer stmt.Close()
	if err == nil {
		row := stmt.QueryRow(args...)
		if err == nil {
			if err = row.Scan(&exists); err == nil {
				//fmt.Println(exists)
				if exists == 1 {
					return true, nil
				}
			}
		}
	}
	return false, err

}
