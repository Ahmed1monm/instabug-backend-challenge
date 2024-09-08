![instabug-task system design](./instabug-task-system-design.svg "System design")
----------------------------


## INSTALLATION
Make sure you have docker installed on your machine.
```bash
    docker-compose up --build
```

> [!CAUTION] 
> Disclaimer: You may see some error logs that the app can't connect to the database, no worries it's because the database container > ?takes some time to be ready, the app will retry to connect to the database until it's ready.

## USAGE
- Attached is a postman collection that you can use to test the API. in the top level directory of the project

## SYSTEM DESIGN

### Components:
- **NGINX** - Reverse proxy server that routes read requests to the Rails API and write requests to the Go API.
- **Rails API** - Handles read requests and the CREATE of `applications`
- **Go API [Writer service]** - Handles write requests of `Chats` and `Messages`
- **Worker** - Worker service that listens to the `Chats` and `Messages` queue to write to the database and index the messages in Elasticsearch. 
- **Elasticsearch** - Used to index the messages for search functionality.
- **MySQL** - Main datastore for the application.
- **Redis** - Used as a message broker for the worker service and got accessed using `asyncq` package.
- **Redis** - Used for caching the `applications` data, it chashes the GET application only [for simplicity]
- **Rake task** - A rake task that used as background job to update the count of the `chats` in the `applications` table and `messages` count in the `chats` table.

### API Endpoints:
```bash
GET /applications
POST /applications
GET /applications/:Token
PUT /applications/:Token

POST /applications/:Token/chats
GET /applications/:Token/chats
GET /applications/:Token/chats/:ChatNumber
PUT /applications/:Token/chats/:ChatNumber


POST /applications/:Token/chats/:ChatNumber/messages
GET /applications/:Token/chats/:ChatNumber/messages
GET /applications/:Token/chats/:ChatNumber/messages/:MessageNumber
PUT /applications/:Token/chats/:ChatNumber/messages/:MessageNumber
```

### Decisions:
- Preventing collisions of `chat_number` per application and `message_number` per chat by using **Compound Indexes** in the database.

- Menimizing the number of requests to the database by caching the `applications` data in Redis and updating the count of the `chats` in the `applications` table and `messages` count in the `chats` table using a rake task.

- Hanling race conditions in the worker service by using the `asyncq` package to access the Redis queue and apply the writes by Transactions.
