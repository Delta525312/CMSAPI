package handler

import (
	"mainPackage/config"
	"mainPackage/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// @summary Get User
// @tags User
// @security ApiKeyAuth
// @id Get User
// @accept json
// @produce json
// @Param start query int false "start" default(0)
// @Param length query int false "length" default(10)
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users [get]
func GetUmUserList(c *gin.Context) {
	logger := config.GetLog()
	orgId := GetVariableFromToken(c, "orgId")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)

	startStr := c.DefaultQuery("start", "0")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 0
	}
	lengthStr := c.DefaultQuery("length", "1000")
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		length = 1000
	}
	logger.Debug(`Query`, zap.Any("start", start))
	logger.Debug(`Query`, zap.Any("length", length))
	query := `SELECT "id","orgId", "displayName", title, "firstName", "middleName", "lastName", "citizenId", bod,
	blood, gender, "mobileNo", address, photo, username, password, email, "roleId", "userType", "empId",
	"deptId", "commId", "stnId", active, "activationToken", "lastActivationRequest", "lostPasswordRequest",
	"signupStamp", islogin, "lastLogin", "createdAt", "updatedAt", "createdBy", "updatedBy" 
	FROM public.um_users 
	WHERE "orgId"=$1 LIMIT $2 OFFSET $3`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err = conn.Query(ctx, query, orgId, length, start)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var errorMsg string
	var u model.Um_User
	var userList []model.Um_User
	rowIndex := 0
	for rows.Next() {
		rowIndex++
		err := rows.Scan(
			&u.ID,
			&u.OrgID,
			&u.DisplayName,
			&u.Title,
			&u.FirstName,
			&u.MiddleName,
			&u.LastName,
			&u.CitizenID,
			&u.Bod,
			&u.Blood,
			&u.Gender,
			&u.MobileNo,
			&u.Address,
			&u.Photo,
			&u.Username,
			&u.Password,
			&u.Email,
			&u.RoleID,
			&u.UserType,
			&u.EmpID,
			&u.DeptID,
			&u.CommID,
			&u.StnID,
			&u.Active,
			&u.ActivationToken,
			&u.LastActivationRequest,
			&u.LostPasswordRequest,
			&u.SignupStamp,
			&u.IsLogin,
			&u.LastLogin,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.CreatedBy,
			&u.UpdatedBy,
		)
		if err != nil {
			logger.Warn("Scan failed", zap.Error(err))
			errorMsg = err.Error()
			response := model.Response{
				Status: "-1",
				Msg:    "Failed",
				Desc:   errorMsg,
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		userList = append(userList, u)
	}
	if errorMsg != "" {
		response := model.Response{
			Status: "-1",
			Msg:    "Failed",
			Desc:   errorMsg,
		}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := model.Response{
			Status: "0",
			Msg:    "Success",
			Data:   userList,
			Desc:   "",
		}
		c.JSON(http.StatusOK, response)
	}
}

// @summary Get User by username
// @tags User
// @security ApiKeyAuth
// @id Get User by username
// @accept json
// @produce json
// @Param username path string true "username"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users/username/{username} [get]
func GetUmUserByUsername(c *gin.Context) {
	logger := config.GetLog()
	username := c.Param("username")
	orgId := GetVariableFromToken(c, "orgId")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	query := `
	SELECT t1."id",t1."orgId", t1."displayName", t1.title, t1."firstName", t1."middleName", t1."lastName",
		t1."citizenId", t1.bod, t1.blood, t1.gender, t1."mobileNo", t1.address, t1.photo, t1.username, t1.password,
		t1.email, t1."roleId", t1."userType", t1."empId", t1."deptId", t1."commId", t1."stnId", t1.active,
		t1."activationToken",t1."lastActivationRequest", t1."lostPasswordRequest", t1."signupStamp",
		t1.islogin, t1."lastLogin",t1."createdAt", t1."updatedAt", t1."createdBy", t1."updatedBy",t2.name,t3."roleName"
		FROM public.um_users t1
		JOIN public.organizations t2 ON t1."orgId" = t2.id
		JOIN public.um_roles t3 ON t1."roleId" = t3.id
		WHERE t1.username=$1 AND t1."orgId"=$2`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query),
		zap.Any("Input", []any{
			username, orgId,
		}))
	rows, err := conn.Query(ctx, query, username, orgId)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var errorMsg string
	var u model.Um_User
	if rows.Next() {
		err = rows.Scan(&u.ID,
			&u.OrgID, &u.DisplayName, &u.Title, &u.FirstName, &u.MiddleName, &u.LastName, &u.CitizenID, &u.Bod, &u.Blood,
			&u.Gender, &u.MobileNo, &u.Address, &u.Photo, &u.Username, &u.Password, &u.Email, &u.RoleID, &u.UserType,
			&u.EmpID, &u.DeptID, &u.CommID, &u.StnID, &u.Active, &u.ActivationToken, &u.LastActivationRequest,
			&u.LostPasswordRequest, &u.SignupStamp, &u.IsLogin, &u.LastLogin, &u.CreatedAt, &u.UpdatedAt,
			&u.CreatedBy, &u.UpdatedBy, &u.OrgName, &u.RoleName,
		)
		if err != nil {
			errorMsg = err.Error()
			logger.Warn("Scan failed", zap.Error(err))
			response := model.Response{
				Status: "-1",
				Msg:    "Failed",
				Desc:   errorMsg,
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}
	response := model.Response{
		Status: "0",
		Msg:    "Success",
		Data:   u,
		Desc:   "",
	}
	c.JSON(http.StatusOK, response)
}

// @summary Get User by Id
// @tags User
// @security ApiKeyAuth
// @id Get User by Id
// @accept json
// @produce json
// @Param id path int true "id"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users/{id} [get]
func GetUmUserById(c *gin.Context) {
	logger := config.GetLog()
	id := c.Param("id")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	orgId := GetVariableFromToken(c, "orgId")
	defer cancel()
	defer conn.Close(ctx)
	query := `SELECT "orgId", "displayName", title, "firstName", "middleName", "lastName", "citizenId", bod, blood, gender, "mobileNo", address, photo, username, password, email, "roleId", "userType", "empId", "deptId", "commId", "stnId", active, "activationToken", "lastActivationRequest", "lostPasswordRequest", "signupStamp", islogin, "lastLogin", "createdAt", "updatedAt", "createdBy", "updatedBy" 
	FROM public.um_users WHERE id=$1 AND "orgId"=$2`

	var u model.Um_User
	logger.Debug(`Query`, zap.String("query", query))
	err := conn.QueryRow(ctx, query, id, orgId).Scan(
		&u.OrgID,
		&u.DisplayName,
		&u.Title,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.CitizenID,
		&u.Bod,
		&u.Blood,
		&u.Gender,
		&u.MobileNo,
		&u.Address,
		&u.Photo,
		&u.Username,
		&u.Password,
		&u.Email,
		&u.RoleID,
		&u.UserType,
		&u.EmpID,
		&u.DeptID,
		&u.CommID,
		&u.StnID,
		&u.Active,
		&u.ActivationToken,
		&u.LastActivationRequest,
		&u.LostPasswordRequest,
		&u.SignupStamp,
		&u.IsLogin,
		&u.LastLogin,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.CreatedBy,
		&u.UpdatedBy,
	)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}

	response := model.Response{
		Status: "0",
		Msg:    "Success",
		Data:   u,
		Desc:   "",
	}
	c.JSON(http.StatusOK, response)

}

