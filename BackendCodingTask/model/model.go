package model

// Model to Tables
type Movies struct {
	Tconst         string  `json:"tconst,omitempty"`
	TitleType      string  `json:"titleType,omitempty"`
	PrimaryTitle   string  `json:"primarytitle,omitempty"`
	RuntimeMinutes int     `json:"runtimeminutes,omitempty"`
	Genres         string  `json:"genres,omitempty"`
	Averagerating  float64 `json:"averagerating,omitempty"`
	Numvotes       int     `json:"numvotes,omitempty"`
}
