// handlers/index_handler.go
package handlers

import (
    "go-page/database"
    "go-page/models"
    "net/http"

    "github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {
    var settings models.Settings
    err := database.DB.Get(&settings, "SELECT * FROM settings WHERE id = 1") // Assuming single settings record
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch settings"})
    }

    var faqs []models.FAQ
    err = database.DB.Select(&faqs, "SELECT * FROM faqs")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch FAQs"})
    }

    // Gabungkan settings dan faqs dalam satu map data
    data := map[string]interface{}{
        "Settings": settings,
        "FAQs":     faqs,
    }

    return c.Render(http.StatusOK, "index.html", data)
}
