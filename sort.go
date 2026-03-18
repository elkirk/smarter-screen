package dispatcher

const (
	Standard = "STANDARD"
	Special  = "SPECIAL"
	Rejected = "REJECTED"
)

// Sort dispatches packages into stacks based on physical properties.
//
// Stacks:
//   - STANDARD: not bulky or heavy
//   - SPECIAL: either bulky or heavy
//   - REJECTED: both bulky and heavy
//
// Definitions:
//   - Bulky: volume >= 1,000,000 cm³ or any dimension >= 150 cm
//   - Heavy: mass >= 20 kg
//
// Units are cm for dimensions and kg for mass.
func Sort(width, height, length, mass float64) string {
	volume := width * height * length
	bulky := volume >= 1_000_000 || width >= 150 || height >= 150 || length >= 150
	heavy := mass >= 20

	if bulky && heavy {
		return Rejected
	}
	if bulky || heavy {
		return Special
	}
	return Standard
}
