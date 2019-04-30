package db_sqlserver

import (
	//"database/sql"
	//"encoding/json"
	"fmt"
	"testing"
	"time"
	//"time"
)

func Init() {

	InitDb("sqlserver",
		"sqlserver://sa:111111@(local)/sql2016?database=teamcall&encrypt=disable",
		100,
		20, true, true)
}

type ProcPagingBase struct {
	RowTotal  int `db:"rowtotal" json:omit`
	RowNumber int `db:"rownumber" json:"omit"`
}

type UserInfo struct {
	ProcPagingBase
	SeqId       int64   `db:"seqid"`
	Salary      float32 `db:"salary"`
	UserId      string  `db:"userid"`
	UserName    string  `db:"username"`
	CompanyName string  `db:"companyname"`
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
	//Init()

	strSql := fmt.Sprintf(`insert into T_NewBilling_UserAccount(
		accountId,userId,financeId,fixedFee,baseUnit,exceedRate,
		updateTime,updateUser) values(
		'%s','%s','%s',%f,%d,%f,getdate(),'%s')`,
		"accountId_test", "userId_test", "financeId",
		123.4, 666, 0.1234, "yanglin_test")

	seqId, err := ExecInsertGetLastId(strSql)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("执行成功,seqId:", seqId)
	}

}

func _Test_ExecQuery(t *testing.T) {
	//Init()
	//	strSql := fmt.Sprintf(`select seqid,accountid,userid,financeid,
	//		fixedfee,baseunit,exceedrate,updatetime,updateuser
	//		from T_NewBilling_UserAccount
	//		`)

	//	p := []WordBook{}
	//	err := ExecQuery(&p, strSql)
	//	if err != nil {
	//		t.Error(err)
	//	} else {
	//		t.Log("执行成功")
	//		t.Log(len(p))
	//		for _, u := range p {
	//			t.Logf("%s,%s,%s", u.AccountId, u.UserId, u.FinanceId)
	//			t.Log(u.UpdateTime)
	//		}
	//	}
}

type WordBook struct {
	SeqId        int64     `db:"seqid"`
	Types        string    `db:"types"`
	Table_fields []byte    `db:"table_fields"`
	Pptr         []byte    `db:"pptr"`
	Code         []byte    `db:"code"`
	Text         []byte    `db:"text"`
	Remark       []byte    `db:"remark"`
	IsSystem     bool      `db:"issystem"`
	UpdateTime   time.Time `db:"updatetime"`
	UpdateUser   []byte    `db:"updateuser"`
}

func _Test_ExecInsertGetLastId(t *testing.T) {
	strSql := fmt.Sprintf(`select  *
		from T_NewBilling_UserAccount
		where seqId=%s`, "2")
	if exist, err := ExecCheckExists(strSql); err == nil {
		t.Log(exist)
	} else {
		t.Log(err)
	}
}

func Test_Paging(t *testing.T) {
	Init()

	var table string = "T_UserInfo"
	var fields string = `T_UserInfo.seqid,
	isnull(T_UserInfo.userid,'') as userid,
	isnull(T_UserInfo.username,'') as username,
	isnull(T_UserInfo.mobile,'') as mobile,
	isnull(T_UserInfo.usertype,0) as usertype,
	isnull(T_UserInfoExt.userappid,'') as userappid`
	//	T_UserInfo.isCharge,
	//	T_UserInfo.regTime,
	var join string = `left join T_UserInfoExt on T_UserInfoExt.userid = T_UserInfo.userid`
	var where string = `T_UserInfo.regTime > '2017-8-1'`
	var strOrder string = `T_UserInfoExt.userAppId asc,T_UserInfo.seqId asc`
	var pageSize int = 100
	var pageIndex int = 2
	var totalcount, pagecount, outPageIndex int
	result := []UserInfo1{}
	var err error
	totalcount, pagecount, outPageIndex, err = QueryByPage(&result, table, fields, where, strOrder, join, pageSize, pageIndex)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("totalcount:%d,pagecount:%d,outpageIndex:%d", totalcount, pagecount, outPageIndex)
		//t.Log(result)
		//j, err := json.Marshal(result)
		if err != nil {
			t.Error(err)
		} else {
			//	t.Log(string(j))
		}
	}
}

type UserInfo1 struct {
	SeqId     int64  `db:"seqid"`
	UserId    string `db:"userid"`
	UserName  string `db:"username"`
	Mobile    string `db:"mobile"`
	UserAppId string `db:"userappid"`
	UserType  string `db:"usertype"`
	Rownumber int64  `db:"rownumber"`
}

type Call struct {
	SeqId        int64  `db:"seqid"`
	SessionId    string `db:"sessionid"`
	CallItemId   string `db:"callitemid"`
	MemberMobile string `db:"membermobile"`
	ShowCaller   string `db:"showcaller"`
	AccountId    string `db:"accountid"`
	BillingUnit  int    `db:"billingunit"`
}
