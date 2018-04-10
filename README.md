# Overview
It is an in-memory distributed key-value (KV) store using GO programming language. 
- The KV store is able to handle data larger than any one node's memory capacity. That is, at any given time, a single node might not have all the data.
- Server program that accepts HTTP get/post/put requests from the clients and returns a valid response. 
- The server will communicate with it's peer processes (spread across the network) to maintain a consistent view of the key-value database. 
- All communication between the HTTP client and this server should be in JSON format.
- Each server is also a proxy/coordinator process keeps track of available servers and data stored in those servers. 
- A client connects to the proxy/coordinator and server will learn the address of server which has the key and 
  forward the request for performing set/get operations. 
- The proxy server also acts as a load-balancer and ensures a uniform workload distribution among various servers.

# How to run
- Give the executable permission to ./build.sh and ./run.sh (e.g chmod +x FILENAME.sh)
- Open the build.sh and change the Go Paths based on your environment
- Run build.sh. It will create a executable server
- run.sh
  - It spawns the servers based on parameters as below
    - prefix(server prefix like 510 for 5100, 5101 and so on)
    - serverId(continuous ids starting from 0. For example, if you want 4 servers, then it will 0,1,2,3)
    - no. of servers to spawn
  - By default the script will create 4 servers on port 5100, 5101, 5102, 5103
- Run ./run.sh
- Once the servers are spawned, each server will have seed keys a and b.

# Endpoints
- SET - http://localhost:PORT/set
  - It creates/updates the key value pair in the store
  - It has following format
  ```json
  {"key": "<key>", "value": "<value>", "encoding": "<encoding>"}
  ```
  - Encoding is a type of value. For example, integer, string, binary etc.
  - Key is always stored as string. See example request:
  ```console
  curl -H "Content-Type: application/json" -d '{"key": "d", "value": "onto", "encoding": "string"}' http://localhost:5100/set- 
  ```
- GET - http://localhost:PORT/get/{key}
  - Key is the parameter for which you want the value
  - It returns the value of the key
  - You can test in browser. For example; once the servers are spawned, try http://localhost:<PORT>/get/a
  
# Attributions
- https://thenewstack.io/make-a-restful-json-api-go/ for api boilerplate of go
- https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go for hash function
