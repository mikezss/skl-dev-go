import {Component, OnInit} from '@angular/core';
import {NzMessageService} from 'ng-zorro-antd';
import {${firstuppermodulename$}Service} from '../${modulename$}.service';
import {MasterService} from '../../master/master.service';
import {LoginService} from '../../login/login.service';
import {
  NavigationExtras,
  Route,
  Router,
  ParamMap,
  ActivatedRoute
} from '@angular/router';

@Component({
  selector: 'app-${componentname$}',
  templateUrl: './${componentname$}.component.html',
  styleUrls: ['./${componentname$}.component.css']
})
export class ${firstuppercomponentname$}Component implements OnInit {
  queryitems: any = {};
  querydata: any = {};
  listdata: any[] = [];
  listcolnames: any[] = [];
  filteredOptions: any[] = [];

  constructor(private ls: LoginService, private ms: MasterService, private cs: ${firstuppermodulename$}Service, private message: NzMessageService, private router: Router) {
  }

  ngOnInit() {
    ${this.queryitems = [];$}
    ${this.listcolnames = [];$}
  }

  doAction(event) {
    console.log(event);
    console.log(this.listdata);
	this.cs.do${componentname$}(this.listdata).subscribe(data => {
        this.message.info('submit==>' + data.status);
    });
  }

  getquery(event) {
    console.log(this.querydata);     
    this.cs.get${componentname$}(this.querydata).subscribe(data => {
      this.listdata = data; 
    });
  }
  reset() {
    this.querydata = {};
	this.listdata = [];
  }
  listdatachange(event){
	console.log(event);
  }
  formdatachange(event){
	console.log(event);
  }
  
}
