# just a test


just a simple test to concurrently getUrls content in golang 1.18.

the idea is simple, based on the number of the url and parallelNo parameters, create a bufferchannel with size of 
ParallelNo which let the process done concurrently with `parallelNo`.

it means if you have 10 urls and `parallelNo` is  5,  the whole 10 urls will be processed in two batch of 5 concurrent go routines.

### installation
After clone the repository run the following command to get required packages.  
`go mod vendor`

### build
before you run the program you need to build it with
`go build -o myhttp`

### execution
after you build the app. you can run it  by 
`./myhttp -parallel N URL1 URL2 URL3 URL4`

**example1:**
with explicit parallel   

`./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com`

**example2:**
with default parallel ( it's 10 )

`./myhttp adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com`

### test
The main test case for th

`go test ./... -v`


note: due to structure I choose form the project to considered as a generic project I would like to have integration test as well to cover the functionality of the website.

#### Improvement
I believe the test coverage is not enough and it could be better + mocking several urls with different required. 



### configuration
there are some options which you can configure

#### logger
at the moment logging level is `info` but you can override it by creating  an `.env` and set `LOG_LEVEL` variable there.  ( you can `.env.dist` and change the values)   

##### note: only zap logger + testify has been used for better logging and simpler assertation.