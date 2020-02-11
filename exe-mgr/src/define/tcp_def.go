package define

type LoginMsg struct {
	Code    int
	Data    string
	UpUrl   string `json:"upUrl"`
	DownUrl string `json:"downUrl"`
}
