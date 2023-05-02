package Personal

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Personal struct {
	Nama   string `json:"Nama"`
	NIM    string `json:"NIM"`
	Alamat string `json:"Alamat"`
}

var personal = []Personal{
	{Nama: "Rifaldi", NIM: "09876", Alamat: "Jl.Soekarno"},
	{Nama: "Indra", NIM: "67890", Alamat: "Jl.Hatta"},
}

func PersonalHandler(w http.ResponseWriter, r *http.Request) {
	count := 0

	if r.Method == "POST" {
		GetNamePersonal := Personal{
			Nama: r.FormValue("Nama"),
		}

		for _, value := range personal {
			if strings.EqualFold(GetNamePersonal.Nama, value.Nama) {
				response, err := json.Marshal(Personal{
					Nama:   value.Nama,
					NIM:    value.NIM,
					Alamat: value.Alamat,
				})

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				count++

				w.Header().Set("Content-Type", "aplication/json")
				w.Write(response)
			}
		}

		if count == 0 {
			http.Error(w, "Nama Tidak Ketemu", http.StatusMethodNotAllowed)
		}
	} else {

		http.Error(w, "Method Tidak Ada", http.StatusMethodNotAllowed)

	}
}

func AllDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		personalJson, _ := json.Marshal(personal)

		w.Header().Set("Content-Type", "aplication/json")
		w.Write(personalJson)
	} else {
		http.Error(w, "Method Tidak Ada", http.StatusMethodNotAllowed)
	}
}
