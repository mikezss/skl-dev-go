import {Component, OnInit} from '@angular/core';
import {NzMessageService} from 'ng-zorro-antd';
import {${firstuppermodulename$}Service} from '../${modulename$}.service';
import {LoginService} from '../../login/login.service';
import {MasterService} from '../../master/master.service';
import {
  NavigationExtras,
  Route,
  Router,
  ParamMap,
  ActivatedRoute
} from '@angular/router';
import {NzTreeNode} from 'ng-zorro-antd';

@Component({
  selector: 'app-${componentname$}',
  templateUrl: './${componentname$}.component.html',
  styleUrls: ['./${componentname$}.component.css']
})
export class ${firstuppercomponentname$}Component implements OnInit {
  queryitems: any[] = [];
  querydata: any = {};
  listdata: any[] = [];
  listcolnames: any[] = [];
  pageindex = 1;
  pagesize = 10;
  total = 1;
  loading = false;

  constructor(private ls: LoginService, private ms: MasterService, private cs: ${firstuppermodulename$}Service, private message: NzMessageService, private router: ActivatedRoute) {
  }

  ngOnInit() {
    ${this.queryitems = [];$}
    ${this.listcolnames = [];$}    
    
    this.cs.get${componentname$}count(this.querydata).subscribe(response => {
      this.total = response.Total;
    });
    this.refreshtable({'Pageindex': 1, 'Pagesize': 10});

  }

  getquery(event) {
    console.log(event);
    if (event == 'search') {
      this.loading = true;
      this.querydata.Pageindex = this.pageindex;
      this.querydata.Pagesize = this.pagesize;
  	  this.querydata.Flowstatus = this.ms.checkboxgroup2string(this.queryitems[3].checkboxgroup);
      this.cs.get${componentname$}count(this.querydata).subscribe(response => {
        this.total = response.Total;
      });
  
      this.cs.get${componentname$}bypageindex(this.querydata).subscribe(data => {
        let respons: any[] = [];
        respons = data; 
  	  for (let data1 of respons) {
          data1.Routerlink = "/${componentname$}";
          data1.QueryParams = {'Mode': 's', 'Flowinstid': data1.Flowinstid};
          data1.Editable = true;
        }     
        this.listdata = respons;
        this.loading = false;
      });
    } else {
      this.reset();
    }
  }

  edit(event) {
    console.log(event);
  }

  refreshtable(event) {

    console.log(event);
    this.pageindex = event.Pageindex;
    this.pagesize = event.Pagesize;
    this.loading = true;
    this.querydata.Pageindex = this.pageindex;
    this.querydata.Pagesize = this.pagesize;
    console.log(this.querydata.Fiid);
    

    this.cs.get${componentname$}bypageindex(this.querydata).subscribe(data => {
      let respons: any[] = [];
      respons = data;
	  for (let data1 of respons) {
        data1.Routerlink = "/${componentname$}";
        data1.QueryParams = {'Mode': 's', 'Flowinstid': data1.Flowinstid};
        data1.Editable = true;
      }      
      this.listdata = respons;
      this.loading = false;
    });
  }

  reset() {
    this.querydata = {};
    this.listdata = [];
    this.loading = false;
    this.pageindex = 1;
    this.pagesize = 10;
    this.cs.get${componentname$}count(this.querydata).subscribe(response => {
      this.total = response.Total;
    });
    this.refreshtable({'Pageindex': 1, 'Pagesize': 10});
  }
  formdatachange(event) {
    console.log(event);
    if (event == 'Orgid') {
      this.querydata.Userid = '';
      this.queryitems[5].options = [];
      this.ms.getuseroptionsbyorgid(this.querydata.Orgid).subscribe(response => {
        this.queryitems[5].options = response;
      });
    }

  }
  listdatachange(event){
	
  }

}
