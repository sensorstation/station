package main

import (
	"fmt"
	"log"
	"time"

	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// MeshNetwork represents a network of devices that have meshed up.
type MeshNetwork struct {
	ID   string
	Pass string
	MeshRouter
	Nodes map[string]*MeshNode
}

func (m *MeshNetwork) GetNode(nid string) (mn *MeshNode) {
	var e bool

	if mn, e = m.Nodes[nid]; !e {
		mn = &MeshNode{Self: nid}
	}
	return mn
}

// MeshRouter is the optional IP router for the mesh network
type MeshRouter struct {
	SSID string
	Pass string
	Host string
}

// MeshNode represents a single node in the ESP-MESH network, this allows us
// to keep track of our inventory fleet.
type MeshNode struct {
	Self     string
	Parent   string
	Children []string
	Layer    int

	Station

	Updated time.Time
}

type MeshMessage struct {
	Addr string `json:"addr"`
	Typ  string `json:"type"`
	Data []byte `json:"data"`
}

type MeshHeartBeat struct {
	Typ    string `json:"type"`   // heartbeat
	Self   string `json:"self"`   // macaddr of advertising node
	Parent string `json:"parent"` // macaddr of parent
	Layer  int    `json:"layer"`  // node layer
}

func NewNode(d map[string]interface{}) *MeshNode {
	self := d["self"].(string)
	parent := d["parent"].(string)
	mn := &MeshNode{
		Self:    self,
		Parent:  parent,
		Layer:   int(d["layer"].(float64)),
		Updated: time.Now(),
	}
	return mn
}

func (n *MeshNode) String() string {
	str := fmt.Sprintf("NODE self - %s :=: parent - %s :=: layer - %d last update: %q\n",
		n.Self, n.Parent, n.Layer, n.Updated)
	if len(n.Children) < 1 {
		return str
	}
	str += "Chilren:\n"
	for _, n := range n.Children {
		str += "\t" + n + "\n"
	}
	return str
}

// This handles the mesh networking part of this system
func mesh_subscribe(client MQTT.Client) {
	topic := "mesh/+/toCloud"
	qos := 0
	if token := client.Subscribe(topic, byte(qos), onMessageToCloud); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("subscribed to 'mesh/+/tocloud'")
}

// onMessageToCloud is call everytime a control message is sent
// to the mesh/+/tocloud channel where the '+' wildcard represents
// the station ID of the sender. The msg["data"] field tells us
// the type of data message (heartbeat).
func onMessageToCloud(client MQTT.Client, message MQTT.Message) {
	if config.Debug {
		log.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	}

	payload := message.Payload()
	var msg map[string]interface{} = make(map[string]interface{})
	err := json.Unmarshal(payload, &msg)
	if err != nil {
		log.Fatal(err)
	}

	data, e := msg["data"]
	if !e {
		log.Fatal("Unknown message format, expected (type): ", msg["data"])
	}

	d := data.(map[string]interface{})
	t, e := d["type"]
	if !e {
		log.Fatal("Unknown message format, expected (type): ", data)
	}

	switch t {
	case "heartbeat":
		mesh.Heartbeat(d)
	default:
		log.Println("Unknown message type ", t)
	}
}

// Heartbeat records the pulse recent sent from the respective station
// if this is the first time recording the station it will be inserted
// with a timer, every new message updates the timer. A cleanup timer
// periodically runs cleaning up all aged stations..
//
// TODO The timeout timer
func (mn MeshNetwork) Heartbeat(data map[string]interface{}) {
	node := NewNode(data)
	if n, e := mn.Nodes[node.Self]; !e {
		mn.Nodes[n.Self] = n
		log.Printf("We have a new node: %s", n.Self)
	} else {
		n.Updated = time.Now()
	}
	fmt.Println(node.String())
}
