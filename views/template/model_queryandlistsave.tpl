package models

import (
	_ "errors"
	_ "fmt"
	"strconv"
	${"time"$}

	"github.com/astaxie/beego/orm"
)

type ${uppercomponentname$} struct {
	${querystruct$}
}

type ${uppercomponentname$}ITEM struct {
	${liststruct$}
}
//获得数据条数
func Get${componentname$}count(u ${uppercomponentname$}) (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "SELECT count(1) as total  from skl_${componentname$}item_tb a  where 1=1 "	 
	${ifcondition$}
	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)

	return page, err
}

//获得分页数据
func Get${componentname$}bypageindex(u ${uppercomponentname$}) (admins []${uppercomponentname$}ITEM, err error) {
	dbtype := Getdbtype()
	admins = make([]${uppercomponentname$}ITEM, 0)
	o := orm.NewOrm()

	sql := "SELECT a.*,c.flowstatusname from skl_${componentname$}item_tb a "
	sql = sql + " inner join fi_flowstatus c on a.flowstatus=c.flowstatus "
	sql = sql + " where 1=1 "
	${ifcondition$}	 
	 
	var limitstr string = " limit "
	if dbtype == "postgres" {
		limitstr = limitstr + strconv.Itoa(u.Pagesize) + " offset " + strconv.Itoa((u.Pageindex-1)*u.Pagesize)

	} else if dbtype == "mysql" {
		limitstr = limitstr + strconv.Itoa((u.Pageindex-1)*u.Pagesize) + "," + strconv.Itoa(u.Pagesize)

	} else {
		limitstr = limitstr + strconv.Itoa((u.Pageindex-1)*u.Pagesize) + "," + strconv.Itoa(u.Pagesize)

	}
	sql = sql + limitstr
	_, err = o.Raw(ConvertSQL(sql, dbtype)).QueryRows(&admins)

	return admins, err
}
func Do${componentname$}(uarr []${uppercomponentname$}ITEM) error {
	o := orm.NewOrm()
	err := o.Begin()
	 
	//todo your business logic 
	 
	err = o.Commit()

	return err
}