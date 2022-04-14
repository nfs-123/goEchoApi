package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type College struct {
	Id          string `json:"id" validate:"required"`
	CollegeName string `json:"collegeName"`
}

// var collegeVar map[string]interface{}

func main() {
	fmt.Println("Listening Server....")
	v := validator.New()
	// english := en.New()
	// uni := ut.New(english, english)
	// trans, _ := uni.GetTranslator("en")
	// _ = enTranslations.RegisterDefaultTranslations(v, trans)
	colleges1 := []College{
		{
			Id:          "1",
			CollegeName: "ABC...",
		},
		{
			Id:          "2",
			CollegeName: "XYZ",
		},
	}
	// colleges := []map[string]interface{}{
	// 	{
	// 		"collegeId":   "1",
	// 		"collegeName": "A.D Joshi Junior College",
	// 	},
	// 	{
	// 		"collegeId":   "2",
	// 		"collegeName": "Shivaji College",
	// 	},
	// }
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/addCollege", func(c echo.Context) error {
		college := College{}
		// mySlice2 := make([]int, 0)
		if err := c.Bind(&college); err != nil {
			fmt.Println("error in binding", err)

		}
		if err := v.Struct(&college); err != nil {
			fmt.Println(err)

			return c.String(http.StatusBadGateway, "Validation is failed..")

		}
		colleges1 = append(colleges1, college)
		// fmt.Println(college)
		return c.JSON(http.StatusCreated, college)

	})
	e.GET("/getColleges", func(c echo.Context) error {
		return c.JSON(http.StatusOK, colleges1)
	})
	e.GET("/getCollegeFromId/:collegeId", func(c echo.Context) error {
		// id := c.Param("collegeId")
		// for _, college := range colleges {
		// 	if college["collegeId"] == c.Param("collegeId") {
		// 		return c.JSON(http.StatusOK, college)

		// 	}
		// }
		for _, college := range colleges1 {
			if college.Id == c.Param("collegeId") {
				return c.JSON(http.StatusOK, college)

			}

		}

		return c.JSON(http.StatusNotFound, "college not found")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
