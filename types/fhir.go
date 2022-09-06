package types

type EhrRecord struct {
	ResourceType string `json:"resourceType"`
	Id           string `json:"id"`
	Meta         struct {
		Profile []string `json:"profile"`
	} `json:"meta"`
	Text struct {
		Status string `json:"status"`
		Div    string `json:"div"`
	} `json:"text"`
	Extension []struct {
		Extension []struct {
			Url         string `json:"url"`
			ValueCoding struct {
				System  string `json:"system"`
				Code    string `json:"code"`
				Display string `json:"display"`
			} `json:"valueCoding,omitempty"`
			ValueString string `json:"valueString,omitempty"`
		} `json:"extension,omitempty"`
		Url       string `json:"url"`
		ValueCode string `json:"valueCode,omitempty"`
	} `json:"extension"`
	Identifier []struct {
		Use  string `json:"use"`
		Type struct {
			Coding []struct {
				System  string `json:"system"`
				Code    string `json:"code"`
				Display string `json:"display"`
			} `json:"coding"`
			Text string `json:"text"`
		} `json:"type"`
		System string `json:"system"`
		Value  string `json:"value"`
	} `json:"identifier"`
	Active bool `json:"active"`
	Name   []struct {
		Family string   `json:"family"`
		Given  []string `json:"given"`
		Period struct {
			Start string `json:"start"`
		} `json:"period"`
	} `json:"name"`
	Telecom []struct {
		System string `json:"system"`
		Value  string `json:"value"`
		Use    string `json:"use,omitempty"`
	} `json:"telecom"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
	Address   []struct {
		Line       []string `json:"line"`
		City       string   `json:"city"`
		State      string   `json:"state"`
		PostalCode string   `json:"postalCode"`
		Country    string   `json:"country"`
		Period     struct {
			Start string `json:"start"`
		} `json:"period"`
	} `json:"address"`
}
