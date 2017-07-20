package models

type Movie struct {
	ApiKey       string `json:api_key`
	Language     string `json:language`
	SortBy       string `json:sort_by`
	IncludeAdult bool   `json:include_adult`
	IncludeVideo bool   `json:include_video`
	Genres       string `json:with_genres`
	Page         int    `json:page`
}
