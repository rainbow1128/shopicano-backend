package services

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/matcornic/hermes/v2"
	"github.com/shopicano/shopicano-backend/app"
	"github.com/shopicano/shopicano-backend/config"
	"github.com/shopicano/shopicano-backend/models"
)

func SendOrderDetailsEmail(name, email string, order *models.OrderDetailsView) error {
	var entryItems [][]hermes.Entry

	for _, v := range order.Items {
		var rows []hermes.Entry
		rows = append(rows, hermes.Entry{
			Key:   "Item",
			Value: v.Name,
		})
		rows = append(rows, hermes.Entry{
			Key:   "Quantity",
			Value: fmt.Sprintf("%d", v.Quantity),
		})
		rows = append(rows, hermes.Entry{
			Key:   "Price",
			Value: fmt.Sprintf("%d", v.Price),
		})
		rows = append(rows, hermes.Entry{
			Key:   "Sub Total",
			Value: fmt.Sprintf("%d", v.SubTotal),
		})

		entryItems = append(entryItems, rows)
	}

	var entryCalculations []hermes.Entry

	entryCalculations = append(entryCalculations, hermes.Entry{
		Key:   "Sub Total",
		Value: fmt.Sprintf("%d", order.SubTotal),
	})

	if order.ShippingCharge != 0 {
		entryCalculations = append(entryCalculations, hermes.Entry{
			Key:   "Shipping Charge",
			Value: fmt.Sprintf("%d", order.ShippingCharge),
		})
	}

	if order.PaymentProcessingFee != 0 {
		entryCalculations = append(entryCalculations, hermes.Entry{
			Key:   "Payment Processing Fee",
			Value: fmt.Sprintf("%d", order.PaymentProcessingFee),
		})
	}

	entryCalculations = append(entryCalculations, hermes.Entry{
		Key:   "Grand Total",
		Value: fmt.Sprintf("%d", order.GrandTotal),
	})

	emailTemplate := hermes.Email{
		Body: hermes.Body{
			Greeting:  "Hello",
			Signature: "With Love",
			Name:      name,
			Intros: []string{
				"Your order has been placed successfully.",
			},
			Table: hermes.Table{
				Data: entryItems,
			},
			Dictionary: entryCalculations,
			Actions: []hermes.Action{
				{
					Instructions: "",
					Button: hermes.Button{
						Link:      fmt.Sprintf("%s/v1/orders/%s", config.EmailService().VerificationUrl, order.ID),
						Color:     "white",
						Text:      "Order Details",
						TextColor: "black",
					},
				},
			},
		},
	}

	body, err := app.Hermes().GenerateHTML(emailTemplate)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s", config.EmailService().FromEmailAddress))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Order has been placed")
	m.SetBody("text/html", body)

	if err := EmailDialer().DialAndSend(m); err != nil {
		return err
	}
	return nil
}
