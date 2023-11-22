package domain

type BetterCatFact struct {
	BestFactEver  string `json:"bestFactEver"`
	UnixTimestamp int64  `json:"unixTimestamp"`
}

type ApiCatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type BetterFactService interface {
	Get() (BetterCatFact, error)
}

type FactApiService interface {
	Get(ApiCatFact, error)
}
