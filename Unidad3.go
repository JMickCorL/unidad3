package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Libro struct {
	gorm.Model
	Nombre      string `gorm:"type:varchar(100);"`
	Descripcion string `gorm:"type:varchar(450);"`
	Autor       string `gorm:"type:varchar(200); FOREIGNKEY"`
	Editorial   string `gorm:"type:varchar(200);"`
	Fecha       string `gorm:"type:varchar(200);"`
}

func main() {
	DB, _ = gorm.Open("mysql", "root:miguel1411@/biblioteca?charset=utf8&parseTime=True&loc=Local")

	defer DB.Close()

	DB.AutoMigrate(&Libro{})

	r := gin.Default()
	r.GET("/biblioteca/v1/books", ObtenerLibros)
	r.GET("/biblioteca/v1/books/:id", ObtenerLibro)
	r.POST("/biblioteca/v1/booksregister", CrearLibro)
	r.PUT("/biblioteca/v1/booksup/:id", ActualizarLibro)
	r.DELETE("/biblioteca/v1/bookdel/:id", EliminarLibro)

	r.Run(":8080")
}

// CrearLibro es la funcion que nos permite crear libros
func CrearLibro(c *gin.Context) {
	var libro Libro
	c.BindJSON(&libro)

	DB.Create(&libro)
	c.JSON(200, libro)
}

// ObtenerLibros es la funcion que nos permite obtener todos los libros
func ObtenerLibros(c *gin.Context) {
	var libro []Libro

	if err := DB.Find(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, libro)
	}
}

// ObtenerLibro es la funcion que nos permite obtener 1 libro mediante su Id
func ObtenerLibro(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro Libro
	if err := DB.Where("id = ?", id).First(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, libro)
	}
}

// EliminarLibro , funcion que permite eliminar un libro mediante su ID
func EliminarLibro(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro Libro
	d := DB.Where("id = ?", id).Delete(&libro)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// ActualizarLibro funcion para actualizar libro mediante id
func ActualizarLibro(c *gin.Context) {

	var libro Libro
	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).First(&libro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&libro)

	DB.Save(&libro)
	c.JSON(200, libro)

}
