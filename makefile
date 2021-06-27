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