// @summary Create User
// @tags User
// @security ApiKeyAuth
// @id Create User
// @accept json
// @produce json
// @param Case body model.UserInput true "User to be created"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users/add [post]
func UserAdd(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)

	var req model.UserInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// now req is ready to use

	var enc string
	var err error
	var id int
	enc, err = encrypt(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	orgId := GetVariableFromToken(c, "orgId")
	username := GetVariableFromToken(c, "username")
	tokenString := GetVariableFromToken(c, "tokenString")
	now := time.Now()
	query := `
		INSERT INTO public.um_users(
		"orgId", "displayName", title, "firstName", "middleName", "lastName", "citizenId", bod, blood, gender,
		"mobileNo", address, photo, username, password, email, "roleId", "userType", "empId", "deptId",
		"commId", "stnId", active, "activationToken", "lastActivationRequest", "lostPasswordRequest",
		"signupStamp", islogin, "lastLogin", "createdAt", "updatedAt", "createdBy", "updatedBy"
			)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7,
		$8, $9, $10, $11,
		$12, $13, $14, $15, $16,
		$17, $18, $19, $20, $21, 
		$22, $23, $24, $25, $26,
		$27, $28, $29, $30, $31, 
		$32, $33
	)
	RETURNING id;
	`
	logger.Debug(`Query`, zap.String("query", query))
	logger.Debug(`request input`, zap.Any("Input", []any{req}))
	logger.Debug(`Encrypt Password :` + enc)
	err = conn.QueryRow(ctx, query,
		orgId, req.DisplayName, req.Title, req.FirstName, req.MiddleName,
		req.LastName, req.CitizenID, req.Bod, req.Blood,
		req.Gender, req.MobileNo, req.Address, req.Photo, req.Username,
		enc, req.Email, req.RoleID, req.UserType, req.EmpID, req.DeptID, req.CommID, req.StnID,
		req.Active, tokenString, req.LastActivationRequest, req.LostPasswordRequest, req.SignupStamp,
		req.IsLogin, req.LastLogin, now, now, username, username,
	).Scan(&id)

	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Create successfully",
	})
}

