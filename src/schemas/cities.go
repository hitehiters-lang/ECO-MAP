package schemas

type Cities struct {
	ID       int    `bun:"id,pk"`
	CityName string `bun:"city_name"`
}
