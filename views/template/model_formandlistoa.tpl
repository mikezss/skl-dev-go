package models

import (
	_ "errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//Adminid    int64 `orm:"pk;auto"` //主键，自动增长
//Remark         string `orm:"size(5000)"`
//Created         time.Time `orm:"index"`
type ${uppercomponentname$} struct {
	Flowinstid int       `orm:"pk;column(flowinstid)"`
	Caller     string    `orm:"column(caller)"`
	Calltime   time.Time `orm:"column(calltime)"`
	Flowstatus string    `orm:"column(flowstatus)"`
	${formstruct$}
}
type ${uppercomponentname$}ITEM struct {
	Id         int     `orm:"auto;column(id)"`
	Flowinstid int     `orm:"column(flowinstid)"`
	${formliststruct$}
}
type ${uppercomponentname$}ANDITEM struct {
	Userid      string
	Currentfiid int
	Currenttiid int
	Actionid    string
	Opinion     string
	Modualid    string	 
	${firstuppercomponentname$}     ${uppercomponentname$}
	${firstuppercomponentname$}item []${uppercomponentname$}ITEM
}

func (u *${uppercomponentname$}) TableName() string {
	return "skl_${componentname$}_tb"
}
func (u *${uppercomponentname$}ITEM) TableName() string {
	return "skl_${componentname$}item_tb"
}

func AddMulti${uppercomponentname$}(opinion string, submitter string, currentfiid int, currenttiid int, actionid string, modualid string, u ${uppercomponentname$}, u2 []${uppercomponentname$}ITEM) (fiid int, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	isend := false
	var amount float64
	for _, u3 := range u2 {
		amount = amount + u3.Amount
	} 
	mdny := MODUALCNTANDMNY{Submitter: submitter, Content: u.Content, Amount: amount, Opinion: opinion}
	m := make([]map[string]string, 0)
	mm := make(map[string]string)
	mm["money"] = strconv.FormatFloat(amount, 'f', -1, 64)
	m = append(m, mm)
	fiid, err,isend = Doflow(o, modualid, currentfiid, currenttiid, actionid, mdny,m)
	if err != nil {
		fmt.Println(err)
		o.Rollback()
		return -1, err
	}
	if actionid == "save" || actionid == "submit" {
		u.Flowinstid = fiid
		if isend {
			u.Flowstatus = "1"
		} else {
			u.Flowstatus = "0"
		}
		u.Calltime = time.Now()
		_, err = o.Delete(&u)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return -1, err
		}
		_, err = o.Insert(&u)
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return -1, err
		}
		for i, _ := range u2 {
			u2[i].Flowinstid = fiid
		}
		sql := "delete from  skl_${componentname$}item_tb where flowinstid=?"
		sql = ConvertSQL(sql, Getdbtype())
		_, err = o.Raw(sql, fiid).Exec()
		if err != nil {
			fmt.Println(err)
			o.Rollback()
			return -1, err
		}
		//_, err = o.InsertMulti(len(u2), &u2)
		sql = "insert into skl_${componentname$}item_tb(flowinstid,${fieldnames$}) values(?,${values$})"
		sql = ConvertSQL(sql, Getdbtype())
		for _, u3 := range u2 {
			_, err = o.Raw(sql, fiid, ${fieldnames2$}).Exec()
			if err != nil {
				fmt.Println(err)
				o.Rollback()
				return -1, err
			}
		}
	}
	err = o.Commit()

	return fiid,err
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

	sql := "select * from skl_${componentname$}_tb where flowinstid=?"
	sql = ConvertSQL(sql, Getdbtype())
	err = o.Raw(sql, e.Flowinstid).QueryRow(&admin)

	return admin, err
}
func GetAll${uppercomponentname$}ITEM(e ${uppercomponentname$}) (admins []${uppercomponentname$}ITEM, err error) {
	admins = make([]${uppercomponentname$}ITEM, 0)
	o := orm.NewOrm()

	sql := "select * from skl_${componentname$}item_tb where flowinstid=?"
	sql = ConvertSQL(sql, Getdbtype())
	_, err = o.Raw(sql, e.Flowinstid).QueryRows(&admins)

	return admins, err
}
