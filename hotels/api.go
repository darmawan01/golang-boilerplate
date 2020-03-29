package hotels

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type HotelsApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *HotelsApi) Register() {
	api.Router.Handle("/hotels", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/hotels", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/hotels/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/hotels/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/hotels/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("HotelsApi registered")
}

func (api *HotelsApi) all(w http.ResponseWriter, r *http.Request) {
	hotels, err := api.allHandler()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", hotels, 0, 0, 0),
	)
}

func (api *HotelsApi) add(w http.ResponseWriter, r *http.Request) {
	body, err := utils.BodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}
	hotel := body.(Hotel)
	if ok, err := utils.ValidateStruct(hotel); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil))
		return

	}

	lastInsertedId, err := api.addHandler(hotel)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return

	}
	utils.RespondwithJSON(
		w,
		http.StatusCreated,
		map[string]interface{}{
			"lastInsertedId": lastInsertedId,
		},
	)

}

func (api *HotelsApi) detail(w http.ResponseWriter, r *http.Request) {

	hotelId := utils.GetIDParam(r)
	if utils.IsEmpty(hotelId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required hotel_id as param", nil),
		)
	}

	id, err := strconv.Atoi(hotelId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of hotel_id is invalid", nil),
		)
	}

	hotel, err := api.detailHandler(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", hotel, 0, 0, 0),
	)

}

func (api *HotelsApi) update(w http.ResponseWriter, r *http.Request) {
	hotelId := utils.GetIDParam(r)
	if utils.IsEmpty(hotelId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required hotel_id as param", nil),
		)
	}

	id, err := strconv.Atoi(hotelId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of hotel_id is invalid", nil),
		)
	}

	body, err := utils.BodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}
	updated := body.(Hotel)
	if ok, err := utils.ValidateStruct(updated); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return

	}

	hotel, err := api.detailHandler(id)

	if !utils.IsEmpty(updated.Name) {
		hotel.Name = updated.Name
	}

	if !utils.IsEmpty(updated.Address) {
		hotel.Address = updated.Address
	}

	if !utils.IsEmpty(updated.Latitute) {
		hotel.Latitute = updated.Latitute
	}

	if !utils.IsEmpty(updated.Longitude) {
		hotel.Longitude = updated.Longitude
	}

	if err = api.updateHandler(hotel); err != nil {
		if !utils.IsEqual(err.Error(), "failed-to-update") {
			utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
			return
		}
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, nil)
}

func (api *HotelsApi) delete(w http.ResponseWriter, r *http.Request) {
	hotelId := utils.GetIDParam(r)
	if utils.IsEmpty(hotelId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required hotel_id as param", nil),
		)
	}

	id, err := strconv.Atoi(hotelId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of hotel_id is invalid", nil),
		)
	}

	if err = api.deleteHandler(id); err != nil {
		if !utils.IsEqual(err.Error(), "failed-to-delete") {
			utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
			return
		}
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, nil)
}
