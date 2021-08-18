package handlers

import (
	"awesomebook/helpers"
	"awesomebook/repositories"
	"awesomebook/requests"
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"net/http"
)

const SecretKey = "sk_test_51JQ98iJhuo3sClIoHe2qqSBNMMoiaEgFlSY2jIzWjOXh7w0z0FSal51bPmLfQNeM83IK8MT2KCw9enkaFrGpeSUT00j0VvQFBp"

func PaymentHandler(writer http.ResponseWriter, request *http.Request) {
	// Get userID from jwt
	// if not login then required login

	// Fix userID = 1
	userID := 1

	userEntity, _ := repositories.FindUserByUserID(userID)

	// Get shopping cartEntity from UserID
	cartEntity, err := repositories.GetCartFromUserID(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.DBConn().Model(&cartEntity).Association("BookEntities").Find(&cartEntity.BookEntities)
	fmt.Println(cartEntity)

	chargeRequest := &requests.ChargeRequest{}
	err = json.NewDecoder(request.Body).Decode(chargeRequest)
	if err != nil {
		sugar.Error(err)
	}

	stripe.Key = SecretKey

	_, err = charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(chargeRequest.Amount),
		Currency:     stripe.String(string(stripe.CurrencyVND)),
		ReceiptEmail: stripe.String(userEntity.Email),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		Description:  stripe.String("Test...."),
	})

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_, _ = writer.Write([]byte("Success!"))
}
