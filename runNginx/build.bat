goto %1
goto end

:file
call:zip F:\gits\up\runNginx\runNginx.7z
call:upw 1 F:\gits\up\runNginx\runNginx.7z
goto end


:do
call:builds F:\gits\up\runNginx\runNginx.go
call:zip F:\gits\up\runNginx\autoNginx
call:up 1 F:\gits\up\runNginx\autoNginx
goto end
:builds
set pathT=%~dp1
cd /d %pathT%
set GOARCH=amd64
set GOOS=linux
go build -o autoNginx -i %1  
goto end

:up 
set svr=flylkl@fly.lkl:/tmp
scp %2 %svr%
goto end

:upw 
set svr=im@im.guiruntang.club:/tmp
scp %2 %svr%
goto end

:zip
set pathT=%~dp1
cd /d %pathT%
del autoNginx
del runNginx.7z
7zr a runNginx.7z . 
:end

