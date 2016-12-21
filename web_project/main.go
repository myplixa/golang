package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	Calldate    string `json:"calldate"`
	Dst         string `json:"dst"`
	Src         string `json:"src"`
	Lastapp     string `json:"lastapp"`
	Channel     string `json:"channel"`
	Dcontext    string `json:"dcontext"`
	Dstchannel  string `json:"dstchannel"`
	Disposition string `json:"disposition"`
}

type Itogo struct {
	Msg string     `json:"action"`
	ANS []Response `json:"data"`
}

func main() {
	var arrAns Itogo
	db, err := sql.Open("mysql", "imedia:Ngfj9HslqmS23KSd@tcp(tellban.com:3306)/dbasterisk?charset=utf8")
	checkErr(err)
	defer db.Close()
	// query
	rows, err := db.Query("SELECT calldate, dst, src, lastapp, channel, dcontext, dstchannel, disposition FROM cdr WHERE src LIKE '8123091684' AND dst LIKE '%' AND DATE_FORMAT(calldate, '%d-%m-%Y') = '21-03-2016'")
	checkErr(err)

	for rows.Next() {
		r := Response{}
		err = rows.Scan(&r.Calldate, &r.Dst, &r.Src, &r.Lastapp, &r.Channel, &r.Dcontext, &r.Dstchannel, &r.Disposition)
		checkErr(err)
		arrAns.ANS = append(arrAns.ANS, r)
	}
	arrAns.Msg = "DST"
	res, _ := json.Marshal(arrAns)
	fmt.Println(string(res))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
