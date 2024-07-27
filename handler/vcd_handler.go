package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"vcd-rental/vcd"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type vcdHandler struct {
	vcdService vcd.Service
}

func (handler *vcdHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome To VCD Rental Back-End App",
	})
}

func VCDHandler(vcdService vcd.Service) *vcdHandler {
	return &vcdHandler{vcdService}
}

func (handler *vcdHandler) CreateVCD(c *gin.Context) {
	var vcdRequest vcd.VCDRequest

	err := c.ShouldBindJSON(&vcdRequest)

	if err != nil {
		var syntaxError *json.SyntaxError
		var validationErrors validator.ValidationErrors

		switch {
		case errors.As(err, &syntaxError):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Request body contains badly-formed JSON (syntax error)",
			})
			return
		case errors.As(err, &validationErrors):
			errorMessages := []string{}
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessages,
			})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	vcd, err := handler.vcdService.Create(vcdRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "VCD Created Successfully",
		"data":    vcd,
	})
}

func (handler *vcdHandler) GetAllVCD(c *gin.Context) {
	vcds, err := handler.vcdService.GetAllVCD()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var vcdsResponse []vcd.VCDResponse

	for _, v := range vcds {
		vcdResponse := convertToResponse(v)

		vcdsResponse = append(vcdsResponse, vcdResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "VCD Data Retrieved Successfully",
		"data":    vcdsResponse,
	})
}

func (handler *vcdHandler) GetOneVCD(c *gin.Context) {
	getId := c.Param("id")
	id, _ := strconv.Atoi(getId)

	vcds, err := handler.vcdService.GetOneVCD(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	vcdResponse := convertToResponse(vcds)

	c.JSON(http.StatusOK, gin.H{
		"message": "VCD Data Retrieved Successfully",
		"data":    vcdResponse,
	})
}

func (handler *vcdHandler) UpdateVCD(c *gin.Context) {
	getId := c.Param("id")
	id, _ := strconv.Atoi(getId)

	var updateVCDRequest vcd.UpdateVCDRequest

	err := c.ShouldBindJSON(&updateVCDRequest)
	if err != nil {
		var syntaxError *json.SyntaxError
		var validationErrors validator.ValidationErrors

		switch {
		case errors.As(err, &syntaxError):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Request body contains badly-formed JSON (syntax error)",
			})
			return
		case errors.As(err, &validationErrors):
			errorMessages := []string{}
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessages,
			})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	vcd, err := handler.vcdService.UpdateVCD(id, updateVCDRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	vcdResponse := convertToResponse(vcd)

	c.JSON(http.StatusOK, gin.H{
		"message": "VCD Data Updated Successfully",
		"data":    vcdResponse,
	})
}

func (handler *vcdHandler) DeleteVCD(c *gin.Context) {
	getId := c.Param("id")
	id, _ := strconv.Atoi(getId)

	vcd, err := handler.vcdService.DeleteVCD(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	vcdResponse := convertToResponse(vcd)

	c.JSON(http.StatusOK, gin.H{
		"message": "VCD Data Deleted Successfully",
		"data":    vcdResponse,
	})
}

func convertToResponse(v vcd.VCD) vcd.VCDResponse {
	var response vcd.VCDResponse
	response.ID = v.ID
	response.Title = v.Title
	response.Price = v.Price
	response.Stock = v.Stock
	response.Description = v.Description
	return response
}
