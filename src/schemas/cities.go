package schemas

type Cities struct {
	ID       int    `bun:"id,pk,autoincrement"`
	CityName string `bun:"city_name"`
}
