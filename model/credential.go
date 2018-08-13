package model

type Credential struct {
	Application Application `json:"application"`
	User        User        `json:"user"`
}
