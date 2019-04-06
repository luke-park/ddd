.PHONY: build clean

build:
	cd .\ddds & go build ddds
	cd .\dddc & go build dddc

clean:
	-del .\ddds\ddds.exe
	-del .\dddc\dddc.exe

run-server:
	cd .\test-environment\server & ..\..\ddds\ddds.exe

run-client:
	cd .\test-environment\client & set "TEST_VAR=11.8" & ..\..\dddc\dddc.exe 127.0.0.1:21059 deployment.toml