// @summary Update User
// @tags User
// @security ApiKeyAuth
// @id Update User
// @accept json
// @produce json
// @Param id path int true "id"
// @param Body body model.UserUpdate true "Data Update"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users/{id} [patch]
func UserUpdate(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)

	var req model.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// now req is ready to use

	var enc string
	var err error
	id := c.Param("id")
	enc, err = encrypt(req.Password)
	if err != nil {
		return
	}
	orgId := GetVariableFromToken(c, "orgId")
	username := GetVariableFromToken(c, "username")
	now := time.Now()
	query := `
	UPDATE public.um_users
	SET "displayName"=$1, title=$2, "firstName"=$3, "middleName"=$4, "lastName"=$5, "citizenId"=$6,
	bod=$7, blood=$8, gender=$9, "mobileNo"=$10, address=$11, photo=$12, username=$13, password=$14, email=$15, "roleId"=$16,
	"userType"=$17, "empId"=$18, "deptId"=$19, "commId"=$20, "stnId"=$21, active=$22,
	"lastActivationRequest"=$23, "lostPasswordRequest"=$24, "signupStamp"=$25, islogin=$26, "lastLogin"=$27,
	"updatedAt"=$28,"updatedBy"=$29 WHERE id = $30 AND "orgId"=$31`

	logger.Debug(`Query`, zap.String("query", query))
	logger.Debug(`request input`, zap.Any("Input", []any{req}))
	logger.Debug(`Encrypt Password :` + enc)
	_, err = conn.Exec(ctx, query,
		req.DisplayName, req.Title, req.FirstName, req.MiddleName,
		req.LastName, req.CitizenID, req.Bod, req.Blood,
		req.Gender, req.MobileNo, req.Address, req.Photo, req.Username,
		enc, req.Email, req.RoleID, req.UserType, req.EmpID, req.DeptID, req.CommID, req.StnID,
		req.Active, req.LastActivationRequest, req.LostPasswordRequest, req.SignupStamp,
		req.IsLogin, req.LastLogin, now, username, id, orgId,
	)

	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Update successfully",
	})
}

// @summary Update User By Username
// @tags User
// @security ApiKeyAuth
// @id Update User By Username
// @accept json
// @produce json
// @Param username path string true "username"
// @param Body body model.UserUpdate true "Data Update"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users/username/{username} [patch]
func UserUpdateByUsername(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)

	var req model.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// now req is ready to use

	var enc string
	var err error
	id := c.Param("username")
	enc, err = encrypt(req.Password)
	if err != nil {
		return
	}
	orgId := GetVariableFromToken(c, "orgId")
	username := GetVariableFromToken(c, "username")
	now := time.Now()
	query := `
	UPDATE public.um_users
	SET "displayName"=$1, title=$2, "firstName"=$3, "middleName"=$4, "lastName"=$5, "citizenId"=$6,
	bod=$7, blood=$8, gender=$9, "mobileNo"=$10, address=$11, photo=$12, username=$13, password=$14, email=$15, "roleId"=$16,
	"userType"=$17, "empId"=$18, "deptId"=$19, "commId"=$20, "stnId"=$21, active=$22,
	"lastActivationRequest"=$23, "lostPasswordRequest"=$24, "signupStamp"=$25, islogin=$26, "lastLogin"=$27,
	"updatedAt"=$28,"updatedBy"=$29 WHERE username = $30 AND "orgId"=$31`

	logger.Debug(`Query`, zap.String("query", query))
	logger.Debug(`request input`, zap.Any("Input", []any{req}))
	logger.Debug(`Encrypt Password :` + enc)
	_, err = conn.Exec(ctx, query,
		req.DisplayName, req.Title, req.FirstName, req.MiddleName,
		req.LastName, req.CitizenID, req.Bod, req.Bod, req.Blood,
		req.Gender, req.MobileNo, req.Address, req.Photo, req.Username,
		enc, req.Email, req.RoleID, req.UserType, req.EmpID, req.DeptID, req.CommID, req.StnID,
		req.Active, req.LastActivationRequest, req.LostPasswordRequest, req.SignupStamp,
		req.IsLogin, req.LastLogin, now, username, id, orgId,
	)

	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusUnauthorized, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Update successfully",
	})
}

