package controllers

import (
	"net/http"
	"testing"

	"github.com/allisonverdam/best-credit-card/daos"
	"github.com/allisonverdam/best-credit-card/services"
	"github.com/allisonverdam/best-credit-card/testdata"
)

func TestCard(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeCardResource(&router.RouteGroup, services.NewCardService(daos.NewCardDAO()))

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	numberRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"number","error":"cannot be blank"}]}`

	runAPITests(t, router, []apiTestCase{
		{"t1 - get an card", "GET", "/cards/1", "", http.StatusOK, `{			"id": 1,			"number": "1234123412341230",			"due_date": "2016-10-12T00:00:00Z",			"expiration_date": "2017-10-12T00:00:00Z",			"cvv": 123,			"limit": 1000,			"person_id": 1	}`},
		{"t1 - get an card", "GET", "/cards/a", "", http.StatusInternalServerError, ""},
		{"t2 - get a nonexisting card", "GET", "/cards/99999", "", http.StatusNotFound, notFoundError},
		{"t3 - create an card", "POST", "/cards", `	{	"number": "1234123412341234",	"due_date": "2016-10-08T00:00:00Z",	"expiration_date": "2017-10-12T00:00:00Z",			"cvv": 123,			"limit": 1000,			"person_id": 1	}`, http.StatusOK, `{    "id": 11,    "number": "1234123412341234",    "due_date": "2016-10-08T00:00:00Z",    "expiration_date": "2017-10-12T00:00:00Z",    "cvv": 123,    "limit": 1000,    "person_id": 1}`},
		{"t4 - create an card with validation error", "POST", "/cards", `{"number":""}`, http.StatusBadRequest, numberRequiredError},
		{"t5 - update an card", "PUT", "/cards/2", `{"number":"1234123412341234", "person_id": 1}`, http.StatusOK, `{"expiration_date":"2017-10-12T00:00:00Z", "cvv":123, "limit":1000, "person_id":1, "id":2, "number":"1234123412341234", "due_date":"2016-10-11T00:00:00Z"}`},
		{"t6 - update an card with validation error", "PUT", "/cards/2", `{"number":""}`, http.StatusBadRequest, numberRequiredError},
		{"t7 - update a nonexisting card", "PUT", "/cards/99999", "{}", http.StatusNotFound, notFoundError},
		{"t8 - delete an card", "DELETE", "/cards/2", ``, http.StatusOK, `{"cvv":123, "limit":1000, "person_id":1, "id":2, "number":"1234123412341234", "due_date":"2016-10-11T00:00:00Z", "expiration_date":"2017-10-12T00:00:00Z"}`},
		{"t9 - delete a nonexisting card", "DELETE", "/cards/99999", "", http.StatusNotFound, notFoundError},
		{"t10 - get a list of cards", "GET", "/cards?page=3&per_page=2", "", http.StatusOK, `{
			"page": 3,
			"per_page": 2,
			"page_count": 5,
			"total_count": 10,
			"items": [{
					"id": 6,
					"number": "1234123412341235",
					"due_date": "2016-10-07T00:00:00Z",
					"expiration_date": "2017-10-12T00:00:00Z",
					"cvv": 123,
					"limit": 1000,
					"person_id": 2
				},
				{
					"number": "1234123412341236",
					"due_date": "2016-10-06T00:00:00Z",
					"expiration_date": "2017-10-12T00:00:00Z",
					"cvv": 123,
					"limit": 1000,
					"person_id": 2,
					"id": 7
				}
			]
		}`},
	})
}
