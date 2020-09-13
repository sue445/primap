package entity

import "time"

const (
	shopCollectionName = "Shops"
)

// ShopEntity represents a shop entity for Firestore
type ShopEntity struct {
	Name       string
	Prefecture string
	Address    string
	Series     []string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (e *ShopEntity) toFirestore() map[string]interface{} {
	data := map[string]interface{}{
		"Name":       e.Name,
		"Prefecture": e.Prefecture,
		"Address":    e.Address,
		"Series":     e.Series,
		"UpdatedAt":  time.Now(),
	}

	if e.CreatedAt == nil {
		data["CreatedAt"] = time.Now()
	}

	return data
}

func fromFirestore(data map[string]interface{}) *ShopEntity {
	var series []string
	rawSeries := data["Series"].([]interface{})

	for _, raw := range rawSeries {
		series = append(series, raw.(string))
	}

	createdAt := data["CreatedAt"].(time.Time)
	updatedAt := data["UpdatedAt"].(time.Time)

	shop := &ShopEntity{
		Name:       data["Name"].(string),
		Prefecture: data["Prefecture"].(string),
		Address:    data["Address"].(string),
		Series:     series,
		CreatedAt:  &createdAt,
		UpdatedAt:  &updatedAt,
	}

	return shop
}
