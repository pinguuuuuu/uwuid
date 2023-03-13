package uwuid

import (
	"math/rand"
	"time"
)

var uwuidList = [6]string{"UwU", "OwO", "TwT", ":3", ">w<", "^w^"}

func generateUUID() string {
	rand.Seed(time.Now().UnixNano())
    uuid := ""
    for i := 0; i < 6; i++ {
        uuid += uwuidList[rand.Intn(len(uwuidList))]
		if i == 0 {
			uuid += uwuidList[rand.Intn(len(uwuidList))]
		}
        if i < 5 {
            uuid += "-"
        } else {
			uuid += uwuidList[rand.Intn(len(uwuidList))]
			uuid += uwuidList[rand.Intn(len(uwuidList))]
		}
    }
    return uuid
}