package fd

import (
	"bufio"
	"encoding/json"
	"fmt"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/samber/lo"
	"os"
)

type StateJson struct {
	StdinArn string
	StdinUrl string
	Pid      *int
}

func ReadConf(f *os.File) (arn *StateJson) {
	L.Debug.Println("reading conf from fd", f.Fd())
	scan := bufio.NewScanner(f)
	scan.Scan()

	d := scan.Bytes()
	state := &StateJson{}
	err := json.Unmarshal(d, state)
	if err != nil {
		L.Debug.Println("read state:", err)
		return nil
	}

	L.Debug.Println("got conf:", string(d))
	return state
}

func WriteConf(f *os.File, conf StateJson) {
	d := lo.Must(json.Marshal(conf))
	L.Debug.Println("writing conf:", string(d))
	lo.Must(fmt.Fprint(f, string(d)+"\n"))
}
