package httphandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2/client"

	querypoint "github.com/lucasalmeron/microc3/querypoints/pkg/querypoints"
	protoqp "github.com/lucasalmeron/microc3/querypoints/pkg/querypoints/proto"
	user "github.com/lucasalmeron/microc3/users/pkg/users"

	errorprovider "github.com/lucasalmeron/microc3/gateway/internal/helper"
)

var (
	queryPointsClient protoqp.QueryPointsService
)

type QueryPointsHandler struct{}

func InitQueryPointsHandler(router *mux.Router) {
	//create gRPC clients//
	queryPointsClient = protoqp.NewQueryPointsService("go.micro.service.querypoints", client.DefaultClient)

	handler := new(QueryPointsHandler)

	router.Path("/querypoints/list").HandlerFunc(handler.GetList).Methods(http.MethodGet, http.MethodOptions)
	router.Path("/querypoints/id/{queryPointID:[0-9a-fA-F]{24}}").HandlerFunc(handler.GetByID).Methods(http.MethodGet, http.MethodOptions)
	router.Path("/querypoints/name/{name}").HandlerFunc(handler.GetByName).Methods(http.MethodGet, http.MethodOptions)

	router.Path("/querypoints/paginated").HandlerFunc(handler.GetPaginated).Methods(http.MethodPost, http.MethodOptions)

	router.Path("/querypoints/create").HandlerFunc(handler.Create).Methods(http.MethodPost, http.MethodOptions)
	router.Path("/querypoints/update").HandlerFunc(handler.Update).Methods(http.MethodPut, http.MethodOptions)

}

func (h QueryPointsHandler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := queryPointsClient.GetList(context.TODO(), &empty.Empty{})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h QueryPointsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryPointID := mux.Vars(r)["queryPointID"]

	response, err := queryPointsClient.GetByID(context.TODO(), &protoqp.RequestQueryPointID{
		Id: queryPointID,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h QueryPointsHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
		name := mux.Vars(r)["name"]

		response, err := queryPointsClient

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
			return
		}
		json.NewEncoder(w).Encode(response)*/
}

func (h QueryPointsHandler) GetPaginated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	/* EXAMPLE
	{
		"pageNumber":1,
		"registersNumber":10,
		"filters":[
			{"field":"email","value":"luko.ar@gmail.com"}
		]
	}
	*/

	var pageOptions user.PageOptions

	if err := decoder.Decode(&pageOptions); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}
	var filters []*protoqp.RequestPageOptions_Filter
	if len(pageOptions.Filters) > 0 {
		for _, f := range pageOptions.Filters {
			filters = append(filters, &protoqp.RequestPageOptions_Filter{
				Field: f.Field,
				Value: f.Value,
			})
		}
	}

	response, err := queryPointsClient.GetPaginated(context.TODO(), &protoqp.RequestPageOptions{
		PageNumber:      pageOptions.PageNumber,
		RegistersNumber: pageOptions.RegistersNumber,
		OrderBy: &protoqp.RequestPageOptions_Filter{
			Field: pageOptions.OrderBy.Field,
			Value: pageOptions.OrderBy.Value,
		},
		Filters: filters,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h QueryPointsHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var querypoint querypoint.QueryPoint

	if err := decoder.Decode(&querypoint); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}

	response, err := queryPointsClient.Create(context.TODO(), &protoqp.RequestCreateQueryPoint{
		Name:       querypoint.Name,
		Phone:      querypoint.Phone,
		Address:    querypoint.Address,
		District:   querypoint.District,
		Department: querypoint.Department,
		Actions:    querypoint.Actions,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h QueryPointsHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var querypoint querypoint.QueryPoint

	if err := decoder.Decode(&querypoint); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}

	response, err := queryPointsClient.Update(context.TODO(), &protoqp.RequestUpdateQueryPoint{
		Id:         querypoint.ID,
		Name:       querypoint.Name,
		Phone:      querypoint.Phone,
		Address:    querypoint.Address,
		District:   querypoint.District,
		Department: querypoint.Department,
		Actions:    querypoint.Actions,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}
