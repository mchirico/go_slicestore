package pkg

import (
	"encoding/json"
	"log"
)

func Process(data []byte) (DS, error)  {

	s := DS{}
	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Printf("Can't Unmarshal in pkg.process")
	}

	return s,err
}
