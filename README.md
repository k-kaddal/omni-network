# Omni Technical Challenge

## To Run:

1. `make build` ---> build ignite cli docker container. As I faced issues with ignite locally, the following make commands are based on this created docker container

2. `make env ETH_RPC=<put_your_ethrpc_url>` ---> setting yout .env file with your ETH_RPC URL

3. `make serve` ---> to start running the blockchain localy

4. Store a state in Omni, here you would need to put:

   - Eth_Address: the required ethereum address to read state (example: 0xF9e3C81871dF06754AbFC054593762DB0bE48c89)

   - Eth_Slot: the slot to which required to read storage (example: 1)

   - Omni_Address: add your omni address to exacute a transaction or add the name of the account (example: alice)

   cmd: `make store-state Eth_Address=<put_ethereum_address> Eth_Slot=<put_ethereum_slot> Omni_Address=<put omni address>`

5. to query the stored state (address, slot, balance, blocknumber, storate), you can curl as follow. The state object identifier is `<address-slot>`
   `curl -X GET http://0.0.0.0:1317/k-kaddal/omni/ethereumbridge/eth_state/<address-slot>`

6. to query all store states
   `curl -X GET http://0.0.0.0:1317/k-kaddal/omni/ethereumbridge/eth_state`

---

# Rooms for improvement:

    -   Seperate the implementation into three modules: EthState, EthMetadata, EthBridge
    -   EthBridge to construct a client to dial ethereum rpc
    -   EthMetada to set and update Metadata of the store EthState
    -   EthState to set and update the required ethereum state to be stored
    -   Verification mechanism could be refined to check and verify by signature
