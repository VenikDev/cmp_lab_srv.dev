echo "Start"
cd D:\Projects\GO\comparison_laboratories/client

echo "Build"
yarn build

echo "go to server folder"
cd ../server/src

echo "Build and run server"
go run main.go
