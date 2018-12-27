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

  formdata: any = {};
  formcolnames: any[] = [];

  constructor(private ls: LoginService, private cs: ${firstuppermodulename$}Service, private message: NzMessageService) {
  }

  ngOnInit() {
    ${this.formcolnames = [];$}
    this.cs.get${componentname$}().subscribe(response => {
      this.formdata = response;
    });
  }

  dosubmit(event) {
    // console.log(event);
    console.log(this.formdata);
    this.cs.save${componentname$}(this.formdata).subscribe(data => {
      console.log(data);
      this.message.info('submit==>' + data.status);
    });

  }
  formdatachange(event){
	console.log(event);
  }

}
