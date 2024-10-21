//handlers/faq_handler/go
package handlers

import (
    "go-page/database"
    "go-page/models"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

// FAQPage menampilkan halaman FAQ
func FAQPage(c echo.Context) error {
    var faqs []models.FAQ
    err := database.DB.Select(&faqs, "SELECT * FROM faqs")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch FAQs"})
    }
    return c.Render(http.StatusOK, "faq.html", faqs)
}

// GetFAQs untuk mendapatkan semua FAQ
func GetFAQs(c echo.Context) error {
    var faqs []models.FAQ
    err := database.DB.Select(&faqs, "SELECT * FROM faqs")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch FAQs"})
    }
    return c.JSON(http.StatusOK, faqs)
}

// CreateFAQ untuk menambahkan FAQ baru dari web
func CreateFAQ(c echo.Context) error {
    question := c.FormValue("question")
    answer := c.FormValue("answer")

    // Validasi input kosong
    if question == "" || answer == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Question and Answer are required"})
    }

    _, err := database.DB.Exec("INSERT INTO faqs (question, answer) VALUES ($1, $2)", question, answer)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create FAQ"})
    }

    return c.Redirect(http.StatusSeeOther, "/faq")
}

// UpdateFAQ untuk mengupdate FAQ dari web
func UpdateFAQ(c echo.Context) error {
    id, err := strconv.Atoi(c.FormValue("id"))
    if err != nil || id <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    question := c.FormValue("question")
    answer := c.FormValue("answer")

    // Validasi input kosong
    if question == "" || answer == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Question and Answer are required"})
    }

    _, err = database.DB.Exec("UPDATE faqs SET question=$1, answer=$2 WHERE id=$3", question, answer, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update FAQ"})
    }

    return c.Redirect(http.StatusSeeOther, "/faq")
}

// DeleteFAQ untuk menghapus FAQ dari web
func DeleteFAQ(c echo.Context) error {
    id, err := strconv.Atoi(c.FormValue("id"))
    if err != nil || id <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    _, err = database.DB.Exec("DELETE FROM faqs WHERE id=$1", id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete FAQ"})
    }

    return c.Redirect(http.StatusSeeOther, "/faq")
}