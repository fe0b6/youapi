package youapi

type Api struct {
	AccessToken string
}

type TokenData struct {
	Code         string
	ClientId     string
	ClientSecret string
	RedirectUri  string
	RefreshToken string
}

type GetTokenAns struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type GetVideoInfoAns struct {
	Items []GetVideoInfoAnsItems `json:"items"`
}

type GetVideoInfoAnsItems struct {
	Id         string                         `json:"id"`
	Snippet    GetVideoInfoAnsItemsSnippet    `json:"snippet"`
	Statistics GetVideoInfoAnsItemsStatistics `json:"statistics"`
}

type GetVideoInfoAnsItemsSnippet struct {
	PublishedAt  string `json:"publishedAt"`
	ChannelId    string `json:"channelId"`
	ChannelTitle string `json:"channelTitle"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type GetVideoInfoAnsItemsStatistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	DislikeCount  string `json:"dislikeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}
