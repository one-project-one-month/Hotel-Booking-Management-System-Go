package auth

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/events"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/models"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/mq"
	"github.com/one-project-one-month/Hotel-Booking-Management-System-Go/pkg/response"
)

type Service struct {
	queue *mq.MQ
}

func (s *Service) Signin(user *SignInUserDto) *response.ServiceResponse {
	reply := s.queue.Publish(&mq.Message{
		AppID: "AuthService",
		Topic: events.USERFINDBYEMAIL,
		Data:  &events.FindByEmailDto{Email: user.Email},
	})

	var foundUser *models.User
	select {
	case resp := <-reply:
		data := resp.(*response.ServiceResponse)
		if data.Error != nil {
			return &response.ServiceResponse{
				AppID:   "AuthService",
				Error:   response.ErrNotFound,
				Message: "Error while finding user",
			}
		}
		foundUser = data.Data.(*models.User)
	case <-time.Tick(2 * time.Second):
		return &response.ServiceResponse{
			AppID:   "AuthService",
			Message: "Timeout!",
		}
	}

	if !checkPassword(foundUser.Password, user.Password) {
		return &response.ServiceResponse{
			AppID:   "AuthService",
			Error:   response.ErrUnauthorized,
			Message: "Invalid password",
		}
	}

	token, _ := newJWTToken(foundUser.Name, false)

	return &response.ServiceResponse{
		AppID: "AuthService",
		Data: &SignInResponseDto{
			Token: token,
		},
		Message: "User signed in successfully",
	}
}

func (s *Service) Signup(user *SignUpUserDto) *response.ServiceResponse {
	emailReply := s.queue.Publish(&mq.Message{
		AppID: "AuthService",
		Topic: events.USERFINDBYEMAIL,
		Data:  &events.FindByEmailDto{Email: user.Email},
	})

	select {
	case resp := <-emailReply:
		data := resp.(*response.ServiceResponse)
		if data.Data != nil {
			return &response.ServiceResponse{
				AppID:   "AuthService",
				Error:   response.ErrConflict,
				Message: fmt.Sprintf("User with email %s already exists", user.Email),
			}
		}
	case <-time.Tick(1 * time.Second):
		return &response.ServiceResponse{
			AppID:   "AuthService",
			Error:   response.ErrInternalServer,
			Message: "Timeout!",
		}
	}

	phoneNumberReply := s.queue.Publish(&mq.Message{
		AppID: "AuthService",
		Topic: events.USERFINDBYPHONENUMBER,
		Data:  &events.FindByPhoneNumberDto{PhoneNumber: user.PhoneNumber},
	})

	select {
	case resp := <-phoneNumberReply:
		data := resp.(*response.ServiceResponse)
		if data.Data != nil {
			return &response.ServiceResponse{
				AppID:   "AuthService",
				Error:   response.ErrConflict,
				Message: fmt.Sprintf("User with phone number %s already exists", user.PhoneNumber),
			}
		}
	case <-time.Tick(1 * time.Second):
		return &response.ServiceResponse{
			AppID:   "AuthService",
			Error:   response.ErrInternalServer,
			Message: "Timeout!",
		}
	}

	password, err := newPassword(user.Password)
	if err != nil {
		return &response.ServiceResponse{
			AppID:   "AuthService",
			Error:   err,
			Message: "Error while hashing password",
		}
	}

	user.Password = password

	userJson, _ := json.Marshal(user)
	reply := s.queue.Publish(&mq.Message{
		AppID: "AuthService",
		Topic: events.USERCREATED,
		Data:  userJson,
	})

	select {
	case resp := <-reply:
		data := resp.(*response.ServiceResponse)
		if data.Error != nil {
			return &response.ServiceResponse{
				AppID:   "AuthService",
				Error:   data.Error,
				Message: "Error while creating user",
			}
		}
	case <-time.Tick(2 * time.Second):
		return &response.ServiceResponse{
			AppID:   "AuthService",
			Error:   response.ErrInternalServer,
			Message: "Timeout!",
		}
	}

	token, _ := newJWTToken(user.Name, false)
	return &response.ServiceResponse{
		AppID: "AuthService",
		Data: &SignUpResponseDto{
			Token: token,
		},
		Message: "User created successfully",
	}
}

func newService(queue *mq.MQ) *Service {
	s := &Service{
		queue: queue,
	}
	return s
}
