package portStorage

var PortMapping = make(map[int]string)

func AddToMap(port int, containerId string) {
	PortMapping[port] = containerId
}

func GetFromMap(port int) string {
	return PortMapping[port]
}