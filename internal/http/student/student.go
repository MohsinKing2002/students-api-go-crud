package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mohsinking2002/students-api-go-crud/internal/types"
	"github.com/mohsinking2002/students-api-go-crud/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err:= json.NewDecoder(r.Body).Decode(&student)

		//empty body error
		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		//general error
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validation
		if err:= validator.New().Struct(student); err != nil{
			//type cast err to validationErrors
			validatorErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validatorErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]any{"success": "User Created", "data":student})
	}
}