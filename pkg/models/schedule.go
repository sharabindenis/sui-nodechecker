package models

type Schedule struct {
	Ip     string `json:"ip,omitempty"`
	Period uint64 `json:"period,omitempty"`
}

func (b *Schedule) CreateSchedule() *Schedule {

	return b
}

//var db *gorm.DB
//func Index(w http.ResponseWriter, r *http.Request) {
//	t, err := template.New("index").ParseFiles("pkg/views/index.html")
//
//	if err != nil {
//		fmt.Fprintf(w, err.Error())
//	}
//
//	t.ExecuteTemplate(w, "index", nil)
//}

//Создается настройка расписания и запуск

//func Start(w http.ResponseWriter, r *http.Request) {
//	Create := &Schedule{}
//	utils.ParseBody(r, CreateBook)
//	b := CreateBook.
//	res, _ := json.Marshal(b)
//	w.WriteHeader(http.StatusOK)
//	w.Write(res)
//
//	//body := r.Body.Read
//	//var ipstr string
//	//ipstr = r.Body.Read([)
//	//persec := r.FormValue("period")
//	//intpersec, _ := strconv.Atoi(persec)
//
//	s.Every(3).Seconds().Do(CheckSpreadTransaction, ipstr)
//	//s.Every(10).Seconds().Do(models.TotalTT)
//	<-s.Start()
//	fmt.Println("Schedule start")
//
//}
