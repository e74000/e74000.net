package main

import (
	"syscall/js"
)

var (
	config = map[string]interface{}{
		"r": map[string]interface{}{
			"name":    "red:",
			"min":     0.0,
			"max":     1.0,
			"initial": 0.0,
			"step":    0.001,
		},
		"g": map[string]interface{}{
			"name":    "green:",
			"min":     0.0,
			"max":     1.0,
			"initial": 0.0,
			"step":    0.001,
		},
		"b": map[string]interface{}{
			"name":    "blue:",
			"min":     0.0,
			"max":     1.0,
			"initial": 0.0,
			"step":    0.001,
		},
	}

	data = map[string]float64{
		"r": 0.0,
		"g": 0.0,
		"b": 0.0,
	}
)

func updateValue(this js.Value, params []js.Value) any {
	id := params[0].String()
	val := params[1].Float()

	data[id[1:]] = val

	document := js.Global().Get("parent").Get("document")

	oId := ""

	switch id[0] {
	case 'n':
		oId = "s" + id[1:]
	case 's':
		oId = "n" + id[1:]
	}

	other := document.Call("getElementById", oId)
	other.Set("value", val)

	return nil
}

func newSlider(data map[string]interface{}, key string) (slider js.Value, number js.Value) {
	document := js.Global().Get("parent").Get("document")

	slider = document.Call("createElement", "input")
	slider.Call("setAttribute", "type", "range")
	slider.Call("setAttribute", "min", data["min"])
	slider.Call("setAttribute", "max", data["max"])
	slider.Call("setAttribute", "value", data["initial"])
	slider.Call("setAttribute", "step", data["step"])
	slider.Call("setAttribute", "id", "s"+key)
	slider.Call("setAttribute", "class", "optionSlider")

	number = document.Call("createElement", "input")
	number.Call("setAttribute", "type", "number")
	number.Call("setAttribute", "value", data["initial"])
	number.Call("setAttribute", "id", "n"+key)
	number.Call("setAttribute", "class", "optionNumber")

	slider.Call("setAttribute", "oninput", "updateValue(this.id, parseFloat(this.value))")
	number.Call("setAttribute", "onchange", "updateValue(this.id, parseFloat(this.value))")

	return slider, number
}

func addSliders() {
	document := js.Global().Get("parent").Get("document")
	optionPanel := document.Call("getElementById", "options")

	js.Global().Get("parent").Set("updateValue", js.FuncOf(updateValue))

	for key, i := range config {
		if val, ok := i.(map[string]interface{}); ok {
			label := document.Call("createElement", "label")
			label.Call("setAttribute", "for", key)
			label.Call("setAttribute", "class", "optionLabel")
			label.Call("appendChild", document.Call("createTextNode", val["name"]))

			slider, number := newSlider(val, key)

			option := document.Call("createElement", "div")
			option.Call("setAttribute", "class", "option")

			inputBox := document.Call("createElement", "div")
			inputBox.Call("setAttribute", "class", "inputBox")

			inputBox.Call("appendChild", number)
			inputBox.Call("appendChild", slider)

			option.Call("appendChild", label)
			option.Call("appendChild", inputBox)
			optionPanel.Call("appendChild", option)
		}

	}
}

func main() {
	addSliders()      // Create Sliders based on date in config
	_main()           // Main can now use the data global variable
	<-make(chan bool) // Stop
}
