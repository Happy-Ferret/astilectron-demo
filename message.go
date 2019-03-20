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
			id := insert(title)
			payload = id
			return
		}
	case "delete":
		var id int
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &id); err != nil {
				payload = err.Error()
				return
			}
			delete(id)
		}
	case "getTodos":
		payload = Todos
		return
	}
	return
}