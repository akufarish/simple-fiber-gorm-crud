package muridController

import (
	"fiber/models"
	"net/http"

	// "net/http"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var dataMurid []models.Murid

	models.DB.Find(&dataMurid)

	return c.JSON(dataMurid)
}

func Show (c *fiber.Ctx) error {
	id := c.Params("id")

	var dataMurid models.Murid

	models.DB.First(&dataMurid, id)

	return c.JSON(dataMurid)
}

func Delete (c *fiber.Ctx) error {
	id := c.Params("id")

	var dataMurid models.Murid

	models.DB.Delete(&dataMurid, id)

	return c.JSON("Data berhasil dihapus")
}

func Store (c *fiber.Ctx) error {
	var dataMurid models.Murid

	c.BodyParser(&dataMurid)

	models.DB.Create(&dataMurid)

	return c.Status(http.StatusCreated).JSON(dataMurid);
}

func Update (c *fiber.Ctx) error {
	id := c.Params("id")
	var dataMurid models.Murid

	if err := c.BodyParser(&dataMurid); err != nil {
		c.JSON(http.StatusBadRequest)
	}

	if models.DB.Model(&dataMurid).Where("id = ?", id).Updates(&dataMurid).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest)
	}

	return c.Status(http.StatusOK).JSON(dataMurid)
}	