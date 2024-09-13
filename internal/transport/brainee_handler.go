package transport

import (
	"br_api/internal/models/dto"
	"br_api/internal/models/mapper"
	"br_api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BraineeHandler struct {
	service  *service.BraineeService
	validate *validator.Validate
}

func NewBraineeHandler(service *service.BraineeService) *BraineeHandler {
	return &BraineeHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *BraineeHandler) CreateBrainee(c *gin.Context) {
	var braineeRequest dto.CreateBraineeRequest
	if err := c.ShouldBindJSON(&braineeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.validate.Struct(&braineeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	braineeEntity := mapper.ToBraineeEntity(&braineeRequest)

	if err := h.service.CreateBrainee(braineeEntity); err != nil {
		if err.Error() == "brainee already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Brainee already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brainee"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Brainee created successfully"})
}

func (h *BraineeHandler) GetBraineeById(c *gin.Context) {
	idStr := c.Param("braineeId")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brainee Id"})
		return
	}

	brainee, err := h.service.GetBraineeById(int(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brainee not found"})
		return
	}

	res := mapper.ToBraineeResponse(brainee)
	c.JSON(http.StatusOK, res)
}

func (h *BraineeHandler) GetAllBrainees(c *gin.Context) {
	brainees, err := h.service.GetAllBrainees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get brainees"})
		return
	}

	var res []dto.BraineeResponse
	res = []dto.BraineeResponse{}
	for _, brainee := range brainees {
		res = append(res, *mapper.ToBraineeResponse(brainee))
	}

	c.JSON(http.StatusOK, res)
}
