package schemas

type Categories struct {
	ID              int    `bun:"id,pk,autoincrement"`
	Name            string `bun:"name"`
	IsContainerYard bool   `bun:"is_container_yard"`
}
