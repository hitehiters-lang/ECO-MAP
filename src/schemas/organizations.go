package schemas

type Organizations struct {
	ID         int         `bun:"id,pk"`
	Name       string      `bun:"name"`
	Category   *Categories `bun:"rel:belongs-to,join:category_id=id"`
	CategoryID int         `bun:"category_id"`
	CityID     int         `bun:"city_id"`
}
