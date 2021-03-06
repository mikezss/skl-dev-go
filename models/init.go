package models

import (
	"github.com/astaxie/beego/orm"
)

// 需要在init中注册定义的model
func init() {

	orm.RegisterModel(new(CMN_USER_TB))
	orm.RegisterModel(new(CMN_MODUAL_TB))
	orm.RegisterModel(new(CMN_ORG_TB))
	orm.RegisterModel(new(CMN_ROLE_TB))
	orm.RegisterModel(new(CMN_GROUP_TB))
	orm.RegisterModel(new(CMN_MODUALTEMPLATE_TB))
	orm.RegisterModel(new(COMPONENT))
	orm.RegisterModel(new(ENUM))
	orm.RegisterModel(new(ENUMITEM))
}
