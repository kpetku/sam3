package sam3

var (
	// Suitable options if you are shuffling A LOT of traffic. If unused, this 
	// will waste your resources.
	Options_Humongous = []string{"inbound.length=3", "outbound.length=3",
	"inbound.lengthVariance=1", "outbound.lengthVariance=1",
 	"inbound.backupQuantity=3", "outbound.backupQuantity=3",
	"inbound.quantity=6", "outbound.quantity=6"}
 	
 	// Suitable for shuffling a lot of traffic.
	Options_Fat = []string{"inbound.length=3", "outbound.length=3",
	"inbound.lengthVariance=1", "outbound.lengthVariance=1",
 	"inbound.backupQuantity=1", "outbound.backupQuantity=1",
	"inbound.quantity=4", "outbound.quantity=4"}

	// Suitable for shuffling medium amounts of traffic.
	Options_Medium = []string{"inbound.length=3", "outbound.length=3",
	"inbound.lengthVariance=1", "outbound.lengthVariance=1",
 	"inbound.backupQuantity=0", "outbound.backupQuantity=0",
	"inbound.quantity=2", "outbound.quantity=2"}

	// Suitable for small and quick dataflows.
	Options_Small = []string{"inbound.length=3", "outbound.length=3",
	"inbound.lengthVariance=1", "outbound.lengthVariance=1",
 	"inbound.backupQuantity=0", "outbound.backupQuantity=0",
	"inbound.quantity=1", "outbound.quantity=1"}

	// Does not use any anonymization, you connect directly to others tunnel
	// endpoints, thus revealing your identity but not theirs. Use this only
	// if you don't care.
	Options_Warning_ZeroHop = []string{"inbound.length=0", "outbound.length=0",
	"inbound.lengthVariance=0", "outbound.lengthVariance=0",
 	"inbound.backupQuantity=0", "outbound.backupQuantity=0",
	"inbound.quantity=2", "outbound.quantity=2"}
)
