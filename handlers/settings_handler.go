package handlers

import (
    "go-page/database"
    "go-page/models"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

// SettingsPage menampilkan halaman settings
func SettingsPage(c echo.Context) error {
    var settings models.Settings
    err := database.DB.Get(&settings, "SELECT * FROM settings WHERE id = 1") // Assuming a single settings record
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch settings"})
    }
    return c.Render(http.StatusOK, "settings.html", settings)
}

// UpdateSettings untuk mengupdate settings dari web
func UpdateSettings(c echo.Context) error {
    id, err := strconv.Atoi(c.FormValue("id"))
    if err != nil || id <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    appname := c.FormValue("appname")
    description := c.FormValue("description")
    about := c.FormValue("about")
    phone := c.FormValue("phone")
    email := c.FormValue("email")
    location := c.FormValue("location")

    // Validasi input kosong
    if appname == "" || description == "" || about == "" || phone == "" || email == "" || location == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
    }

    _, err = database.DB.Exec("UPDATE settings SET appname=$1, description=$2, about=$3, phone=$4, email=$5, location=$6 WHERE id=$7",
        appname, description, about, phone, email, location, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update settings"})
    }

    return c.Redirect(http.StatusSeeOther, "/settings")
}
