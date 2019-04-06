cd ./ddds
set "GOOS=linux"
go build -ldflags "-X main.buildVersionString=v%1" ddds
cd ..

cd ./dddc
set "GOOS=linux"
go build -ldflags "-X main.buildVersionString=v%1" dddc
cd ..

mkdir .\bin
move .\ddds\ddds .\bin\ddds
move .\dddc\dddc .\bin\dddc