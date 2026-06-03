package schemas

type Users struct {
	ID          int    `bun:"id,pk"`
	Login       string `bun:"login"`
	Password    string `bun:"password"`
	CityID      int    `bun:"city_id"`
	IsSuperuser bool   `bun:"is_superuser"`

	City *Cities `bun:"rel:belongs-to,join:city_id=id"`
}
