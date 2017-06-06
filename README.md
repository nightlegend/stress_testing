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

> Running by docker

<pre>
	1. Docker >= 1.9(if you have not install docker, please reference https://docs.docker.com/v1.12/).
	2. cd ${work directory}/stress_testing
	3. docker build -t stress_test:latest
	4. docker-compose up -d   (if you have not install docker-compose, please reference to https://docs.docker.com/v1.12/compose/install/)
</pre>

> Add a mount point in docker-compose file

<pre>
	For convenient to running and find result,better add a volumes.
	volumes:
  	 - /home/csapp/stress_test/tmp:/go/src/stress_test/tmp
     - /home/csapp/stress_test/conf/conf.yaml:/go/src/stress_test/conf/conf.yml
</pre>

<hr>

<h3>
	Result filed explanation:
</h3>

<pre>
latencies:      Latencies holds computed request latency metrics.
earliest:       First is the earliest timestamp in a Result set.
latest:         Latest is the latest timestamp in a Result set.
end:            End is the latest timestamp in a Result set plus its latency.
duration:       Duration is the duration of the attack.
wait:           Wait is the extra time waiting for responses from targets.
requests:       Requests is the total number of requests executed.
rate:           Rate is the rate of requests per second.
success:        Success is the percentage of non-error responses.
status_codes:   StatusCodes is a histogram of the responses' status codes.
errors:         Errors is a set of unique errors returned by the targets during the attack.

</pre>

<small>Keep update to here for latest changed. Thanks for you love it.</small>

