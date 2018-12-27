CREATE TABLE cmn_admingroup_tb (
  adminid varchar(50) NOT NULL,
  groupid varchar(50) NOT NULL,
  PRIMARY KEY (adminid,groupid)
);
CREATE TABLE cmn_adminorg_tb (
  adminid varchar(50) NOT NULL,
  orgid varchar(50) NOT NULL,
  PRIMARY KEY (adminid,orgid)
);
CREATE TABLE cmn_adminrole_tb (
  adminid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (adminid,roleid)
);


CREATE TABLE cmn_grouplevel_tb (
  upperid varchar(50) NOT NULL,
  lowerid varchar(50) NOT NULL,
  PRIMARY KEY (upperid,lowerid)
);
CREATE TABLE cmn_grouprole_tb (
  groupid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (groupid,roleid)
);
CREATE TABLE cmn_orgleader_tb (
  orgid varchar(50) NOT NULL,
  userid varchar(50) NOT NULL,
  leadertype char(1) DEFAULT NULL,
  PRIMARY KEY (orgid,userid)
);
CREATE TABLE cmn_orglevel_tb (
  upperid varchar(50) NOT NULL,
  lowerid varchar(50) NOT NULL,
  PRIMARY KEY (upperid,lowerid)
);
CREATE TABLE cmn_orgrole_tb (
  orgid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (orgid,roleid)
);
CREATE TABLE cmn_roleprivilege_tb (
  roleid varchar(50) NOT NULL,
  modualid varchar(50) NOT NULL,
  PRIMARY KEY (roleid,modualid)
);
CREATE TABLE cmn_usergroup_tb (
  userid varchar(50) NOT NULL,
  groupid varchar(50) NOT NULL,
  expireddate date DEFAULT NULL,
  PRIMARY KEY (userid,groupid)
);
CREATE TABLE cmn_userrole_tb (
  userid varchar(50) NOT NULL,
  roleid varchar(50) NOT NULL,
  PRIMARY KEY (userid,roleid)
);
CREATE TABLE dev_componentdetail_tb (
  componentname varchar(255) NOT NULL,
  seq int(11) NOT NULL DEFAULT '0',
  controlname varchar(255) NOT NULL DEFAULT '',
  controldisplayname varchar(255) DEFAULT '',
  controltype varchar(255) NOT NULL DEFAULT '',
  rows int(11) DEFAULT '0',
  filetype varchar(255) DEFAULT '',
  ismultiple tinyint(1) DEFAULT '0',
  filesize int(11) DEFAULT '0',
  islimit tinyint(1) DEFAULT '0',
  limitfileqty int(11) NOT NULL DEFAULT '0',
  minvalues int(11) NOT NULL DEFAULT '0',
  maxvalues int(11) NOT NULL DEFAULT '0',
  stepvalue int(11) NOT NULL DEFAULT '0',
  icon varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (componentname,seq)
);
CREATE TABLE sequence (
  seqname varchar(50) NOT NULL,
  currentValue int(11) NOT NULL,
  increment int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (seqname)
);
INSERT INTO cmn_org_tb VALUES ('root', '', '', null, -1, null, null);
INSERT INTO cmn_user_tb VALUES ('devzss', 'devzss', null, 0, 666666, '', '0001-01-01 00:00:00', null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, '', '', null, null, null, null, null, '', 0, 0, null, null, null, null, null, '', null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null, null);
INSERT INTO cmn_userrole_tb VALUES ('devzss', 'super');
INSERT INTO cmn_role_tb VALUES ('root', 'root', '', -1, null);
INSERT INTO cmn_role_tb VALUES ('super', '超级管理员', 'root', 0, null);
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'create-component');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', '项目管理');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'managesystem');
INSERT INTO cmn_roleprivilege_tb VALUES ('super', 'rolemanage');

