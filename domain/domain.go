package domain

type Todo struct {
	//primarykey
	ID       uint   `gorm:"primary" json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

type UpdatedTodo struct {
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}
