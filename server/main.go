package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

func init() {
	// ランダムデータのシード。これが無いとランダムにならない
	rand.Seed(time.Now().UnixNano())
}

func main() {
	port := flag.String("p", "8888", "待ち受けのPort(デフォルト8888)")
	wait := flag.Duration("w", 1*time.Second, "クライアントへ応答を返すまでの待ち時間")
	flag.Parse()

	// Listen(待ち受け)
	l, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		panic(err)
	}
	log.Println("Listen成功")

	for {
		log.Println("クライアントからの接続を待っています")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		log.Println(conn.RemoteAddr().String(), "が接続しました")

		// クライアントからのデータを受け取る
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		name := string(buf[:n])
		log.Println(conn.RemoteAddr().String(), "から受信:", name)

		log.Println(wait, "間の間待機します")
		time.Sleep(*wait)

		// クライアントにデータを送信
		_, err = conn.Write([]byte(getPhrase(name)))
		if err != nil {
			panic(err)
		}
		log.Println(conn.RemoteAddr().String(), "にデータを送信しました")

		// コネクションの終了
		err = conn.Close()
		if err != nil {
			panic(err)
		}
		log.Println(conn.RemoteAddr().String(), "との接続を終了しました")
	}

	// ここは到達しなくなる
	log.Println("プログラムを終了します")
}

func getPhrase(name string) string {
	// 末尾の改行コードを削除する
	name = strings.TrimRight(name, "\r\n")
	// ランダムで返す語句をいくつか定義する
	// %sの部分には、引数で指定したnameが入る
	phrases := []string{
		"サバになれ!%s!!\n",
		"君なら出来る!!%s!!\n",
		"君は太陽だ!!!%s!!\n",
	}
	// 0 <= n < len(phrases) の整数取得する
	//   lenは配列の長さ(length)を返す
	//   rand.Intnは0から引数n未満のランダム(randはrandomのrand)な整数(IntはIntegerのInt)を返す
	n := rand.Intn(len(phrases))

	// fmt.Sprintfはフォーマット(phrases)に引数の変数を埋め込んだ文字列(string)を作ってくれる
	return fmt.Sprintf(phrases[n], name)
}
