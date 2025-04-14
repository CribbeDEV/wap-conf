package types

import "os"

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

func (w WlcConf) GetMainWLC() string {
	name, ok := WLC_HOSTS[w.Main]
	if !ok {
		panic("Main WLC ip not known")
	}
	return name + " " + w.Main
}

func (w WlcConf) GetSecondaryWLC() (string, bool) {
	name, ok := WLC_HOSTS[w.Secondary]
	if !ok {
		return "", false
	}
	return name + " " + w.Secondary, true
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

func (t TemplateLocations) GetMain() string {
	if _, err := os.Stat(t.Main); os.IsNotExist(err) {
		// Return link to RAW
		return "https://raw.githubusercontent.com/CribbeDEV/wap-conf/refs/heads/main/templates/main_template.txt"
	} else {
		return t.Main
	}
}
func (t TemplateLocations) GetSite() string {
	if _, err := os.Stat(t.Site); os.IsNotExist(err) {
		// Return link to RAW
		return "https://raw.githubusercontent.com/CribbeDEV/wap-conf/refs/heads/main/templates/site_template.txt"
	} else {
		return t.Site
	}
}
func (t TemplateLocations) GetReset() string {
	if _, err := os.Stat(t.Reset); os.IsNotExist(err) {
		// Return link to RAW
		return "https://raw.githubusercontent.com/CribbeDEV/wap-conf/refs/heads/main/templates/ap_reset.txt"
	} else {
		return t.Reset
	}
}

type UserVariables struct {
	Templates       TemplateLocations
	OutputDirectory string
}
