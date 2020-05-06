package professions

type Actor struct {
	*Employee    `json:"employee"`
	ScopesList   []Scope       `json:"scopes"`
	Achievements []Achievement `json:"achievements,omitempty"`
	Biography    string        `json:"biography,omitempty"`
}

type Scope string

const (
	Movies     Scope = "movies"
	Theatre    Scope = "theatre"
	Television Scope = "television"
	Radio      Scope = "radio"
)

type Achievement string

const (
	NationalArtist Achievement = "national artist"
	HonoredArtist  Achievement = "honored artist"
)
