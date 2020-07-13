cd frontend
call npm run build
cd ..\backend
call go build -o go-vue-wiki.exe
cd ..
mkdir image >nul 2>nul
xcopy /y backend\go-vue-wiki.exe image
xcopy /y /s /i frontend\dist image\dist
