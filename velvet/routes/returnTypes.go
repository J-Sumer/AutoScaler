package routes

type CreateContainerReturn struct {
	ID string `json:"container_id"`
	Port string `json:"port"`
}