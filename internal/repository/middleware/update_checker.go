package middleware

import (
	"fmt"
	"strings"
	restapimediasoft "tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft"
)

func UpdateChecker(input restapimediasoft.UpdateListInput) (setQuery string, count int, arg []interface{}) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Manufacturer != nil {
		setValues = append(setValues, fmt.Sprintf("manufacturer=$%d", argId))
		args = append(args, *input.Manufacturer)
		argId++
	}
	if input.Height != nil {
		setValues = append(setValues, fmt.Sprintf("height=$%d", argId))
		args = append(args, *input.Height)
		argId++
	}
	if input.Width != nil {
		setValues = append(setValues, fmt.Sprintf("width=$%d", argId))
		args = append(args, *input.Width)
		argId++
	}
	if input.Lenght != nil {
		setValues = append(setValues, fmt.Sprintf("lenght=$%d", argId))
		args = append(args, *input.Lenght)
		argId++
	}
	Query := strings.Join(setValues, ", ")
	return Query, argId, args
}
