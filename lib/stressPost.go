package lib

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"os"
	"stress_test/conf"
	"time"
)

func Post(config conf.Configure) {

	rate := uint64(config.Rate) // per second
	duration := config.Times * time.Second
	urlStr := "http://" + config.Host + ":" + config.Port + config.RequestPath
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    urlStr,
		Body:   []byte(config.RequestData),
	})

	attacker := vegeta.NewAttacker()
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
