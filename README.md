# Introduction
I wanted to build an application that had a Vue.js front-end and a Go Lang back-end. This is the result and will be a useful blueprint for my future projects.

I'm extremely impressed by vue.js and Go. Both have done away with all the boilerplate code and make it incredibly easy to just concentrate on the code that matters. I found Visual Studio Code to be the perfect IDE and, with some auto build cleverness in place (see my Notes in How to Run the Server section) both the front and back end were rebuilding and redeploying within less than half a second of me making changes to the source files. This is especially useful when playing with fiddly CSS as you can see your results pretty instantly and adjust as necessary.

The choice of router for Vue was easy as there is an official one called Vue Router funnily enough. There are more choices for Go but I settled on Gorilla Mux as it seemed to be the most popular and fully featured one.

The Go server serves up the app files (which are generated using npm) as well as exposing the endpoints needed by the app.
For the auto building of the app I found it was much faster to let npm serve the app files and, as Go was running on port 8080 I had to use port 8081 for npm. This meant I needed to enable CORS on the GO server so the app could still use the endpoints. I only enabled CORS for localhost as its only needed during development and it all seemed to work well.

I also created a couple of scripts (one for windows and one for Linux) to build and package up the app and server into a neat distributable package.

Finally, I tried developing on both Windows and Ubuntu Linux (VSCode is available for both) and both were equally easy.

I'll definitely be using Vue.js and Go for many of my future projects, they were a delight to learn and use.

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
On Ubuntu you need to install Go by typing:
```
apt install golang-go
```
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
