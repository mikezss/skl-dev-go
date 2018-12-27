package models

import (
	_ "errors"
	"fmt"
	"html/template"
	"strconv"
	"strings"
	_ "time"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//流程定义fi_template_tb表
type COMPONENT struct {
	Componentname  string `orm:"pk;column(componentname)"`
	Parentid       string `orm:"column(parentid)"`
	Title          string `orm:"column(title)"`
	Buttons        string `orm:"column(buttons);null"`
	Style          string `orm:"column(style);null"`
	Gutter         string `orm:"column(gutter);null"`
	Colcount       string `orm:"column(colcount);null"`
	Componentlevel string `orm:"column(componentlevel);null"`
	Godirectory    string `orm:"column(godirectory);null"`
	Ngdirectory    string `orm:"column(ngdirectory);null"`
}

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//流程定义fi_template_tb表
type COMPONENTDETAIL struct {
	Componentname      string `orm:"pk;column(componentname)"`
	Seq                int    `orm:"pk;column(seq)"`
	Controlname        string `orm:"column(controlname)"`
	Controldisplayname string `orm:"column(controldisplayname)"`
	Controltype        string `orm:"column(controltype)"`
	Rows               int    `orm:"column(rows)"`
	Filetype           string `orm:"column(filetype)"`
	Ismultiple         bool   `orm:"column(ismultiple)"`
	Filesize           int    `orm:"column(filesize)"`
	Islimit            bool   `orm:"column(islimit)"`
	Limitfileqty       int    `orm:"column(limitfileqty)"`
	Minvalues          int    `orm:"column(minvalues)"`
	Maxvalues          int    `orm:"column(maxvalues)"`
	Stepvalue          int    `orm:"column(stepvalue)"`
	Icon               string `orm:"column(icon)"`
}
type COMPONENT_AND_DETAIL struct {
	Component COMPONENT
	Detail    []COMPONENTDETAIL
}

//机构变量表
func (u *COMPONENT) TableName() string {
	return "dev_component_tb"
}

//机构变量表
func (u *COMPONENTDETAIL) TableName() string {
	return "dev_componentdetail_tb"
}

// 多字段唯一键
func (u *COMPONENTDETAIL) TableUnique() [][]string {
	return [][]string{
		[]string{"Componentname", "Seq"},
	}
}
func GetComponent(u COMPONENT) (cad COMPONENT_AND_DETAIL, err error) {
	cm := COMPONENT{}
	cmdetail := make([]COMPONENTDETAIL, 0)
	o := orm.NewOrm()
	sql := "select * from dev_component_tb where componentname=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, u.Componentname).QueryRow(&cm)
	detailsql := "select * from dev_componentdetail_tb where componentname=?"
	detailsql = ConvertSQL(detailsql, Getdbtype())
	_, err = o.Raw(detailsql, u.Componentname).QueryRows(&cmdetail)
	cad = COMPONENT_AND_DETAIL{cm, cmdetail}
	return cad, err

}
func GetComponentbyparentid(u COMPONENT) (cad []COMPONENT_AND_DETAIL, err error) {
	cmarr := make([]COMPONENT, 0)

	o := orm.NewOrm()
	sql := "select * from dev_component_tb where parentid=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.Componentname).QueryRows(&cmarr)

	for _, cm := range cmarr {
		cmdetail := make([]COMPONENTDETAIL, 0)
		detailsql := "select a.* from dev_componentdetail_tb a where componentname=?"
		detailsql = ConvertSQL(detailsql, Getdbtype())
		_, err = o.Raw(detailsql, cm.Componentname).QueryRows(&cmdetail)
		cad = append(cad, COMPONENT_AND_DETAIL{cm, cmdetail})
	}

	return cad, err

}
func DeleteComponent(u COMPONENT) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	deletesql := "delete from dev_componentdetail_tb where componentname=? or componentname in (select componentname from dev_component_tb where parentid=?)"
	deletesql = ConvertSQL(deletesql, Getdbtype())
	_, err = o.Raw(deletesql, u.Componentname, u.Componentname).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	deletesql = "delete from dev_component_tb where componentname=? or parentid=?"
	deletesql = ConvertSQL(deletesql, Getdbtype())
	_, err = o.Raw(deletesql, u.Componentname, u.Componentname).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}

	err = o.Commit()
	return err
}
func GetAllComponent() (u []COMPONENT, err error) {
	o := orm.NewOrm()
	sql := "select * from dev_component_tb"
	_, err = o.Raw(sql).QueryRows(&u)
	return u, err
}
func GetAllComponentoptions() (u []OPTIONS, err error) {
	o := orm.NewOrm()
	sql := "select componentname as value,title as label from dev_component_tb"
	_, err = o.Raw(sql).QueryRows(&u)
	return u, err
}
func AddCOMPONENT(u COMPONENT, udetail []COMPONENTDETAIL) error {
	dbtype := Getdbtype()
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Delete(&u)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	deletesql := "delete from dev_componentdetail_tb where componentname=?"
	deletesql = ConvertSQL(deletesql, dbtype)
	_, err = o.Raw(deletesql, u.Componentname).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	//_, err = o.Insert(&u)
	if u.Parentid != "-1" {
		m, _ := GetComponentbyid(COMPONENT{Componentname: u.Parentid})
		componentlevel, err := strconv.Atoi(m.Componentlevel)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
		strcomponentlevel := strconv.Itoa(componentlevel + 1)
		u.Componentlevel = strcomponentlevel
	} else {
		u.Componentlevel = "0"
	}
	sql := "insert into dev_component_tb(componentname,parentid,title,buttons,style,gutter,colcount,componentlevel,godirectory,ngdirectory) values(?,?,?,?,?,?,?,?,?,?)"
	_, err = o.Raw(ConvertSQL(sql, dbtype), u.Componentname, u.Parentid, u.Title, u.Buttons, u.Style, u.Gutter, u.Colcount, u.Componentlevel, u.Godirectory, u.Ngdirectory).Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	sql = "insert into dev_componentdetail_tb(componentname,seq,controlname,controldisplayname,controltype,rows," +
		"filetype,ismultiple,filesize,islimit,limitfileqty,minvalues,maxvalues,stepvalue,icon) values(" +
		"?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	sql = ConvertSQL(sql, dbtype)
	for _, ud := range udetail {
		_, err = o.Raw(sql, u.Componentname, ud.Seq, ud.Controlname, ud.Controldisplayname, ud.Controltype, ud.Rows, ud.Filetype, ud.Ismultiple, ud.Filesize, ud.Islimit, ud.Limitfileqty, ud.Minvalues, ud.Maxvalues, ud.Stepvalue, ud.Icon).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
		}
	}
	err = o.Commit()

	return err
}
func GetComponentbyid(mt COMPONENT) (admin COMPONENT, err error) {

	o := orm.NewOrm()
	dbtype := Getdbtype()
	sql := "select * from dev_component_tb where componentname=?"
	sql = ConvertSQL(sql, dbtype)
	err = o.Raw(sql, mt.Componentname).QueryRow(&admin)

	return admin, err
}
func GetProjectpath() (admin COMPONENT, err error) {

	o := orm.NewOrm()
	dbtype := Getdbtype()
	sql := "select * from dev_component_tb where componentlevel='0'"
	sql = ConvertSQL(sql, dbtype)
	err = o.Raw(sql).QueryRow(&admin)

	return admin, err
}
func Unescaped(x string) interface{} {
	return template.HTML(x)
}
func Unescapedjs(x string) interface{} {
	return template.JS(x)
}
func UnescapedJSStr(x string) interface{} {
	return template.JSStr(x)
}
func Mod(s, m int) int {
	return s % m
}
func OutputFN(s string) string {
	return "function edit" + s + "(){"
}
func Toupper(s string) string {
	return strings.ToUpper(s)
}
func Calculate(source int, inc int) int {
	return source + inc
}
func Tofirstupper(s string) string {
	arrs := strings.Split(s, "")
	fmt.Println(arrs)
	ups := strings.ToUpper(arrs[0])
	fmt.Println(ups)
	lasts := strings.Join(arrs[1:], "")
	fmt.Println(lasts)
	return ups + lasts
}
func Tolower(s string) string {
	return strings.ToLower(s)
}
func Replace(s string, oldstr string, newstr string) string {
	return strings.Replace(s, oldstr, newstr, -1)
}
func Hastime(cadarr []COMPONENT_AND_DETAIL) bool {
	hastime := false
	for _, cad := range cadarr {
		if hastime {
			break
		}
		for _, caddtl := range cad.Detail {
			if caddtl.Controltype == "datepicker" || caddtl.Controltype == "timepicker" || caddtl.Controltype == "yearpicker" || caddtl.Controltype == "monthpicker" || caddtl.Controltype == "rangepicker" || caddtl.Controltype == "weekpicker" {
				break
				hastime = true
			}
		}

	}
	return hastime
}
func Outputifcondition(cad COMPONENTDETAIL) string {
	ifcondition := ""

	switch cad.Controltype {
	case "number":
		ifcondition = ifcondition + "if u." + cad.Controlname + " != -9999 {\n"
		ifcondition = ifcondition + "sql = sql + \" and " + cad.Controlname + "=\" + strconv.Itoa(u." + cad.Controlname + ")"
		ifcondition = ifcondition + "\n}\n"
		break
		//	case "checkbox":
		//		ifcondition = ifcondition + "if u." + cad.Controlname + " != false {\n"
		//		ifcondition = ifcondition + "sql = sql + \" and " + cad.Controlname + "=\"+u." + cad.Controlname
		//		ifcondition = ifcondition + "\n}\n"
		//	case "checkboxgroup":
		//		ifcondition = ifcondition + "if u." + cad.Controlname + " != false {\n"
		//		ifcondition = ifcondition + "sql = sql + \" and " + cad.Controlname + "=\"+u." + cad.Controlname
		//		ifcondition = ifcondition + "\n}\n"
	case "radio":
		ifcondition = ifcondition + "if u." + cad.Controlname + " != false {\n"
		ifcondition = ifcondition + "sql = sql + \" and " + cad.Controlname + "=\"+u." + cad.Controlname
		ifcondition = ifcondition + "\n}\n"
		break
	case "radiogroup":
		ifcondition = ifcondition + "if u." + cad.Controlname + " != false {\n"
		ifcondition = ifcondition + "sql = sql + \" and " + cad.Controlname + "=\"+u." + cad.Controlname
		ifcondition = ifcondition + "\n}\n"
		break
	case "datepicker":
		ifcondition = ifcondition + "if u." + cad.Controlname + ".Format(\"2006-01-02\") != \"9999-01-01\" && u." + cad.Controlname + ".Format(\"2006-01-02\") != \"0001-01-01\"{\n"
		ifcondition = ifcondition + "sql = sql + \" and DATE_FORMAT(" + Tolower(cad.Controlname) + ",'%Y-%m-%d')>='\"+u." + cad.Controlname + ".Format(\"2006-01-02\")+\"'\""
		ifcondition = ifcondition + "\n}\n"
		break
	case "timepicker":
		ifcondition = ifcondition + "if u." + cad.Controlname + ".Format(\"2006-01-02\") != \"9999-01-01\" && u." + cad.Controlname + ".Format(\"2006-01-02\") != \"0001-01-01\"{\n"
		ifcondition = ifcondition + "sql = sql + \" and DATE_FORMAT(" + Tolower(cad.Controlname) + ",'%Y-%m-%d')>='\"+u." + cad.Controlname + ".Format(\"2006-01-02\")+\"'\""
		ifcondition = ifcondition + "\n}\n"
		break
	default:
		ifcondition = ifcondition + "if u." + cad.Controlname + "!=\"\"{\n"
		ifcondition = ifcondition + " sql=sql+ \" and " + Tolower(cad.Controlname) + "='\"+" + "u." + cad.Controlname + "+\"'\""
		ifcondition = ifcondition + "\n}\n"
	}
	return ifcondition
}
func Outputdatatype(cd COMPONENTDETAIL) string {
	//
	switch cd.Controltype {
	case "number":
		if cd.Controlname == "Amount" {
			return " float64 "
		} else {
			return " int "
		}
	case "datepicker":
		return " time.Time "
	case "timepicker":
		return " time.Time "
	default:
		return " string "
	}
	return " string "
}
func Outputfieldtype(cd COMPONENTDETAIL) string {
	//
	switch cd.Controltype {
	case "number":
		if cd.Controlname == "Amount" {
			return " float64 "
		} else {
			return " int "
		}
	case "datepicker":
		return " ;type(date) "
	case "timepicker":
		return " ;type(datetime) "
	default:
		return " string "
	}
	return ""
}
func Outputcontrol(cd COMPONENTDETAIL) string {
	//{'Controlname': 'Planid', 'Controltype': 'textbox'}
	switch cd.Controltype {
	case "textarea":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'textarea','rows':" + strconv.Itoa(cd.Rows) + "}"
	case "checkbox":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'checkbox'}"
	case "checkboxgroup":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'checkboxgroup','checkboxgroup': [],'datasource':''}"
	case "radio":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'radio'}"
	case "radiogroup":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'radiogroup','radiogroup': [],'datasource':''}"
	case "number":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'number'" + ", 'minvalue': " + strconv.Itoa(cd.Minvalues) + ", 'maxvalue':" + strconv.Itoa(cd.Maxvalues) + ",'stepvalue': " + strconv.Itoa(cd.Stepvalue) + "}"
	case "select":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'select','options': [], 'nzMode': 'default','datasource':''}"
	case "dataselect":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'select','options': [], 'nzMode': 'default'}"
	case "treeselect":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'treeselect', 'nodes': []}"
	case "linkAction":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'linkAction'}"
	case "datepicker":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'datepicker'}"
	case "timepicker":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'timepicker'}"
	case "atcomplete":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'autocomplete','datasource':[]}"
	case "label":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'label'}"
	case "upload":
		//",'filetype':" + cd.Filetype +
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'upload'" + ",'ismultiple':" + Bool2string(cd.Ismultiple) + ",'limit':" + Bool2string(cd.Islimit) + ",'limitfileqty':" + strconv.Itoa(cd.Limitfileqty) + ",'filesize':" + strconv.Itoa(cd.Filesize) + ",'fileList':[]}"
	case "routerLink":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'routerLink'}"
	case "icon":
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'icon'}"
	default:
		return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'textbox'}"
	}
	return "{'Controlname': '" + cd.Controlname + "', 'Controltype': 'textbox'}"
}
func Bool2string(bvalue bool) string {
	if bvalue {
		return "true"
	}
	return "false"
}
