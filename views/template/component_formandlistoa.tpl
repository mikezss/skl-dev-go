import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs';
import {UploadFile} from 'ng-zorro-antd';
import {MasterService} from '../../master/master.service';
import {NzMessageService} from 'ng-zorro-antd';
import {${firstuppermodulename$}Service} from '../${modulename$}.service';
import {FlowService} from '../../flow/flow.service';
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
  m = 'a';
  actionids: any[] = [];   
  modualid = '${componentname$}';
  currentfiid = 0;
  currenttiid = 0;
  formcolnames: any[] = [];
  listcolnames: any[] = [];
  formdata: any = {};
  listdata: any[] = [];

  constructor(private ls: LoginService, private fs: FlowService, private es: ${firstuppermodulename$}Service, private ms: MasterService, private message: NzMessageService, private router: ActivatedRoute, private router2: Router) {
  }


  ngOnInit() {

    ${this.formcolnames = [];$}
    ${this.listcolnames = [];$}	
	

    this.router.queryParams.subscribe(params => {
        console.log(params);
        if (params.Flowinstid != null && params.Flowinstid != '' && params.Flowinstid != 'undefined' && params.Flowinstid != 'NaN') {
          this.currentfiid = parseInt(params.Flowinstid);
          this.currenttiid = parseInt(params.Tiid);
        }
        if (params.Mode != '' && params.Mode != 'undefined') {
          this.m = params.Mode;
          if (this.m != 'a') {
            this.formcolnames[0].NotDisplayed = false;
            this.formcolnames[0].NotEditable = true;
          } else {
            this.formcolnames[0].NotDisplayed = true;
            this.formcolnames[0].NotEditable = false;
          }

        }

        this.fs.gettaskinfo(this.modualid, this.currentfiid, this.currenttiid).subscribe(data => {
          this.actionids = [];
          for (const data1 of data) {
            this.actionids.push(data1.Action);
          }
          console.log(this.currentfiid);
          console.log(this.currenttiid);

        });
        console.log('params.Flowinstid=======>');
        console.log(params.Flowinstid);
        console.log(params);
        this.es.get${componentname$}byid({'Flowinstid': parseInt(params.Flowinstid)})
          .subscribe(data => {
            this.formdata = data;
          });
        this.es.get${componentname$}item({'Flowinstid': parseInt(params.Flowinstid)})
          .subscribe(data => {
            this.listdata = data;
          });
        this.es.get${componentname$}filesbyid({'Fiid': parseInt(params.Flowinstid)})
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
    console.log(event);
    console.log(this.formdata);
    console.log(this.listdata);
    let navigationExtras: NavigationExtras = {
      queryParamsHandling: 'preserve',
      preserveFragment: true
    };
    this.es.save${componentname$}(this.formdata.Opinion, this.ls.userid, this.modualid, this.currentfiid, this.currenttiid, event, this.formdata, this.listdata).subscribe(data => {
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


            formData.append('filepath', 'static/files/' + fiid + '/');
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
            console.log(this.listdata[i].Attachment);
            if (this.listdata[i].Attachment != 'undefined' && this.listdata[i].Attachment.length > 0) {
              const formData = new FormData();
              let istr = '';
              istr = String(i + 1);
              console.log(this.listdata[i].Attachment);
              for (let j = 0; j < this.listdata[i].Attachment.length; j++) {
                console.log(this.listdata[i].Attachment[j]);
                console.log(this.listdata[i].Attachment[j].originFileObj);
                formData.append('filelist', this.listdata[i].Attachment[j].originFileObj);
              }
              formData.append('filepath', 'static/files/' + fiid + '/' + istr + '/');
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
        if (event == 'save' || event == 'submit') {
          navigationExtras = {queryParams: {'Mode': 'a'}};
          this.router2.navigate(['/myflow'], navigationExtras);
        } else {
          navigationExtras = {queryParams: {'Mode': 'a'}};
          this.router2.navigate(['/todo'], navigationExtras);
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
    console.log(this.formdata);

  }
  listdatachange(event){
	console.log(event);
  }
  formdatachange(event){
	console.log(event);
  }

}
