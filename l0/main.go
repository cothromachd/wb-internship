package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/stan.go"
)

type modelStruct struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

var cache = make(map[string]modelStruct, 10)

func main() {
	r := gin.Default()

	dbPool, err := pgxpool.New(context.Background(), "postgres://khalidmagnificent:190204@localhost:5432/l0_task")

	rows, err := dbPool.Query(context.Background(), `SELECT * FROM model;`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var ouid string
	var bjson []byte
	modelObject := modelStruct{}
	for rows.Next() {
		err = rows.Scan(&ouid, &bjson)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(bjson, &modelObject)
		if err != nil {
			log.Fatal(err)
		}

		cache[modelObject.OrderUid] = modelObject
	}

	sc, err := stan.Connect("test-cluster", "client-123")
	if err != nil {
		log.Fatal(err)
	}

	go sc.Publish("foo", []byte(`{ "order_uid": "b563feb7b2b84b6test", "track_number": "WBILMTESTTRACK", "entry": "WBIL", "delivery": { "name": "Test Testov", "phone": "+9720000000", "zip": "2639809", "city": "Kiryat Mozkin", "address": "Ploshad Mira 15", "region": "Kraiot", "email": "test@gmail.com" }, "payment": { "transaction": "b563feb7b2b84b6test", "request_id": "", "currency": "USD", "provider": "wbpay", "amount": 1817, "payment_dt": 1637907727, "bank": "alpha", "delivery_cost": 1500, "goods_total": 317, "custom_fee": 0 }, "items": [ { "chrt_id": 9934930, "track_number": "WBILMTESTTRACK", "price": 453, "rid": "ab4219087a764ae0btest", "name": "Mascaras", "sale": 30, "size": "0", "total_price": 317, "nm_id": 2389212, "brand": "Vivienne Sabo", "status": 202 } ], "locale": "en", "internal_signature": "", "customer_id": "test", "delivery_service": "meest", "shardkey": "9", "sm_id": 99, "date_created": "2021-11-26T06:22:19Z", "oof_shard": "1" }`))

	sub, err := sc.Subscribe("foo", func(msg *stan.Msg) {
		model := modelStruct{}
		err = json.Unmarshal(msg.Data, &model)

		log.Println("======")
		log.Println(model.OrderUid)
		log.Println("======")

		if err != nil {
			log.Println(err)
			return
		}

		cache[model.OrderUid] = model

		_, err = dbPool.Exec(context.Background(), `INSERT INTO model(ouid, data)
					 VALUES ($1, $2);`, model.OrderUid, msg.Data)
		if err != nil {
			log.Println(err)
			return
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Fatal(err)
	}

	sub.Unsubscribe()

	sc.Close()

	r.GET("/foo/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		result, ok := cache[id]
		if ok {
			resultJSON, err := json.MarshalIndent(result, "", "    ")
			log.Println(string(resultJSON))
			if err != nil {
				log.Fatal(err)
			}

			r.LoadHTMLFiles("index.html")

			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"data": string(resultJSON),
			})
			return
		}

		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "data not found"})
	})

	r.Run(":8080")
}
