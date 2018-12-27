package models

import (
	_ "errors"
	"fmt"
	"strconv"
	${"time"$}

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
//{{Unescapedjs .uppercomponentname}}
type ${uppercomponentname$} struct {
	${formstruct$}
}

type ${uppercomponentname$}INDEX struct {
	Pageindex int
	Pagesize  int
	${formstruct$}	
}

func (u *${uppercomponentname$}) TableName() string {
	return "skl_${componentname$}_tb"
}


func Add${uppercomponentname$}(u *${uppercomponentname$}) error {
	o := orm.NewOrm()
	err := o.Begin()
	_, err = o.Insert(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err
}

func AddMulti${uppercomponentname$}(u []${uppercomponentname$}) error {
	o := orm.NewOrm()
	err := o.Begin()
	sql := "select count(1) as ncount from skl_${componentname$}_tb where ${pkfield$}=?"
	updatesql := "update skl_${componentname$}_tb set ${updatefieldnames$} where ${pkfield$}=?"
	
	insertsql := "insert into skl_${componentname$}_tb(${insertfieldnames$}) values(${values$})"
	
	insertsql = ConvertSQL(insertsql, Getdbtype())
	for _, u1 := range u {
		ncount := 0
		err = o.Raw(sql, u1.${firstupperpkfield$}).QueryRow(&ncount)
		if ncount > 0 {
			_, err = o.Raw(updatesql, ${updatefieldvalue$}).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		} else {
			_, err = o.Raw(insertsql, ${insertfieldvalue$}).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return err
			}
		}

	}
	err = o.Commit()
	return err
}

func GetAll${uppercomponentname$}() (admins []${uppercomponentname$}, err error) {
	admins = make([]${uppercomponentname$}, 0)
	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}
//获得数据条数
func Get${componentname$}count(u ${uppercomponentname$}) (page PAGE, err error) {

	o := orm.NewOrm()

	sql := "SELECT count(1) as total  from skl_${componentname$}_tb a  where 1=1 "	 
	${ifcondition$}
	
	err = o.Raw(ConvertSQL(sql, Getdbtype())).QueryRow(&page)

	return page, err
}
//获得分页数据
func Get${componentname$}bypageindex(u ${uppercomponentname$}INDEX) (admins []${uppercomponentname$}, err error) {
	dbtype := Getdbtype()
	admins = make([]${uppercomponentname$}, 0)
	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}_tb where 1=1 "
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

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func GetAll${uppercomponentname$}options() (admins []OPTIONS, err error) {
	admins = make([]OPTIONS, 0)
	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}_tb"

	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func Get${uppercomponentname$}(u *${uppercomponentname$}) (admins []${uppercomponentname$}, err error) {
	admins = make([]${uppercomponentname$}, 0)
	o := orm.NewOrm()
	sql := "select * from skl_${componentname$}_tb where 1=1 "
	 
	_, err = o.Raw(sql).QueryRows(&admins)

	return admins, err
}

func Delete${uppercomponentname$}(u *${uppercomponentname$}) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}

func Update${uppercomponentname$}(u *${uppercomponentname$}) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Update(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
	} else {
		err = o.Commit()
	}
	return err

}
