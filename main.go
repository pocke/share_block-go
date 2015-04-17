package main

import (
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("CK"))
	anaconda.SetConsumerSecret(os.Getenv("CS"))
	fromAPI := anaconda.NewTwitterApi(os.Getenv("FROM_AT"), os.Getenv("FROM_AS"))
	toAPI := anaconda.NewTwitterApi(os.Getenv("TO_AT"), os.Getenv("TO_AS"))

	fromIDs, err := getAllBlockIDs(fromAPI)
	if err != nil {
		panic(err)
	}
	toIDs, err := getAllBlockIDs(toAPI)
	if err != nil {
		panic(err)
	}
	fromIDs.Diff(toIDs)

	for _, id := range fromIDs {
		v := url.Values{}
		v.Add("user_id", strconv.FormatInt(id, 10))
		_, err := toAPI.Block(v)
		log.Printf("Blocked %s.", id)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Minute)
	}

}

func getAllBlockIDs(api *anaconda.TwitterApi) (ids IDs, err error) {
	ids = make([]int64, 0)
	c := anaconda.Cursor{ // dummy
		Next_cursor:     -1,
		Next_cursor_str: "-1",
	}

	for c.Next_cursor != 0 {
		v := url.Values{}
		v.Add("cursor", c.Next_cursor_str)
		c, err = api.GetBlocksIds(v)
		if err != nil {
			return ids, err
		}
		ids = append(ids, c.Ids...)
	}

	sort.Sort(ids)
	return ids, nil
}
