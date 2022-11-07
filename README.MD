# :100: Service counter.

## Overview 🔍
🌐 Web application to work with counter object :100:. Designed with Golang (net/http package).

## 🏗️ Install & Run🏃

Download or copy repository:

```bash
git clone https://github.com/Pythonyan3/counter.git
```

Run service without docker:
```bash
# build up exectable file
go build -o ./cmd/main ./cmd/main.go
# run service
./cmd/main
```

Using docker:

```bash
# run only counter service
docker build . -t counter
docker container run -p 8000:8080 -it counter

# run counter service and postgres db service
docker-compose build
docker-compose up

# run counter service with docker swarm
docker swarm init
docker stack deploy -c docker-compose.yaml counter
```

## API Doc 📚

Service allow to work with ``Counter`` entity.

### List of API endpoints:

1. `/ (GET)` - retrieve counter current value;
2. `/stat (GET)` - increment counter value by 1 and retrieve new value;
3. `/about (GET)` - retrieve static html page with author name.

### Some examples of usage

#### `/` (GET) - response example:
```json
{
	"counter": 0
}
```

#### `/stat` (GET) - response example:
```json
{
	"counter": 1
}
```

#### `/about` (GET) - response example:
```
<h3> Hello, Vitalii Manoilo.</h3>
```