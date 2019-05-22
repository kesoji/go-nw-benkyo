package main

import (
	"flag"
	"log"
	"net"
	"os"
)

func main() {
	// プログラム引数の処理
	ip := flag.String("i", "127.0.0.1", "接続先のIP(デフォルト127.0.0.1)")
	port := flag.String("p", "8888", "接続先のPort(デフォルト8888)")
	name := flag.String("n", "名無し", "送信する名前(デフォルト名無し)")
	flag.Parse()

	// サーバへ接続
	conn, err := net.Dial("tcp", *ip+":"+*port)
	if err != nil {
		log.Println("サーバに接続できませんでした")
		os.Exit(1)
	}
	log.Println("サーバへ接続しました")

	// サーバへデータの送信
	_, err = conn.Write([]byte(*name))
	if err != nil {
		log.Println("サーバにデータを送信できませんでした")
		os.Exit(1)
	}
	log.Println("サーバへデータを送信しました。送信したデータ:", *name)

	// サーバからデータの受け取り、表示
	buf := make([]byte, 1000)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("サーバからデータを受信できませんでした")
		os.Exit(1)
	}
	log.Println("サーバからデータを受信しました。受信したバイト数:", n, "受信したデータ:", string(buf[:n]))

	log.Println("プログラムを終了します")
}
