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
	ids := DiffInt64(fromIDs, toIDs)

	for _, id := range ids {
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

func getAllBlockIDs(api *anaconda.TwitterApi) (ids []int64, err error) {
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

	sortedIDs := IDs(ids)
	sort.Sort(sortedIDs)
	ids = []int64(sortedIDs)
	return ids, nil
}

func DiffInt64(a, b []int64) []int64 {
	if len(b) == 0 {
		res := make([]int64, len(a))
		copy(res, a)
		return res
	}

	res := make([]int64, 0, len(a))
	i := 0
	for _, av := range a {
		for j := i; len(b) != j; j++ {
			if b[j] == av {
				break
			} else if b[j] > av {
				res = append(res, av)
				j--
				break
			}
		}
	}

	return res
}
