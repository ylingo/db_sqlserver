package db_sqlserver

import (
	"fmt"
)

//"fmt"
//"github.com/jmoiron/sqlx"

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
	//fmt.Println(argsToString(args...))
	sqlStmt, err := database.Prepare(strStmt)
	defer checkerr()
	defer sqlStmt.Close()
	if err != nil {
		errorinfo("ExecInsertGetLastId1,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
		return 0, err
	}
	sqlResult, err := sqlStmt.Exec(args...)
	if err != nil {
		errorinfo("ExecInsertGetLastId2,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
		return 0, err
	}
	id, err := sqlResult.LastInsertId()
	if err != nil {
		errorinfo("ExecInsertGetLastId3,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
		return 0, err
	}
	//loginfo("ExecInsertGetLastId,sql:%s,param:%s", strStmt, "", argsToString(args...))
	return id, nil
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

	//方法二
	//	stmt, err := database.Prepare(strStmt)
	//	if err != nil {
	//		errorinfo("ExecQuery,1,sql:%s,param:%s,err:%s", strStmt,  argsToString(args...), err)
	//		return err
	//	}
	//	rows, err := stmt.Query(args...)
	//	defer rows.Close()
	//	if err != nil {
	//		errorinfo("ExecQuery,2,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
	//		return err
	//	}
	//	if err := sqlx.StructScan(rows, dest); err != nil {
	//		errorinfo("ExecQuery,4,sql:%s,param:%s,err:%s", strStmt, argsToString(args...), err)
	//		return err
	//	}
	//	loginfo("ExecQuery,sql:%s,param:%s", strStmt, argsToString(args...))
	//	return nil
}

func ExecCheckExists(strStmt string, args ...interface{}) (bool, error) {
	//defer checkerr()
	strSql := "select case when exists(" + strStmt + ") then 1 else 0 end as tt"

	var exists int

	stmt, err := database.Prepare(strSql)
	if err == nil {
		row := stmt.QueryRow(args...)
		if err == nil {
			if err = row.Scan(&exists); err == nil {
				fmt.Println(exists)
				if exists == 1 {
					return true, nil
				}
			}
		}
	}
	return false, err

	//	tmp := []checkExistsTmp{}
	//	if err := ExecQuery(&tmp, strSql, args...); err != nil {
	//		return false, err
	//	}
	//	if len(tmp) == 1 && tmp[0].Exist == 1 {
	//		return true, nil
	//	}
	//	return false, nil
}

//type checkExistsTmp struct {
//	Exist int64 `db:"tt"`
//}
