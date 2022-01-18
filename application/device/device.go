package device

var DeviceCode = map[string]string{
	"CONSUMER": "c",
	"KEYBOARD": "k",
	"MOUSE":    "m",
}

var KeyCode = map[string]string{
	"JOG_PRESS":   "1p",
	"JOG_RELEASE": "1r",
	"SW1_PRESS":   "1p", //Alias
	"SW1_RELEASE": "1r", // Alias
	"SW2_PRESS":   "2p",
	"SW2_RELEASE": "2r",
	"SW3_PRESS":   "3p",
	"SW3_RELEASE": "3r",
	"SW4_PRESS":   "4p",
	"SW4_RELEASE": "4r",
	"SW5_PRESS":   "5p",
	"SW5_RELEASE": "5r",
	"SW6_PRESS":   "6p",
	"SW6_RELEASE": "6r",
	"SW7_PRESS":   "7p",
	"SW7_RELEASE": "7r",
	"SW8_PRESS":   "8p",
	"SW8_RELEASE": "8r",
	"SW9_PRESS":   "9p",
	"SW9_RELEASE": "9r",
	"JOG_CW":      "R+",
	"JOG_CCW":     "R-",
	"JOG":         "R", // Only for callbacks!
}