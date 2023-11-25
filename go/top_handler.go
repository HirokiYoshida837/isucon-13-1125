package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TagModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type TagsResponse struct {
	Tags []*Tag `json:"tags"`
}

func LoadTagCache() error {

	ctx := context.Background()

	var tagModels []*TagModel
	if err := dbConn.SelectContext(ctx, &tagModels, "SELECT * FROM tags"); err != nil {
		return err
	}

	//if err := dbConn.Commit(); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, "failed to commit: "+err.Error())
	//}

	tags := make([]*Tag, 0)
	for _, item := range tagModels {
		at := &Tag{
			ID:   item.ID,
			Name: item.Name,
		}
		tags = append(tags, at)
	}

	tagCache = tags

	return nil

}

func getTagHandler(c echo.Context) error {
	//ctx := c.Request().Context()

	if len(tagCache) == 0 {
		LoadTagCache()
	}

	tags := tagCache

	c.Response().Header().Set("Cache-Control", "max-age=36000000")
	return c.JSON(http.StatusOK, &TagsResponse{
		Tags: tags,
	})

}

// 配信者のテーマ取得API
// GET /api/user/:username/theme
func getStreamerThemeHandler(c echo.Context) error {
	ctx := c.Request().Context()

	if err := verifyUserSession(c); err != nil {
		// echo.NewHTTPErrorが返っているのでそのまま出力
		c.Logger().Printf("verifyUserSession: %+v\n", err)
		return err
	}

	username := c.Param("username")

	/*	tx, err := dbConn.BeginTxx(ctx, nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to begin transaction: "+err.Error())
		}
		defer tx.Rollback()*/

	userModel := UserModel{}
	err := dbConn.GetContext(ctx, &userModel, "SELECT id FROM users WHERE name = ?", username)
	if errors.Is(err, sql.ErrNoRows) {
		return echo.NewHTTPError(http.StatusNotFound, "not found user that has the given username")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user: "+err.Error())
	}

	themeModel := ThemeModel{}
	if err := dbConn.GetContext(ctx, &themeModel, "SELECT `id`, `dark_mode` FROM themes WHERE user_id = ?", userModel.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user theme: "+err.Error())
	}

	//if err := tx.Commit(); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, "failed to commit: "+err.Error())
	//}

	theme := Theme{
		ID:       themeModel.ID,
		DarkMode: themeModel.DarkMode,
	}

	c.Response().Header().Set("Cache-Control", "max-age=36000000")

	return c.JSON(http.StatusOK, theme)
}
