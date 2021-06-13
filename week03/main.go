package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// error group ctx
	group, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "index.")
		log.Print("index.")
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//  添加服务启动函数
	group.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Print("http server error: %v", err.Error())
		}
		return err
	})

	// 添加信号检测函数
	group.Go(func()  error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		select {
		case <-sigs:
			//如果有信号关闭server
			server.Shutdown(ctx)
		case <- ctx.Done():
			//如果server因错误退出
			log.Print("server exit")
		}
		return nil
	})

	if err := group.Wait(); err != nil && err != http.ErrServerClosed{
		log.Print("ErrorGroup wait err %v", err)
	}
	log.Print("Not error server exit")
}
