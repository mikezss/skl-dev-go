import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs';
import {UploadFile} from 'ng-zorro-antd';
import {NzMessageService} from 'ng-zorro-antd';
import {MasterService} from '../../master/master.service';
import {LoginService} from '../../login/login.service';
import {${firstuppermodulename$}Service} from '../${modulename$}.service';
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
  m = 'a';  
  formcolnames: any[] = [];
  listcolnames: any[] = [];
  formdata: any = {};
  listdata: any[] = [];

  constructor(private cs: ${firstuppermodulename$}Service,private ls: LoginService,private ms: MasterService,private message: NzMessageService, private router: ActivatedRoute, private router2: Router) {
  }

  ngOnInit() {
	${this.formcolnames = [];$}
    ${this.listcolnames = [];$}
    
    this.router.queryParams.subscribe(params => {
        console.log(params);
       
        this.cs.get${componentname$}byid({'${pkfield$}': params.${pkfield$}})
          .subscribe(data => {
            this.formdata = data;
          });
        this.cs.get${componentname$}item({'${pkfield$}': params.${pkfield$}})
          .subscribe(data => {
            this.listdata = data;
          });
        this.cs.get${componentname$}filesbyid({'${pkfield$}': params.${pkfield$}})
          .subscribe(data => {
            const flist: any[] = [];
            for (const data1 of data) {
              flist.push({'uid': data1.Uid, 'name': data1.Name, 'url': this.ls.api_url + data1.Url});
            }
            this.formcolnames[0].fileList = flist;
          });
      }
    );

  }

  dosubmit(event) {
   
    let navigationExtras: NavigationExtras = {
      queryParamsHandling: 'preserve',
      preserveFragment: true
    };
    this.cs.save${componentname$}(this.formdata, this.listdata).subscribe(data => {
      console.log(data);

      if (data.status == 'ok') {
        let fiid = '';
        fiid = <string>data.result;
        if (this.formcolnames[0].fileList.length > 0) {

          if (event == 'save' || event == 'submit') {
            const formData = new FormData();           
            for (let i = 0; i < this.formcolnames[0].fileList.length; i++) {
              formData.append('filelist', this.formcolnames[0].fileList[i]);
            }

            formData.append('filepath', '');
            formData.append('renamefilename', 'false');

            this.ms.uploadfile(formData).subscribe(data => {
              if (data.status == 'ok') {
                this.message.info(data.status);
              } else {
                this.message.info(data.result);
              }

            });
          }

          for (let i = 0; i < this.listdata.length; i++) {
            
            if (this.listdata[i].Attachment != 'undefined' && this.listdata[i].Attachment.length > 0) {
              const formData = new FormData();
              let istr = '';
              istr = String(i + 1);
              for (let j = 0; j < this.listdata[i].Attachment.length; j++) {
                
                formData.append('filelist', this.listdata[i].Attachment[j].originFileObj);
              }
              formData.append('filepath', '');
              formData.append('renamefilename', 'false');
              this.ms.uploadfile(formData).subscribe(data => {
                if (data.status == 'ok') {
                  this.message.info(data.status);
                } else {
                  this.message.info(data.result);
                }
              });
            }

          }
        }
        
      } else {
        this.message.info(data.result);
      }
    });
  }

  uploadcheck(event) {
    console.log(event);
    for (const colname of this.formcolnames) {
      if (colname.Controltype == 'upload') {
        if (event.size / 1024 > colname.filesize) {
          this.message.info('permitted file size is :' + colname.filesize);
          break;
        }

        if (colname.fileList.length < colname.limitfileqty) {
          colname.fileList.push(event);
          break;
        }
      }
    }    

  }
  listdatachange(event){
	console.log(event);
  }
  formdatachange(event){
	console.log(event);
  }

}
