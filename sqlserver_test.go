package db_sqlserver

import (
	//"database/sql"
	//"encoding/json"
	"fmt"
	"testing"
	//"time"
)

func Init() {

	InitDb("mssql", "sqlserver://sa:111111@(local)?database=statute&encrypt=disable",
		100, 20, true, true)
	// InitDb("odbc",
	// 	"DRIVER={sql server};server=(local);uid=sa;pwd=111111;database=statute;",
	// 	100, 20, true, true)
}

func _Test_ExecQueryPaging(t *testing.T) {
	//	InitDb("mssql",
	//		"sqlserver://sa:111111@(local)/sql2016?database=teamcall&encrypt=disable",
	//		100,
	//		20, true, true)
	//	dest := []UserInfo{}
	//	totalcount, totalpage, err := ExecQueryPaging(&dest, `
	//		select a.seqid,a.userid,a.username,b.companyname,c.salary
	//		from t_userinfo a left join t_userinfoext b on a.userid=b.userid
	//		inner join t_usersalary c on a.userid=c.userid
	//		where a.seqid>0
	//		and b.companyname like '联友电讯%' and a.username like '罗%' and c.salary<>0 `, `salary desc,userid desc
	//	`, 30, 1)
	//	if err != nil {
	//		t.Error("totalcount:", totalcount, "totalpage:", totalpage, "err:", err)
	//	} else {
	//		t.Log("totalcount:", totalcount, "totalpage:", totalpage)
	//		for i := range dest {
	//			t.Log(dest[i])
	//		}
	//	}
}

func _Test_ExecNoQuery(t *testing.T) {
	Init()

	strSql := "insert into tbl_test(c1,c2,c3) values(?,?,?)"

	err := ExecNoQuery(strSql, "aaaaa", "bbbbb", "ccccc")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("执行成功")
	}

}

func _Test_ExecQuery(t *testing.T) {
	Init()
	strSql := fmt.Sprintf(`select seqid,c1,c2,c3
			from tbl_test
			`)

	p := []WordBook{}
	err := ExecQuery(&p, strSql)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("执行成功")
		t.Log(len(p))
		for _, u := range p {
			t.Logf("%s,%s,%s", u.C1, u.C2, u.C3)
		}
	}
}

type WordBook struct {
	SeqId int64  `db:"seqid"`
	C1    string `db:"c1"`
	C2    string `db:"c2"`
	C3    string `db:"c3"`
}

func Test_ExecInsertGetLastId(t *testing.T) {
	Init()
	strSql := fmt.Sprintf(`insert into tbl_test(c1,c2,c3) values(?,?,?)`)
	if seqId, err := ExecInsertGetLastId(strSql, "aaa2", "bbb2", "ccc2"); err == nil {
		t.Log(seqId)
	} else {
		t.Log(err)
	}
}

func Test_CheckExist(t *testing.T) {
	Init()
	strSql := `select seqid from tbl_test where seqid=?`
	if exist, err := ExecCheckExists(strSql, 1); err == nil {
		if exist {
			t.Log("exist")
		} else {
			t.Log("not exist")
		}
	} else {
		t.Error(err)
	}
}

type UserInfo struct {
	RowNum   int64  `db:"rownumber"`
	Seqid    int64  `db:"seqid"`
	DocId    string `db:"doc_id"`
	DocTitle string `db:"doc_h1_text"`
	DocGroup string `db:"doc_group"`
}

func Test_Paging(t *testing.T) {
	Init()

	var table string = "doc_base as t"
	var fields string = `t.seqid,
	isnull(t.doc_id,'') as doc_id,
	isnull(t.doc_h1_text,'') as doc_h1_text,
	isnull(t.doc_group,'') as doc_group`
	//	T_UserInfo.isCharge,
	//	T_UserInfo.regTime,
	var join string = ``
	var where string = ``
	var strOrder string = `t.seqId asc`
	var pageSize int = 100
	var pageIndex int = 1
	var totalcount, pagecount, outPageIndex int
	result := []UserInfo{}
	var err error
	totalcount, pagecount, outPageIndex, err = QueryByPage(&result, table, fields, where, strOrder, join, pageSize, pageIndex)
	if err != nil {
		t.Error(err)
	} else {
		// for i := range result {
		// 	t.Log(result[i])
		// }
		t.Logf("totalcount:%d,pagecount:%d,outpageIndex:%d", totalcount, pagecount, outPageIndex)

	}
}
