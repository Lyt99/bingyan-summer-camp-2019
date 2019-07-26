package controller

import (
	"final_task/serializer"
	"final_task/service"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

// POST /pics 上传图片
func PicsPost(c echo.Context) (e error) {
	s := service.PicsRegisterService{}
	if e = c.Bind(s); e != nil {
		return c.JSON(http.StatusBadRequest, serializer.Response{
			Success: false,
			Error:   "请求格式不正确",
			Data:    nil,
		})
	}
	pic, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := pic.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("./pics/" + s.Name)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	if errResponse := s.Register(); errResponse != nil {
		return c.JSON(http.StatusServiceUnavailable, errResponse)
	}
	return c.JSON(http.StatusOK, serializer.PicsRegisterResponse{
		Url: "https://pics/" + s.Name + ".jpg",
	})
}
