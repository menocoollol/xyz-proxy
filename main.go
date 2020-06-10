package main
import (
	"github.com/BurntSushi/toml"
	"inet.af/tcpproxy"
	"io/ioutil"
	"log"
)

type proxy struct {
	LocalPort string
	RemoteHost string
	RemotePort string
}
type proxies struct {
	Proxy []proxy
}

func main() {
	log.Println("Starting xyz-proxy...");

	// Create a proxy
	var p tcpproxy.Proxy

	// Register routes
	dat, err := ioutil.ReadFile("proxies.toml")
	if err != nil {
		log.Println("Error reading config: ")
		log.Fatal(err);
	}

	var pList proxies;
	if _, err := toml.Decode(string(dat), &pList); err != nil {
		log.Println("Error paring config: ")
		log.Fatal(err)
	}

	for _, e := range pList.Proxy {
		log.Println("Registering proxy :" + e.LocalPort + " to " +  e.RemoteHost + ":" + e.RemotePort)
		p.AddRoute(":" + e.LocalPort, tcpproxy.To(e.RemoteHost + ":" + e.RemotePort))
	}

	// Start the proxy
	log.Println("Liftoff 🚀");
	log.Fatal(p.Run())

}
