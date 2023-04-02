package contracts

type NodeAddRequest struct {
	NodeIp *string `json:"address"`
	NodeId *string `json:"node_id"`
}

type NodeDeleteRequest struct {
	NodeIp *string `json:"address"`
	NodeId *string `json:"node_id"`
}