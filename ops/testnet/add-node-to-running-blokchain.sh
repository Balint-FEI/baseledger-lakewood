first_node_tendermint_p2p_port=26655
tendermint_internal_grpc_port=26657
node1_id=d172fc55469d6b093d4144a9699403818282ebfb
node1_ip=46.163.116.182

docker exec baseledger-node_blockchain_app_1 baseledgerd init <node_id> --chain-id baseledger
docker exec baseledger-node_blockchain_app_1 baseledgerd keys add <node_id>_address_1 --keyring-backend test
docker exec baseledger-node_blockchain_app_1 baseledgerd keys add <node_id>_address_2 --keyring-backend test
docker exec baseledger-node_blockchain_app_1 baseledgerd keys add <node_id>_address_3 --keyring-backend test
docker exec baseledger-node_blockchain_app_1 baseledgerd keys add <node_id>_address_4 --keyring-backend test
docker exec baseledger-node_blockchain_app_1 baseledgerd keys add <node_id>_address_5 --keyring-backend test
docker exec baseledger-node_blockchain_app_1 sed -i 's/persistent_peers = ""/persistent_peers = "'${node1_id}'@'${node1_ip}':'${first_node_tendermint_p2p_port}'"/' ~/.baseledger/config/config.toml
docker exec baseledger-node_blockchain_app_1 sed -i 's/addr_book_strict = true/addr_book_strict = false/' ~/.baseledger/config/config.toml
docker exec baseledger-node_blockchain_app_1 sed -i 's/allow_duplicate_ip = false/allow_duplicate_ip = true/' ~/.baseledger/config/config.toml
docker exec baseledger-node_blockchain_app_1 sed -i 's/timeout_commit = "5s"/timeout_commit = "9s"/' ~/.baseledger/config/config.toml
docker exec baseledger-node_blockchain_app_1 sed -i 's@laddr = "tcp://127.0.0.1:'${tendermint_internal_grpc_port}'"@laddr = "tcp://0.0.0.0:'${tendermint_internal_grpc_port}'"@' ~/.baseledger/config/config.toml
docker exec baseledger-node_blockchain_app_1 sed -i 's/enable = false/enable = true/' ~/.baseledger/config/app.toml

docker cp ./genesis.json baseledger-node_blockchain_app_1:/root/.baseledger/config/genesis.json

# docker exec baseledger-node_blockchain_app_1 baseledgerd start

