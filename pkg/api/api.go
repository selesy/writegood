package api

type Suggestion struct {
	SuggesterName string
	RuleName      string
	Description   string
	StartPos      int
	EndPos        int
	StartTok      int
	EndTok        int
}

type Suggester func(txt string) []Suggestion
