rpc:
	protoc.exe --go_out=. --go-grpc_out=. .\rpc_service\auth_rpc.proto
b:
	go build -o ./build/cicoyi_user_rpc.exe ./cmd/cicoyi_user_rpc.go

clean:
	cd build && del cicoyi_user_rpc.exe && cd ..

run:
	.\build\cicoyi_user_rpc.exe

git:
	git add .
	git commit -m "cicoyi_user_rpc"
	git push