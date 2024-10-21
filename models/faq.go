//models/faq.go
package models

type FAQ struct {
    ID       int    `db:"id" json:"id"`
    Question string `db:"question" json:"question"`
    Answer   string `db:"answer" json:"answer"`
}
