# Build Front-End
```
cd frontend
vue build
```
### Notes:
Must have vue cli + service installed:
```
npm install -g @vue/cli @vue/cli-service-global
```
Will build into dist folder so server must be configured to serve root (static files) from frontend/dist

# Build Back-End
```
cd backend
go build
```
### Notes:
Will build .exe in current folder.
If you want the .exe in the go/bin folder (already on path), use *go install* instead of *go build*.

# Run Server
```
cd backend
backend.exe
```
### Notes:
During development you can use the 'refresh server' shortcut instead.
This will monitor the backend folder for source file changes and will automatically build and run the server after any change.
    
To install refresh.exe (written in go) use:
```    
go get github.com/markbates/refresh
```
