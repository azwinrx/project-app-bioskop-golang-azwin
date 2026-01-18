package dto

// Cinema DTOs
type CinemaResponse struct {
ID      int    `json:"id"`
Name    string `json:"name"`
City    string `json:"city"`
Address string `json:"address"`
}

type CinemaListResponse struct {
Cinemas    []CinemaResponse `json:"cinemas"`
Pagination Pagination       `json:"pagination"`
}
