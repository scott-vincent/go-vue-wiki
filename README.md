# Introduction
This is my first application with a Vue.js front-end and a Go Lang back-end.

# How To Build The Front-End
```
cd frontend
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
go build
```
### Notes:
Will build .exe in current folder.
If you want the .exe in the go/bin folder (already on path), use *go install* instead of *go build*.

# How To Run The Server
```
cd backend
backend.exe
```
### Notes:
During development you can use the *refresh app* and *refresh server* shortcuts instead.
*Refresh app* will monitor the frontend folder for source file changes and will automatically build the Vue app after any change.
*Refresh server* will monitor the backend folder for source file changes and will automatically build and run the Go server after any change.
*Refresh npm app* just uses *npm run serve* to serve the Vue app and monitor for changes. Note that this uses port 8081 and the server uses port 8080.
The npm refresh is much faster than *refresh app* but will only work if the server port is hard-coded into your URLs, otherwise it will try to call the backend on port 8081.
    
To install refresh.exe (written in go) use:
```    
go get github.com/scott-vincent/refresh
```
