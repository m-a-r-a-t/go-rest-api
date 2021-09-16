package core



type User struct {
    Id      string `db:"id"`
    Balance float64 `db:"balance"`
}