package event

type Contract struct {
	Abi      string
	Event    string
	EventHex string
}

var (
	StartupContract = Contract{
		Event:    "Created",
		EventHex: "0x4185e1bc3b938a449a979ff25265998b2865ec30afd0f10d3060a55cdfe43a28",
	}
)
