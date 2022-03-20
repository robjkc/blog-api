# https://blog.akbuluteren.com/blog/multi-stage-docker-build
FROM golang:1.17
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app .

FROM alpine
COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
