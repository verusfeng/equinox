package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

func get_list() {
	func() {
		var allList []string
		for pageID := 1; ; pageID++ {
			errCnt := 0
			respStr, err := func(pageID int) (string, error) {
				apiURL := fmt.Sprintf("http://api.live.bilibili.com/room/v3/area/getRoomList?platform=web&parent_area_id=1&cate_id=0&area_id=199&sort_type=online&page=%d&page_size=50&tag_version=1", pageID)

				resp, err := http.Get(apiURL)
				if err != nil {
					return "", errors.New("http Get error.....e")
				}
				defer resp.Body.Close()
				// http get success
				rbyte, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return "", errors.New("Response read error.....e")
				}
				// read success
				return string(rbyte), nil
			}(pageID)
			if err == nil {
				tmpList := func(rstr string) []string {
					msgData, messageData, dataData :=
						func(s string) (string, string, string) {
							return gjson.Get(s, "msg").String(),
								gjson.Get(s, "message").String(),
								gjson.Get(s, "data").String()
						}(rstr)

					if msgData == "success" && messageData == "success" {
						listArray := gjson.Get(dataData, "list").Array()
						if len(listArray) == 0 {
							return []string{}
						}
						return func(l []gjson.Result) []string {
							var retList []string
							for _, value := range l {
								retList = append(retList, value.String())
							}
							return retList
						}(listArray)

					}
					//fmt.Println("Fail ...")
					return []string{}

				}(respStr)
				if len(tmpList) == 0 {
					log.Println("End of page!")
					break
				}
				allList = append(allList, tmpList...)
			} else {
				errCnt++
				if errCnt > 10 {
					log.Fatal("too many http.Get error.....")
				}
			}
			time.Sleep(time.Millisecond * 10)
			log.Println(fmt.Sprintf("page:%d", pageID), err)
		}
		// fmt.Println(allList)
		fmt.Println(len(allList))
		if len(allList) == 0 {
			log.Fatal("room list is 0 .... ")
		}
		// TODO : handle the room list ,
		// allList = myTest()
		// for _, value := range allList {
		for index := len(allList) - 1; index >= 0; index-- {
			value := allList[index]
			roomID, userName, onlineCNT := func(ss string) (string, string, string) {
				return gjson.Get(ss, "roomid").String(), gjson.Get(ss, "uname").String(), gjson.Get(ss, "online").String()
			}(value)

			urlAddress := "https://live.bilibili.com/" + roomID

			title := gjson.Get(value, "title").String()
			userName = userName + "    ==" + title
			if len(userName) > 70 {
				userName = userName[:70]
			}
			fmt.Printf("%-35s -%7s -%-15s \n", urlAddress, onlineCNT, userName)

			// if i, _ := strconv.Atoi(onlineCNT); i >= 2000 {
			//  fmt.Printf("%10s  %-15s  %-s \n", onlineCNT, userName[:15], urlAddress)
			// }
		}
	}()
}
