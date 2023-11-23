package domain

type BetterCatFact struct {
	BestFactEver  string `json:"bestFactEver"`
	UnixTimestamp int64  `json:"unixTimestamp"`
}

type ApiCatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// more idiomatic approach
// interface is used only in handlers, can be moved there
type BetterFactService interface {
	Get() (BetterCatFact, error)
}

// type FactApiService interface {
// 	Get() (ApiCatFact, error)
// }
