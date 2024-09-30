package models

import "Golang_/araha/constants"

type Subscription struct {
	ID               string                     `json:"id"`
	SubscriptionType constants.SubscriptionType `json:"SubscriptionType"`
	Description      string                     `json:"description"`
}
