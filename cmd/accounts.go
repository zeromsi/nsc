package cmd

import (
	"github.com/labstack/echo"
	"log"
)

type Account struct{}

func AccountRouter(g *echo.Group) {
	account := Account{}
	g.POST("", account.Save)
}

func (a Account) Save(context echo.Context) error {
	_, _, err := ExecuteCmd(CreateAddAccountCmd(), "-n", "abc")
	if err != nil {
		log.Println(err.Error())
		return GenerateErrorResponse(context, nil, err.Error())
	}
	return GenerateSuccessResponse(context, nil, "Operation successful!")
}
