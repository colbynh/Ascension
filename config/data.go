package config
// Config files need to be created per vendor.
// E.g. ~/ascension/philips/config.json

// VendorConfig is the master config struct that contains
// all configurations for all vendors.
type VendorConfig struct {
	Philips PhilipsConfig
}

// PhilipsConfig struct holds data for philips hue configurations.
type PhilipsConfig struct {
	BridgeIP string `json:"bridgeip"`
	APIKey string `json:"apikey"`
}