package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = API("BearSighting", func() {
	Title("Bear API")
	Description("クマ目撃情報の記録を行うAPI")
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("bear", func() {
	BasePath("/bear")

	// クマ目撃情報を追加するアクション
	Action("add", func() {
		Routing(POST("/"))
		Description("新しいクマ目撃情報を追加する")
		Payload(func() {
			Attribute("latitude", Number, func() {
				Description("目撃場所の緯度")
				Example(35.6895)
			})
			Attribute("longitude", Number, func() {
				Description("目撃場所の経度")
				Example(139.6917)
			})
			Attribute("city", String, func() {
				Description("目撃があった市区町村")
				Example("東京都")
			})
			Attribute("address", String, func() {
				Description("目撃場所の具体的な住所")
				Example("東京都渋谷区")
			})
			Attribute("date", String, func() {
				Description("目撃日 (ISO 8601フォーマット)")
				Example("2024-10-12")
			})
			Attribute("time", String, func() {
				Description("目撃時刻 (24時間形式 HH:MM)")
				Example("09:30")
			})
			Attribute("details", String, func() {
				Description("目撃の追加詳細や説明")
				Example("体重約200kgのクマがゆっくりと移動していた。")
			})

			Required("latitude", "longitude", "city", "address", "date", "time", "details")
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
	Description("クマ目撃情報")
	Attributes(func() {
		Attribute("id", String, "一意の目撃情報ID")
		Attribute("latitude", Number, "目撃場所の緯度")
		Attribute("longitude", Number, "目撃場所の経度")
		Attribute("city", String, "目撃があった市区町村")
		Attribute("address", String, "目撃場所の具体的な住所")
		Attribute("date", String, "目撃日")
		Attribute("time", String, "目撃時刻")
		Attribute("details", String, "目撃の追加詳細や説明")
		Required("id", "latitude", "longitude", "city", "address", "date", "time", "details")
	})
	View("default", func() {
		Attribute("id")
		Attribute("latitude")
		Attribute("longitude")
		Attribute("city")
		Attribute("address")
		Attribute("date")
		Attribute("time")
		Attribute("details")
	})
})
