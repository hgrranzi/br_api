package mapper

import (
	"br_api/internal/models/dto"
	"br_api/internal/models/entity"
)

func ToBraineeEntity(braineeRequest *dto.CreateBraineeRequest) *entity.Brainee {
	return &entity.Brainee{
		Text:   braineeRequest.Text,
		Author: braineeRequest.Author,
		Brand:  braineeRequest.Brand,
	}
}

func ToBraineeResponse(brainee *entity.Brainee) *dto.BraineeResponse {
	return &dto.BraineeResponse{
		Id:     brainee.Id,
		Text:   brainee.Text,
		Author: brainee.Author,
		Brand:  brainee.Brand,
	}
}
