package schemas

type Organizations struct {
	ID          int     `bun:"id,pk,autoincrement"`
	Name        string  `bun:"name"`
	CategoryID  int     `bun:"category_id"`
	CityID      int     `bun:"city_id"`
	Description string  `bun:"description"`
	CoordX      float64 `bun:"coord_x"`
	CoordY      float64 `bun:"coord_y"`

	Category *Categories `bun:"rel:belongs-to,join:category_id=id"`
	City     *Cities     `bun:"rel:belongs-to,join:city_id=id"`
}
