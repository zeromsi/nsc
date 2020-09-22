package cmd

import (
	"github.com/labstack/echo"
	"log"
)

type Operator struct {
}

func OperatorRouter(g *echo.Group) {
	operator := Operator{}
	g.POST("", operator.Save)
}

func (o Operator) Save(context echo.Context) error {
	_, _, err := ExecuteCmd(CreateAddOperatorCmd(), "-n", "abc")
	if err != nil {
		log.Println(err.Error())
		return GenerateErrorResponse(context, nil, err.Error())
	}
	return GenerateSuccessResponse(context, nil, "Operation successful!")
}
