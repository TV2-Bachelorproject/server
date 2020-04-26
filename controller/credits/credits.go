package credits

import (
	"encoding/json"
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

type CreateRequest struct {
	ProgramID     uint
	SeasonID      uint
	SerieID       uint
	CreditGroupID uint
	Persons       []uint
}

func Create(w http.ResponseWriter, r *http.Request) {
	var data CreateRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var people public.People
	db.Model(&public.Person{}).Where("id IN (?)", data.Persons).Find(&people)

	credit := public.Credit{
		ProgramID: data.ProgramID,
		SeasonID:  data.SeasonID,
		SerieID:   data.SerieID,
		Persons:   people,
	}

	db.Create(&credit)
}

type DeleteRequest struct {
	CreditID uint
	PersonID uint
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var data DeleteRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Exec("DELETE FROM credit_persons WHERE credit_id = ? AND person_id = ?", data.CreditID, data.PersonID)
}

type AcceptRequest struct {
	ProgramID uint
}

func Accept(w http.ResponseWriter, r *http.Request) {
	var data AcceptRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var credits public.Credits
	db.Where("program_id = ?", data.ProgramID).Find(&credits)

	for _, credit := range credits {
		credit.Accepted = true
		db.Save(&credit)
	}
}