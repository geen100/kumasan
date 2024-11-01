package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = API("BearSighting", func() {
	Title("Bear Sighting API")
	Description("API for recording and retrieving bear sightings")
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("sightings", func() {
	BasePath("/sightings")

	// クマ目撃情報を追加するアクション
	Action("add", func() {
		Routing(POST("/"))
		Description("Add a new bear sighting")
		Payload(func() {
			Attribute("latitude", Number, func() {
				Description("Latitude of the sighting location")
				Example(35.6895)
			})
			Attribute("longitude", Number, func() {
				Description("Longitude of the sighting location")
				Example(139.6917)
			})
			Attribute("city", String, func() {
				Description("City or municipality of the sighting")
				Example("Tokyo")
			})
			Attribute("area", String, func() {
				Description("Specific area or district where the sighting occurred")
				Example("Shinjuku")
			})
			Attribute("date", String, func() {
				Description("Date of the sighting (ISO 8601 format)")
				Example("2024-10-12")
			})
			Attribute("time", String, func() {
				Description("Time of the sighting (24-hour format HH:MM)")
				Example("09:30")
			})
			Attribute("situation", String, func() {
				Description("Details about the sighting situation")
				Example("Bear was seen near the trail, moving towards the forest.")
			})
			Attribute("details", String, func() {
				Description("Additional details or description")
				Example("Bear was approximately 200kg, seen moving slowly.")
			})
			Attribute("classification", String, func() {
				Description("Classification of sighting")
				Enum("目撃", "痕跡", "噂")
				Example("目撃")
			})
			Required("latitude", "longitude", "city", "area", "date", "time", "situation", "details", "classification")
		})
		Response(Created, func() {
			Media(Sighting)
		})
		Response(BadRequest)
		Response(InternalServerError)
	})
})

// クマ目撃情報のメディアタイプ定義
var Sighting = MediaType("application/vnd.sighting+json", func() {
	ContentType("application/json")
	Description("Bear Sighting")
	Attributes(func() {
		Attribute("id", Integer, "Unique sighting ID")
		Attribute("latitude", Number, "Latitude of the sighting location")
		Attribute("longitude", Number, "Longitude of the sighting location")
		Attribute("city", String, "City or municipality of the sighting")
		Attribute("area", String, "Specific area or district where the sighting occurred")
		Attribute("date", String, "Date of the sighting")
		Attribute("time", String, "Time of the sighting")
		Attribute("situation", String, "Details about the sighting situation")
		Attribute("details", String, "Additional details or description")
		Attribute("classification", String, "Classification of sighting")
		Required("id", "latitude", "longitude", "city", "area", "date", "time", "situation", "classification")
	})
	View("default", func() {
		Attribute("id")
		Attribute("latitude")
		Attribute("longitude")
		Attribute("city")
		Attribute("area")
		Attribute("date")
		Attribute("time")
		Attribute("situation")
		Attribute("details")
		Attribute("classification")
	})
})
