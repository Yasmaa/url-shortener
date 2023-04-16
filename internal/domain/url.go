package domain

type URL struct {
	LongUrl  string `json:"long_url""`
	Alias string `json:"alias"`

}


type URLRequest struct {
	LongUrl  string `json:"long_url""`
}

type AliasRequest struct {
	Alias  string `json:"alias""`
}