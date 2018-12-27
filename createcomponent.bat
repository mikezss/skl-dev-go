rem ng generate module login --module=app
rem ng generate service home/home --module=home
rem ng generate component skl-core/skl-query-list-save

set ngpath=%1
set gopath=%2
set modulename=%4

cd %ngpath%
d:


if ""%3"" == ""module"" goto createmodule
if ""%3"" == ""component"" goto createcomponent
if ""%3"" == ""service"" goto createservice

:createmodule
ng generate module %modulename% --module=app && ng generate service %modulename%/%modulename% --module=%modulename%  

:createservice
ng generate service %modulename%/%modulename% --module=%modulename%

:createcomponent
set componentname=%5
ng generate component %modulename%/%componentname%