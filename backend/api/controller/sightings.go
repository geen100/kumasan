package api

import (
	"backend/backend/api/app"
	goa "github.com/shogo82148/goa-v1"
)

// SightingsController implements the sightings resource.
type SightingsController struct {
	*goa.Controller
}

// NewSightingsController creates a sightings controller.
func NewSightingsController(service *goa.Service) *SightingsController {
	return &SightingsController{Controller: service.NewController("SightingsController")}
}

// Add runs the add action.
func (c *SightingsController) Add(ctx *app.AddSightingsContext) error {
	// SightingsController_Add: start_implement

	// Put your logic here

	return nil
	// SightingsController_Add: end_implement
}
