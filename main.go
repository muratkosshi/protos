package main

//
//type loggingRoundTripper struct {
//	logger io.Writer
//	next   http.RoundTripper
//}
//
//func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
//	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
//	return l.next.RoundTrip(r)
//}

func main() {
	//client := &http.Client{
	//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//		fmt.Println(req.Response.Status)
	//		fmt.Println("REDIRECT")
	//		return nil
	//	},
	//	Transport: &loggingRoundTripper{
	//		logger: os.Stdout,
	//		next:   http.DefaultTransport,
	//	},
	//	Timeout: time.Second * 10,
	//}
	//resp, err := client.Get("https://api.coincap.io/v2/assets")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer resp.Body.Close()
	//fmt.Println("Respone status", resp.StatusCode)
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(body))
}

//func print(message []string) error {
//	if len(message) == 0 {
//
//	}
