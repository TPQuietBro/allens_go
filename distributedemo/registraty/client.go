package registraty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)

	if err != nil {
		return err
	}

	res, err := http.Post(ServicesUrl, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service: %v, code: %d", r.ServiceName, res.StatusCode)
	}
	return nil
}
