package rooms

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type RoomsApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *RoomsApi) Register() {
	api.Router.Handle("/order-guests", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/order-guests", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/order-guests/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("RoomsApi registered")
}

func (api *RoomsApi) all(w http.ResponseWriter, r *http.Request) {
	rooms, err := api.allHandler()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", rooms, 0, 0, 0),
	)
}

func (api *RoomsApi) add(w http.ResponseWriter, r *http.Request) {
	body, err := utils.BodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}
	room := body.(Room)
	if ok, err := utils.ValidateStruct(room); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil))
		return

	}

	lastInsertedId, err := api.addHandler(room)
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

func (api *RoomsApi) detail(w http.ResponseWriter, r *http.Request) {

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

	room, err := api.detailHandler(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", room, 0, 0, 0),
	)

}

func (api *RoomsApi) update(w http.ResponseWriter, r *http.Request) {
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
	updated := body.(Room)
	if ok, err := utils.ValidateStruct(updated); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return

	}

	room, err := api.detailHandler(id)

	if !utils.IsEmpty(updated.Name) {
		room.Name = updated.Name
	}

	if updated.Quantity != room.Quantity {
		room.Quantity = updated.Quantity
	}

	if updated.Price != room.Price {
		room.Price = updated.Price
	}

	if updated.HotelId != room.HotelId {
		room.HotelId = updated.HotelId
	}

	if err = api.updateHandler(room); err != nil {
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

func (api *RoomsApi) delete(w http.ResponseWriter, r *http.Request) {
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
