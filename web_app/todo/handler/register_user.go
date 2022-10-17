package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/kokoichi206/go-expert/web/todo/entity"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var b struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	// Reader の性質か何か、1回読み取ると次は無理そう（EOF）
	// body, _ := ioutil.ReadAll(r.Body)
	// fmt.Println(string(body))

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	fmt.Println(b)

	if err := ru.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	u, err := ru.Service.RegisterUser(ctx, b.Name, b.Password, b.Role)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
