package filmdto

type FilmResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"image" validate:"required"`
	Year          string `json:"year" form:"year"`
	LinkFilm      string `json:"linkFilm"`
	CategoryID    string `json:"category_id"`
	Category      string `json:"category"`
	Description   string `json:"description" form:"description" validate:"required"`
}
