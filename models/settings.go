package models

type Settings struct {
    ID          int    `db:"id" json:"id"`
    AppName     string `db:"appname" json:"appname"`
    Description string `db:"description" json:"description"`
    About       string `db:"about" json:"about"`
    Phone       string `db:"phone" json:"phone"`
    Email       string `db:"email" json:"email"`
    Location    string `db:"location" json:"location"`
}
