package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	productsdto "waysbeans_be/dto/product"
	dto "waysbeans_be/dto/result"
	"waysbeans_be/models"
	"waysbeans_be/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// var path_file = "http://localhost:5000/uploads/"

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// for i, p := range products {
	// 	products[i].Image = path_file + p.Image
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: products})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// product.Image = path_file + product.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: convertResponseProduct(product)})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	dataContex := c.Get("dataFile").(string)
	stock, _ := strconv.Atoi(c.FormValue("stock"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	request := productsdto.CreateProductRequest{
		Name:        c.FormValue("name"),
		Stock:       stock,
		Price:       price,
		Description: c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// Declare Context Background, Cloud Name, API Key, API Secret ...
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataContex, uploader.UploadParams{Folder: "waysbeans"})
	if err != nil {
		fmt.Println(err.Error())
	}

	product := models.Product{
		Name:        request.Name,
		Stock:       request.Stock,
		Price:       request.Price,
		Description: request.Description,
		Image:       resp.SecureURL,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: product})
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataContex := c.Get("dataFile")
	filename := dataContex.(string)

	stock, _ := strconv.Atoi(c.FormValue("stock"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	request := productsdto.UpdateProductRequest{
		Name:        c.FormValue("name"),
		Stock:       stock,
		Price:       price,
		Description: c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	product, _ := h.ProductRepository.GetProduct(id)

	product.Name = request.Name
	product.Stock = request.Stock
	product.Price = request.Price
	product.Description = request.Description
	product.Image = filename

	product, err = h.ProductRepository.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: product})
}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	deleteProduct, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := dto.SuccessResult{Status: "Success", Data: deleteProduct}
	return c.JSON(http.StatusOK, response)
}

func convertResponseProduct(u models.Product) productsdto.ProductResponse {
	return productsdto.ProductResponse{
		Name:        u.Name,
		Stock:       u.Stock,
		Price:       u.Price,
		Description: u.Description,
		Image:       u.Image,
	}
}