// @summary Delete User
// @tags User
// @security ApiKeyAuth
// @id Delete User
// @accept json
// @produce json
// @Param id path int true "id"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users/{id} [delete]
func UserDelete(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	id := c.Param("id")
	orgId := GetVariableFromToken(c, "orgId")
	query := `DELETE FROM public."um_users" WHERE id = $1 AND "orgId"=$2`
	logger.Debug("Query", zap.String("query", query), zap.Any("id", id))
	_, err := conn.Exec(ctx, query, id, orgId)
	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Update failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Delete successfully",
	})
}

// @summary Get User with skills
// @tags User
// @security ApiKeyAuth
// @id Get User with skills
// @accept json
// @produce json
// @Param start query int false "start" default(0)
// @Param length query int false "length" default(10)
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_skills [get]
func GetUserWithSkills(c *gin.Context) {
	logger := config.GetLog()
	orgId := GetVariableFromToken(c, "orgId")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)

	startStr := c.DefaultQuery("start", "0")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 0
	}
	lengthStr := c.DefaultQuery("length", "1000")
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		length = 1000
	}

	query := `SELECT "orgId", "userName", "skillId", active, "createdAt", "updatedAt", "createdBy", "updatedBy" 
	FROM public.um_user_with_skills WHERE "orgId"=$1 LIMIT $2 OFFSET $3`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err = conn.Query(ctx, query, orgId, length, start)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var errorMsg string
	var u model.UserSkill
	var userList []model.UserSkill
	rowIndex := 0
	for rows.Next() {
		rowIndex++
		err := rows.Scan(
			&u.OrgID,
			&u.UserName,
			&u.SkillID,
			&u.Active,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.CreatedBy,
			&u.UpdatedBy,
		)

		if err != nil {
			logger.Warn("Scan failed", zap.Error(err))
			response := model.Response{
				Status: "-1",
				Msg:    "Failed",
				Desc:   errorMsg,
			}
			c.JSON(http.StatusInternalServerError, response)
		}
		userList = append(userList, u)
	}
	if errorMsg != "" {
		response := model.Response{
			Status: "-1",
			Msg:    "Failed",
			Desc:   errorMsg,
		}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := model.Response{
			Status: "0",
			Msg:    "Success",
			Data:   userList,
			Desc:   "",
		}
		c.JSON(http.StatusOK, response)
	}
}

