package tododto

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Files       string `json:"files"`
	// Sublist     []models.SubListResponse `json:"sublist"`
}
