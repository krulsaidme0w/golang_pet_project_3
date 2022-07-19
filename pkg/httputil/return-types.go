package httputil

type Error struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type ReturnType struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
