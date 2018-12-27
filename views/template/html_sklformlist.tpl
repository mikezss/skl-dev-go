<app-skl-list 
	listtitle="{{'${componentname$}title'|translate}}" 
	[listdata]="listdata" 
	[listcolnames]="listcolnames"
	(listdatachange)="listdatachange($event)" 
	[zScroll]="{ x: '400px' }"
	tablesize="small"
    [mode]="m"
>
</app-skl-list>
