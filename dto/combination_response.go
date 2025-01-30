package dto

import "github.com/PandaX185/computer-alchemy-api/models"

type CombinationResponse struct {
	FirstElement     *models.Element
	SecondElement    *models.Element
	ResultingElement *models.Element
}
