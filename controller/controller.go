package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimseokgis/backend-ai/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// GetUsers returns all users
func GetUsers(c *fiber.Ctx) error {
	response, err := json.Marshal(Response)
	var users []model.User
	config.DB.Find(&users)
	return c.JSON(users)
	}

// GetUser returns a user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return helper.ErrorResponse(c, "User not found")
	}
	return c.JSON(user)
}

func NotFound(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Message = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, resp)
}

func Comment(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	comment := new(model.Comment)
	resp.Status = false
	conn := helper.SetConnection()
	err := json.NewDecoder(req.Body).Decode(comment)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
		helper.WriteJSON(respw, http.StatusNotAcceptable, resp)
	}
	insID := helper.InsertComment(conn, *comment)
	resp.Status = true
	resp.Message = fmt.Sprintf("Data berhasil diinsert %s", insID)
	helper.WriteJSON(respw, http.StatusOK, resp)
	return
}
