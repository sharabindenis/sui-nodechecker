# Simple sui fullnode checker

### Telegram bot<br/>
Open botfather.<br/>
Create bot.<br/>
Grab the token.<br/>
Place it in pkg/controllers/telegram-controller.go<br/>
Change interval in CreateScheduleByBot<br/>


### Docker<br/>
```
docker
docker build -t sch/dev:latest .
docker run -p 8080:8080 -d sch/dev:latest
```
### Test data</br>
```
/ip 5.23.48.96
/ip 45.8.147.68
/ip 159.223.119.225
```
