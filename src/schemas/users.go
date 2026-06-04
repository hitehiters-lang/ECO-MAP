package schemas

type Users struct {
	ID          int    `bun:"id,pk,autoincrement"`
	Login       string `bun:"login"`
	Password    string `bun:"password"`
	CityID      int    `bun:"city_id"`
	IsSuperuser bool   `bun:"is_superuser"`
	Name        string `bun:"name"`

	City *Cities `bun:"rel:belongs-to,join:city_id=id"`
}
