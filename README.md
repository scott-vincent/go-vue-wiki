# Introduction
This is my first application with a Vue.js front-end and a Go Lang back-end.

# How To Build The Front-End
```
cd frontend
npm install     (only needed once to download dependencies)
npm run build
```
### Notes:
Must have vue cli + service installed:
```
npm install -g @vue/cli @vue/cli-service-global
```
Will build into dist folder so server must be configured to serve root (static files) from frontend/dist

# How To Build The Back-End
```
cd backend
go build -o go-vue-wiki.exe
```
### Notes:
On Ubuntu you need to install Go by typing: apt install golang-go
On Windows, download the binary from the GoLang website.
Will build the .exe in the current folder.
If you want the .exe in the go/bin folder (already on path), use *go install* instead of *go build*.

# How To Run The Server
```
cd backend
go-vue-wiki.exe
```
### Notes:
The app is normally served from the server but during development you can use the *refresh app* and *refresh server* shortcuts instead to instantly apply updates when you make source code changes.
*Refresh server* will monitor the backend folder for source file changes and will automatically build and run the Go server after any change.
*Refresh app* just uses *npm run serve* to serve the Vue app on port 8081 and monitor for changes. As the server uses port 8080 it must allow CORS from http://localhost:8081 for this to work successfully.
    
To install refresh.exe (written in go) use:
```    
go get github.com/scott-vincent/refresh
```
