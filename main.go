package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Version 現在のバージョン
const (
	Version = "latest"
)

func main() {
	fmt.Println("=== Application Starting!!")
	http.HandleFunc("/hc", healthHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/fibo", fiboHandler)
	http.HandleFunc("/down", downHandler)
	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- healthHandler")
	fmt.Fprint(w, "OK")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- infoHandler")
	// インスタンスIDの取得
	sess := session.Must(session.NewSession())
	svc := ec2metadata.New(sess)
	doc, _ := svc.GetInstanceIdentityDocument()
	instanceID := doc.InstanceID
	// レスポンス
	fmt.Fprint(w, "version:"+Version+"\ninstanceID: "+instanceID)
}

func fiboHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- fiboHandler")
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		fmt.Fprint(w, "ERROR")
		return
	}
	fmt.Fprint(w, strconv.Itoa(n)+"番目のフィボナッチ数は、"+strconv.Itoa(fibo(n)))
}

func downHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--- downHandler")
	log.Fatal("DOWN!!!")
}

func fibo(n int) int {
	if n < 2 {
		return 1
	}
	return fibo(n-2) + fibo(n-1)
}
