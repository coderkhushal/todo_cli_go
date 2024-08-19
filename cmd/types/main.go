package types

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Reset  = "\033[0m" // Use "\033[0m" to reset all attributes, not "\033[37m"
)

type TodoData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}
