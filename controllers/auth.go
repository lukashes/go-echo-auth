package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/parnurzeal/gorequest"
	"github.com/lukashes/go-echo-auth/utils"
)

type AuthParams struct {
	UserId      int    `json:"user_id"      validate:"required"`
	AccessToken string `json:"access_token" validate:"required"`
}

type VkUserInfo struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func AuthRegister(c *echo.Context) error {
	var auth AuthParams
	c.Bind(&auth)

	err := utils.Validator.Struct(auth)
	if err != nil {
		return err
	}

	check := fmt.Sprintf("https://api.vk.com/method/users.get?v=5.37&access_token=%s", auth.AccessToken)
	_, body, errs := gorequest.New().Get(check).End()
	if errs != nil {
		return errs[0]
	}
	var resp map[string]*json.RawMessage
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return err
	}
	if _, ok := resp["response"]; !ok {
		return fmt.Errorf("Data structure is broken")
	}
	var users []VkUserInfo
	err = json.Unmarshal(*resp["response"], &users)
	if err != nil {
		return err
	}
	if users[0].Id != auth.UserId {
		return fmt.Errorf("User id is wrong")
	}
	return c.JSON(http.StatusOK, users[0])
}
