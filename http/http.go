package http

import (
	"github.com/signmem/etcd3pool/etcd"
	"github.com/signmem/etcd3pool/g"
	"log"
	"net/http"
	"strings"
	"encoding/json"
)

type Dto struct {
	Msg  string              `json:"msg"`
	Data     interface{}     `json:"data"`
}


func init() {
	showMetric()
	showNow()
	healthCheck()
}

func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Dto{Msg: "success", Data: data})
}

func RenderMsgJson(w http.ResponseWriter, msg string) {
	RenderJson(w, map[string]string{"msg": msg})
}

func AutoRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		RenderMsgJson(w, err.Error())
		return
	}

	RenderDataJson(w, data)
}


func Start() {
	if ! g.Config().Http.Enabled {
		return
	}

	addr := g.Config().Http.Listen

	if addr == "" || strings.Split(addr, ":")[1] == "" {
		log.Printf("[ERROR] add error.")
		return
	}

	log.Printf("[INFO] http start with %s\n", addr)
	s := &http.Server {
		Addr:  addr,
		MaxHeaderBytes: 1 << 30,
	}
	log.Printf("[INFO] listening: %s\n", addr)
	log.Fatalf("[INFO] %s", s.ListenAndServe())
}

func showMetric() {

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metric := etcd.READ
		RenderDataJson(w,  map[string]interface{} {
			"etcdwrite" : metric,
		})
	})
}

func showNow() {

	http.HandleFunc("/now", func(w http.ResponseWriter, r *http.Request) {
		metric := etcd.WRITE
		RenderDataJson(w,  map[string]interface{} {
			"etcdwriting" : metric,
		})
	})
}

func healthCheck() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w,map[string]interface{} {
			"version":  g.Version,
		})
	})
}