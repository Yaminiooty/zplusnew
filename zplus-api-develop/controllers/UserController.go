package controllers

import (
	"net/http"
	"sec-tool/dto"
	"sec-tool/logger"
	"sec-tool/services"
	"sec-tool/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// @Summary Register User
// @Tags User
// @Param user body  dto.CreteNewUserDTO true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /users [post]
func (uc *UserController) CreateNewUser(ctx *gin.Context) {
	logger.Info("UserController", "CreateNewUser", "Create new user process stared", ctx.Request.Header.Get("X-Request-ID"))
	var newUserDto dto.CreteNewUserDTO
	err := ctx.ShouldBindJSON(&newUserDto)
	if err != nil {
		logger.Error("UserController", "CreateNewUser", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		ErrorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, ErrorResponseBody, ctx)
		return
	}
	logger.Debug("UserController", "CreateNewUser", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.CreateNewUser(newUserDto, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "CreateNewUser", "Finished processing create user request", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary User Login
// @Tags User
// @Param user body  dto.UserDTO true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 401 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /login [post]
func (uc *UserController) LoginWithEmailAndPassword(ctx *gin.Context) {
	logger.Info("UserController", "LoginWithEmailAndPassword", "Login with email and password process stared", ctx.Request.Header.Get("X-Request-ID"))
	var inputUser dto.UserDTO
	err := ctx.ShouldBindJSON(&inputUser)
	if err != nil {
		logger.Error("UserController", "LoginWithEmailAndPassword", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		ErrorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, ErrorResponseBody, ctx)
		return
	}
	logger.Debug("UserController", "LoginWithEmailAndPassword", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.LoginWithEmailAndPassword(inputUser, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "LoginWithEmailAndPassword", "Finished processing login request", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary List of all registered users
// @Tags User
// @Security Bearer
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /users [get]
func (uc *UserController) FindAll(ctx *gin.Context) {
	logger.Info("UserController", "FindAll", "FindAll process stared", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.FindAll(ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "FindAll", "FindAll process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Find User By Id
// @Tags User
// @Security Bearer
// @Param id path  string true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /users/{id} [get]
func (uc *UserController) FindById(ctx *gin.Context) {
	logger.Info("UserController", "FindById", "FindById process stared", ctx.Request.Header.Get("X-Request-ID"))
	objectId := ctx.Param("id")
	Message, StatusCode, Data, Error := uc.UserService.FindById(objectId, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "FindById", "FindById process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Delete User By Id
// @Tags User
// @Security Bearer
// @Param id path  string true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /users/{id} [delete]
func (uc *UserController) DeleteById(ctx *gin.Context) {
	logger.Info("UserController", "DeleteById", "DeleteById process stared", ctx.Request.Header.Get("X-Request-ID"))
	objectId := ctx.Param("id")
	Message, StatusCode, Data, Error := uc.UserService.DeleteById(objectId, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "DeleteById", "DeleteById process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Resetting the password when the user selects forget the password.
// @Tags User
// @Param reset_password body  dto.ResetPasswordDTO true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /reset_password [post]
func (uc *UserController) ResetPassword(ctx *gin.Context) {
	logger.Info("UserController", "ResetPassword", "ResetPassword process stared", ctx.Request.Header.Get("X-Request-ID"))
	var resetPasswordInput dto.ResetPasswordDTO
	err := ctx.ShouldBindJSON(&resetPasswordInput)
	if err != nil {
		logger.Debug("UserController", "ResetPassword", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		errorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, errorResponseBody, ctx)
		return
	}
	logger.Debug("UserController", "ResetPassword", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.ResetPassword(resetPasswordInput, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "ResetPassword", "ResetPassword process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Updates the user password. Used internally by the reset password functionality.
// @Tags User
// @Param update_password body  dto.UpdatePasswordDTO true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /update_password [patch]
func (uc *UserController) UpdatePassword(ctx *gin.Context) {
	logger.Info("UserController", "UpdatePassword", "UpdatePassword process stared", ctx.Request.Header.Get("X-Request-ID"))
	var updatePasswordInput dto.UpdatePasswordDTO
	err := ctx.ShouldBindJSON(&updatePasswordInput)
	if err != nil {
		logger.Debug("UserController", "UpdatePassword", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		errorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, errorResponseBody, ctx)
		return
	}
	logger.Debug("UserController", "UpdatePassword", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.UpdatePassword(updatePasswordInput, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "UpdatePassword", "UpdatePassword process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary request new access token.
// @Tags User
// @Security Bearer
// @Success 200 {object} utils.ApiResponse
// @Failure 401 {object} utils.ApiResponse
// @Router /access_token [get]
func (uc *UserController) NewAccessToken(ctx *gin.Context) {
	logger.Info("UserController", "NewAccessToken", "Getting new access token process started", ctx.Request.Header.Get("X-Request-ID"))
	AuthorizationHeader := ctx.Request.Header["Authorization"]
	Message, StatusCode, Data, Error := uc.UserService.NewAccessToken(AuthorizationHeader[0], ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "NewAccessToken", "NewAccessToken process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Account verification post registration.
// @Tags User
// @Param verify_account body  dto.VerifyAccountDTO true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /verify_account [post]
func (uc *UserController) VerifyAccount(ctx *gin.Context) {
	logger.Info("UserController", "VerifyAccount", "Verify account process started", ctx.Request.Header.Get("X-Request-ID"))
	var verifyAccountDTO dto.VerifyAccountDTO
	err := ctx.ShouldBindJSON(&verifyAccountDTO)
	if err != nil {
		logger.Debug("UserController", "VerifyAccount", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		errorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, errorResponseBody, ctx)
		return
	}
	logger.Debug("UserController", "VerifyAccount", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.VerifyAccount(verifyAccountDTO, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "VerifyAccount", "VerifyAccount process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Get verfication code on email.
// @Tags User
// @Param verification_code body  dto.GetVerificationCodeDTO true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /get_verification_code [post]
func (uc *UserController) GetVerificationCode(ctx *gin.Context) {
	logger.Info("UserController", "GetVerificationCode", "GetVerificationCode process stared", ctx.Request.Header.Get("X-Request-ID"))
	var getVerificationCodeInput dto.GetVerificationCodeDTO
	err := ctx.ShouldBindJSON(&getVerificationCodeInput)
	if err != nil {
		logger.Debug("UserController", "GetVerificationCode", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		errorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, errorResponseBody, ctx)
		return
	}
	logger.Debug("UserController", "GetVerificationCode", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.GetVerificationCode(getVerificationCodeInput, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "GetVerificationCode", "GetVerificationCode process finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Logout the user.
// @Tags User
// @Security Bearer
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /logout [get]
func (uc *UserController) Logout(ctx *gin.Context) {
	logger.Info("UserController", "Logout", "Logout Started", ctx.Request.Header.Get("X-Request-ID"))
	AuthorizationHeader := ctx.Request.Header["Authorization"]
	jwtTokenStr := strings.Split(AuthorizationHeader[0], "Bearer ")[1]
	Message, StatusCode, Data, Error := uc.UserService.Logout(jwtTokenStr, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "Logout", "Logout Finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Get logged in user details.
// @Tags User
// @Security Bearer
// @Success 200 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /user [get]
func (uc *UserController) GetUserDetails(ctx *gin.Context) {
	logger.Info("UserController", "GetUserDetails", "Get user details started", ctx.Request.Header.Get("X-Request-ID"))
	user := ctx.Request.Header.Get("user")
	Message, StatusCode, Data, Error := uc.UserService.GetUserDetails(user, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "GetUserDetails", "Get user details finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary Update the name and phone of the user.
// @Tags User
// @Security Bearer
// @Param update_user body  dto.UpdateUserDetails true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /update_user_details [patch]
func (uc *UserController) UpdateUserDetails(ctx *gin.Context) {
	logger.Info("UserController", "GetUserDetails", "Get user details started", ctx.Request.Header.Get("X-Request-ID"))
	user := ctx.Request.Header.Get("user")
	var updatedUserData dto.UpdateUserDetails
	err := ctx.ShouldBindJSON(&updatedUserData)
	if err != nil {
		logger.Debug("UserController", "GetVerificationCode", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		errorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, errorResponseBody, ctx)
		return
	}
	if updatedUserData.Phone != "" {
		if !utils.CheckPhone(updatedUserData.Phone) {
			utils.SendJSONResponse("Invalid Phone.", http.StatusBadRequest, nil, nil, ctx)
			return
		}
	}
	logger.Debug("UserController", "GetVerificationCode", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.UpdateUserDetails(user, updatedUserData, ctx.Request.Header.Get("X-Request-ID"))
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "GetUserDetails", "Get user details finished", ctx.Request.Header.Get("X-Request-ID"))
}

// @Summary update the user password.
// @Tags User
// @Security Bearer
// @Param update_user body  dto.UpdateUserDetails true " "
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /change-password [patch]
func (uc *UserController) ChangeUserPassword(ctx *gin.Context) {
	logger.Info("UserController", "ChangeUserPassword", "Change user password started", ctx.Request.Header.Get("X-Request-ID"))
	user := ctx.Request.Header.Get("user")
	var changePasswordDTO dto.ChangePasswordDTO
	err := ctx.ShouldBindJSON(&changePasswordDTO)
	if err != nil {
		logger.Debug("UserController", "ChangeUserPassword", "Payload verification failed", ctx.Request.Header.Get("X-Request-ID"))
		errorResponseBody := utils.ParseBindingErrors(err)
		utils.SendJSONResponse("Invalid Payload.", http.StatusBadRequest, nil, errorResponseBody, ctx)
		return
	}

	logger.Debug("UserController", "ChangeUserPassword", "Payload verification successful", ctx.Request.Header.Get("X-Request-ID"))
	Message, StatusCode, Data, Error := uc.UserService.ChangeUserPassword(user, changePasswordDTO)
	utils.SendJSONResponse(Message, StatusCode, Data, Error, ctx)
	logger.Info("UserController", "ChangeUserPassword", "Change user password finished", ctx.Request.Header.Get("X-Request-ID"))
}
