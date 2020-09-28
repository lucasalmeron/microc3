package errorprovider

import (
	"encoding/json"
	"log"
)

type HttpError struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type gRPCError struct {
	ID     string `json:"id"`
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func ConvertToJSON(gRPCErr error) interface{} {
	rawIn := json.RawMessage(gRPCErr.Error())
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		log.Print(err)
	}

	var e gRPCError
	json.Unmarshal(bytes, &e)

	return e
}
