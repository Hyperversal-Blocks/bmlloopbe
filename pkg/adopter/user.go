package adopter

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/bmlloopbe/pkg/store"
)

type profile struct {
	db      store.Store
	logger  *logrus.Logger
	address common.Address
}

func (p *profile) Create(adopter *Adopter) error {
	adopterJSON, err := json.Marshal(adopter)
	if err != nil {
		return fmt.Errorf("unable to marshal adopter object: %w", err)
	}

	err = p.db.Put(p.address.String(), adopterJSON)
	if err != nil {
		return fmt.Errorf("unable to create adopter object: %w", err)
	}

	return nil
}

func (p *profile) Get() (*Adopter, error) {
	// Retrieve the Adopter data from BadgerDB
	obj, err := p.db.Get(p.address.String())
	if err != nil {
		return nil, fmt.Errorf("unable to fetch adopter object: %w", err)
	}

	// Unmarshal the JSON data back into an Adopter struct
	var user Adopter
	err = json.Unmarshal(obj, &user)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal adopter object: %w", err)
	}
	return &user, nil
}

func New(db store.Store, logger *logrus.Logger, address common.Address) Service {
	return &profile{
		db:      db,
		logger:  logger,
		address: address,
	}
}

type Adopter struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Gender  string  `json:"gender,omitempty"`
	DOB     string  `json:"dob"`
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
	Type    Type    `json:"type"`
	Wallet  string  `json:"wallet"`
	PubKey  string  `json:"pubKey"`
}

type Address struct {
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}

type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Type struct {
	MBTI      string `json:"mbti,omitempty"`
	Horoscope string `json:"horoscope,omitempty"`
}

type Service interface {
	Create(user *Adopter) error
	Get() (*Adopter, error)
}
