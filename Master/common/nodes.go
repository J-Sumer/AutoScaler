package common

type Node struct {
	NodeIp *string `json:"node_ip"`
	NodeId *string `json:"node_id"`
}

func AddNode(n Node) error {
	NodeList[*n.NodeIp] = n
	return nil
}

func DeleteNode(n Node) error {
	delete(NodeList, *n.NodeIp)
	return nil
}