// @summary Get User with skills by id
// @tags User
// @security ApiKeyAuth
// @id Get User with skills by id
// @Param id path int true "id"
// @accept json
// @produce json
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_skills/{id} [get]
func GetUserWithSkillsById(c *gin.Context) {
	logger := config.GetLog()
	id := c.Param("id")
	orgId := GetVariableFromToken(c, "orgId")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	query := `SELECT "orgId", "userName", "skillId", active, "createdAt", "updatedAt", "createdBy", "updatedBy" 
	FROM public.um_user_with_skills WHERE id=$1 AND "orgId"=$2`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err := conn.Query(ctx, query, id, orgId)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var u model.UserSkill
	err = rows.Scan(
		&u.OrgID,
		&u.UserName,
		&u.SkillID,
		&u.Active,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.CreatedBy,
		&u.UpdatedBy,
	)

	if err != nil {
		logger.Warn("Scan failed", zap.Error(err))
		response := model.Response{
			Status: "-1",
			Msg:    "Failed",
			Desc:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := model.Response{
		Status: "0",
		Msg:    "Success",
		Data:   u,
		Desc:   "",
	}
	c.JSON(http.StatusOK, response)

}

// @summary Get User with skills by skill id
// @tags User
// @security ApiKeyAuth
// @id Get User with skills by skillId
// @Param skillId path string true "skillId"
// @accept json
// @produce json
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_skills/skillId/{skillId} [get]
func GetUserWithSkillsBySkillId(c *gin.Context) {
	logger := config.GetLog()
	skillId := c.Param("skillId")
	orgId := GetVariableFromToken(c, "orgId")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	query := `SELECT "orgId", "userName", "skillId", active, "createdAt", "updatedAt", "createdBy", "updatedBy" 
	FROM public.um_user_with_skills WHERE "skillId" = $1 AND "orgId" = $2`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err := conn.Query(ctx, query, skillId, orgId)
	var userList []model.UserSkill
	var u model.UserSkill
	var errorMsg string
	rowIndex := 0
	for rows.Next() {
		rowIndex++
		err := rows.Scan(
			&u.OrgID,
			&u.UserName,
			&u.SkillID,
			&u.Active,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.CreatedBy,
			&u.UpdatedBy,
		)

		if err != nil {
			logger.Warn("Scan failed", zap.Error(err))
			response := model.Response{
				Status: "-1",
				Msg:    "Failed",
				Desc:   errorMsg,
			}
			c.JSON(http.StatusInternalServerError, response)
		}
		userList = append(userList, u)
	}
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	// 	logger.Warn("Scan failed", zap.Error(err))
	// 	response := model.Response{
	// 		Status: "-1",
	// 		Msg:    "Failed",
	// 		Desc:   err.Error(),
	// 	}
	// 	c.JSON(http.StatusInternalServerError, response)
	// 	return
	// }

	response := model.Response{
		Status: "0",
		Msg:    "Success",
		Data:   userList,
		Desc:   "",
	}
	c.JSON(http.StatusOK, response)

}

// @summary Create User with skill
// @id Create User with skill
// @security ApiKeyAuth
// @tags User
// @accept json
// @produce json
// @param Body body model.UserSkillInsert true "Create Data"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_skills/add [post]
func InsertUserWithSkills(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()
	username := GetVariableFromToken(c, "username")
	var req model.UserSkillInsert
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// now req is ready to use
	now := time.Now()
	var id int
	query := `
	INSERT INTO public."um_user_with_skills"(
	"orgId", "userName", "skillId", active, "createdAt", "updatedAt", "createdBy", "updatedBy")
	VALUES ($1, $2, $3, $4, $5, $6, $7,$8)
	RETURNING id ;
	`

	err := conn.QueryRow(ctx, query,
		req.OrgID, req.UserName, req.SkillID, req.Active, now,
		now, username, username).Scan(&id)

	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Create successfully",
	})

}

// @summary Update User with skill
// @id Update User with skill
// @security ApiKeyAuth
// @accept json
// @produce json
// @tags User
// @Param id path int true "id"
// @param Body body model.UserSkillUpdate true "Update data"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_skills/{id} [patch]
func UpdateUserWithSkills(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()

	id := c.Param("id")

	var req model.UserSkillUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Update failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.UpdateCaseResponse{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
			ID:     ToInt(id),
		})
		return
	}
	username := GetVariableFromToken(c, "username")
	now := time.Now()
	query := `UPDATE public."um_user_with_skills"
	SET 
    "skillId"=$2,
    active=$3,
	"updatedAt"=$4,
	"updatedBy"=$5
	WHERE id = $1 `
	_, err := conn.Exec(ctx, query,
		id, req.SkillID, req.Active,
		now, username, username,
	)
	logger.Debug("Update Case SQL Args",
		zap.String("query", query),
		zap.Any("Input", []any{
			id, req.SkillID, req.Active,
			now, username,
		}))
	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Update failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Update successfully",
	})
}

