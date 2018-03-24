package philips

// ColorHue struct is storage for phillips smart hue color lights
type ColorHue struct {
	Modelid   string `json:"modelid"`
	Name      string `json:"name"`
	Swversion string `json:"swversion"`
	State     struct {
		Xy        []float32  `json:"xy"`
		Ct        int    `json:"ct"`
		Alert     string `json:"alert"`
		Sat       int    `json:"sat"`
		Effect    string `json:"effect"`
		Bri       int    `json:"bri"`
		Hue       int    `json:"hue"`
		Colormode string `json:"colormode"`
		Reachable bool   `json:"reachable"`
		On        bool   `json:"on"`
	} `json:"state"`
	Type        string `json:"type"`
	Pointsymbol struct {
		Num1 string `json:"1"`
		Num2 string `json:"2"`
		Num3 string `json:"3"`
		Num4 string `json:"4"`
		Num5 string `json:"5"`
		Num6 string `json:"6"`
		Num7 string `json:"7"`
		Num8 string `json:"8"`
	} `json:"pointsymbol"`
	Uniqueid string `json:"uniqueid"`
	Index    string 
}

// LightGroup stores hue groups data
type LightGroup struct {
	Index	string
	Name  string `json:"name"`
	Lights  []string `json:"lights"`
	Type  string `json:"type"`
	Action LightAction `json:"action"`
}

// LightAction stores action elements of a group
type LightAction struct {
	On  bool `json:"on"`
	Bri  int `json:"bri"`
	Hue int `json:"hue"`
	Sat int `json:"sat"`
	Effect string `json:"effect"`
	Xy []float32 `json:"xy"`
	Ct int `json:"ct"`
	Alert string `json:"alert"`
	Colormode string `json:"colormode"`
}