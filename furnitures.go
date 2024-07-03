package restapimediasoft

type Furnitures_list struct {
	id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name" binding:"required`
	Manufacturer int    `json:"manufacturer" db:"manufacturer" binding:"required"`
	Height       int    `json:"height" db:"height"`
	Width        int    `json:"width" db:"width"`
	Lenght       int    `json:"lenght" db:"lenght"`
}

type UpdateListInput struct {
	Name         *string `json:"name"`
	Manufacturer *int    `json:"manufacturer"`
	Height       *int    `json:"height"`
	Width        *int    `json:"width"`
	Lenght       *int    `json:"lenght"`
}
