build:
	docker build -f Dockerfile . -t omni_i

version:
	docker run --rm -it omni_i ignite version 

scaffold:
	docker run --rm -it -v $(shell pwd):/omni -w /omni omni_i ignite scaffold chain github.com/k-kaddal/omni

#Try not to use serve with shell
serve: 
# chmod +x serve.sh && ./serve.sh
	docker run --rm -it -v $(shell pwd):/omni -w /omni -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name omni omni_i ignite chain serve

status: 
	docker exec -it omni omnid status

module: 
	docker exec -it omni ignite scaffold module EthState

# list: 
# # docker exec -it omni ignite scaffold list EthereumState Address StorageSlot StateStorage StateBalance --module EthBridgeTest
# 	docker exec -it omni ignite scaffold list EthereumStateNoMessage Address StorageSlot StateStorage StateBalance --module EthBridge --no-message

map:
# docker exec -it omni ignite scaffold map metadata address timestamp blockNumber --module EthBridge --no-message
	docker exec -it omni ignite scaffold map state slot balance data metadata:Metadata --module EthBridge --index address --no-message

message:
	docker exec -it omni ignite scaffold message store-state address slot --response state:State --module EthBridge

store-state:
	docker exec -it omni omnid tx ethbridge store-state $(Address) $(Slot) --from alice

# delete-all:
# 	rm -rf omni
# 	docker rmi omni_i