// @summary Delete User with skill
// @id Delete User with skill
// @security ApiKeyAuth
// @accept json
// @tags User
// @produce json
// @Param id path int true "id"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_skills/{id} [delete]
func DeleteUserWithSkills(c *gin.Context) {

	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()

	id := c.Param("id")
	query := `DELETE FROM public."um_user_with_skills" WHERE id = $1 `
	logger.Debug("Query", zap.String("query", query), zap.Any("id", id))
	_, err := conn.Exec(ctx, query, id)
	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Update failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Delete successfully",
	})
}

// @summary Get User with contacts
// @tags User
// @security ApiKeyAuth
// @id Get User with contacts
// @accept json
// @produce json
// @Param start query int false "start" default(0)
// @Param length query int false "length" default(10)
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_contacts [get]
func GetUserWithContacts(c *gin.Context) {
	logger := config.GetLog()

	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	orgId := GetVariableFromToken(c, "orgId")
	defer cancel()
	defer conn.Close(ctx)

	startStr := c.DefaultQuery("start", "0")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 0
	}
	lengthStr := c.DefaultQuery("length", "1000")
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		length = 1000
	}

	query := `SELECT  "orgId", username, "contactName", "contactPhone", "contactAddr", "createdAt", "updatedAt", "createdBy", "updatedBy" 	
	FROM public.um_user_contacts WHERE "orgId"=$1 LIMIT $2 OFFSET $3`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err = conn.Query(ctx, query, orgId, length, start)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var errorMsg string
	var u model.UserContact
	var userList []model.UserContact
	rowIndex := 0
	for rows.Next() {
		rowIndex++
		err := rows.Scan(
			&u.OrgID,
			&u.Username,
			&u.ContactName,
			&u.ContactPhone,
			&u.ContactAddr,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.CreatedBy,
			&u.UpdatedBy,
		)
		if err != nil {
			logger.Warn("Scan failed", zap.Error(err))
			response := model.Response{
				Status: "-1",
				Msg:    "Failed",
				Desc:   errorMsg,
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		userList = append(userList, u)
	}
	if errorMsg != "" {
		response := model.Response{
			Status: "-1",
			Msg:    "Failed",
			Desc:   errorMsg,
		}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := model.Response{
			Status: "0",
			Msg:    "Success",
			Data:   userList,
			Desc:   "",
		}
		c.JSON(http.StatusOK, response)
	}
}

// @summary Get User with contacts by id
// @tags User
// @security ApiKeyAuth
// @id Get User with contacts by id
// @Param id path int true "id"
// @accept json
// @produce json
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_contacts/{id} [get]
func GetUserWithContactsById(c *gin.Context) {
	logger := config.GetLog()
	id := c.Param("id")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	orgId := GetVariableFromToken(c, "orgId")
	query := `SELECT  "orgId", username, "contactName", "contactPhone", "contactAddr", "createdAt", "updatedAt", "createdBy", "updatedBy" 
	FROM public.um_user_contacts WHERE id=$1 AND "orgId"=$2`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err := conn.Query(ctx, query, id, orgId)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var u model.UserContact

	err = rows.Scan(
		&u.OrgID,
		&u.Username,
		&u.ContactName,
		&u.ContactPhone,
		&u.ContactAddr,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.CreatedBy,
		&u.UpdatedBy,
	)

	if err != nil {
		logger.Warn("Scan failed", zap.Error(err))
		response := model.Response{
			Status: "-1",
			Msg:    "Failed",
			Desc:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := model.Response{
		Status: "0",
		Msg:    "Success",
		Data:   u,
		Desc:   "",
	}
	c.JSON(http.StatusOK, response)

}

// @summary Create User with contacts
// @id Create User with contacts
// @security ApiKeyAuth
// @tags User
// @accept json
// @produce json
// @param Body body model.UserContactInsert true "Create Data"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_contacts/add [post]
func InsertUserWithContacts(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()

	var req model.UserContactInsert
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	now := time.Now()
	username := GetVariableFromToken(c, "username")
	var id int
	query := `
	INSERT INTO public."um_user_contacts"(
	"orgId", username, "contactName", "contactPhone", "contactAddr", "createdAt", "updatedAt", "createdBy", "updatedBy")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id ;
	`

	err := conn.QueryRow(ctx, query,
		req.OrgID, req.Username, req.ContactName, req.ContactPhone, req.ContactAddr, now,
		now, username, username).Scan(&id)

	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Create successfully",
	})

}

// @summary Update User with contacts
// @id Update User with contacts
// @security ApiKeyAuth
// @accept json
// @produce json
// @tags User
// @Param id path int true "id"
// @param Body body model.UserContactInsertUpdate true "Update data"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_contacts/{id} [patch]
func UpdateUserWithContacts(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()

	id := c.Param("id")

	var req model.UserContactInsertUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Update failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	username := GetVariableFromToken(c, "username")
	orgId := GetVariableFromToken(c, "orgId")
	now := time.Now()
	query := `UPDATE public."um_user_contacts"
	SET "contactName"=$2, "contactPhone"=$3, "contactAddr"=$4,"updatedAt"=$5,"updatedBy"=$6
	WHERE id = $1 AND "orgId"=$7`
	_, err := conn.Exec(ctx, query, id,
		req.ContactName, req.ContactPhone, req.ContactAddr, now, username, orgId)

	logger.Debug("Update Case SQL Args",
		zap.String("query", query),
		zap.Any("Input", []any{
			id, req.ContactName, req.ContactPhone, req.ContactAddr, now, username, orgId,
		}))
	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Update failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Update successfully",
	})
}

