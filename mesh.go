package main

import (
	"fmt"
	"log"
	"time"
	"strings"
	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// MeshNetwork represents a network of devices that have meshed up.
type MeshNetwork struct {
	ID   string
	Pass string
	MeshRouter

	RootId	string				// Id of the root node
	Nodes map[string]*MeshNode
}

func (m *MeshNetwork) GetNode(nid string) (mn *MeshNode) {
	var e bool

	if mn, e = m.Nodes[nid]; !e {
		mn = &MeshNode{Self: nid}
	}
	return mn
}

func (m *MeshNetwork) UpdateRoot(rootid string) {

	// TODO create a fully configured node and schedule network topology updates.
	// log.Printf("%s.%s %s[%.0f]: rootid: %s, self: %s, parent: %s\n",
	//	addr, typ, msgtype, layer, rootid, self, parent);
	if (mesh.RootId != rootid) {
		// we have a change of roots
		log.Println("Root Node has changed from ", mesh.RootId, " to ", rootid)
		mesh.RootId = rootid
	}
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
	Parent   *MeshNode
	Children map[string]*MeshNode
	Layer    int

	Station

	Updated time.Time
}

func NewNode(d map[string]interface{}) *MeshNode {
	self := d["self"].(string)
	parent := d["parent"].(string)
	pnode := mesh.GetNode(parent)

	mn := &MeshNode{
		Self:    self,
		Parent:  pnode,
		Layer:   int(d["layer"].(float64)),
		Updated: time.Now(),
	}
	return mn
}

func (n *MeshNode) UpdateParent(p *MeshNode) {
	if (n.Parent != p) {
		log.Printf("n.Parent has changed from %s -> %s\n", n.Parent, p)
	}
	n.Parent = p		
}

func (n *MeshNode) UpdateChild(c *MeshNode) {
	log.Print("Parent ", n.Self)
	if (n.Children == nil) {
		n.Children = make(map[string]*MeshNode)
	}

	if _, e := n.Children[c.Self]; e {
		log.Println(" update existing child ")
	} else {
		log.Println(" ADDING NEW child ")
	}
	n.Children[c.Self] = c
	log.Println(c.Self)
}


func (n *MeshNode) String() string {
	str := fmt.Sprintf("NODE self - %s :=: parent - %s :=: layer - %d last update: %q\n",
		n.Self, n.Parent, n.Layer, n.Updated)
	if len(n.Children) < 1 {
		return str
	}
	str += "Chilren:\n"
	for _, n := range n.Children {
		str += "\t" + n.Self + "\n"
	}
	return str
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

func (mn MeshNetwork) MsgRecv(topic string, payload []byte) {
	
	var m ToCloudMsg
	err := json.Unmarshal(payload, &m)
	if err != nil {
		log.Fatal("Failed to unmarshal payload")
	}

	// unravel the json message and verify our current node information
	paths := strings.Split(topic, "/");
	if len(paths) != 3 {
		log.Fatal("Error unsupported path")
	}

	rootid := paths[1]
	//addr := m.Addr
	//typ  := m.Type
	data := m.Data
	msgtype, _ := data["type"]

	switch (msgtype) {
	case "heartbeat":
		
		self, _ := data["self"]
		parent, _ := data["parent"]
		layer, _ := data["layer"]

		mesh.UpdateRoot(rootid)
		node := mesh.GetNode(self.(string))
		if node == nil {
			log.Fatalln("GetNode returned nil for ", self)
		}

		if (node.Layer != layer) {
			log.Println("Node has changed layers from %d -> %d ", node.Layer, layer)
		}

		pnode := mesh.GetNode(parent.(string))
		if pnode == nil {
			log.Fatalln("GetParent returned nil for ", parent)
		}

		pnode.UpdateChild(node)
		node.UpdateParent(pnode)

	default:
		log.Fatalln("Unknown message type: ", msgtype)
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
