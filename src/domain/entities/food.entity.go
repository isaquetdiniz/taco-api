package entities

import "fmt"

type Food struct {
	*Entity
	Category
	*NutritionalInformations
}

func NewFood(
	name string,
	category Category,
	nutritionalInformations *NutritionalInformations,
	options ...Option,
) *Food {
	f := &Food{
		Entity:                  NewEntity(name, options...),
		Category:                category,
		NutritionalInformations: nutritionalInformations,
	}

	return f
}

func (e Food) String() string {
	return fmt.Sprintf("[ID]: %s\n[Name]: %s\n[Category]: %s\n[Protein]: %d\n[Carb]: %d\n[Lipid]: %d\n[CreatedAt]: %s\n[UpdatedAt]: %s\n", e.ID, e.Name, e.Category.Name, e.NutritionalInformations.Protein, e.NutritionalInformations.Carbohydrate, e.NutritionalInformations.Lipid, e.CreatedAt, e.UpdatedAt)
}
