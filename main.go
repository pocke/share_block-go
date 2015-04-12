package main

import (
	"net/url"
	"os"
	"sort"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("CK"))
	anaconda.SetConsumerSecret(os.Getenv("CS"))
	fromAPI := anaconda.NewTwitterApi(os.Getenv("FROM_AT"), os.Getenv("FROM_AS"))
	toAPI := anaconda.NewTwitterApi(os.Getenv("TO_AT"), os.Getenv("TO_AS"))

	getAllBlockIDs(fromAPI)
	getAllBlockIDs(toAPI)
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

	sortedIDs := int64s(ids)
	sort.Sort(sortedIDs)
	ids = []int64(sortedIDs)
	return ids, nil
}
