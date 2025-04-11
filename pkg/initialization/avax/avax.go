package avax

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/pubsub"
	"github.com/gorilla/websocket"
	"github.com/qiniu/x/log"
)

func Init() {
	var count = 0
	for {
		log.Info("------------Avax Init:", count)
		ListenForAvax()
		count++
		time.Sleep(5 * time.Second)
	}
}

func ListenForAvax() (err error) {
	dialer := websocket.Dialer{
		NetDial: func(netw, addr string) (net.Conn, error) {
			return net.Dial(netw, addr)
		},
	}

	const avaxPubApi = "wss://api.avax-test.network/ext/bc/X/events"
	log.Info("------------avax.Dial: ", avaxPubApi)
	httpHeader := http.Header{}
	conn, _, err := dialer.Dial(avaxPubApi, httpHeader)
	if err != nil {
		log.Warn("------------", err)
		return
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	readMsg := func() {
		defer waitGroup.Done()

		for {
			log.Info("------------conn.ReadMessage()...")
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				log.Info("------------", err)
				return
			}
			log.Info("------------messageType:", mt, "   msg:", string(msg))
			switch mt {
			case websocket.TextMessage:
				log.Info("------------", string(msg))
			default:
				log.Info("------------", mt, string(msg))
			}
		}
	}

	log.Info("------------before readMsg()")
	go readMsg()
	log.Info("------------after readMsg()")

	cmd := &pubsub.Command{NewSet: &pubsub.NewSet{}}
	cmdmsg, err := json.Marshal(cmd)
	if err != nil {
		log.Warn("------------", err)
		return
	}
	log.Info("------------cmdmsg:", string(cmdmsg))
	err = conn.WriteMessage(websocket.TextMessage, cmdmsg)
	if err != nil {
		log.Warn("------------", err)
		return
	}

	var addresses []string
	//addresses = append(addresses, "0x7E94572BCc67B6eDa93DBa0493b681dC0ae9E964")
	addresses = append(addresses, "X-fuji193h7kk79amswl697lhuexpef282q24khlxfgrm")
	//addresses = append(addresses, "X-fuji132sa7p9nmv6gx5qg45l777kr6ct0cjkzz2vpz")

	cmd = &pubsub.Command{AddAddresses: &pubsub.AddAddresses{JSONAddresses: api.JSONAddresses{Addresses: addresses}}}
	cmdmsg, err = json.Marshal(cmd)
	if err != nil {
		log.Warn(err)
		return
	}
	log.Info("------------address cmdmsg:", string(cmdmsg))

	err = conn.WriteMessage(websocket.TextMessage, cmdmsg)
	if err != nil {
		log.Warn("------------", err)
		return
	}
	log.Info("------------conn.WriteMessage done")

	waitGroup.Wait()
	log.Info("------------waitGroup.Wait() after")
	conn.Close()
	log.Info("------------conn.Close() after")
	return
}
