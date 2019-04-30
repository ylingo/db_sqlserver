package db_sqlserver

//import (
//	"bytes"
//	"fmt"
//	"strings"
//	"time"
//	//"github.com/jmoiron/sqlx"
//)

//type PARAM_DIRECT string

//const (
//	output PARAM_DIRECT = "output"
//	input  PARAM_DIRECT = "input"
//)

//type PROCEDURE_PARAM struct {
//	Name       string
//	Direct     PARAM_DIRECT
//	Value      interface{}
//	ParamPoint interface{}
//}

////该功能暂时不考虑输出参数
//func ExecProcedureQuery(dest interface{}, sqlProcedure string, params ...PROCEDURE_PARAM) error {
//	defer checkerr()
//	//	outParams := []PROCEDURE_PARAM{}
//	//	var outParamsPoint []interface{}

//	var _sb bytes.Buffer

//	_sb.WriteString("exec ")
//	_sb.WriteString(sqlProcedure)
//	_sb.WriteString(" ")
//	for i := range params {
//		if i != 0 {
//			_sb.WriteString(",")
//		}
//		//		if params[i].Direct == output {
//		//			_sb.WriteString("@")
//		//			_sb.WriteString(params[i].Name)
//		//			_sb.WriteString(" output")
//		//			outParamsPoint = append(outParamsPoint, params[i].ParamPoint)
//		//			outParams = append(outParams, params[i])
//		//		} else {
//		if params[i].Value == nil {
//			_sb.WriteString("null")
//		} else {
//			switch params[i].Value.(type) {
//			case int, int8, int16, int32, int64:
//				_sb.WriteString(fmt.Sprintf("%d", params[i].Value))
//			case float32, float64:
//				_sb.WriteString(fmt.Sprintf("%f", params[i].Value))
//			case string:
//				_sb.WriteString("\"")
//				_sb.WriteString(fmt.Sprintf("%s", strings.Replace(params[i].Value.(string), "\\", "\\\\", -1)))
//				_sb.WriteString("\"")
//			case time.Time:
//				_sb.WriteString("\"")
//				_sb.WriteString(params[i].Value.(time.Time).Format("2006-01-02 15:04:05.999"))
//				_sb.WriteString("\"")
//			case bool:
//				if params[i].Value.(bool) {
//					_sb.WriteString("true")
//				} else {
//					_sb.WriteString("false")
//				}
//			}
//		}
//		//}
//	}
//	_sb.WriteString("; ")
//	fmt.Println(_sb.String())
//	err := database.Select(dest, _sb.String())
//	if err != nil {
//		errorinfo("ExecProcedureQuery, %s ,err:%s", _sb.String(), err)
//		return err
//	}

//	//	var _sb1 bytes.Buffer
//	//	if len(outParams) > 0 {
//	//		_sb1.WriteString("select ")
//	//		for i := range outParams {
//	//			if i != 0 {
//	//				_sb1.WriteString(",")
//	//			}
//	//			_sb1.WriteString("@")
//	//			_sb1.WriteString(outParams[i].Name)
//	//			_sb1.WriteString(" as ")
//	//			_sb1.WriteString(outParams[i].Name)
//	//		}
//	//		_sb1.WriteString(";") //try append semicolon
//	//		rows1 := database.QueryRow(_sb1.String())
//	//		if err = rows1.Scan(outParamsPoint...); err != nil {
//	//			errorinfo("ExecProcedureQuery,%s %s,err:%s", sqlProcedure, _sb1.String(), err)
//	//			return err
//	//		}
//	//	}
//	loginfo("ExecProcedureQuery, %s", _sb.String())

//	return nil
//}

////该功能暂时不支持输出参数
//func ExecProcedureNoQuery(sqlProcedure string, params ...PROCEDURE_PARAM) error {
//	defer checkerr()
//	//	outParams := []PROCEDURE_PARAM{}
//	//	var outParamsPoint []interface{}

//	var _sb bytes.Buffer

//	_sb.WriteString("exec ")
//	_sb.WriteString(sqlProcedure)
//	_sb.WriteString(" ")
//	for i := range params {
//		if i != 0 {
//			_sb.WriteString(",")
//		}
//		//		if params[i].Direct == output {
//		//			_sb.WriteString("@")
//		//			_sb.WriteString(params[i].Name)
//		//			_sb.WriteString(" output")
//		//			outParamsPoint = append(outParamsPoint, params[i].ParamPoint)
//		//			outParams = append(outParams, params[i])
//		//		} else {
//		if params[i].Value == nil {
//			_sb.WriteString("null")
//		} else {
//			switch params[i].Value.(type) {
//			case int, int8, int16, int32, int64:
//				_sb.WriteString(fmt.Sprintf("%d", params[i].Value))
//			case float32, float64:
//				_sb.WriteString(fmt.Sprintf("%f", params[i].Value))
//			case string:
//				_sb.WriteString("\"")
//				_sb.WriteString(fmt.Sprintf("%s", strings.Replace(params[i].Value.(string), "\\", "\\\\", -1)))
//				_sb.WriteString("\"")
//			case time.Time:
//				_sb.WriteString("\"")
//				_sb.WriteString(params[i].Value.(time.Time).Format("2006-01-02 15:04:05.999"))
//				_sb.WriteString("\"")
//			case bool:
//				if params[i].Value.(bool) {
//					_sb.WriteString("true")
//				} else {
//					_sb.WriteString("false")
//				}
//			}
//		}
//		//}
//	}
//	_sb.WriteString("; ")
//	//fmt.Println(_sb.String())
//	_, err := database.Exec(_sb.String())
//	if err != nil {
//		errorinfo("ExecProcedureNoQuery, %s ,err:%s", _sb.String(), err)
//		return err
//	}

//	//	var _sb1 bytes.Buffer
//	//	if len(outParams) > 0 {
//	//		_sb1.WriteString("select ")
//	//		for i := range outParams {
//	//			if i != 0 {
//	//				_sb1.WriteString(",")
//	//			}
//	//			_sb1.WriteString("@")
//	//			_sb1.WriteString(outParams[i].Name)
//	//			_sb1.WriteString(" as ")
//	//			_sb1.WriteString(outParams[i].Name)
//	//		}
//	//		rows1 := database.QueryRow(_sb1.String())
//	//		if err = rows1.Scan(outParamsPoint...); err != nil {
//	//			errorinfo("ExecProcedureNoQuery,%s  %s,err:%s", sqlProcedure, _sb1.String(), err)
//	//			return err
//	//		}
//	//	}
//	loginfo("ExecProcedureNoQuery, %s", _sb.String())
//	return nil
//}

//func GetProcSql(proc, declare, in, out string, outparas ...string) string {
//	_sql := fmt.Sprintf("%v;exec %v %v", declare, proc, in)

//	var outparam string = ""
//	for _, outp := range outparas {
//		outparam = fmt.Sprintf("%v,%v=%v OUTPUT", outparam, outp, outp)
//	}
//	outparam = fmt.Sprintf("%v;", outparam)
//	if out != "" {
//		_sql = fmt.Sprintf("%v%vselect %v;", _sql, outparam, out)
//	} else {
//		_sql = fmt.Sprintf("%v%v", _sql, outparam)
//	}
//	return _sql
//}
