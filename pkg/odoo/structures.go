package odoo

import (
	"encoding/xml"
)

type OdooContext map[string]interface{}
type OdooReadResult []map[string]interface{}
type OdooNameGetResult [][]interface{}
type X2ManyResult map[string]map[int]string
type OdooSearchResult []int
type OdooFieldInfo struct {
	Description      string        `mapstructure:"string"`
	Type             string        `mapstructure:"type"`
	Relation         string        `mapstructure:"relation"`
	RelationField    string        `mapstructure:"relation_field"`
	Manual           bool          `mapstructure:"manual"`
	Required         bool          `mapstructure:"required"`
	Readonly         bool          `mapstructure:"readonly"`
	Selection        []interface{} `mapstructure:"selection"`
	Store            bool          `mapstructure:"store"`
	CompanyDependent bool          `mapstructure:"company_dependent"`
	Searchable       bool          `mapstructure:"searchable"`
	Sortable         bool          `mapstructure:"sortable"`
}
type OdooMetadataResult struct {
	ID int `mapstructure:"id"`
}
type OdooFieldsResult map[string]OdooFieldInfo
type OdooFieldsViewGetResult struct {
	Arch      string `mapstructure:"arch"`
	BaseModel string `mapstructure:"base_model"`
	Model     string `mapstructure:"model"`
	Name      string `mapstructure:"name"`
	Type      string `mapstructure:"type"`
	ViewID    int    `mapstructure:"view_id"`
}
type OdooResponseData struct {
	Name      string        `json:"name"`
	Message   string        `json:"message"`
	Debug     string        `json:"debug"`
	Arguments []interface{} `json:"arguments"`
	Context   OdooContext   `json:"context"`
}

type OdooUserResult struct {
	UID           int         `mapstructure:"uid"`
	IsAdmin       bool        `mapstructure:"is_admin"`
	UserContext   OdooContext `mapstructure:"user_context"`
	ServerVersion string      `mapstructure:"server_version"`
}
type OdooError struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    OdooResponseData `json:"data"`
}
type OdooResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Error   OdooError   `json:"error"`
	Result  interface{} `json:"result"`
	Cookies map[string]string
}
type OdooTree struct {
	XMLName xml.Name `xml:"tree"`
	Text    string   `xml:",chardata"`
	String  string   `xml:"string,attr"`
	Field   []struct {
		Text      string `xml:",chardata"`
		Name      string `xml:"name,attr"`
		String    string `xml:"string,attr"`
		OnChange  string `xml:"on_change,attr"`
		Modifiers string `xml:"modifiers,attr"`
		Invisible string `xml:"invisible,attr"`
		Class     string `xml:"class,attr"`
		Optional  string `xml:"optional,attr"`
		Widget    string `xml:"widget,attr"`
		Domain    string `xml:"domain,attr"`
		CanCreate string `xml:"can_create,attr"`
		CanWrite  string `xml:"can_write,attr"`
		Readonly  string `xml:"readonly,attr"`
		Options   string `xml:"options,attr"`
	} `xml:"field"`
}
