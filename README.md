## Redis
- SADD: 
  - add new items for set by key
  - SADD <key> <item1> [<item2> ...]
- SMEMBERS
  - Retrieve a set by key
  - SMEMBERS <key>
- SREM:
  - Delete one or more item of a set
  - SREM <key> <item1> [<item2> ...]
- DEL:
  - Delete a set
  - DEL users
- SScan:
  - SSCAN <key> cursor [MATCH <pattern>] [COUNT <count>]
  - 

- Pipelining:
  - The techinique is used for improving performance
  - Send multi requests without waiting fỏ response from the server util all requests have been sent

- SET:
  - SET <key> <value> [EX <seconds>] [PX <milliseconds>] [NX|XX]
  - NX: not exist
  - XX: existed (update)

- GET:
  - GET <key>
- MGET:
  - MGET <key1> [<key2> ...]


## TODO
- env - v
- optimize run app server - v
- write response not success
- context value
- 
