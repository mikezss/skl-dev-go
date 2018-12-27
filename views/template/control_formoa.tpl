package controllers

import (
	"encoding/json"
	"${modulename$}/models"

	_ "fmt"
	"strconv"
	"github.com/astaxie/beego"
)

// Operations about Users
type ${uppercomponentname$}Controller struct {
	//beego.Controller
	BASEController
}
// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *${uppercomponentname$}Controller) Save${componentname$}() {
	var status = ""
	var fiid int
	ob := models.${uppercomponentname$}ANDITEM{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob.${firstuppercomponentname$}.Caller = ob.Userid
	fiid, err = models.AddMulti${uppercomponentname$}(ob.Opinion, ob.Userid, ob.Currentfiid, ob.Currenttiid, ob.Actionid, ob.Modualid, ob.${firstuppercomponentname$})

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status, "result": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"	
		fiidstr := strconv.Itoa(fiid)	
		ctl.Data["json"] = map[string]string{"status": status, "result": fiidstr}
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *${uppercomponentname$}Controller) Get${componentname$}() {
	var status = ""

	ob, err := models.GetAll${uppercomponentname$}()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = ob
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *${uppercomponentname$}Controller) Get${componentname$}byid() {
	var status = ""
	ob := models.${uppercomponentname$}{}

	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}
	ob2, err := models.Get${uppercomponentname$}BYID(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = ob2
		ctl.ServeJSON()
	}

}
