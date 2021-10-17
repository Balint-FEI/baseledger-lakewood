Linux requirements Docker: BuildKit (https://docs.docker.com/develop/develop-images/build_enhancements/#to-enable-buildkit-builds)
COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_BUILDKIT=1

# Setup

1. Navigate to proxy_app/ops folder and edit **Dockerfile**
2. Set **API_UB_USER**=< your_api_user > and **API_UB_PWD**=< your_api_password > parameters as basic authantication credentials 
3. Set **DB_UB_USER**=< your_db_user > and **DB_UB_USER**=< your_db_password > parameters as DB  access credentials
4. Navigate to proxy_app/ops/migrations and edit **000001_initial.up.sql**
5. Replace < pass > with < your_db_password >
6. Navigate to ops/local folder and edit **docker-compose.yml**
7. Set **POSTGRES_USER**=< your_pg_user > and **POSTGRES_PASSWORD**=< your_pg_password > parameters and replace **nats: command: < nats-token >** with < your_nats_token> as Postgres credentials

# Running the full blockchain locally

1. Navigate to repo root/ops/local folder
2. Run: (WSL)
```
 ./run_blockchain.sh 
```
 or (Non root user on Linux)
```
sudo sh run_blockchain.sh
``` 
3. In two different terminals run:
```
# first terminal

 docker exec first_node_blockchain_app_1 baseledgerd start

# second terminal

 docker exec second_node_blockchain_app_1 baseledgerd start
```

# Cleanup

1. Navigate to repo root/ops/local folder
2. Run: (WSL)
```
 ./clear_blockchain.sh 
 ```
 or (Non root user on Linux)
 ```
 sudo sh clear_blockchain.sh 
```

# Postman collection

After setting up a local blockchain, you can play around with requests between two proxies by following the steps:

1. Import proxy_app/misc/baseledger demo proxy.postman_collection
2. Make sure endpoints have correct user and pass for basic auth previously set(API_UB_USER, API_UB_PWD), and are set to target localhost
3. Make sure workgroup id, organization id and organization nats endpoints have correct values set (as provided in docker-compose)
4. If request could not connect, check if the **proxy_app** container of the node is running:
```
docker ps -a
```
if not, start it:
```
docker start <container_id>
```