cd frontend
npm run build
cd ../backend
go build -o go-vue-wiki
cd ..
mkdir image >/dev/null 2>&1
cp backend/go-vue-wiki image
cp frontend/dist image/dist
