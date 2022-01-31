package internal

import (
	"encoding/json"
	"fmt"
	"github.com/joshinjohnson/authentication-engine/pkg/api"
	"github.com/joshinjohnson/authentication-engine/pkg/models"
	"github.com/joshinjohnson/authentication-manager/pkg/errors"
	outputModel "github.com/joshinjohnson/authentication-manager/pkg/models"
	"github.com/joshinjohnson/authentication-manager/tokenengine"
	"net/http"
	"time"
)

const (
	emailField   = "Email"
	passwordHash = "Password-Hash"
	fnameField   = "First-Name"
	lnameField   = "Last-Name"
	dobField     = "Date-Of-Birth"
	timeLayout   = "Jan 1, 2000 at 3:04pm (MST)"
	tokenField   = "Token"
	privateKey   = "j0sh19"
)

var (
	claimsMap = map[string]string{
		"user": "joshin.johnson",
		"exp":  fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}
)

type Manager struct {
	AuthenticationEngine api.AuthenticationEngine
	TokenGeneratorEngine tokenengine.TokenGeneratorEngine
}

func (m *Manager) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email, emailOK := r.Header[emailField]
	pass, passOK := r.Header[passwordHash]

	if !emailOK || !passOK || len(email) == 0 || len(pass) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errors.ErrBadRequest))
		return
	}

	cred := models.UserCredential{
		Email:    email[0],
		Password: pass[0],
	}
	_, err := m.AuthenticationEngine.Authenticate(cred)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf(errors.ErrLogin+": %v", err.Error())))
		return
	}
	token := m.getToken(cred)
	if token == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errors.ErrInternalServer))
		return
	}

	w.WriteHeader(http.StatusOK)
	msg, _ := json.Marshal(outputModel.LoginSuccessResponse{
		Message: "user logged in",
		Token:   token,
	})
	w.Write(msg)
}

func (m *Manager) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email, emailOK := r.Header[emailField]
	pass, passOK := r.Header[passwordHash]
	fname, fnameOK := r.Header[fnameField]
	lname, lnameOK := r.Header[lnameField]

	if !emailOK || !passOK || !fnameOK || !lnameOK ||
		len(email) == 0 || len(pass) == 0 || len(fname) == 0 || len(lname) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errors.ErrBadRequest))
		return
	}

	cred := models.UserCredential{
		Email:    email[0],
		Password: pass[0],
	}
	//dobT, _ := time.Parse(timeLayout, dob[0])
	detail := models.UserDetails{
		FirstName: fname[0],
		LastName:  lname[0],
		//DateOfBirth: dobT,
	}

	if err := m.AuthenticationEngine.Register(cred, detail); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(errors.ErrRegistration+": %v", err.Error())))
		return
	}
	token := m.getToken(cred)
	if token == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errors.ErrInternalServer))
		return
	}

	w.WriteHeader(http.StatusOK)
	msg, _ := json.Marshal(outputModel.LoginSuccessResponse{
		Message: "user registered",
		Token:   token,
	})
	w.Write(msg)
}

func (m *Manager) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	msg, _ := json.Marshal(outputModel.HomeSuccessResponse{
		Message: "Home",
	})
	w.Write(msg)
}

func (m Manager) VerifyTokenMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header[tokenField]
		if ok {
			if isAuthenticToken, err := m.TokenGeneratorEngine.VerifyToken(token[0], privateKey); err != nil || !isAuthenticToken {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(errors.ErrInvalidToken))
				return
			}
		}
		handler.ServeHTTP(w, r)
	})
}

func (m Manager) getToken(cred models.UserCredential) string {
	return m.TokenGeneratorEngine.TokenGeneratorFunc()(cred, claimsMap, privateKey)
}
