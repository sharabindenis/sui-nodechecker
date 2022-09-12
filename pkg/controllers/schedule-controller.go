package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/sharabindenis/sui-nodechecker/pkg/models"
	"io"
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
var S = gocron.NewScheduler()
var Sch = &models.Schedule{}

func (b *SuiTask) CreateSuiTask() *SuiTask {

	return b
}

func CreateScheduleByBot(ip string, chtid int64) {
	fmt.Println("Node Ip", ip)
	//fmt.Println("Period, seconds", Sch.Period)
	//fmt.Println()
	//holder := SuiTaskHolder

	//err := S.Every(Sch.Period).Seconds().Do(TotalTT, Sch.Ip)
	job := S.Every(3).Seconds()
	err := job.Do(TotalTT, ip, chtid)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("1")
	go S.Start()
	fmt.Println(S.Jobs())
	var forhold = SuiTask{ip, job}
	SuiTaskHolder[Sch.Ip] = forhold
	fmt.Println(forhold)

}

func TotalTT(ip string, chtid int64) {
	fullnodett := &SuiTotalTransaction{}
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	//body, err := ioutil.ReadAll(res.Body) // response body is []byte
	err = json.NewDecoder(res.Body).Decode(fullnodett)

	fmt.Println(fullnodett.Result)
	TelegramBotAlert(chtid, fullnodett.Result, ip)
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

//OLD REST

//func CreateSchedule(w http.ResponseWriter, r *http.Request) {
//	utils.ParseBody(r, Sch)
//	fmt.Println("Node Ip", Sch.Ip)
//	fmt.Println("Period, seconds", Sch.Period)
//	//fmt.Println()
//	//holder := SuiTaskHolder
//	task, ok := SuiTaskHolder[Sch.Ip]
//	fmt.Println(task.Job)
//	if ok {
//		fmt.Println(task)
//	} else {
//		//err := S.Every(Sch.Period).Seconds().Do(TotalTT, Sch.Ip)
//		job := S.Every(Sch.Period).Seconds()
//		err := job.Do(TotalTT, Sch.Ip)
//		//job.Tag("EtoNoviiJob")
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("1")
//		go S.Start()
//		var dopizdi = SuiTask{Sch.Ip, job}
//		SuiTaskHolder[Sch.Ip] = dopizdi
//		fmt.Println(dopizdi)
//	}
//	//fmt.Println(SuiTaskHolder)
//	w.WriteHeader(http.StatusOK)
//}

//func ShowTasks(w http.ResponseWriter, r *http.Request) {
//	items := make([]SuiTask, len(SuiTaskHolder))
//	var i int
//	for _, v := range SuiTaskHolder {
//		items[i] = v
//		i++
//	}
//	fmt.Println(items)
//	res, err := json.Marshal(items)
//	if err != nil {
//		fmt.Println(err)
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//	//return items
//}
//
//func StopSchedule(w http.ResponseWriter, r *http.Request) {
//	stp := &SuiTask{}
//	utils.ParseBody(r, stp)
//	task := SuiTaskHolder[stp.Ip]
//	S.RemoveByRef(task.Job)
//	delete(SuiTaskHolder, stp.Ip)
//	w.WriteHeader(http.StatusOK)
//}
