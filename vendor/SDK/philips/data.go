package lights

type Data struct {

}

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

type LightState struct {
	Index     string 
	Type      string `json:"type"`
	Modelid   string `json:"modelid"`
	Name      string `json:"name"`
	Swversion string `json:"swversion"`
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
}
