package lib

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"net/http"
	"os"
	"stress_test/conf"
	"time"
)

func Post(config conf.Configure) {
	rate := uint64(config.Rate) // per second
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	duration := config.Times * time.Second
	timeout := config.Timeout * time.Second
	urlStr := "http://" + config.Host + ":" + config.Port + config.RequestPath
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    urlStr,
		Body:   []byte(config.RequestData),
		Header: headers,
	})

	attacker := vegeta.NewAttacker(vegeta.Timeout(timeout))

	var metrics vegeta.Metrics
	var results vegeta.Results
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
		results.Add(res)
	}
	metrics.Close()
	results.Close()

	report := vegeta.NewJSONReporter(&metrics)
	htmlreport := vegeta.NewPlotReporter("title", &results)

	jsonResultFile, err := os.Create(config.ResultPath + config.TestResulName + ".json")
	htmlResultFile, err := os.Create(config.ResultPath + config.TestResulName + ".html")
	if err != nil {
		fmt.Println(err)
	}
	report.Report(jsonResultFile)
	htmlreport.Report(htmlResultFile)

}
