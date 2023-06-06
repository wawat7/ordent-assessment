# Ordent Assessment

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">



## Specification

- framework Gin Gonic
- database MongoDB 

## Requirement

- Docker
- Go 1.17 or newer
- makefiles (optional)

### How to run with Docker

- setup your env, if you want copy file .env.example to .env
```
cp .env.example .env
```
- run app with makefiles (optional)
```
make up-build
```
- run app with docker command
```
docker-compose up -d --build
```


## Swagger Documentation

access your app with url `http://{backend}/docs/index.html`

for example
```
http://localhost:4000/docs/index.html
```

## Concept Architecture App

<img src="https://i.ibb.co/JmnQ2Ls/Screen-Shot-2023-06-06-at-23-03-12.png">
