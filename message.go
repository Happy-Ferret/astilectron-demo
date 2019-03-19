package main

import (

	"encoding/json"
	
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"

)

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "create":
		// Unmarshal payload
		var title string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &title); err != nil {
				payload = err.Error()
				return
			}
			insert(title)
		}
	}
	return
}