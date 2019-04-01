<app-skl-list 
	listtitle="{{'${componentname$}title'|translate}}" 
	[listdata]="listdata" 	
	[listcolnames]="listcolnames" 
	(action)="save${componentname$}($event)"
	(listdatachange)="listdatachange($event)"
	[(pageIndex)]="pageindex" 
	[(pageSize)]="pagesize" 
	[total]="total" 
	[loading]="loading" 
	(refreshtable)="refreshtable($event)"
	[buttons]="[{'name':'save','icon':'save'}]">
</app-skl-list>
