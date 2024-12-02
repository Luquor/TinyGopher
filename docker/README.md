# Environment variables
To setup the database, you need to provide a `.env` in the same directory as the `compose.yml` that contains your secrets, based on the following file structure:
```bash
MYSQL_ROOT_PASSWORD=
MYSQL_USER=
MYSQL_PASSWORD=
```

# SQL script
## Table structure
```SQL
CREATE TABLE IF NOT EXISTS shortened_urls (
    uuid INT AUTO_INCREMENT PRIMARY KEY,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```
- `uuid`: `INT AUTO_INCREMENT PRIMARY KEY` Unique identifier for each shortened URL.

- `original_url`: `TEXT NOT NULL` The original (long) URL to which the shortened URL redirects.

- `created_at`: `TIMESTAMP DEFAULT CURRENT_TIMESTAMP` Records the timestamp when the entry is created.

- `expires_at`: `TIMESTAMP NULL` (Optional) Specifies when the shortened URL will expire; can be NULL.

# Docker compose
The `compose.yml` file handles the MariaDB database. You need to have a `.env` file like said before to provide passwords.  
The SQL script is mount in a volume to the `docker-entrypoint` to populate the database.

## Run & destroy
```bash
docker compose up -d
```
```bash
docker exec -it tinygopher_mariadb mariadb -uroot -p
```
```bash
docker compose down
```