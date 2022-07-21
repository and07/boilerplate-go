package httpserver

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/service/mail"
	"github.com/and07/boilerplate-go/pkg/utils"
)

// RefreshToken handles refresh token request
func (ah *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(UserKey{}).(data.User)
	accessToken, err := ah.jwtManager.GenerateAccessToken(&user)
	if err != nil {
		ah.logger.Error("unable to generate access token", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(&GenericError{Error: err.Error()}, w)
		data.ToJSON(&GenericResponse{Status: false, Message: "Unable to generate access token.Please try again later"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	// data.ToJSON(&TokenResponse{AccessToken: accessToken}, w)
	data.ToJSON(&GenericResponse{
		Status:  true,
		Message: "Successfully generated new access token",
		Data:    &TokenResponse{AccessToken: accessToken},
	}, w)
}

// Greet request greet request
func (ah *AuthHandler) Greet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(UserIDKey{}).(string)
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("hello, " + userID))
	data.ToJSON(&GenericResponse{
		Status:  true,
		Message: "hello," + userID,
	}, w)
}

// Profile ...
func (ah *AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(UserIDKey{}).(string)
	user, err := ah.repo.GetUserByID(context.Background(), userID)
	if err != nil {
		ah.logger.Error("unable to get user", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericResponse{Status: false, Message: "Unable to get user. Please try again later"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	data.ToJSON(&GenericResponse{
		Status:  true,
		Message: "success",
		Data: struct {
			ID         string `json:"id"`
			Email      string `json:"email"`
			Username   string `json:"username"`
			IsVerified bool   `json:"isverified"`
		}{
			ID:         user.ID,
			Email:      user.Email,
			Username:   user.Username,
			IsVerified: user.IsVerified,
		},
	}, w)
}

// GeneratePassResetCode generate a new secret code to reset password.
func (ah *AuthHandler) GeneratePassResetCode(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(UserIDKey{}).(string)

	user, err := ah.repo.GetUserByID(context.Background(), userID)
	if err != nil {
		ah.logger.Error("unable to get user to generate secret code for password reset", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericResponse{Status: false, Message: "Unable to send password reset code. Please try again later"}, w)
		return
	}
	//TODO
	// Send verification mail
	from := "vikisquarez@gmail.com"
	to := []string{user.Email}
	subject := "Password Reset"
	mailType := mail.PassReset
	mailData := &mail.Data{
		Username: user.Username,
		Code:     utils.GenerateRandomString(8),
	}

	mailReq := ah.mailService.NewMail(from, to, subject, mailType, mailData)
	err = ah.mailService.SendMail(mailReq)
	if err != nil {
		ah.logger.Error("unable to send mail", "error", err)
		//TODO
		//w.WriteHeader(http.StatusInternalServerError)
		//data.ToJSON(&GenericResponse{Status: false, Message: "Unable to send password reset code. Please try again later"}, w)
		//return
	}

	// store the password reset code to db
	verificationData := &data.VerificationData{
		Email:     user.Email,
		Code:      mailData.Code,
		Type:      data.PassReset,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(ah.configs.PassResetCodeExpiration)),
	}

	err = ah.repo.StoreVerificationData(context.Background(), verificationData)
	if err != nil {
		ah.logger.Error("unable to store password reset verification data", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericResponse{Status: false, Message: "Unable to send password reset code. Please try again later"}, w)
		return
	}

	ah.logger.Debug("successfully mailed password reset code")
	w.WriteHeader(http.StatusOK)
	data.ToJSON(&GenericResponse{Status: true, Message: "Please check your mail for password reset code"}, w)
}

// GoogleOAuth ...
func (ah *AuthHandler) GoogleOAuth(w http.ResponseWriter, r *http.Request) {
	URL, err := url.Parse(ah.oauthConfGl.Endpoint.AuthURL)
	if err != nil {
		ah.logger.Error("Parse: " + err.Error())
	}
	ah.logger.Info(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", ah.oauthConfGl.ClientID)
	parameters.Add("scope", strings.Join(ah.oauthConfGl.Scopes, " "))
	parameters.Add("redirect_uri", ah.oauthConfGl.RedirectURL)
	parameters.Add("response_type", "code")
	//TODO state
	//parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	ah.logger.Info(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GoogleLogin ...
func (ah *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {

	ah.logger.Info("Callback-gl..")

	//TODO state

	code := r.FormValue("code")
	ah.logger.Info(code)

	if code == "" {
		ah.logger.Warn("Code not found..")
		w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User has denied Permission.."))
		}
		return
	}

	token, err := ah.oauthConfGl.Exchange(context.Background(), code)
	if err != nil {
		ah.logger.Error("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
		data.ToJSON(&GenericResponse{Status: false, Message: err.Error()}, w)
		return
	}
	ah.logger.Info("TOKEN>> AccessToken>> " + token.AccessToken)
	ah.logger.Info("TOKEN>> Expiration Time>> " + token.Expiry.String())
	ah.logger.Info("TOKEN>> RefreshToken>> " + token.RefreshToken)

	resp, err := http.Get(endpointProfile + "?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		ah.logger.Error("Get: " + err.Error() + "\n")
		data.ToJSON(&GenericResponse{Status: false, Message: err.Error()}, w)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ah.logger.Error("ReadAll: " + err.Error() + "\n")
		data.ToJSON(&GenericResponse{Status: false, Message: err.Error()}, w)
		return
	}

	var u googleUser
	if err = json.Unmarshal(response, &u); err != nil {
		ah.logger.Error("Unmarshal: " + err.Error() + "\n")
		data.ToJSON(&GenericResponse{Status: false, Message: err.Error()}, w)
		return
	}

	ah.logger.Info("parseResponseBody: " + string(response) + "\n")

	reqUser := data.User{
		Email:      u.Email,
		Username:   u.Name,
		IsVerified: u.VerifiedEmail,
	}

	user, err := ah.repo.GetUserByEmail(context.Background(), reqUser.Email)
	if err != nil {
		ah.logger.Error("error fetching the user", "error", err)
		errMsg := err.Error()
		if strings.Contains(errMsg, PgNoRowsMsg) {
			reqUser.TokenHash = utils.GenerateRandomString(15)
			err = ah.repo.Create(context.Background(), &reqUser)
			if err != nil {
				ah.logger.Error("unable to insert user to database", "error", err)
				errMsg := err.Error()
				if strings.Contains(errMsg, PgDuplicateKeyMsg) {
					w.WriteHeader(http.StatusBadRequest)
					// data.ToJSON(&GenericError{Error: ErrUserAlreadyExists}, w)
					data.ToJSON(&GenericResponse{Status: false, Message: ErrUserAlreadyExists}, w)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
					// data.ToJSON(&GenericError{Error: errMsg}, w)
					data.ToJSON(&GenericResponse{Status: false, Message: UserCreationFailed}, w)
				}
				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			// data.ToJSON(&GenericError{Error: err.Error()}, w)
			data.ToJSON(&GenericResponse{Status: false, Message: "Unable to retrieve user from database.Please try again later"}, w)
			return
		}

	} else {
		reqUser.ID = user.ID
		reqUser.TokenHash = user.TokenHash
	}

	accessToken, err := ah.jwtManager.GenerateAccessToken(&reqUser)
	if err != nil {
		ah.logger.Error("unable to generate access token", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(&GenericError{Error: err.Error()}, w)
		data.ToJSON(&GenericResponse{Status: false, Message: "Unable to login the user. Please try again later"}, w)
		return
	}
	refreshToken, err := ah.jwtManager.GenerateRefreshToken(&reqUser)
	if err != nil {
		ah.logger.Error("unable to generate refresh token", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(&GenericError{Error: err.Error()}, w)
		data.ToJSON(&GenericResponse{Status: false, Message: "Unable to login the user. Please try again later"}, w)
		return
	}

	ah.logger.Debug("successfully generated token", "accesstoken", accessToken, "refreshtoken", refreshToken)
	w.WriteHeader(http.StatusOK)
	// data.ToJSON(&AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Username}, w)
	data.ToJSON(&GenericResponse{
		Status:  true,
		Message: "Successfully logged in",
		Data:    &AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: reqUser.Username},
	}, w)

}
