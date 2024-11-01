package controller

import (
	"kumasan/api/app"
	"kumasan/api/models"

	goa "github.com/shogo82148/goa-v1"
)

// SightingsController implements the sightings resource.
type SightingsController struct {
	*goa.Controller
	Repo models.SQLSightingRepository
}

// NewSightingsController creates a sightings controller.
func NewSightingsController(service *goa.Service, repo models.SQLSightingRepository) *SightingsController {
	return &SightingsController{
		Controller: service.NewController("SightingsController"),
		Repo:       repo}
}

// Add runs the add action.
func (c *SightingsController) Add(ctx *app.AddSightingsContext) error {
	// SightingsController_Add: start_implement
	latitude := ctx.Payload.Latitude
	longitude := ctx.Payload.Longitude
	city := ctx.Payload.City
	area := ctx.Payload.Area
	date := ctx.Payload.Date
	time := ctx.Payload.Time
	situation := ctx.Payload.Situation
	details := ctx.Payload.Details
	classification := ctx.Payload.Classification

	sighting := models.Sighting{
		Latitude:       latitude,
		Longitude:      longitude,
		City:           city,
		Area:           area,
		Date:           date,
		Time:           time,
		Situation:      situation,
		Details:        details,
		Classification: classification,
	}
	// Put your logic here
	id, err := c.Repo.AddSighting(ctx, &sighting)
	if err != nil {
		goa.LogError(ctx, "InternalError:failed to add a sighting", "error", err.Error())
	}
	goa.LogInfo(ctx, "successfully add sighting", "sightingid", id)
	return ctx.Created(&app.Sighting{
		ID:             int(id),
		Latitude:       sighting.Latitude,
		Longitude:      sighting.Longitude,
		City:           sighting.City,
		Area:           sighting.Area,
		Date:           sighting.Date,
		Time:           sighting.Time,
		Situation:      sighting.Situation,
		Details:        &sighting.Details,
		Classification: sighting.Classification,
	})

	// SightingsController_Add: end_implement
}
