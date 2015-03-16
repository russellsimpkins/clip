package clip

import (
	"hash/crc32"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func GetApplicationFeatures(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Printf("We have this many %v", vars)
}

/*
 * Taken right from the docs. This is how you populate a map of maps
 */
func add(m map[string]map[string]int, env, feature string, val int) {
    mm, ok := m[env]
    if !ok {
        mm = make(map[string]int)
        m[env] = mm
    }
    mm[feature] = val
}


func main() {
	key := "d81697b6322a81a7fb19e0ef1141f534da0634244e76b5590332a1a186c7c4a9"
	short := crc32.ChecksumIEEE([]byte(key))
	fmt.Printf("This is my value %d\n", short)
	r := mux.NewRouter()
	r.HandleFunc("/svc/clip/{application}", GetApplicationFeatures)
	var sample map[string]map[string]int
	
	sample = make(map[string]map[string]int)
	add(sample, "dev", "feature", 1)
	fmt.Printf("The feature is on? %d\n", sample["dev"]["feature"])
}
