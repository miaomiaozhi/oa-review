package v1

type Application struct {
	Context      string `json:"Context"`
	ReviewStatus bool   `json:"ReviewStatus"`
}

type UserGetInfoResponse struct {
	Id           int64          `json:"Id"`
	Name         string         `json:"Name"`
	Applications []*Application `json:"Applications"`
	Priority     int32          `json:"Priority"`
}
