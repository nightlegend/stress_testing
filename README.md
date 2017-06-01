# Stress test: 

stress test is build by golang and vegeta. Hope find a good solution privode to you about stress testing.<br>


<a href="https://golang.org/">golang</a>: Who use, who know.

<a href="https://github.com/tsenart/vegeta">vegeta</a>: It`s a good testing tools.

# Stress test support methods

> Now only support GET and POST methods.

## GET

### conf.yml

<pre>
	host: 127.0.0.1
	port: 8012
	times: 4
	rate: 4
	timeout: 150
	resultPath: D:/SelfStudy/GoProject/data/
	testResulName: testing
	requestType: get
	requestPath: /
	requestData: ''
</pre>

## POST

### conf.yml

<pre>
	host: 127.0.0.1
	port: 8012
	times: 4
	rate: 4
	timeout: 150
	resultPath: D:/SelfStudy/GoProject/data/
	testResulName: testing
	requestType: post
	requestPath: /form_post
	requestData: '{"id":"PERM-25-a811-4951-8fef-3057091d8992"}'
</pre>


<h2>How to run?</h2>

<div>

1. prepare
<pre>
    install golang >= 1.6
</pre>


2. startup
a. update you configure:
<pre>
	cd $workdir/conf
	update conf.yml
</pre>

b. build and start
<pre>
 	cd workdir
 	go get
 	go build main.go
 	go run main.go
</pre>

</div>

<hr>

<h3>
	some result filed explanation:
</h3>
<pre>
latencies:          Latencies holds computed request latency metrics.
earliest:           First is the earliest timestamp in a Result set.
latest:   			Latest is the latest timestamp in a Result set.
end:       			End is the latest timestamp in a Result set plus its latency.
duration:           Duration is the duration of the attack.
wait:      			Wait is the extra time waiting for responses from targets.
requests:           Requests is the total number of requests executed.
rate:      			Rate is the rate of requests per second.
success:            Success is the percentage of non-error responses.
status_codes:    	StatusCodes is a histogram of the responses' status codes.
errors:   			Errors is a set of unique errors returned by the targets during the attack.

</pre>

<small>Keep update to here for latest changed. Thanks for you love it.</small>

