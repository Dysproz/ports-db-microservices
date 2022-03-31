FROM golang:1.17.1 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /clientAPI cmd/clientAPI.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /portDomainService cmd/portDomainService.go

FROM alpine:3.13.6 as clientAPI
WORKDIR /root/
COPY --from=builder /clientAPI ./
EXPOSE 5000
ENTRYPOINT ["./clientAPI", "-log-level", "debug"]

FROM alpine:3.13.6 as portDomainService
WORKDIR /root/
COPY --from=builder /portDomainService ./
EXPOSE 5000
ENTRYPOINT ["./portDomainService",  "-log-level", "debug"]
