package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"strconv"
	"strings"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/hc", healthHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/fibo", fiboHandler)
	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	// インスタンスIDの取得
	sess := session.Must(session.NewSession())
	svc := ec2metadata.New(sess)
	doc, _ := svc.GetInstanceIdentityDocument()
	instanceId := doc.InstanceID
	// コンテナIDの取得
	containerId, _ := os.Hostname()
	// タスクの取得
	resp, err := http.Get(os.Getenv("ECS_CONTAINER_METADATA_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var metadata interface{}
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		log.Fatal(err)
	}
	taskArn := metadata.(map[string]interface{})["Labels"].(map[string]interface{})["com.amazonaws.ecs.task-arn"].(string)
	task := strings.Split(taskArn, "/")[1]
	// レスポンス
	fmt.Fprint(w, "instanceId: "+instanceId+"\ntask: "+task+"\ncontainerId: "+containerId)
}

func fiboHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		fmt.Fprint(w, "ERROR")
		return
	}
	fmt.Fprint(w, strconv.Itoa(n)+"番目のフィボナッチ数は、"+strconv.Itoa(fibo(n)))
}

func fibo(n int) int {
	if n < 2 {
		return 1
	} else {
		return fibo(n-2) + fibo(n-1)
	}
}
