package controllers

import (
	"encoding/json"
	"github.com/mikezss/skl-dev-go/models"

	_ "bufio"
	_ "errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Users
type DEVController struct {
	//beego.Controller
	BASEController
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *DEVController) Savecomponent() {
	var status = ""
	ob := models.COMPONENT_AND_DETAIL{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.AddCOMPONENT(ob.Component, ob.Detail)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *DEVController) Getcomponent() {
	var status = ""
	ob := models.COMPONENT{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	cp, err := models.GetComponent(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = cp
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *DEVController) Deletecomponent() {
	var status = ""
	ob := models.COMPONENT{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteComponent(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *DEVController) Getallcomponent() {
	var status = ""

	cp, err := models.GetAllComponent()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = cp
		ctl.ServeJSON()
	}
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *DEVController) Getallcomponentoptions() {
	var status = ""

	cp, err := models.GetAllComponentoptions()

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = cp
		ctl.ServeJSON()
	}
}
func (ctl *DEVController) Getcomponenttreejson() {

	coompenttreeson := models.CreateComponentTreeJson()

	ctl.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	content := []byte(stringsToJson(coompenttreeson))
	ctl.Ctx.Output.Body(content)
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (ctl *DEVController) Createcomponent() {
	var status = ""
	ob := models.COMPONENT{}
	beego.Debug(string(ctl.Ctx.Input.RequestBody))
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error(err)
	}

	err = ctl.createModuleOrComponent(ob)

	if err != nil {

		status = "false"
		ctl.Data["json"] = map[string]string{"status": err.Error()}
		ctl.ServeJSON()

	} else {
		status = "ok"
		ctl.Data["json"] = map[string]string{"status": status}
		ctl.ServeJSON()
	}
}
func (this *DEVController) registerrouter(uppercomponentname string, gopath string) {
	filePath := gopath + "/routers/router.go"
	includedstr := "beego.AutoRouter(&controllers." + uppercomponentname + "Controller{})"
	if this.IsFileIncludestring(filePath, includedstr) {
		return
	}
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	_, err = f.Seek(-2, 2)
	if err != nil {
		fmt.Println("Seek")
		fmt.Println(err)
	}
	// 获取文件指针当前位置
	cur_offset, _ := f.Seek(0, os.SEEK_CUR)
	fmt.Printf("current offset is %d\n", cur_offset)
	_, err = f.WriteString("    beego.AutoRouter(&controllers." + uppercomponentname + "Controller{})\n")
	_, err = f.WriteString("}\n")
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
	}

	defer f.Close()

}
func (this *DEVController) registermodel(uppercomponentname string, gopath string) {
	filePath := gopath + "/models/init.go"
	includedstr := "orm.RegisterModel(new(" + uppercomponentname + "))"
	if this.IsFileIncludestring(filePath, includedstr) {
		return
	}

	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)

	_, err = f.Seek(-2, 2)
	if err != nil {
		fmt.Println("Seek")
		fmt.Println(err)
	}
	// 获取文件指针当前位置
	cur_offset, _ := f.Seek(0, os.SEEK_CUR)
	fmt.Printf("current offset is %d\n", cur_offset)
	_, err = f.WriteString("    orm.RegisterModel(new(" + uppercomponentname + "))\n")
	_, err = f.WriteString("}\n")
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
	}

	defer f.Close()

}
func (this *DEVController) IsFileIncludestring(filePath string, includedstr string) bool {

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	s := string(b)
	return strings.Contains(s, includedstr)
}

//stype：views，controllers，models
//componentname
//componentnamemap
//gopath :D:\goproject\src\skl-api
//tplnames :*.tpl
func (this *DEVController) createMVC(stype string, componentname string, componentnamemap map[string]string, gopath string, tplnames string) {
	var filepaths, filename string
	filepaths = gopath + "/" + stype + "/"

	filename = componentname + "model.go"
	this.Data["firstuppercomponentname"] = models.Tofirstupper(componentname)
	this.Data["componentname"] = componentname
	this.Data["uppercomponentname"] = strings.ToUpper(componentname)
	for key, value := range componentnamemap {
		cadarr, _ := models.GetComponentbyparentid(models.COMPONENT{Componentname: value})
		if tplnames == "template/singletablelistmodel.tpl" {
			this.Data[key] = cadarr[0]
		} else {
			this.Data[key] = cadarr
		}

	}
	this.TplName = tplnames
	bt, err := this.RenderBytes()
	if err != nil {
		fmt.Println("RenderBytes")
		fmt.Println(err)
	}

	f, err := os.Create(filepaths + filename)
	if err != nil {
		fmt.Println("Create")
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.Write(bt)
	if err != nil {
		fmt.Println("Write")
		fmt.Println(err)
	}

}
func (this *DEVController) createModuleOrComponent(mt models.COMPONENT) error {

	commonctrl := COMMONController{}
	currentDirectory := commonctrl.GetCurrentDirectory()

	project, _ := models.GetProjectpath()
	ngpath := project.Ngdirectory
	gopath := project.Godirectory
	//create NG module and NG service
	if mt.Componentlevel == "1" {
		modulefilename := ngpath + "/src/app/" + mt.Componentname + "/" + mt.Componentname + ".module.ts"
		if commonctrl.IsFDexists(modulefilename) {
			//return errors.New(modulefilename + " has exists!")
		}
		//createcomponent.bat D:\angular\skl D:\goproject\src\skl-api module testmodule3
		err := this.Execmd(currentDirectory+"/createcomponent.bat", ngpath, gopath, "module", mt.Componentname)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = this.insertmodule(modulefilename, mt.Componentname)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = commonctrl.Replacefilecontent(modulefilename, "import { CommonModule } from '@angular/common';", "")
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = commonctrl.Replacefilecontent(modulefilename, `\s{1,4}CommonModule`, "    SklCoreModule, SklCommonModule")
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = commonctrl.Replacefilecontent(modulefilename, `declarations: \[\]`, "declarations: [],\n  providers: ["+models.Tofirstupper(mt.Componentname)+"Service]")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	if mt.Componentlevel == "2" {
		componentdirname := ngpath + "/src/app/" + mt.Parentid + "/" + mt.Componentname
		if commonctrl.IsFDexists(componentdirname) {
			//return errors.New(componentdirname + " has exists!")
		} else {
			err := this.Execmd(currentDirectory+"/createcomponent.bat", ngpath, gopath, "component", mt.Parentid, mt.Componentname)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		//this.createMVC("models", mt.Componentname, map[string]string{"fieldlist": mt.Componentname}, gopath, "template/singletablelistmodel.tpl")
		cadarr, err := models.GetComponentbyparentid(models.COMPONENT{Componentname: mt.Componentname})
		if err != nil {
			fmt.Println(err)
			return err
		}
		this.createmodel(mt.Componentname, gopath, currentDirectory)
		this.createcontrol(mt.Componentname, gopath, currentDirectory)
		this.registerrouter(strings.ToUpper(mt.Componentname), gopath)
		if len(cadarr) == 1 {
			style1 := cadarr[0].Component.Style
			switch style1 {
			case "form":
				this.registermodel(strings.ToUpper(mt.Componentname), gopath)
			case "oaform":
				this.registermodel(strings.ToUpper(mt.Componentname), gopath)
			case "singletablelist":
				this.registermodel(strings.ToUpper(mt.Componentname), gopath)
			case "listandsave":
				this.registermodel(strings.ToUpper(mt.Componentname), gopath)
			}

		} else {
			style1 := cadarr[0].Component.Style
			style2 := cadarr[1].Component.Style
			if style1+style2 == "formformlist" || style1+style2 == "formlistform" {
				this.registermodel(strings.ToUpper(mt.Componentname), gopath)
				this.registermodel(strings.ToUpper(mt.Componentname)+"ITEM", gopath)
			}
			if style1+style2 == "oaformformlist" || style1+style2 == "formlistoaform" {
				this.registermodel(strings.ToUpper(mt.Componentname), gopath)
				this.registermodel(strings.ToUpper(mt.Componentname)+"ITEM", gopath)
			}
			if style1+style2 == "querysingletablelist" || style1+style2 == "singletablelistquery" {
				this.registermodel(strings.ToUpper(mt.Componentname), gopath)
			}
		}

		this.updateNGservice(mt.Parentid, mt.Componentname, ngpath, gopath, currentDirectory)
		this.updateNGrouter(mt.Parentid, mt.Componentname, ngpath)
		this.updateNGcomponent(mt.Parentid, mt.Componentname, ngpath, gopath, currentDirectory)
		this.updateNGcomponenthtml(mt.Parentid, mt.Componentname, ngpath, gopath, currentDirectory)
	}
	return nil
}
func (this *DEVController) createmodel(componentname string, gopath string, currentDirectory string) error {

	tmpfilename := ""
	cadarr, err := models.GetComponentbyparentid(models.COMPONENT{Componentname: componentname})
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(cadarr) == 1 {
		if cadarr[0].Component.Style == "form" {
			tmpfilename = "model_form.tpl"
		}
		if cadarr[0].Component.Style == "oaform" {
			tmpfilename = "model_formoa.tpl"
		}
		if cadarr[0].Component.Style == "singletablelist" {
			tmpfilename = "model_sgltblist.tpl"
		}
	} else {
		style1 := cadarr[0].Component.Style
		style2 := cadarr[1].Component.Style
		if style1+style2 == "querylist" || style1+style2 == "listquery" {
			tmpfilename = "model_queryandlist.tpl"
		}
		if style1+style2 == "querylistandsave" || style1+style2 == "listandsavequery" {
			tmpfilename = "model_queryandlistsave.tpl"
		}
		if style1+style2 == "formformlist" || style1+style2 == "formlistform" {
			tmpfilename = "model_formandlist.tpl"
		}
		if style1+style2 == "oaformformlist" || style1+style2 == "formlistoaform" {
			tmpfilename = "model_formandlistoa.tpl"
		}
		if style1+style2 == "querysingletablelist" || style1+style2 == "singletablelistquery" {
			tmpfilename = "model_querysgltblist.tpl"
		}
	}
	commonctrl := COMMONController{}
	f, err2 := os.OpenFile(gopath+"/models/"+componentname+"model.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	filecontent, err := commonctrl.Readfile2string(currentDirectory+"/views/template/"+tmpfilename, "utf8")
	if err != nil {
		fmt.Println(err)
		return err
	}
	hastime := ""
	if models.Hastime(cadarr) {
		hastime = "\"time\""
	}
	reg := regexp.MustCompile(`\$\{\"time\"\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, hastime)

	reg = regexp.MustCompile(`\$\{firstuppercomponentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(componentname))

	reg = regexp.MustCompile(`\$\{componentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, componentname)

	reg = regexp.MustCompile(`\$\{uppercomponentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Toupper(componentname))
	//form:${formstruct$} ${pkfield$}
	//singletablelist:${formstruct$}  ${fieldnames$} ${values$} ${fieldnames2$}
	//formandlist:${formstruct$} ${formliststruct$} ${fieldnames$}  ${values$} ${fieldnames2$} ${pkfield$}
	//queryandlist:${querystruct$} ${liststruct$} ${ifcondition$}

	//Postid   string `orm:"pk;column(postid)"`
	if len(cadarr) == 1 {
		if cadarr[0].Component.Style == "form" {
			formstuct := ""
			pkfield := ""
			spacestring := ""
			for idx, cad := range cadarr[0].Detail {
				if idx > 0 {
					spacestring = "    "
				}
				formstuct = formstuct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""
				if idx == 0 {
					formstuct = formstuct + "pk;"
					pkfield = cad.Controlname
				}
				formstuct = formstuct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
			}
			reg = regexp.MustCompile(`\$\{formstruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formstuct)
			reg = regexp.MustCompile(`\$\{pkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tolower(pkfield))
			reg = regexp.MustCompile(`\$\{firstupperpkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(pkfield))
		}
		if cadarr[0].Component.Style == "oaform" {
			formstuct := ""
			spacestring := ""
			for idx, cad := range cadarr[0].Detail {
				if idx == 0 {
					spacestring = ""
				} else {
					spacestring = "    "
				}
				formstuct = formstuct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""

				formstuct = formstuct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
			}
			reg = regexp.MustCompile(`\$\{formstruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formstuct)
		}
		if cadarr[0].Component.Style == "singletablelist" {
			pkfield := ""
			formstuct := ""
			insertfieldnames := ""
			updatefieldnames := ""
			values := ""
			insertfieldvalue := ""
			updatefieldvalue := ""
			spacestring := ""
			for idx, cad := range cadarr[0].Detail {
				if idx > 0 {
					spacestring = "    "
				}
				formstuct = formstuct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""
				if idx == 0 {
					formstuct = formstuct + "pk;"
					pkfield = cad.Controlname
				}
				formstuct = formstuct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
				insertfieldnames = insertfieldnames + models.Tolower(cad.Controlname)
				if idx > 0 {
					updatefieldnames = updatefieldnames + models.Tolower(cad.Controlname) + "=?"
				}
				if idx > 0 && idx < len(cadarr[0].Detail)-1 {
					updatefieldnames = updatefieldnames + ","
				}
				if idx < len(cadarr[0].Detail)-1 {
					insertfieldnames = insertfieldnames + ","
				}
				values = values + "?"
				if idx < len(cadarr[0].Detail)-1 {
					values = values + ","
				}
				insertfieldvalue = insertfieldvalue + "u1." + cad.Controlname
				if idx < len(cadarr[0].Detail)-1 {
					insertfieldvalue = insertfieldvalue + ","
				}
				if idx > 0 {
					updatefieldvalue = updatefieldvalue + "u1." + cad.Controlname + ","
				}
				if idx == len(cadarr[0].Detail)-1 {
					updatefieldvalue = updatefieldvalue + "u1." + models.Tofirstupper(pkfield)
				}
			}
			reg = regexp.MustCompile(`\$\{formstruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formstuct)
			reg = regexp.MustCompile(`\$\{insertfieldnames\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, insertfieldnames)
			reg = regexp.MustCompile(`\$\{updatefieldnames\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, updatefieldnames)
			reg = regexp.MustCompile(`\$\{values\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, values)
			reg = regexp.MustCompile(`\$\{insertfieldvalue\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, insertfieldvalue)
			reg = regexp.MustCompile(`\$\{updatefieldvalue\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, updatefieldvalue)
			reg = regexp.MustCompile(`\$\{pkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tolower(pkfield))
			reg = regexp.MustCompile(`\$\{firstupperpkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(pkfield))
		}
	} else {
		style1 := cadarr[0].Component.Style
		style2 := cadarr[1].Component.Style
		x := 0
		y := 1
		if style1+style2 == "querysingletablelist" || style1+style2 == "singletablelistquery" {
			if style1 == "query" {
				x = 0
				y = 1
			} else {
				x = 1
				y = 0
			}
			querystruct := ""
			ifcondition := ""
			pkfield := ""
			formstuct := ""
			insertfieldnames := ""
			updatefieldnames := ""
			values := ""
			insertfieldvalue := ""
			updatefieldvalue := ""
			spacestring := ""
			for idx, cad := range cadarr[x].Detail {
				if idx > 0 {
					spacestring = "    "
				}
				querystruct = querystruct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "\n"
				ifcondition = ifcondition + models.Outputifcondition(cad)
			}
			for idx, cad := range cadarr[y].Detail {
				if idx > 0 {
					spacestring = "    "
				} else {
					spacestring = ""
				}
				formstuct = formstuct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""
				if idx == 0 {
					formstuct = formstuct + "pk;"
					pkfield = cad.Controlname
				}
				formstuct = formstuct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
				insertfieldnames = insertfieldnames + models.Tolower(cad.Controlname)
				if idx > 0 {
					updatefieldnames = updatefieldnames + models.Tolower(cad.Controlname) + "=?"
				}
				if idx > 0 && idx < len(cadarr[y].Detail)-1 {
					updatefieldnames = updatefieldnames + ","
				}
				if idx < len(cadarr[y].Detail)-1 {
					insertfieldnames = insertfieldnames + ","
				}
				values = values + "?"
				if idx < len(cadarr[y].Detail)-1 {
					values = values + ","
				}
				insertfieldvalue = insertfieldvalue + "u1." + cad.Controlname
				if idx < len(cadarr[y].Detail)-1 {
					insertfieldvalue = insertfieldvalue + ","
				}
				if idx > 0 {
					updatefieldvalue = updatefieldvalue + "u1." + cad.Controlname + ","
				}
				if idx == len(cadarr[y].Detail)-1 {
					updatefieldvalue = updatefieldvalue + "u1." + models.Tofirstupper(pkfield)
				}
			}
			reg = regexp.MustCompile(`\$\{querystruct\$\}`)
			querystruct = querystruct + "    Pageindex int\n"
			querystruct = querystruct + "    Pagesize int\n"
			filecontent = reg.ReplaceAllString(filecontent, querystruct)
			reg = regexp.MustCompile(`\$\{ifcondition\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, ifcondition)

			reg = regexp.MustCompile(`\$\{formstruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formstuct)
			reg = regexp.MustCompile(`\$\{insertfieldnames\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, insertfieldnames)
			reg = regexp.MustCompile(`\$\{updatefieldnames\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, updatefieldnames)
			reg = regexp.MustCompile(`\$\{values\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, values)
			reg = regexp.MustCompile(`\$\{insertfieldvalue\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, insertfieldvalue)
			reg = regexp.MustCompile(`\$\{updatefieldvalue\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, updatefieldvalue)
			reg = regexp.MustCompile(`\$\{pkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tolower(pkfield))
			reg = regexp.MustCompile(`\$\{firstupperpkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(pkfield))

		}
		if style1+style2 == "querylist" || style1+style2 == "listquery" || style1+style2 == "querylistandsave" || style1+style2 == "listandsavequery" {
			//queryandlist:${querystruct$} ${liststruct$} ${ifcondition$}
			if style1 == "query" {
				x = 0
				y = 1
			} else {
				x = 1
				y = 0
			}
			querystruct := ""
			liststruct := ""
			ifcondition := ""
			spacestring := ""
			for idx, cad := range cadarr[x].Detail {
				if idx > 0 {
					spacestring = "    "
				}
				querystruct = querystruct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "\n"
				ifcondition = ifcondition + models.Outputifcondition(cad)
			}
			for idx, cad := range cadarr[y].Detail {
				if idx > 0 {
					spacestring = "    "
				} else {
					spacestring = ""
				}
				liststruct = liststruct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "\n"
			}
			reg = regexp.MustCompile(`\$\{querystruct\$\}`)
			querystruct = querystruct + "    Pageindex int\n"
			querystruct = querystruct + "    Pagesize int\n"
			filecontent = reg.ReplaceAllString(filecontent, querystruct)
			reg = regexp.MustCompile(`\$\{liststruct\$\}`)
			liststruct = liststruct + "    Pageindex int\n"
			liststruct = liststruct + "    Pagesize int\n"
			filecontent = reg.ReplaceAllString(filecontent, liststruct)
			reg = regexp.MustCompile(`\$\{ifcondition\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, ifcondition)

		}
		if style1+style2 == "formformlist" || style1+style2 == "formlistform" {
			//formandlist:${formstruct$} ${formliststruct$} ${fieldnames$}  ${values$} ${fieldnames2$} ${pkfield$}
			if style1 == "form" {
				x = 0
				y = 1
			} else {
				x = 1
				y = 0
			}
			formstruct := ""
			formliststruct := ""
			fieldnames := ""
			values := ""
			fieldnames2 := ""
			pkfield := ""
			spacestring := ""
			for idx, cad := range cadarr[x].Detail {
				if idx > 0 {
					spacestring = "    "
				}
				formstruct = formstruct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""
				if idx == 0 {
					formstruct = formstruct + "pk;"
				}
				formstruct = formstruct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
				if idx == 0 {
					pkfield = cad.Controlname
				}
			}
			for idx, cad := range cadarr[y].Detail {
				if idx > 0 {
					spacestring = "    "
				} else {
					spacestring = ""
				}
				formliststruct = formliststruct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""
				if idx == 0 {
					formliststruct = formliststruct + "pk;"
				}
				formliststruct = formliststruct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
				fieldnames = fieldnames + models.Tolower(cad.Controlname)
				if idx < len(cadarr[y].Detail)-1 {
					fieldnames = fieldnames + ","
				}
				values = values + "?"
				if idx < len(cadarr[y].Detail)-1 {
					values = values + ","
				}
				fieldnames2 = fieldnames2 + "u3." + cad.Controlname
				if idx < len(cadarr[y].Detail)-1 {
					fieldnames2 = fieldnames2 + ","
				}
			}
			reg = regexp.MustCompile(`\$\{formstruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formstruct)
			reg = regexp.MustCompile(`\$\{formliststruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formliststruct)
			reg = regexp.MustCompile(`\$\{fieldnames\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, fieldnames)
			reg = regexp.MustCompile(`\$\{values\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, values)
			reg = regexp.MustCompile(`\$\{fieldnames2\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, fieldnames2)
			reg = regexp.MustCompile(`\$\{pkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tolower(pkfield))
			reg = regexp.MustCompile(`\$\{firstupperpkfield\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(pkfield))
		}
		if style1+style2 == "oaformformlist" || style1+style2 == "formlistoaform" {
			//formandlist:${formstruct$} ${formliststruct$} ${fieldnames$}  ${values$} ${fieldnames2$} ${pkfield$}
			if style1 == "oaform" {
				x = 0
				y = 1
			} else {
				x = 1
				y = 0
			}
			formstruct := ""
			formliststruct := ""
			fieldnames := ""
			values := ""
			fieldnames2 := ""
			spacestring := "    "
			for idx, cad := range cadarr[x].Detail {
				if idx == 0 {
					spacestring = ""
				} else {
					spacestring = "    "
				}
				formstruct = formstruct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""
				formstruct = formstruct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
			}
			for idx, cad := range cadarr[y].Detail {
				if idx == 0 {
					spacestring = ""
				} else {
					spacestring = "    "
				}
				formliststruct = formliststruct + spacestring + cad.Controlname + models.Outputdatatype(cad) + "`orm:\""
				formliststruct = formliststruct + "column(" + models.Tolower(cad.Controlname) + ")\"`\n"
				fieldnames = fieldnames + models.Tolower(cad.Controlname)
				if idx < len(cadarr[y].Detail)-1 {
					fieldnames = fieldnames + ","
				}
				values = values + "?"
				if idx < len(cadarr[y].Detail)-1 {
					values = values + ","
				}
				fieldnames2 = fieldnames2 + "u3." + cad.Controlname
				if idx < len(cadarr[y].Detail)-1 {
					fieldnames2 = fieldnames2 + ","
				}
			}
			reg = regexp.MustCompile(`\$\{formstruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formstruct)
			reg = regexp.MustCompile(`\$\{formliststruct\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, formliststruct)
			reg = regexp.MustCompile(`\$\{fieldnames\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, fieldnames)
			reg = regexp.MustCompile(`\$\{values\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, values)
			reg = regexp.MustCompile(`\$\{fieldnames2\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, fieldnames2)
		}
	}

	_, err = f.WriteString(filecontent)
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}

	return nil
}

//style:form,formlist,singletablelist,query,list,listandsave
func (this *DEVController) updateNGcomponenthtml(modulename string, componentname string, ngpath string, gopath string, currentDirectory string) error {
	cadarr, err := models.GetComponentbyparentid(models.COMPONENT{Componentname: componentname})
	if err != nil {
		fmt.Println(err)
		return err
	}
	commonctrl := COMMONController{}
	componenthtmlfilepath := ngpath + "/src/app/" + modulename + "/" + componentname + "/" + componentname + ".component.html"
	f, err3 := os.OpenFile(componenthtmlfilepath, os.O_RDWR|os.O_TRUNC, 0644)
	defer f.Close()
	if err3 != nil {
		fmt.Println(err3)
		return err3
	}
	for _, cad := range cadarr {
		htmltplfilename := ""
		switch cad.Component.Style {
		case "oaform":
			htmltplfilename = "html_sklformoa.tpl"
		case "form":
			htmltplfilename = "html_sklform.tpl"
		case "formlist":
			htmltplfilename = "html_sklformlist.tpl"
		case "singletablelist":
			htmltplfilename = "html_sgltblist.tpl"
		case "query":
			htmltplfilename = "html_sklquery.tpl"
		case "list":
			htmltplfilename = "html_sklquerylist.tpl"
		case "listandsave":
			htmltplfilename = "html_sklquerylistsave.tpl"
		}

		htmltplfilepath := currentDirectory + "/views/template/" + htmltplfilename
		filecontent, err := commonctrl.Readfile2string(htmltplfilepath, "utf8")
		if err != nil {
			fmt.Println(err)
			return err
		}
		buttons := "["
		barry := strings.Split(cad.Component.Buttons, ",")

		for idx, cdtl := range barry {
			//{'Controlname': 'Postid', 'Controltype': 'textbox'},

			buttons = buttons + "'" + cdtl + "'"
			if idx < len(barry)-1 {
				buttons = buttons + ","
			}

		}
		buttons = buttons + "]"

		switch cad.Component.Style {
		case "oaform":
			reg := regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Componentname)
			reg = regexp.MustCompile(`\$\{gutter\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Gutter)
			reg = regexp.MustCompile(`\$\{colcount\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Colcount)
			reg = regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Componentname)

		case "form":
			reg := regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Componentname)
			reg = regexp.MustCompile(`\$\{gutter\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Gutter)
			reg = regexp.MustCompile(`\$\{colcount\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Colcount)
			reg = regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Componentname)
			reg = regexp.MustCompile(`\$\{buttons\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, buttons)
		case "formlist":
			reg := regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Componentname)
		case "query":
			reg := regexp.MustCompile(`\$\{colcount\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Colcount)
			reg = regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Componentname)
		case "list":
			reg := regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, cad.Component.Componentname)
		case "singletablelist":
			reg := regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, componentname)
		case "listandsave":
			reg := regexp.MustCompile(`\$\{componentname\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, componentname)
			reg = regexp.MustCompile(`\$\{buttons\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, buttons)
		}

		_, err = f.WriteString(filecontent)
		if err != nil {
			fmt.Println("WriteString")
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (this *DEVController) updateNGcomponent(modulename string, componentname string, ngpath string, gopath string, currentDirectory string) error {
	cadarr, err := models.GetComponentbyparentid(models.COMPONENT{Componentname: componentname})
	if err != nil {
		fmt.Println(err)
		return err
	}
	componentfilepath := ngpath + "/src/app/" + modulename + "/" + componentname + "/" + componentname + ".component.ts"
	f, err3 := os.OpenFile(componentfilepath, os.O_RDWR|os.O_TRUNC, 0644)
	defer f.Close()
	if err3 != nil {
		fmt.Println(err3)
		return err3
	}

	componentfiletmppath := ""
	if len(cadarr) == 2 {
		style1 := cadarr[0].Component.Style
		style2 := cadarr[1].Component.Style
		if style1+style2 == "oaformformlist" || style1+style2 == "formlistoaform" {
			componentfiletmppath = "/views/template/component_formandlistoa.tpl"
		}
		if style1+style2 == "formformlist" || style1+style2 == "formlistform" {
			componentfiletmppath = "/views/template/component_formandlist.tpl"
		}
		if style1+style2 == "querylist" || style1+style2 == "listquery" {
			componentfiletmppath = "/views/template/component_queryandlist.tpl"
		}
		if style1+style2 == "querylistandsave" || style1+style2 == "listandsavequery" {
			componentfiletmppath = "/views/template/component_queryandlistsave.tpl"
		}
		if style1+style2 == "querysingletablelist" || style1+style2 == "singletablelistquery" {
			componentfiletmppath = "/views/template/component_querysgltblist.tpl"
		}

	} else {
		switch cadarr[0].Component.Style {
		case "form":
			componentfiletmppath = "/views/template/component_form.tpl"
		case "oaform":
			componentfiletmppath = "/views/template/component_formoa.tpl"
		case "singletablelist":
			componentfiletmppath = "/views/template/component_sgltblist.tpl"

		}
	}
	commonctrl := COMMONController{}

	filecontent, err := commonctrl.Readfile2string(currentDirectory+componentfiletmppath, "utf8")
	if err != nil {
		fmt.Println(err)
		return err
	}
	reg := regexp.MustCompile(`\$\{firstuppermodulename\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(modulename))

	reg = regexp.MustCompile(`\$\{modulename\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, modulename)

	reg = regexp.MustCompile(`\$\{uppermodulename\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Toupper(modulename))

	reg = regexp.MustCompile(`\$\{componentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, componentname)

	reg = regexp.MustCompile(`\$\{firstuppercomponentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(componentname))

	reg = regexp.MustCompile(`\$\{uppercomponentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Toupper(componentname))

	for _, cad := range cadarr {
		pkfield := ""
		listcolnames := ""
		if cad.Component.Style == "query" {
			listcolnames = "this.queryitems =["
		}
		if cad.Component.Style == "form" {
			listcolnames = "this.formcolnames =["
			pkfield = cad.Detail[0].Controlname
		}
		if cad.Component.Style == "oaform" {
			listcolnames = "this.formcolnames =["
			listcolnames = listcolnames + "{'Controlname': 'Opinion', 'Controltype': 'textbox', 'NotDisplayed': true},\n"

			pkfield = cad.Detail[0].Controlname
		}
		if cad.Component.Style == "list" || cad.Component.Style == "listandsave" || cad.Component.Style == "formlist" || cad.Component.Style == "singletablelist" {
			listcolnames = "this.listcolnames =["
		}
		for idx, cdtl := range cad.Detail {
			listcolnames = listcolnames + models.Outputcontrol(cdtl)
			if idx < len(cad.Detail)-1 {
				listcolnames = listcolnames + ",\n"
			}
		}
		listcolnames = listcolnames + "];"
		fmt.Println(listcolnames)
		if cad.Component.Style == "query" {
			reg = regexp.MustCompile(`\$\{this.queryitems = \[\];\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, listcolnames)
		}
		if cad.Component.Style == "form" {
			reg = regexp.MustCompile(`\$\{this.formcolnames = \[\];\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, listcolnames)
		}
		if cad.Component.Style == "oaform" {
			reg = regexp.MustCompile(`\$\{this.formcolnames = \[\];\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, listcolnames)
		}
		if cad.Component.Style == "list" || cad.Component.Style == "listandsave" || cad.Component.Style == "formlist" || cad.Component.Style == "singletablelist" {
			reg = regexp.MustCompile(`\$\{this.listcolnames = \[\];\$\}`)
			filecontent = reg.ReplaceAllString(filecontent, listcolnames)
		}
		reg = regexp.MustCompile(`\$\{pkfield\$\}`)
		filecontent = reg.ReplaceAllString(filecontent, models.Tolower(pkfield))

	}
	_, err = f.WriteString(filecontent)
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}
	return nil

}

//update angular home-routing
func (this *DEVController) updateNGrouter(modulename string, componentname string, ngpath string) error {
	//import {CreateComponentComponent} from './skl-dev/create-component/create-component.component';
	//{path: 'create-component', component: CreateComponentComponent}
	commonctrl := COMMONController{}
	homeroutfilepath := ngpath + "/src/app/" + "home-routing.module.ts"
	filecontent, err := commonctrl.Readfile2string(homeroutfilepath, "utf8")
	if err != nil {
		fmt.Println(err)
		return err
	}
	//avoid to import duplicated
	detcontent1 := "import {" + models.Tofirstupper(componentname) + "Component} from './" + modulename + "/" + componentname + "/" + componentname + ".component';"
	reg := regexp.MustCompile(detcontent1)
	if reg.FindString(filecontent) != "" {
		return nil
	}

	srccontent := "import {HomeComponent} from './home/home/home.component';\n"
	reg = regexp.MustCompile(`import {HomeComponent} from './home/home/home.component';`)
	detcontent := "import {" + models.Tofirstupper(componentname) + "Component} from './" + modulename + "/" + componentname + "/" + componentname + ".component';\n"
	filecontent = reg.ReplaceAllString(filecontent, detcontent+srccontent)

	reg = regexp.MustCompile(`\{path: 'lang', component: LangComponent\}`)
	srccontent = "     {path: 'lang', component: LangComponent}\n"
	detcontent = "     {path: '" + componentname + "', component: " + models.Tofirstupper(componentname) + "Component},\n"
	filecontent = reg.ReplaceAllString(filecontent, detcontent+srccontent)

	f, err3 := os.OpenFile(homeroutfilepath, os.O_RDWR|os.O_TRUNC, 0644)
	defer f.Close()
	if err3 != nil {
		fmt.Println(err3)
		return err3
	}
	_, err = f.WriteString(filecontent)
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}
	return nil
}

//update angular service
//modulename:
//componentname:
//ngpath:
//gopath:
func (this *DEVController) updateNGservice(modulename string, componentname string, ngpath string, gopath string, currentDirectory string) error {
	cadarr, err := models.GetComponentbyparentid(models.COMPONENT{Componentname: componentname})
	if err != nil {
		fmt.Println(err)
		return err
	}
	//constructor()-->constructor(private ls: LoginService, private http: HttpClient)
	commonctrl := COMMONController{}
	servicefilepath := ngpath + "/src/app/" + modulename + "/" + modulename + ".service.ts"
	filecontent, err1 := commonctrl.Readfile2string(servicefilepath, "utf8")
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}
	reg := regexp.MustCompile(`constructor\(\)`)
	filecontent = reg.ReplaceAllString(filecontent, "constructor(private ls: LoginService, private http: HttpClient)")

	//insert service_save.tpl
	savecontent := ""
	servicesavefilepath := ""
	nothassavecontent := true

	if len(cadarr) == 1 {
		if cadarr[0].Component.Style == "singletablelist" || cadarr[0].Component.Style == "form" {
			servicesavefilepath = "/views/template/" + "service_sgltblist.tpl"
			//save${componentname$}
			reg = regexp.MustCompile("save" + componentname)
			if reg.FindString(filecontent) != "" {
				nothassavecontent = false
			}
		}
		if cadarr[0].Component.Style == "oaform" {
			servicesavefilepath = "/views/template/" + "service_formoa.tpl"
			//save${componentname$}
			reg = regexp.MustCompile("save" + componentname)
			if reg.FindString(filecontent) != "" {
				nothassavecontent = false
			}
		}
	} else {
		style1 := cadarr[0].Component.Style
		style2 := cadarr[1].Component.Style
		if style1+style2 == "oaformformlist" || style1+style2 == "formlistoaform" {
			servicesavefilepath = "/views/template/" + "service_formandlistoa.tpl"
			//save${componentname$}
			reg = regexp.MustCompile("save" + componentname)
			if reg.FindString(filecontent) != "" {
				nothassavecontent = false
			}
		}
		if style1+style2 == "formformlist" || style1+style2 == "formlistform" {
			servicesavefilepath = "/views/template/" + "service_formandlist.tpl"
			//save${componentname$}
			reg = regexp.MustCompile("save" + componentname)
			if reg.FindString(filecontent) != "" {
				nothassavecontent = false
			}
		}
		if style1+style2 == "querylist" || style1+style2 == "listquery" {
			servicesavefilepath = "/views/template/" + "service_queryandlist.tpl"
			//get${componentname$}count
			reg = regexp.MustCompile("get" + componentname + "count")
			if reg.FindString(filecontent) != "" {
				nothassavecontent = false
			}
		}
		if style1+style2 == "querysingletablelist" || style1+style2 == "singletablelistquery" {
			servicesavefilepath = "/views/template/" + "service_querysgltblist.tpl"
			//save${componentname$}
			reg = regexp.MustCompile("save" + componentname)
			if reg.FindString(filecontent) != "" {
				nothassavecontent = false
			}
		}
		if style1+style2 == "querylistandsave" || style1+style2 == "listandsavequery" {
			servicesavefilepath = "/views/template/" + "service_queryandlistsave.tpl"
			//get${componentname$}count
			reg = regexp.MustCompile("get" + componentname + "count")
			if reg.FindString(filecontent) != "" {
				nothassavecontent = false
			}
		}
	}
	if nothassavecontent {
		savecontent, err = commonctrl.Readfile2string(currentDirectory+servicesavefilepath, "utf8")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	reg = regexp.MustCompile(`\$\{componentname\$\}`)
	savecontent = reg.ReplaceAllString(savecontent, componentname)

	//insert service_import.tpl
	importcontent := ""
	reg = regexp.MustCompile(`import \{HttpHeaders\} from '@angular/common/http';`)
	if reg.FindString(filecontent) == "" {

		serviceimportfilepath := "/views/template/" + "service_import.tpl"
		importcontent, err = commonctrl.Readfile2string(currentDirectory+serviceimportfilepath, "utf8")
		if err != nil {
			fmt.Println(err)
			return err
		}

	}
	f, err3 := os.OpenFile(servicefilepath, os.O_RDWR|os.O_TRUNC, 0644)
	defer f.Close()
	if err3 != nil {
		fmt.Println(err3)
		return err3
	}
	_, err = f.WriteString(importcontent + filecontent)
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}
	_, err = f.Seek(-3, 2)
	if err != nil {
		fmt.Println("Seek")
		fmt.Println(err)
		return err
	}
	_, err = f.WriteString("\n" + savecontent + "}\n")
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}
	return nil
}
func (this *DEVController) createcontrol(componentname, gopath, currentDirectory string) error {
	cadarr, err := models.GetComponentbyparentid(models.COMPONENT{Componentname: componentname})
	if err != nil {
		fmt.Println(err)
		return err
	}
	commonctrl := COMMONController{}
	f, err2 := os.OpenFile(gopath+"/controllers/"+componentname+"control.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	controlfiletmppath := ""
	if len(cadarr) == 2 {
		style1 := cadarr[0].Component.Style
		style2 := cadarr[1].Component.Style
		if style1+style2 == "oaformformlist" || style1+style2 == "formlistoaform" {
			controlfiletmppath = "/views/template/control_formandlistoa.tpl"
		}
		if style1+style2 == "formformlist" || style1+style2 == "formlistform" {
			controlfiletmppath = "/views/template/control_formandlist.tpl"
		}
		if style1+style2 == "querylist" || style1+style2 == "listquery" {
			controlfiletmppath = "/views/template/control_queryandlist.tpl"
		}
		if style1+style2 == "querylistandsave" || style1+style2 == "listandsavequery" {
			controlfiletmppath = "/views/template/control_queryandlistsave.tpl"
		}
		if style1+style2 == "querysingletablelist" || style1+style2 == "singletablelistquery" {
			controlfiletmppath = "/views/template/control_querysgltblist.tpl"
		}

	} else {
		switch cadarr[0].Component.Style {
		case "form":
			controlfiletmppath = "/views/template/control_form.tpl"
		case "oaform":
			controlfiletmppath = "/views/template/control_formoa.tpl"
		case "singletablelist":
			controlfiletmppath = "/views/template/control_sgltblist.tpl"

		}
	}

	filecontent, err := commonctrl.Readfile2string(currentDirectory+controlfiletmppath, "utf8")
	if err != nil {
		fmt.Println(err)
		return err
	}

	reg := regexp.MustCompile(`\$\{firstuppercomponentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Tofirstupper(componentname))

	reg = regexp.MustCompile(`\$\{componentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, componentname)

	reg = regexp.MustCompile(`\$\{uppercomponentname\$\}`)
	filecontent = reg.ReplaceAllString(filecontent, models.Toupper(componentname))

	reg = regexp.MustCompile(`\$\{modulename\$\}`)

	filecontent = reg.ReplaceAllString(filecontent, commonctrl.Getgoprojectnamebygopath(gopath))

	_, err = f.WriteString(filecontent)
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}

	return nil
}
func (this *DEVController) insertmodule(modulefilename string, modulename string) error {
	//import { ExpenseService } from './expense.service';
	srvstr := "import { " + models.Tofirstupper(modulename) + "Service } from './" + modulename + ".service';\n"
	includedstrarr := "import { SklCoreModule } from '../skl-core/skl-core.module';\nimport { SklCommonModule } from '../common/common.module';\nimport { TranslateModule} from '@ngx-translate/core';\nimport { RouterModule} from '@angular/router';\n"
	commonctrl := COMMONController{}
	filecontent, err := commonctrl.Readfile2string(modulefilename, "utf8")
	if err != nil {
		fmt.Println(err)
		return err
	}
	f, err2 := os.OpenFile(modulefilename, os.O_RDWR, 0644)
	defer f.Close()
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	// seek sets the offset for the next Read or Write on file to offset, interpreted
	// according to whence: 0 means relative to the origin of the file, 1 means
	// relative to the current offset, and 2 means relative to the end.
	// It returns the new offset and an error, if any.

	// 获取文件指针当前位置
	//cur_offset, _ := f.Seek(0, os.SEEK_CUR)
	_, err = f.Seek(0, 0)
	if err != nil {
		fmt.Println("Seek")
		fmt.Println(err)
		return err
	}

	_, err = f.WriteString(includedstrarr + srvstr + filecontent)
	if err != nil {
		fmt.Println("WriteString")
		fmt.Println(err)
		return err
	}

	return nil
}
func (this *DEVController) Execmd(cmdstr string, arg ...string) error {
	cmd := &exec.Cmd{}
	cmd = exec.Command(cmdstr, arg...)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return err
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return err
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return err
	}
	fmt.Printf("stdout:\n\n %s", bytes)
	return nil
}
