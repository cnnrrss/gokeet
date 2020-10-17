package swagger

import (
	"encoding/json"
	"io/ioutil"
)

type File struct {
	version string
}

func Read() (map[string]interface{}, error) {
	var data map[string]interface{}

	raw, err := ioutil.ReadFile("./examples/mockSwagger.json")
	if err != nil {
		return data, err
	}


	if err = json.Unmarshal(raw, &data); err != nil {
		return data, err
	}

	return data, nil
}
