package main

import (
	"fmt"
	"log"
	"os"

	"github.com/processone/gox/xmpp"
)

func main() {
	options := xmpp.Options{Address: "localhost:5222", Jid: "id@localhost", Password: "123456", PacketLogger: os.Stdout}

	var client *xmpp.Client
	var err error
	if client, err = xmpp.NewClient(options); err != nil {
		log.Fatal("Error: ", err)
	}

	var session *xmpp.Session
	if session, err = client.Connect(); err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println("Stream opened, we have streamID = ", session.StreamId)

	// Iterator to receive packets coming from our XMPP connection
	for packet := range client.Recv() {
                text := ""
		switch packet := packet.(type) {
		case *xmpp.ClientMessage:
			fmt.Fprintf(os.Stdout, "Body = %s - from = %s\n", packet.Body, packet.From)
                        text=packet.Body
                        if(packet.Body=="cmd"){
                           text="cmd too"
                        }
			reply := xmpp.ClientMessage{Packet: xmpp.Packet{To: packet.From}, Body: text}
			client.Send(reply.XMPPFormat())
		default:
			fmt.Fprintf(os.Stdout, "Ignoring packet: %T\n", packet)
		}
	}
}
