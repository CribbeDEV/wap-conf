package types

type WlcConf struct {
	Main      string
	Secondary string
}

type UserConfiguration struct {
	Location    string
	CountryCode string
	Site        string
	Wlc         WlcConf
}

type Configuration struct {
	OriginalName string
	NewName      string
	UserConfiguration
}

var WLC_HOSTS = map[string]string{
	"10.0.2.38":   "EMEAWLC01",
	"10.182.2.26": "APACWLC01",
	"10.0.2.200":  "WLC-CL",
}

type TemplateLocations struct {
	Main  string `json:"main"`
	Site  string `json:"site"`
	Reset string `json:"reset"`
}

type UserVariables struct {
	Templates       TemplateLocations
	OutputDirectory string
}
