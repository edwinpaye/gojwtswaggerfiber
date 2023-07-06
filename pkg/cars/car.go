package cars

type Car struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Model       string `json:"model"`
	Placa      string `json:"placa"`
	Description string `json:"description"`
}
