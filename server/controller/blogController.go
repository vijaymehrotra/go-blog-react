package controller

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vijaymehrotra/blog/database"
	"github.com/vijaymehrotra/blog/model"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	time.Sleep(time.Millisecond * 1000)

	db := database.DBConn
	var records []model.Blog
	res := db.Find(&records)
	if res.Error != nil {
		log.Println("Error Fetching the Records")
		return res.Error
	}

	context["statusText"] = "Ok"
	context["msg"] = "Successfully Get All the Blogs"
	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func BlogDetail(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record Not Found")
		c.Status(400)
		return c.JSON(context)
	}

	context["record"] = record
	context["statusText"] = "Ok"
	context["msg"] = "Blog Detail"

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	record := new(model.Blog)

	if err := c.BodyParser(record); err != nil {
		log.Println("Error in Parsing Request")
		context["statusText"] = ""
		context["msg"] = "Something Went Wrong"
	}

	//file upload
	file , err := c.FormFile("file")
	 if err != nil {
		log.Println("Error in file Upload",err)
		return err
	 } 

	if file.Size > 0 {
		fileName := "./static/uploads/" + file.Filename 
		if err := c.SaveFile(file,fileName); err != nil {
			log.Println("Error in saving file. . . ",err)
		} 
		record.Image = fileName	
	}  

	result := database.DBConn.Create(record)
	if result.Error != nil || record.ID == 0 {
		log.Println("Error in serving data")
		context["statusText"] = ""
		context["msg"] = "Something Went Wrong"
		return result.Error
	}
	context["msg"] = "Record is Saved Successfully"
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update a  Blog for the given id",
	}
	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record Not Found")
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in Parsing Request")
	}

	result := database.DBConn.Save(record)
	if result.Error != nil {
		log.Println("Error in Parsing Request")
	}

	context["data"] = record
	context["Msg"] = "Record Updated Successfully"

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{}
	id := c.Params("id")
	var record model.Blog
	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record Not Found")
		c.Status(400)
		return c.JSON(fiber.Map{
			"Error": "Record Not Found",
		})
	}

	//remove image

	fileName := record.Image

	err := os.Remove(fileName)
	if err != nil {
		log.Println("Error Deleting file",err)
	}

	result := database.DBConn.Delete(&record, id)
	if result.Error != nil {
		log.Println("Error Deleting Record")
		return result.Error
	}

	context["msg"] = "Record Deleted Successfully"

	c.Status(200)
	return c.JSON(context)
}
