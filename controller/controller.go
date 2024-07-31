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
	
	w.Write(response)
	return
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
