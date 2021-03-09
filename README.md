# scaffold-gin-gonic

Scaffold to build API using gin gonic framework

## Usage

You must change the main package directory for example:

When you need use the settings package, the import sentence would be 
```sh
import "gin-gonic/pkg/setting"
```
Change for this:

```sh
import "your-dir/pkg/setting"
```

## Build - DEV

Run the following command:

```sh
docker build -t app-name -f dev.Dockerfile .
```

When the build is done, run:

```sh
docker run --rm -it -p port:port app-name go run main.go api
```
## Build

Run the following command:

```sh
docker build -t app-name .
```

When the build is done, run:

```sh
docker run --rm -it -p port:port app-name main api
```
## environment variables

```sh
ENV CLIENT_ID p@ssw0rd!
ENV CONN_STRING mongodb+srv://user:pass@127.0.0.1/bdname
ENV DB_HOST localhost
ENV DB_NAME dbname
ENV DB_PASS password
ENV DB_PORT 5432
ENV DB_TYPE postgres
ENV DB_USER user
ENV DEBUG true
ENV JWT_SECRET Aech7eepaesi8goo8phu8laech8aet4yie1kahsa4phohLuiHu9aeph6oa9Eoth7
ENV PAGE_SIZE 25
ENV PORT 8080
ENV READ_TIMEOUT 60
ENV REDIS_URL 172.0.17.2
ENV RUN_MODE debug
ENV WRITE_TIMEOUT 60
```

## Next steps

This scaffold will be upgraded to support gracefully restart to speed up the development process. The Dockerfile will change to compile the app and run the binary.  
