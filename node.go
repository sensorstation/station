package station

var (
	mesh Mesh
)

type Mesh struct {
	Root	string				// root id
	Nodes	map[string]*Node
}

func (m *Mesh) GetNode(nid string) (n *Node) {
	var ex bool
	if n, ex = m.Nodes[nid]; !ex {
		n = &Node{nid, "", nil}
	}
	return n
}

func (m *Mesh) String() (s string) {
	s = "Root: " + m.Root + " nodes: "
	for _, n := range m.Nodes {
		s += "\t" + n.String() + "\n"
	}
	return s
}

type Node struct {
	Id		string
	Parent	string
	Children []string
}

func (n *Node) String() (s string) {
	s = "Node: " + n.Id + " parent: " + n.Parent + " children: "
	for _, c := range n.Children {
		s += c + ", "
	}
	return s
}