// @summary Delete User with contacts
// @id Delete User with contacts
// @security ApiKeyAuth
// @accept json
// @tags User
// @produce json
// @Param id path int true "id"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_contacts/{id} [delete]
func DeleteUserWithContacts(c *gin.Context) {

	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()

	id := c.Param("id")
	orgId := GetVariableFromToken(c, "orgId")
	query := `DELETE FROM public."um_user_contacts" WHERE id = $1 AND "orgId"=$2`
	logger.Debug("Query", zap.String("query", query), zap.Any("id", id))
	_, err := conn.Exec(ctx, query, id, orgId)
	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Update failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Delete successfully",
	})
}

// @summary Get User with socials
// @tags User
// @security ApiKeyAuth
// @id Get User with socials
// @accept json
// @produce json
// @Param start query int false "start" default(0)
// @Param length query int false "length" default(10)
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_socials [get]
func GetUserWithSocials(c *gin.Context) {
	logger := config.GetLog()

	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	orgId := GetVariableFromToken(c, "orgId")
	startStr := c.DefaultQuery("start", "0")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 0
	}
	lengthStr := c.DefaultQuery("length", "1000")
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		length = 1000
	}

	query := `SELECT  "orgId", username, "socialType", "socialId", "socialName", "createdAt", "updatedAt", "createdBy", "updatedBy" 	
	FROM public.um_user_with_socials WHERE "orgId"=$1 
	LIMIT $2 OFFSET $3`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err = conn.Query(ctx, query, orgId, length, start)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var errorMsg string
	var u model.UserSocial
	var userList []model.UserSocial
	rowIndex := 0
	for rows.Next() {
		rowIndex++
		err := rows.Scan(
			&u.OrgID,
			&u.Username,
			&u.SocialType,
			&u.SocialID,
			&u.SocialName,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.CreatedBy,
			&u.UpdatedBy,
		)
		if err != nil {
			logger.Warn("Scan failed", zap.Error(err))
			response := model.Response{
				Status: "-1",
				Msg:    "Failed",
				Desc:   errorMsg,
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		userList = append(userList, u)
	}
	if errorMsg != "" {
		response := model.Response{
			Status: "-1",
			Msg:    "Failed",
			Desc:   errorMsg,
		}
		c.JSON(http.StatusInternalServerError, response)
	} else {
		response := model.Response{
			Status: "0",
			Msg:    "Success",
			Data:   userList,
			Desc:   "",
		}
		c.JSON(http.StatusOK, response)
	}
}

// @summary Get User with Socials by id
// @tags User
// @security ApiKeyAuth
// @id Get User with Socials by id
// @Param id path int true "id"
// @accept json
// @produce json
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_socials/{id} [get]
func GetUserWithSocialsById(c *gin.Context) {
	logger := config.GetLog()
	id := c.Param("id")
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	orgId := GetVariableFromToken(c, "orgId")
	query := `SELECT  "orgId", username, "socialType", "socialId", "socialName", "createdAt", "updatedAt", "createdBy", "updatedBy" 	
	FROM public.um_user_with_socials WHERE id=$1 AND "orgId"=$2`

	var rows pgx.Rows
	logger.Debug(`Query`, zap.String("query", query))
	rows, err := conn.Query(ctx, query, id, orgId)
	if err != nil {
		logger.Warn("Query failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	defer rows.Close()
	var u model.UserSocial
	err = rows.Scan(
		&u.OrgID,
		&u.Username,
		&u.SocialType,
		&u.SocialID,
		&u.SocialName,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.CreatedBy,
		&u.UpdatedBy,
	)

	if err != nil {
		logger.Warn("Scan failed", zap.Error(err))
		response := model.Response{
			Status: "-1",
			Msg:    "Failed",
			Desc:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
	}

	response := model.Response{
		Status: "0",
		Msg:    "Success",
		Data:   u,
		Desc:   "",
	}
	c.JSON(http.StatusOK, response)

}

// @summary Create User with socials
// @id Create User with socials
// @security ApiKeyAuth
// @tags User
// @accept json
// @produce json
// @param Body body model.UserSocialInsert true "Create Data"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_socials/add [post]
func InsertUserWithSocials(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()

	var req model.UserSocialInsert
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}
	username := GetVariableFromToken(c, "username")
	// now req is ready to use
	now := time.Now()
	var id int
	query := `
	INSERT INTO public."um_user_with_socials"(
	"orgId", username, "socialType", "socialId", "socialName", "createdAt", "updatedAt", "createdBy", "updatedBy")
	VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9)
	RETURNING id ;
	`

	err := conn.QueryRow(ctx, query,
		req.OrgID, req.Username, req.SocialType, req.SocialID, req.SocialName, now,
		now, username, username).Scan(&id)

	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Insert failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Create successfully",
	})

}

// @summary Update User with socials
// @id Update User with socials
// @security ApiKeyAuth
// @accept json
// @produce json
// @tags User
// @Param id path int true "id"
// @param Body body model.UserSocialUpdate true "Update data"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_socials/{id} [patch]
func UpdateUserWithSocials(c *gin.Context) {
	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()

	id := c.Param("id")

	var req model.UserSocialUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Update failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		return
	}
	now := time.Now()
	username := GetVariableFromToken(c, "username")
	orgId := GetVariableFromToken(c, "orgId")
	query := `UPDATE public."um_user_with_socials"
	SET "orgId"=$2, username=$3, "socialType"=$4, "socialId"=$5, "socialName"=$6, "updatedAt"=$7, "updatedBy"=$8
	WHERE id = $1 AND "orgId"=$9`
	_, err := conn.Exec(ctx, query,
		id, req.OrgID, req.Username, req.SocialType, req.SocialID, req.SocialName,
		now, username, orgId,
	)
	logger.Debug("Update Case SQL Args",
		zap.String("query", query),
		zap.Any("Input", []any{
			id,
			req.OrgID, req.Username, req.SocialType, req.SocialID, req.SocialName,
			now, username, orgId,
		}))
	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Update failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Update successfully",
	})
}

// @summary Delete User with socials
// @id Delete User with socials
// @security ApiKeyAuth
// @accept json
// @tags User
// @produce json
// @Param id path int true "id"
// @response 200 {object} model.Response "OK - Request successful"
// @Router /api/v1/users_with_socials/{id} [delete]
func DeleteUserWithSocials(c *gin.Context) {

	logger := config.GetLog()
	conn, ctx, cancel := config.ConnectDB()
	if conn == nil {
		return
	}
	defer cancel()
	defer conn.Close(ctx)
	defer cancel()
	orgId := GetVariableFromToken(c, "orgId")
	id := c.Param("id")
	query := `DELETE FROM public."um_user_with_socials" WHERE id = $1 AND "orgId"=$2`
	logger.Debug("Query", zap.String("query", query), zap.Any("id", id))
	_, err := conn.Exec(ctx, query, id, orgId)
	if err != nil {
		// log.Printf("Insert failed: %v", err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status: "-1",
			Msg:    "Failure",
			Desc:   err.Error(),
		})
		logger.Warn("Update failed", zap.Error(err))
		return
	}

	// Continue logic...
	c.JSON(http.StatusOK, model.Response{
		Status: "0",
		Msg:    "Success",
		Desc:   "Delete successfully",
	})
}
