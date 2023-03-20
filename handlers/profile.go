package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	profilesdto "waysbeans_be/dto/profile"
	dto "waysbeans_be/dto/result"
	"waysbeans_be/models"
	"waysbeans_be/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) FindProfiles(c echo.Context) error {
	userInfo := c.Get("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	profiles, err := h.ProfileRepository.FindProfile(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// for i, p := range profiles {
	// 	profiles[i].Image = path_file + p.Image
	// }

	response := dto.SuccessResult{Status: "Success", Data: profiles}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerProfile) GetProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	// profile.Image = path_file + profile.Image

	response := dto.SuccessResult{Status: "Success", Data: convertResponseProfile(profile)}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerProfile) CreateProfile(c echo.Context) error {
	userInfo := c.Get("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	files := form.File["file"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "no file uploaded"})
	}

	file, err := files[0].Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	defer file.Close()

	// Declare Context Background, Cloud Name, API Key, API Secret ...
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "waysbeans"})
	if err != nil {
		fmt.Println(err.Error())
	}

	request := profilesdto.CreateProfileRequest{
		Address:  c.FormValue("address"),
		Postcode: c.FormValue("postcode"),
		Phone:    c.FormValue("phone"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	profile := models.Profile{
		Image:    resp.SecureURL,
		Address:  request.Address,
		Postcode: request.Postcode,
		Phone:    request.Phone,
		UserID:   userId,
	}

	profile, err = h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	profile, _ = h.ProfileRepository.GetProfile(profile.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: profile})
}

func (h *handlerProfile) UpdateProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataContex := c.Get("dataFile")
	filename := dataContex.(string)

	request := profilesdto.UpdateProfileRequest{
		Address:  c.FormValue("address"),
		Postcode: c.FormValue("postcode"),
		Phone:    c.FormValue("phone"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	profile, _ := h.ProfileRepository.GetProfile(id)

	profile.Address = request.Address
	profile.Postcode = request.Postcode
	profile.Phone = request.Phone

	if filename != "false" {
		profile.Image = filename
	}

	profile, err = h.ProfileRepository.UpdateProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: profile})
}

func (h *handlerProfile) DeleteProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	deleteProfile, err := h.ProfileRepository.DeleteProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "Success", Data: deleteProfile})
}

func convertResponseProfile(u models.Profile) profilesdto.ProfileResponse {
	return profilesdto.ProfileResponse{
		ID:       u.ID,
		Image:    u.Image,
		Address:  u.Address,
		Postcode: u.Postcode,
		Phone:    u.Phone,
		UserID:   u.User.ID,
	}
}
