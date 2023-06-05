# Build stage
FROM golang:1.20.4-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

ARG MONGO_URI=''
ARG MONGO_DATABASE=''
ARG MONGO_POOL_MIN=10
ARG MONGO_POOL_MAX=100
ARG MONGO_MAX_IDLE_TIME_SECOND=60


RUN echo "MONGO_URI=${MONGO_URI}"  >> ".env"
RUN echo "MONGO_DATABASE=${MONGO_DATABASE}"  >> ".env"
RUN echo "MONGO_POOL_MIN=${MONGO_POOL_MIN}"  >> ".env"
RUN echo "MONGO_POOL_MAX=${MONGO_POOL_MAX}"  >> ".env"
RUN echo "MONGO_MAX_IDLE_TIME_SECOND=${MONGO_MAX_IDLE_TIME_SECOND}"  >> ".env"


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final stage
FROM alpine:3.14.2

RUN apk --no-cache add curl

WORKDIR /root/

COPY --from=build /app/app .
COPY --from=build /app/.env .

EXPOSE 4000

HEALTHCHECK --interval=20s --start-period=5s CMD curl -f localhost:4000/health

CMD ["./app"]