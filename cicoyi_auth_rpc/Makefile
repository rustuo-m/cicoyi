rpc:
	protoc.exe --go_out=. --go-grpc_out=. .\rpc_service\auth_rpc.proto
b:
	go build -o ./build/auth_rpc.exe ./cmd/auth_rpc.go

clean:
	cd build && del auth_rpc.exe && cd ..

run:
	.\build\auth_rpc.exe

git:
	git add .
	git commit -m "cicoyi_auth_rpc"
	git push