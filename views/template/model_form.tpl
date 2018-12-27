package models

import (
	_ "errors"
	"fmt"
	_ "strconv"
	${"time"$}

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
type ${uppercomponentname$} struct {
	${formstruct$}
}

func (u *${uppercomponentname$}) TableName() string {
	return "skl_${componentname$}_tb"
}


func AddMulti${uppercomponentname$}(u ${uppercomponentname$}) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	 
	_, err = o.Delete(&u)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	_, err = o.Insert(&u)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
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
func Get${uppercomponentname$}BYID(e ${uppercomponentname$}) (admin ${uppercomponentname$}, err error) {

	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}_tb where ${pkfield$}=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.${firstupperpkfield$}).QueryRow(&admin)

	return admin, err
}
func Delete${uppercomponentname$}(u *${uppercomponentname$}) error {

	o := orm.NewOrm()

	err := o.Begin()
	_, err = o.Delete(u)

	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	
	err = o.Commit()
	 
	return err
}
