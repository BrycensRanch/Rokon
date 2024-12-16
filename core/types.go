package core

import "encoding/xml"

type Root struct {
	XMLName     xml.Name    `xml:"root"`
	SpecVersion SpecVersion `xml:"specVersion"`
	Device      Device      `xml:"device"`
}

type SpecVersion struct {
	Major int `xml:"major"`
	Minor int `xml:"minor"`
}

type Device struct {
	DeviceType       string      `xml:"deviceType"`
	FriendlyName     string      `xml:"friendlyName"`
	Manufacturer     string      `xml:"manufacturer"`
	ManufacturerURL  string      `xml:"manufacturerURL"`
	ModelDescription string      `xml:"modelDescription"`
	ModelName        string      `xml:"modelName"`
	ModelNumber      string      `xml:"modelNumber"`
	ModelURL         string      `xml:"modelURL"`
	SerialNumber     string      `xml:"-"` // Acknowledging this field exists yet shouldn't be serialized.
	UDN              string      `xml:"UDN"`
	IconList         IconList    `xml:"iconList"`
	ServiceList      ServiceList `xml:"serviceList"`
}

type IconList struct {
	Icons []Icon `xml:"icon"`
}

type Icon struct {
	MimeType string `xml:"mimetype"`
	Width    int    `xml:"width"`
	Height   int    `xml:"height"`
	Depth    int    `xml:"depth"`
	URL      string `xml:"url"`
}

type ServiceList struct {
	Services []Service `xml:"service"`
}

type Service struct {
	ServiceType string `xml:"serviceType"`
	ServiceID   string `xml:"serviceId"`
	ControlURL  string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
	SCPDURL     string `xml:"SCPDURL"`
}

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}
