package models

type WordDefinition []struct {
	Meta struct {
		ID        string   `json:"id"`
		UUID      string   `json:"uuid"`
		Sort      string   `json:"sort"`
		Src       string   `json:"src"`
		Section   string   `json:"section"`
		Stems     []string `json:"stems"`
		Offensive bool     `json:"offensive"`
	} `json:"meta"`
	Hom int `json:"hom,omitempty"`
	Hwi struct {
		Hw  string `json:"hw"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
	} `json:"def"`
	Et       [][]string `json:"et,omitempty"`
	Date     string     `json:"date,omitempty"`
	Shortdef []string   `json:"shortdef"`
	Ins      []struct {
		Ifc string `json:"ifc"`
		If  string `json:"if"`
	} `json:"ins,omitempty"`
	Uros []struct {
		Ure string `json:"ure"`
		Prs []struct {
			Mw    string `json:"mw"`
			Sound struct {
				Audio string `json:"audio"`
				Ref   string `json:"ref"`
				Stat  string `json:"stat"`
			} `json:"sound"`
		} `json:"prs"`
		Fl string `json:"fl"`
	} `json:"uros,omitempty"`
}
