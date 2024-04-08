package transport

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/Corray333/dating/internal/domains/user/types"
	"github.com/Corray333/dating/pkg/server/auth"
	"github.com/go-chi/chi/v5"
)

type Storage interface {
	InsertUser(user types.User, agent string) (int, string, error)
	LoginUser(user types.User, agent string) (int, string, error)
	CheckAndUpdateRefresh(id int, refresh string) (string, error)
	SelectUser(id string) (types.User, error)
	UpdateUser(user types.User) error
}

// SignUp is an HTTP handler function that signs up a new user.
// It expects a JSON body with user details (email, password, etc.).
// If the user details are valid and the user is successfully created, it returns a JSON response with an access token, a refresh token, and user details.
// If there's an error reading the request body or unmarshalling the JSON, it returns a 400 Bad Request status.
// If there's an error inserting the user into the database or creating the access token, it returns a 500 Internal Server Error status.
//
//	@Summary		User signup
//	@Description	Sign up a new user and return an access token, a refresh token, and user details.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		types.User	true	"User details"
//	@Success		200		{object}	LogInResponse
//	@Failure		400		{string}	string	"Bad Request. Failed to read or unmarshal request body."
//	@Failure		500		{string}	string	"Internal Server Error. Failed to insert user, create token, or encode response."
//	@Router			/users/signup [post]
func SignUp(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := types.User{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			slog.Error("Failed to read request body: " + err.Error())
			return
		}
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
			slog.Error("Failed to unmarshal request body: " + err.Error())
			return
		}
		id, refresh, err := store.InsertUser(user, r.UserAgent())
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			slog.Error("Failed to insert user: " + err.Error())
			return
		}
		user.ID = id

		token, err := auth.CreateToken(user.ID, auth.AccessTokenLifeTime)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			slog.Error("Failed to create token: " + err.Error())
			return
		}
		user.Password = ""
		if err := json.NewEncoder(w).Encode(LogInResponse{Authorization: token,
			Refresh: refresh,
			User:    user,
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			slog.Error("Failed to send response: " + err.Error())
			return
		}
	}
}

type LogInResponse struct {
	Authorization string     `json:"authorization"`
	Refresh       string     `json:"refresh"`
	User          types.User `json:"user"`
}

// LogIn is an HTTP handler function that logs in a user.
// It expects a JSON body with user credentials (email and password).
// If the credentials are valid, it returns a JSON response with an access token, a refresh token, and user details.
// If the credentials are invalid, it returns a 403 Forbidden status.
// If there's an error reading the request body or unmarshalling the JSON, it returns a 400 Bad Request status.
// If there's an error creating the access token, it returns a 500 Internal Server Error status.
//
//	@Summary		User login
//	@Description	Log in a user and return an access token, a refresh token, and user details.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		types.User		true	"User credentials"
//	@Success		200		{object}	LogInResponse	"Successful login"
//	@Failure		400		{string}	string			"Bad Request. Failed to read or unmarshal request body."
//	@Failure		403		{string}	string			"Forbidden. Wrong password or email."
//	@Failure		500		{string}	string			"Internal Server Error. Failed to create token or encode response."
//	@Router			/users/login [post]
func LogIn(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := types.User{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			slog.Error("Failed to read request body: " + err.Error())
			return
		}
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
			slog.Error("Failed to unmarshal request body: " + err.Error())
			return
		}
		id, refresh, err := store.LoginUser(user, r.UserAgent())
		if err != nil {
			http.Error(w, "Wrong password or email", http.StatusForbidden)
			slog.Error("Failed to login user: " + err.Error())
			return
		}
		user.ID = id

		token, err := auth.CreateToken(user.ID, auth.AccessTokenLifeTime)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			slog.Error("Failed to create token: " + err.Error())
			return
		}
		user.Password = ""
		if err := json.NewEncoder(w).Encode(LogInResponse{Authorization: token,
			Refresh: refresh,
			User:    user,
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			slog.Error("Failed to send response: " + err.Error())
			return
		}
	}
}

type RefreshAccessTokenResponse struct {
	Authorization string `json:"authorization"`
	Refresh       string `json:"refresh"`
}

