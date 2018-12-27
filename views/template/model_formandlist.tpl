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
type ${uppercomponentname$}ITEM struct {
	${formliststruct$}
}
type ${uppercomponentname$}ANDITEM struct {	 
	${firstuppercomponentname$}     ${uppercomponentname$}
	${firstuppercomponentname$}item []${uppercomponentname$}ITEM
}

func (u *${uppercomponentname$}) TableName() string {
	return "skl_${componentname$}_tb"
}
func (u *${uppercomponentname$}ITEM) TableName() string {
	return "skl_${componentname$}item_tb"
}

func AddMulti${uppercomponentname$}(u ${uppercomponentname$}, u2 []${uppercomponentname$}ITEM) (err error) {
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
		 
	sql := "delete from  skl_${componentname$}item_tb"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return err
	}
	//_, err = o.InsertMulti(len(u2), &u2)
	sql = "insert into skl_${componentname$}item_tb(${fieldnames$}) values(${values$})"
	sql = ConvertSQL(sql, Getdbtype())
	for _, u3 := range u2 {
		_, err = o.Raw(sql,${fieldnames2$}).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return err
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
func Get${uppercomponentname$}BYID(e ${uppercomponentname$}) (admin ${uppercomponentname$}, err error) {

	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}_tb where ${pkfield$}=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.${firstupperpkfield$}).QueryRow(&admin)

	return admin, err
}
func GetAll${uppercomponentname$}ITEM(e ${uppercomponentname$}) (admins []${uppercomponentname$}ITEM, err error) {
	admins = make([]${uppercomponentname$}ITEM, 0)
	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}item_tb where ${pkfield$}=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, e.${firstupperpkfield$}).QueryRows(&admins)

	return admins, err
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
	sql := "delete from skl_${componentname$}item_tb where ${pkfield$}=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, u.${firstupperpkfield$}).Exec()
	err = o.Commit()
	 
	return err
}
func Get${uppercomponentname$}(e ${uppercomponentname$}) (admin ${uppercomponentname$}, err error) {

	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}_tb where ${pkfield$}=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.${firstupperpkfield$}).QueryRow(&admin)

	return admin, err
}
