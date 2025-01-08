package model

type BearID string

// 目撃情報
type Bear struct {
	ID        BearID  //ユニークなid
	Latitude  float64 // 緯度
	Longitude float64 // 経度
	City      string  // 都道府県
	Address   string  // 住所
	Date      string  // 西暦年月日
	Time      string  //  時間
	Details   string  // 詳細
}
