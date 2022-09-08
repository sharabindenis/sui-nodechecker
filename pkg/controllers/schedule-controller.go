package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/sharabindenis/sui-nodechecker/pkg/models"
	"github.com/sharabindenis/sui-nodechecker/pkg/utils"
	"net/http"
	_ "path/filepath"
	"strings"
)

// TODO basicauth

var Schedule *models.Schedule
var s = gocron.NewScheduler()

type TotalTransaction struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  int    `json:"result"`
	ID      int    `json:"id"`
}

//func Index(w http.ResponseWriter, r *http.Request) {
//	p := path.Dir("/index.html")
//	fmt.Println(p)
//	// set header
//	w.Header().Set("Content-type", "text/html")
//	http.ServeFile(w, r, p)
//}

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	sch := &models.Schedule{}
	utils.ParseBody(r, sch)
	fmt.Println("Node Ip", sch.Ip)
	fmt.Println("Period, seconds", sch.Period)
	s.Every(sch.Period).Seconds().Do(TotalTT, sch.Ip)
	<-s.Start()
	w.WriteHeader(http.StatusOK)
}

func TotalTT(ip string) {
	//fmt.Println("start")
	fullnodett := &TotalTransaction{}
	//формирование запроса
	var url string
	url = ip
	method := "POST"
	payload := strings.NewReader(`{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "sui_getTotalTransactionNumber"
}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	//body, err := ioutil.ReadAll(res.Body) // response body is []byte
	err = json.NewDecoder(res.Body).Decode(fullnodett)

	//check spread of transactions
	// TODO mssg model
	if err == nil {
		url := fmt.Sprint("https://api.telegram.org/bot5436979593:AAEEoh_5A6-OeUEdXEdnTBeRcMbdCMk1hXk/sendMessage?chat_id=-725815999&text=Total%20%20", fullnodett.Result, " ", ip) //paste your credentials
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
	}
}

//func CheckSpreadTransaction(ipstr string) {
//	fullnodett := &TotalTransaction{} //создается
//	//формирование запроса
//	url := "https://fullnode.devnet.sui.io/"
//	method := "POST"
//	payload := strings.NewReader(`{
//   "jsonrpc": "2.0",
//   "id": 1,
//   "method": "sui_getTotalTransactionNumber"
//}`)
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, payload)
//	if err != nil {
//		fmt.Println(err)
//	}
//	req.Header.Add("Content-Type", "application/json")
//
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer res.Body.Close()
//
//	//body, err := ioutil.ReadAll(res.Body) // response body is []byte
//	err = json.NewDecoder(res.Body).Decode(fullnodett)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(fullnodett.Result)
//
//	mynodett := &TotalTransaction{} //создается
//	//формирование запроса
//	myurl := ipstr
//	myclient := &http.Client{}
//	myreq, err := http.NewRequest(method, myurl, payload)
//	if err != nil {
//		fmt.Println(err)
//	}
//	myreq.Header.Add("Content-Type", "application/json")
//
//	myres, err := myclient.Do(myreq)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer myres.Body.Close()
//	//body, err := ioutil.ReadAll(res.Body) // response body is []byte
//	err = json.NewDecoder(myres.Body).Decode(mynodett)
//	fmt.Println(mynodett.Result)
//
//	alert := mynodett.Result - fullnodett.Result
//	fmt.Println(alert)
//
//	//check spread of transactions
//	if alert == 0 {
//		fmt.Println("check")
//		defer res.Body.Close()
//	} else if alert < -10 {
//		url := fmt.Sprint("https://api.telegram.org/bot5436979593:AAEEoh_5A6-OeUEdXEdnTBeRcMbdCMk1hXk/sendMessage?chat_id=-725815999&text=Расхождение%20транзакций%20%20", alert) //paste your credentials
//		method := "GET"
//
//		client := &http.Client{}
//		req, err := http.NewRequest(method, url, nil)
//
//		if err != nil {
//			fmt.Println(err)
//		}
//		res, err := client.Do(req)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer res.Body.Close()
//	}
//}

func Stop(w http.ResponseWriter, r *http.Request) {
	//s.Remove(CheckSpreadTransaction)
	jobs := s.Jobs()
	s.Remove(TotalTT)
	fmt.Println("Schedule stop", jobs)
}

func ShowJobs(w http.ResponseWriter, r *http.Request) {
	jobs := s.Jobs()
	fmt.Println("Now works ", jobs)
}
