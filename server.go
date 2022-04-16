package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type College struct {
	Id          string `json:"id" validate:"required"`
	CollegeName string `json:"collegeName"`
}

// var collegeVar map[string]interface{}
func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func upload(c echo.Context) error {
	// Read form fields
	// name := c.FormValue("name")
	// email := c.FormValue("email")

	//-----------
	// Read file
	//-----------
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(err)

		return err
	}

	defer src.Close()
	// var dir = "Photos/ipload"
	dirPath := "D:\\Photos1\\"

	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		// log.Fatal(err)
		fmt.Println("Error in creating directory: ", err)
	}
	filePath := dirPath + file.Filename

	// Destination
	dst, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)

		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)

		return err
	}

	return c.String(http.StatusOK, filePath)
}

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
	// e.Use(middleware.CORS())

	e.File("/getHtml", "public/index.html")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/upload", upload)
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
