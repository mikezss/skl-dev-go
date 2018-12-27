import {Component, OnInit} from '@angular/core';
import {${firstuppermodulename$}Service} from '../${modulename$}.service';
import {NzMessageService} from 'ng-zorro-antd';
import {LoginService} from '../../login/login.service';
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

  constructor(private ls: LoginService, private cs: ${firstuppermodulename$}Service, private message: NzMessageService) {
  }

  ngOnInit() {
	${this.queryitems = [];$}
    ${this.listcolnames = [];$}
    
	this.cs.get${componentname$}count(this.querydata).subscribe(response => {
      this.total = response.Total;
    });
    this.refreshtable({'Pageindex': 1, 'Pagesize': 10});
  }

  save${componentname$}(event) {
    console.log(this.listdata);
    this.cs.save${componentname$}(this.listdata).subscribe(data => {
      console.log(data);
      this.message.info('submit==>' + data.status);
    });
	this.cs.get${componentname$}count(this.querydata).subscribe(response => {
      this.total = response.Total;
    });
    this.refreshtable({'Pageindex': 1, 'Pagesize': 10});
  }
  listdatachange(event){
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
      this.listdata = data;
      this.loading = false;
    });
  }
  getquery(event) {
    console.log(event);
    this.loading = true;
    this.querydata.Pageindex = this.pageindex;
    this.querydata.Pagesize = this.pagesize;
	this.querydata.Flowstatus = this.ms.checkboxgroup2string(this.queryitems[3].checkboxgroup);
    this.cs.get${componentname$}count(this.querydata).subscribe(response => {
      this.total = response.Total;
    });

    this.cs.get${componentname$}bypageindex(this.querydata).subscribe(data => {
      
      this.listdata = data;
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

}
