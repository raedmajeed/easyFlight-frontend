package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/frontend/config"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Data struct {
		UserID     int    `json:"user_id"`
		TotalFare  int    `json:"total_fare"`
		BookingRef string `json:"booking_reference"`
		Email      string `json:"email"`
		OrderID    string `json:"order_id"`
	} `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func OnlinePaymentRender(ctx *gin.Context, cf *config.Parameters) {
	bookingRef := ctx.Query("bookingRef")
	searchToken := ctx.Query("searchToken")

	params := url.Values{}
	params.Set("bookingRef", bookingRef)
	params.Set("searchToken", searchToken)

	baseURL := fmt.Sprintf("http://%v/api/v1/user/booking/confirm/online/payment", cf.BACKENDPORT)
	reqUrl := baseURL + "?" + params.Encode()
	response, err := http.Get(reqUrl)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "failed.html", gin.H{})
	}

	if response.Status >= "300" {
		log.Println("error, status 300+")
		return
	}

	byteVal, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "failed.html", gin.H{})
	}
	defer response.Body.Close()
	var res Response
	if err := json.Unmarshal(byteVal, &res); err != nil {
		ctx.HTML(http.StatusBadRequest, "failed.html", gin.H{})
	}
	ctx.HTML(http.StatusOK, "app.html", gin.H{
		"bookingRef":  bookingRef,
		"orderId":     res.Data.OrderID,
		"TotalFare":   res.Data.TotalFare,
		"searchToken": searchToken,
		"email":       res.Data.Email,
	})
}

func OnlinePaymentSuccess(ctx *gin.Context, cf *config.Parameters) {
	bookingReference := ctx.DefaultQuery("booking_reference", "")
	paymentId := ctx.DefaultQuery("payment_id", "")
	token := ctx.DefaultQuery("search_token", "")

	params := url.Values{}
	params.Set("booking_reference", bookingReference)
	params.Set("payment_id", paymentId)
	params.Set("search_token", token)
	baseURL := fmt.Sprintf("http://%v/api/v1/user/booking/confirm/online/payment/success", cf.BACKENDPORT)

	reqUrl := baseURL + "?" + params.Encode()
	response, err := http.Get(reqUrl)
	byteVal, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "failed.html", gin.H{})
	}
	defer response.Body.Close()
	var res SuccessResponse
	if err := json.Unmarshal(byteVal, &res); err != nil {
		ctx.HTML(http.StatusBadRequest, "failed.html", gin.H{})
	}
	fmt.Println(res, "response")
	ctx.JSON(http.StatusOK, gin.H{"status": true})
}
func OnlinePaymentSuccessRender(ctx *gin.Context, cf *config.Parameters) {
	ctx.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": ctx.Query("booking_reference"),
	})
}
