bf:
	go build -o front -tags front
bs:
	go build -o server -tags server
df:
	env GOOS=linux go build -o dfront -tags front
ds:
	env GOOS=linux go build -o dserver -tags server

bdf: df
	docker build -f ./deployments/front/Dockerfile -t obsidianio/front .
	rm dfront

bds: ds
	docker build -f ./deployments/server/Dockerfile -t obsidianio/server .
	rm dserver
bpb:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	pb/counting.proto

# rpc error: code = Unavailable desc = upstream connect error or disconnect/reset before headers. reset reason: connection termination