--- 

# SQL Repository Service

The SQL Repository Service is a part of the Message Generator and Router Project. It integrates with MySQL and Redis to store and index data, and it uses Kafka for message brokering.

## Components

### SQL Repository (`zg_sql_repo`)
This component manages the storage and indexing of data using MySQL and Redis.

#### Docker Compose Configuration
```yaml
version: '3.8'

networks:
  local-net:
    external: true

services:

  zg_sql_repo:
    build:
      context: .
      dockerfile: ./Dockerfile
    container_name: zg_sql_repo
    env_file:
      - .env-docker
    networks:
      - local-net
    volumes:
      - .:/app
    depends_on:
      - zg_sql_repo_redis
      - zg_mysql_db_1
      - zg_mysql_db_2

  zg_sql_repo_redis:
    image: redis:latest
    container_name: zg_sql_repo_redis
    env_file:
      - .env-docker
    ports:
      - '16379:6379'
    networks:
      - local-net

  zg_mysql_db_1:
    image: mysql:5.7
    container_name: zg_mysql_db_1
    env_file:
      - .env-docker
    ports:
      - "${MYSQL_DB_1_PORT}:3306"
    volumes:
      - mysql_db_1_data:/var/lib/mysql
    networks:
      - local-net

  zg_mysql_db_2:
    image: mysql:5.7
    container_name: zg_mysql_db_2
    env_file:
      - .env-docker
    ports:
      - "${MYSQL_DB_2_PORT}:3306"
    volumes:
      - mysql_db_2_data:/var/lib/mysql
    networks:
      - local-net

volumes:
  mysql_db_1_data:
  mysql_db_2_data:
```

#### Environment Variables (`.env-docker`)
```env
ZG_KAFKA_ADDRESS=kafka:29092
MYSQL_DB_1_HOST=zg_mysql_db_1
MYSQL_DB_2_HOST=zg_mysql_db_2
MYSQL_DB_1_PORT=3306
MYSQL_DB_2_PORT=3307
MYSQL_ROOT_PASSWORD=rootpassword
MYSQL_DATABASE=db
MYSQL_USER=user
MYSQL_PASSWORD=password
ZG_REDIS_URL=zg_sql_repo_redis:16379
ZG_REDIS_DB=0
ZG_REDIS_CACHE_DB=1
ZG_REDIS_EXP_TIME=3600s
LOGSTASH_URL=http://logstash:5000
```

#### Configuration File (`config.yaml`)
```yaml
kafka:
  address: ${ZG_KAFKA_ADDRESS}
  group_id: sql_repo
  user: guest
  password: guest
  topic: processing_1

dbs:
  - host: ${MYSQL_DB_1_HOST}
    port: ${MYSQL_DB_1_PORT}
    database: ${MYSQL_DATABASE}
    user: ${MYSQL_USER}
    password: ${MYSQL_PASSWORD}
    migrations_path: ./db/migrations

  - host: ${MYSQL_DB_2_HOST}
    port: ${MYSQL_DB_2_PORT}
    database: ${MYSQL_DATABASE}
    user: ${MYSQL_USER}
    password: ${MYSQL_PASSWORD}
    migrations_path: ./db/migrations

cache:
  address: ${ZG_REDIS_URL}
  db: ${ZG_REDIS_CACHE_DB}
  exp_time: ${ZG_REDIS_EXP_TIME}

kvdb:
  address: ${ZG_REDIS_URL}
  db: ${ZG_REDIS_DB}

logstash:
  url: ${LOGSTASH_URL}
```

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Running the SQL Repository Service
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/message-generator.git
   cd message-generator/sql-repo
   ```
2. Build and run the Docker containers:
   ```bash
   docker-compose up --build
   ```

### Environment Variables
Ensure to set the following environment variables in the `.env-docker` file:
- `ZG_KAFKA_ADDRESS`: Address of the Kafka server (e.g., `kafka:29092`).
- `MYSQL_DB_1_HOST`: Hostname of the first MySQL instance (e.g., `zg_mysql_db_1`).
- `MYSQL_DB_2_HOST`: Hostname of the second MySQL instance (e.g., `zg_mysql_db_2`).
- `MYSQL_DB_1_PORT`: Port of the first MySQL instance (e.g., `3306`).
- `MYSQL_DB_2_PORT`: Port of the second MySQL instance (e.g., `3307`).
- `MYSQL_ROOT_PASSWORD`: Root password for MySQL instances (e.g., `rootpassword`).
- `MYSQL_DATABASE`: Database name (e.g., `db`).
- `MYSQL_USER`: Username for MySQL instances (e.g., `user`).
- `MYSQL_PASSWORD`: Password for MySQL instances (e.g., `password`).
- `ZG_REDIS_URL`: URL of the Redis server (e.g., `zg_sql_repo_redis:16379`).
- `ZG_REDIS_DB`: Redis database number for KV store (e.g., `0`).
- `ZG_REDIS_CACHE_DB`: Redis database number for caching (e.g., `1`).
- `ZG_REDIS_EXP_TIME`: Expiration time for Redis cache (e.g., `3600s`).
- `LOGSTASH_URL`: URL of the Logstash server (e.g., `http://logstash:5000`).

## License
This project is licensed under the MIT License.

---
