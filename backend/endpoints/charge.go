package endpoints

import (
	"fmt"
	"time"

	log "github.com/cihub/seelog"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"

	"github.com/1backend/1backend/backend/config"
	"github.com/1backend/1backend/backend/domain"
)

const pricePer100k = 900 // in cents

// taken from https://stripe.com/docs/charges

func (e Endpoints) Charge(accessToken, paymentToken string, amount uint64) error {
	tk, err := domain.NewAccessTokenDao(e.db).GetByToken(accessToken)
	if err != nil {
		return err
	}
	serviceToken := domain.Token{}
	err = e.db.Where("user_id = ? AND name = 'default'", tk.UserId).First(&serviceToken).Error
	if err != nil {
		return fmt.Errorf("Can't find default token: %v", err.Error())
	}

	stripe.Key = config.C.StripeKey

	// token is created using Checkout or Elements!
	// token is the payment token ID submitted by the form

	// Charge the user's card:
	params := &stripe.ChargeParams{
		Amount:   amount,
		Currency: "usd",
		Desc:     fmt.Sprintf("%v00k quota added", amount/pricePer100k), // @todo hardcoded price
	}
	params.SetSource(paymentToken)

	log.Infof("About to charge user with id %v", tk.UserId)
	_, err = charge.New(params)
	if err != nil {
		log.Errorf("Charging user withd id %v was unsuccessful: %v", tk.UserId, err)
		return err
	}
	// intentionally not using transactions here
	ch := &domain.Charge{
		UserId:    tk.UserId,
		Id:        domain.Sid.MustGenerate(),
		Amount:    int(amount),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = e.db.Create(ch).Error
	if err != nil {
		log.Errorf("Failed to save charge for user widh id %v: %v", tk.UserId, err)
	}
	err = e.db.Table("users").Where("user_id = ?", tk.UserId).Update("premium", true).Error
	if err != nil {
		log.Errorf("Could set user to premium: %v", err)
	}
	err = e.state.IncrementBy(serviceToken.Token, 100000*int64(amount/pricePer100k))
	if err != nil {
		log.Errorf("Failed to increment quita for user %v, charge id %v", tk.UserId, ch.Id)
	}
	return err
}
