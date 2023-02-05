package server

import (
	"k8s.io/klog/v2"
	"net/http"
)

func StartServer() {
	RouterSetup()
	server := &http.Server{
		Addr:    ":8080",
		Handler: R,
	}

	klog.Info("starting server...")

	err := server.ListenAndServe()
	if err != nil {
		klog.Errorf(err.Error())
	}
	//err := http.ListenAndServe(":8080",R)
}
