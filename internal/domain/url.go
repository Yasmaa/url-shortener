package domain

type Url struct {
	LongUrl  string `json:"long_url""`
	ShortUrl string `json:"short_url"`
	Alias string `json:"alias"`

}


type URLRequest struct {
	LongUrl  string `json:"long_url""`
}

type AliasRequest struct {
	Alias  string `json:"alias""`
}