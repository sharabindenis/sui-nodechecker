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

type SuiTotalTransaction struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  int    `json:"result"`
	ID      int    `json:"id"`
}

type SuiTask struct {
	//gorm.Model
	Ip  string      `json:"ip"`
	Job *gocron.Job `json:"job"`
}

var SuiTaskHolder = map[string]SuiTask{}
var s = gocron.NewScheduler()

func (b *SuiTask) CreateSuiTask() *SuiTask {

	return b
}

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	sch := &models.Schedule{}
	utils.ParseBody(r, sch)
	fmt.Println("Node Ip", sch.Ip)
	fmt.Println("Period, seconds", sch.Period)
	//fmt.Println()
	//holder := SuiTaskHolder
	task, ok := SuiTaskHolder[sch.Ip]
	fmt.Println(task.Job)
	if ok {
		fmt.Println(task)
	} else {
		//err := s.Every(sch.Period).Seconds().Do(TotalTT, sch.Ip)
		job := s.Every(sch.Period).Seconds()
		err := job.Do(TotalTT, sch.Ip)
		//job.Tag("EtoNoviiJob")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("1")
		go s.Start()
		var dopizdi = SuiTask{sch.Ip, job}
		SuiTaskHolder[sch.Ip] = dopizdi
		fmt.Println(dopizdi)
	}
	//fmt.Println(SuiTaskHolder)
	w.WriteHeader(http.StatusOK)
}

func ShowTasks(w http.ResponseWriter, r *http.Request) {
	items := make([]SuiTask, len(SuiTaskHolder))
	var i int
	for _, v := range SuiTaskHolder {
		items[i] = v
		i++
	}
	fmt.Println(items)
	res, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	//return items
}

func StopSchedule(w http.ResponseWriter, r *http.Request) {
	stp := &SuiTask{}
	utils.ParseBody(r, stp)
	task := SuiTaskHolder[stp.Ip]
	s.RemoveByRef(task.Job)
	delete(SuiTaskHolder, stp.Ip)
	w.WriteHeader(http.StatusOK)
}

func TotalTT(ip string) {
	//fmt.Println("start")
	fullnodett := &SuiTotalTransaction{}
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

	fmt.Println(fullnodett.Result)
	//check spread of transactions
	// TODO mssg model
	//if err == nil {
	//	url := fmt.Sprint("https://api.telegram.org/bot5436979593:AAEEoh_5A6-OeUEdXEdnTBeRcMbdCMk1hXk/sendMessage?chat_id=-725815999&text=Total%20%20", fullnodett.Result, " ", ip) //paste your credentials
	//	method := "GET"
	//
	//	client := &http.Client{}
	//	req, err := http.NewRequest(method, url, nil)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	res, err := client.Do(req)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	defer res.Body.Close()
	//}
}

//func CheckSpreadTransaction(ipstr string) {
//	fullnodett := &SuiTotalTransaction{} //создается
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
//	mynodett := &SuiTotalTransaction{} //создается
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
