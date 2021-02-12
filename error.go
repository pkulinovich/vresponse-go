package response

type Error struct {
	Title   string `json:"title"`
	Details string `json:"details"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}
