package device

var Command = map[string]string{
	"ASSIGN":  "A",
	"DISPLAY": "D",
	"LED":     "L",
	"REFRESH": "R",
	"INFO":    "I",
}

// func Event(device, keycode, value) {
//     if device == DELAY:
//         return device + str(keycode)
//     elif type(value) == int:
//         return device.value + str(keycode.value) + "i" + str(value)
//     elif type(value) == ActionCode:
//         return device.value + str(keycode.value) + value.value
//     else:
//         return device.value + str(keycode.value)
// }

// #DelayCode
// DELAY = "d"