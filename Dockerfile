FROM ubuntu:latest
WORKDIR /go/src/github.com/meloalright/guora
COPY . .
RUN apt-get update && apt-get -y install redis-server --no-install-recommends && rm -rf /var/lib/apt/lists/*
CMD sh -c '/etc/init.d/redis-server start && mkdir -p /etc/guora && cp ./configuration.example.yaml /etc/guora/configuration.yaml && ./guora -init'
