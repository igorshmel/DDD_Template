package rest

import (
	"default_ddd/app/cmd/default_ddd/middleware"
	"default_ddd/app/internal/adapters/port"
	"default_ddd/app/pkg/dto"
	"default_ddd/app/pkg/errs"
	"default_ddd/app/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProductEndpoint --
type CreateProductEndpoint struct {
	log     logger.Logger
	usecase port.CreateProductUseCase
}

// NewCreateProductEndpoint --
func NewCreateProductEndpoint(usecase port.CreateProductUseCase, log logger.Logger) port.CreateProductEndpoint {
	return CreateProductEndpoint{
		log:     log,
		usecase: usecase,
	}
}

// ExecuteCreateProductEndpoint is handler
func (ths CreateProductEndpoint) ExecuteCreateProductEndpoint(ctx *gin.Context) {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("CreateProduct endpoint")

	req := dto.NewCreateProductRequest()

	// request parse
	if err := req.Parse(ctx); err != nil {
		log.Error("unable to parse a request: %s", err)
		ctx.JSON(http.StatusInternalServerError, errs.ErrParseRequest)
		return
	}

	// validate request
	if err := req.Validate(); err != nil {
		log.Error("error of validation: %s", err)
		ctx.JSON(http.StatusInternalServerError, errs.ErrSyntax)
		return
	}

	// call usecase
	err := ths.usecase.Execute(ctx, req)
	if err != nil {
		log.Error("%s", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
