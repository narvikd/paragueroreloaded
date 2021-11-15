package andalu

type Andalu struct {
	Spanish string `json:"spanish"`
	Andaluh string `json:"andaluh"`
	Rules   struct {
		Vaf         string `json:"vaf"`
		Vvf         string `json:"vvf"`
		EscapeLinks bool   `json:"escapeLinks"`
	} `json:"rules"`
}
