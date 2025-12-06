package utils

type Key []string

var (
	KeyUp    Key = []string{"A", "w", "up"}
	KeyDown  Key = []string{"B", "s", "down"}
	KeyRight Key = []string{"C", "d", "right"}
	KeyLeft  Key = []string{"D", "a", "left"}
	Enter    Key = []string{"enter"}
	Space    Key = []string{" "}
)

func KeyMatchesInput(input string, key Key) bool {
	matches := false
	for _, v := range key {
		if v == input {
			return true
		}
	}

	return matches
}
