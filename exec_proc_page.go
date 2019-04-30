package db_sqlserver

import (
	"errors"
	//"context"
	"bytes"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

//分页查询，pageSize大于0小于等于1000，否则会自动重置为100
//pageIndex索引从1开始，如果索引大于总页数，则pageIndex=totalPageCount
//orderBy,不能有关键字 order by, 可以是: [field] asc,[field] desc
//where,不能有关键字where
//join,必须有完整的jion关系，如:left join [table1] on [table1].[f] = [table].f
func QueryByPage(dest_ptr interface{},
	table, fields, where, orderBy, join string,
	pageSize, pageIndex int) (totalCount, totalPageCount, outPageIndex int, err error) {
	totalCount = 0
	totalPageCount = 0
	outPageIndex = 0

	if table == "" {
		return 0, 0, 0, errors.New("查询的主表不能为空")
	}
	if fields == "" {
		return 0, 0, 0, errors.New("需要查询的字段不能为空")
	}
	if orderBy == "" {
		return 0, 0, 0, errors.New("分页查询必须有一个排序字段")
	}

	defer checkerr()
	table = strings.Replace(table, "'", "''", -1)
	fields = strings.Replace(fields, "'", "''", -1)
	where = strings.Replace(where, "'", "''", -1)
	orderBy = strings.Replace(orderBy, "'", "''", -1)
	join = strings.Replace(join, "'", "''", -1)

	var _sb bytes.Buffer
	_sb.WriteString("declare @totalCount int=0;")
	_sb.WriteString("declare @totalPageCount int=0;")
	_sb.WriteString("declare @pageIndex int = " + fmt.Sprintf("%d", pageIndex) + ";")
	_sb.WriteString("execute sp_paging ")
	_sb.WriteString("N'" + table + "',")
	_sb.WriteString("N'" + fields + "',")
	_sb.WriteString("N'" + orderBy + "',")
	if join == "" {
		_sb.WriteString("'',")
	} else {
		_sb.WriteString("N'" + join + "',")
	}
	if where == "" {
		_sb.WriteString("'',")
	} else {
		_sb.WriteString("N'" + where + "',")
	}
	_sb.WriteString(fmt.Sprintf("%d", pageSize) + ",")
	_sb.WriteString("@pageIndex OUTPUT,")
	_sb.WriteString("@totalCount OUTPUT,")
	_sb.WriteString("@totalPageCount OUTPUT;")
	_sb.WriteString("select @pageIndex,@totalCount,@totalPageCount;")
	procSql := _sb.String()
	//输出多个结果集，仅go1.8+支持
	rows, err := database.Query(procSql)
	defer rows.Close()
	if err != nil {
		return 0, 0, 0, err
	}
	//if rows.Next() {
	//if err = rows.Scan(dest); err != nil {
	if rows != nil {
		if err = sqlx.StructScan(rows, dest_ptr); err != nil {
			return 0, 0, 0, errors.New(err.Error())
		}
		if rows.NextResultSet() {
			if rows.Next() {
				if err = rows.Scan(&outPageIndex, &totalCount, &totalPageCount); err != nil {
					return 0, 0, 0, errors.New(err.Error())
				}
			}
		}
	}
	return totalCount, totalPageCount, outPageIndex, nil
}

//func QueryByPage(dest interface{},
//	tables, fields, where, join, primaryKey, strOrder string,
//	pageSize, pageIndex int) (totalcount, pagecount, outPageIndex int, err error) {

//	totalcount = 0
//	pagecount = 0
//	defer checkerr()

//	tables = strings.Replace(tables, "'", "''", -1)
//	fields = strings.Replace(fields, "'", "''", -1)
//	where = strings.Replace(where, "'", "''", -1)
//	join = strings.Replace(join, "'", "''", -1)
//	primaryKey = strings.Replace(primaryKey, "'", "''", -1)
//	strOrder = strings.Replace(strOrder, "'", "''", -1)

//	var _sb1 bytes.Buffer
//	_sb1.WriteString("select count(*) as totalcount from ")
//	_sb1.WriteString(tables)
//	if join != "" {
//		_sb1.WriteString("  " + join)
//	}
//	row := database.QueryRow(_sb1.String())
//	row.Scan(&totalcount)

//	//计算
//	pagecount = (int)(math.Ceil(float64(totalcount) / float64(pageSize)))

//	if pageIndex > pagecount {
//		pageIndex = pagecount
//	}
//	outPageIndex = pageIndex

//	var _sb bytes.Buffer
//	_sb.WriteString("declare @totalcount int;")
//	_sb.WriteString("execute sp_paging ")
//	_sb.WriteString("'" + tables + "'")
//	if primaryKey == "" {
//		_sb.WriteString(",''")
//	} else {
//		_sb.WriteString(",'" + primaryKey + "'")
//	}
//	if strOrder == "" {
//		_sb.WriteString(",''")
//	} else {
//		_sb.WriteString(",'" + strOrder + "'")
//	}
//	_sb.WriteString("," + fmt.Sprintf("%d", pageIndex))
//	_sb.WriteString("," + fmt.Sprintf("%d", pageSize))
//	_sb.WriteString(",'" + fields + "'")
//	if join == "" {
//		_sb.WriteString(",''")
//	} else {
//		_sb.WriteString(",' " + join + "'")
//	}
//	if where == "" {
//		_sb.WriteString(",''")
//	} else {
//		_sb.WriteString(",'" + where + "'")
//	}
//	_sb.WriteString(",'',")
//	_sb.WriteString("@totalcount output;")

//	fmt.Println(_sb.String())
//	err = database.Select(dest, _sb.String())
//	if err != nil {
//		return
//	}

//	return
//}
