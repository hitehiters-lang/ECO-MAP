package schemas

type Categories struct {
	ID              int    `bun:"id,pk"`
	Name            string `bun:"name"`
	IsContainerYard bool   `bun:"is_container_yard"`
}
