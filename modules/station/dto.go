package station

type Station struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type StationResponse struct {
	ID     int64  `json:"id"`
	UserId int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type StationListResponse struct {
	Stations []Station `json:"stations"`
}
