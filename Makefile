build:
	docker build -f Dockerfile . -t omni_i

env:
	echo "ETH_RPC=$(ETH_RPC)" > .env

version:
	docker run --rm -it omni_i ignite version 

serve: 
	docker run --rm -it -v $(shell pwd):/omni -w /omni -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name omni omni_i ignite chain serve

store-state:
	docker exec -it omni omnid tx ethereumbridge store-state $(Eth_Address) $(Eth_Slot) --from $(Omni_Address)

reset: 
	docker run --rm -it -v $(shell pwd):/omni -w /omni -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name omni omni_i ignite chain reset