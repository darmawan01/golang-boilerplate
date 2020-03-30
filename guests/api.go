package guests

import (
	"kodingworks/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
)

type GuestsApi struct {
	Router *mux.Router
	Db     *pgx.ConnPool
}

func (api *GuestsApi) Register() {
	api.Router.Handle("/guests", http.HandlerFunc(api.all)).Methods("GET")
	api.Router.Handle("/guests", http.HandlerFunc(api.add)).Methods("POST")
	api.Router.Handle("/guests/{id}", http.HandlerFunc(api.detail)).Methods("GET")
	api.Router.Handle("/guests/{id}", http.HandlerFunc(api.update)).Methods("PUT")
	api.Router.Handle("/guests/{id}", http.HandlerFunc(api.delete)).Methods("DELETE")

	log.Println("GuestsApi registered")
}

func (api *GuestsApi) all(w http.ResponseWriter, r *http.Request) {
	guests, err := api.allHandler()
	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
		return
	}

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", guests, 0, 0, 0),
	)
}

func (api *GuestsApi) add(w http.ResponseWriter, r *http.Request) {
	guest, err := api.bodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Invalid body", nil),
		)
		return
	}

	if ok, err := utils.ValidateStruct(guest); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil))
		return

	}

	lastInsertedId, err := api.addHandler(guest)
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

func (api *GuestsApi) detail(w http.ResponseWriter, r *http.Request) {

	guestId := utils.GetIDParam(r)
	if utils.IsEmpty(guestId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required guest_id as param", nil),
		)
	}

	id, err := strconv.Atoi(guestId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of guest_id is invalid", nil),
		)
	}

	guest, err := api.detailHandler(id)

	utils.RespondwithJSON(
		w,
		http.StatusOK,
		utils.DataFormat("Success !", guest, 0, 0, 0),
	)

}

func (api *GuestsApi) update(w http.ResponseWriter, r *http.Request) {
	guestId := utils.GetIDParam(r)
	if utils.IsEmpty(guestId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required guest_id as param", nil),
		)
	}

	id, err := strconv.Atoi(guestId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of guest_id is invalid", nil),
		)
	}

	updated, err := api.bodyToStruct(r.Body)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Invalid body", nil),
		)
		return
	}

	if ok, err := utils.ValidateStruct(updated); !ok && err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return

	}

	guest, err := api.detailHandler(id)

	if !utils.IsEmpty(updated.Name) {
		guest.Name = updated.Name
	}

	if !utils.IsEmpty(updated.Email) {
		guest.Email = updated.Email
	}

	if !utils.IsEmpty(updated.PhoneNumber) {
		guest.PhoneNumber = updated.PhoneNumber
	}

	if !utils.IsEmpty(updated.IdentificationId) {
		guest.IdentificationId = updated.IdentificationId
	}

	if err = api.updateHandler(guest); err != nil {
		if !utils.IsEqual(err.Error(), "failed-to-update") {
			utils.RespondwithJSON(w, http.StatusInternalServerError, nil)
			return
		}
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat(err.Error(), nil),
		)
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, map[string]interface{}{
		"messages": "Success !",
	})
}

func (api *GuestsApi) delete(w http.ResponseWriter, r *http.Request) {
	guestId := utils.GetIDParam(r)
	if utils.IsEmpty(guestId) {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Required guest_id as param", nil),
		)
	}

	id, err := strconv.Atoi(guestId)
	if err != nil {
		utils.RespondwithJSON(w, http.StatusBadRequest,
			utils.ErrFormat("Type of guest_id is invalid", nil),
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

	utils.RespondwithJSON(w, http.StatusOK, map[string]interface{}{
		"messages": "Success !",
	})
}
