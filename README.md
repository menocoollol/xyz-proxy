# xyz-proxy
Supply a list of ports to proxy, and you'll have proxied ports! This is a simple TCP proxy written in Go that will 
map your local ports to some remote ip+port.

> I started learning Go, and I needed a simple TCP proxy, and so I wrote this little thing!

## Usage
Get dependencies:
```bash
go get inet.af/tcpproxy
go get github.com/BurntSushi/toml
```
Compile:
```bash
go build -o xyz-proxy main.go 
```

Edit the proxies.toml file. Here are the example contents:
```toml
[[proxy]]
localport = "8080"
remotehost = "192.168.0.134"
remoteport = "80"

[[proxy]]
localport = "8081"
remotehost = "google.com"
remoteport = "80"
```

(This will forward local port 8080 to 192.168.0.134:80 and local port 8081 to google.com:80)

Run:
```bash
./xyz-proxy
```