// RefreshAccessToken is an HTTP handler function that refreshes an access token.
// It expects a "Refresh" header with the refresh token.
// If the refresh token is valid, it returns a JSON response with a new access token and a new refresh token.
// If there's an error refreshing the access token, it returns a 500 Internal Server Error status.
// If there's an error encoding the response, it also returns a 500 Internal Server Error status.
//
//	@Summary		Refresh access token
//	@Description	Refresh an access token and return a new access token and a new refresh token.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			Refresh	header		string						true	"Refresh token"
//	@Success		200		{object}	RefreshAccessTokenResponse	"Successful refresh"
//	@Failure		500		{string}	string						"Internal Server Error. Failed to refresh token or encode response."
//	@Router			/users/refresh [post]
func RefreshAccessToken(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh := r.Header.Get("Refresh")
		access, refresh, err := auth.RefreshAccessToken(store, refresh)
		if err != nil {
			http.Error(w, "Failed to refresh token", http.StatusInternalServerError)
			slog.Error("Failed to refresh token: " + err.Error())
			return
		}
		if err := json.NewEncoder(w).Encode(RefreshAccessTokenResponse{
			Authorization: access,
			Refresh:       refresh,
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			slog.Error("Failed to encode response: " + err.Error())
			return
		}
	}
}

type GetUserResponse struct {
	User types.User `json:"user"`
}

// GetUser is an HTTP handler function that retrieves a user's details.
// It expects a URL parameter "id" with the user's ID.
// If the user ID is valid and the user is found, it returns a JSON response with the user's details.
// If there's an error getting the user or encoding the response, it returns a 500 Internal Server Error status.
//
//	@Summary		Get user details
//	@Description	Retrieve a user's details.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{object}	GetUserResponse	"Successful retrieval"
//	@Failure		500	{string}	string			"Internal Server Error. Failed to get user or encode response."
//	@Router			/users/{id} [get]
func GetUser(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "id")
		user, err := store.SelectUser(userId)
		if err != nil {
			http.Error(w, "Failed to get user", http.StatusInternalServerError)
			slog.Error("Failed to get user: " + err.Error())
			return
		}
		if err := json.NewEncoder(w).Encode(GetUserResponse{
			User: user,
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			slog.Error("Failed to encode response: " + err.Error())
			return
		}
	}
}

// UpdateUser is an HTTP handler function that updates a user's details.
// It expects an "Authorization" header with the user's credentials, a form field "username" with the new username, and an optional form file "avatar" with the new avatar image.
// If the credentials are valid and the user is successfully updated, it returns a 200 OK status.
// If there's an error extracting the credentials, reading the avatar file, or updating the user, it returns a 400 Bad Request or 500 Internal Server Error status.
//
//	@Summary		Update user details
//	@Description	Update a user's username and optionally their avatar.
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			Authorization	header		string	true	"User credentials"
//	@Param			username		formData	string	true	"New username"
//	@Param			avatar			formData	file	false	"New avatar image"
//	@Success		200				{string}	string	"User updated successfully."
//	@Failure		400				{string}	string	"Bad Request. Failed to extract credentials or read file."
//	@Failure		500				{string}	string	"Internal Server Error. Failed to get user, create file, write file, or update user."
//	@Router			/users/{id} [post]
func UpdateUser(store Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		creds, err := auth.ExtractCredentials(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Failed to extract credentials", http.StatusBadRequest)
			slog.Error("Failed to extract credentials: " + err.Error())
			return
		}
		user, err := store.SelectUser(strconv.Itoa(int(creds.ID)))
		if err != nil {
			http.Error(w, "Failed to get user", http.StatusInternalServerError)
			slog.Error("Failed to get user: " + err.Error())
			return
		}
		file, _, err := r.FormFile("avatar")
		if err != nil && err.Error() != "http: no such file" {
			http.Error(w, "Failed to read file", http.StatusBadRequest)
			slog.Error("Failed to read file: " + err.Error())
			return
		}
		if file != nil {
			newFile, err := os.Create("../files/images/avatars/avatar" + strconv.Itoa(int(user.ID)) + ".png")
			if err != nil {
				http.Error(w, "Failed to create file", http.StatusInternalServerError)
				slog.Error("Failed to create file: " + err.Error())
				return
			}
			data, err := io.ReadAll(file)
			if err != nil {
				http.Error(w, "Failed to read file", http.StatusInternalServerError)
				slog.Error("Failed to read file: " + err.Error())
				return
			}
			if _, err := newFile.Write(data); err != nil {
				http.Error(w, "Failed to write file", http.StatusInternalServerError)
				slog.Error("Failed to write file: " + err.Error())
				return
			}
			user.Avatar = "http://localhost:3001/images/avatars/avatar" + strconv.Itoa(int(user.ID)) + ".png"
		}
		user.Username = r.FormValue("username")
		if err := store.UpdateUser(user); err != nil {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			slog.Error("Failed to update user: " + err.Error())
			return
		}
	}
}