INSERT INTO cmn_modual_tb VALUES ('root', 'root', '', null, null);
INSERT INTO cmn_modual_tb VALUES (0, 'managesystem', 'root', '', '');
INSERT INTO cmn_modual_tb VALUES ('modualmanage', '模块管理', 'systemmanage', '/modual', '');
INSERT INTO cmn_modual_tb VALUES ('orgmanage', '机构管理', 'systemmanage', '/org', '');
INSERT INTO cmn_modual_tb VALUES ('passwordchange', '密码变更', 'commonmanage', '/passwordchange', '');
INSERT INTO cmn_modual_tb VALUES ('rolemanage', '角色管理', 'systemmanage', '/role', '');
INSERT INTO cmn_modual_tb VALUES ('systemmanage', '系统管理', 0, '', '');
INSERT INTO cmn_modual_tb VALUES ('usergroup', '用户组管理', 'systemmanage', '/usergroup', '');
INSERT INTO cmn_modual_tb VALUES ('usermanage', '用户管理', 'systemmanage', '/usermanage', '');
INSERT INTO cmn_modual_tb VALUES ('项目管理', '项目管理', 0, '', '');
INSERT INTO cmn_modual_tb VALUES ('create-component', '组件开发', '项目管理', '/create-component', '');

INSERT INTO dev_componentdetail_tb VALUES ('claimantform', 1, 'Emergency', '', 'select', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantform', 2, 'Content', '', 'textbox', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantform', 3, 'Currencyid', '', 'select', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantform', 4, 'Remark', '', 'textarea', 2, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantform', 5, 'Hasticket', '', 'select', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantform', 6, 'Ticketcomment', '', 'textarea', 2, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantform', 7, 'Attchment', '', 'upload', 0, '', 1, 10240, 1, 3, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantformlist', 1, 'Paycontent', '', 'select', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantformlist', 2, 'Payee', '', 'atcomplete', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantformlist', 3, 'Bankaccount', '', 'select', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantformlist', 4, 'Amount', '', 'number', 0, '', 0, 0, 0, 0, 1, 999999, 1, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantformlist', 5, 'Ticketnum', '', 'number', 0, '', 0, 0, 0, 0, 1, 999, 1, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantformlist', 6, 'Remark', '', 'textbox', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantformlist', 7, 'Attchment', '', 'upload', 0, '', 1, 999999, 1, 3, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquery', 1, 'Flowinstid', '', 'textbox', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquery', 2, 'Calltimefrom', '', 'datepicker', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquery', 3, 'Orgid', '', 'select', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquery', 4, 'Calltimeto', '', 'datepicker', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquery', 5, 'Caller', '', 'select', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquery', 6, 'Flowstatus', '', 'checkboxgroup', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquerylist', 1, 'Flowinstid', '', 'label', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquerylist', 2, 'Content', '', 'routerLink', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquerylist', 3, 'Amount', '', 'label', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquerylist', 4, 'Caller', '', 'label', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquerylist', 5, 'Calltime', '', 'label', 0, '', 0, 0, 0, 0, 0, 0, 0, '');
INSERT INTO dev_componentdetail_tb VALUES ('claimantquerylist', 6, 'Flowstatus', '', 'label', 0, '', 0, 0, 0, 0, 0, 0, 0, '');

INSERT INTO dev_component_tb VALUES ('root', -1, '***项目', '', '', '', '', 0, 'D:\goproject\src\skl-api', 'D:\angular\skl');
INSERT INTO dev_component_tb VALUES ('claimant', 'root', '索赔模块', ',save', '', '', '', 1, '', '');
INSERT INTO dev_component_tb VALUES ('claimantapply', 'claimant', '索赔申请', ',save', 'oaform', 40, 1, 2, '', '');
INSERT INTO dev_component_tb VALUES ('claimantform', 'claimantapply', '索赔申请表单', ',save', 'oaform', 40, 1, 3, '', '');
INSERT INTO dev_component_tb VALUES ('claimantformlist', 'claimantapply', '索赔申请列表', ',save', 'formlist', 40, 1, 3, '', '');
INSERT INTO dev_component_tb VALUES ('claimantsearch', 'claimant', '索赔申请查询', ',save', 'oaform', 40, 1, 2, '', '');
INSERT INTO dev_component_tb VALUES ('claimantquery', 'claimantsearch', '索赔申请查询', ',save', 'query', 40, 2, 3, '', '');
INSERT INTO dev_component_tb VALUES ('claimantquerylist', 'claimantsearch', '索赔申请查询列表', ',save', 'list', 40, 1, 3, '', '');
INSERT INTO cmn_group_tb VALUES ('root', '', 'root', -1, null);
INSERT INTO cmn_group_tb VALUES ('chuna', 'root', '出纳', -1, '');
insert into sequence(seqname,currentValue,increment) values('fiid_sequence',201800001,1);
insert into sequence(seqname,currentValue,increment) values('tiid_sequence',1,1);