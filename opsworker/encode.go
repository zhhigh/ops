package opsworker

import(
	"github.com/mikespook/gearman-go/worker"
	"ops/algorithm"
	"fmt"
)

const(
	EncryptString = "zhhigh@163.com"
)


func Encrypt(job *worker.Job)([]byte, error){
	alg := algorithm.NewAlg()
	alg.SetKeyt(EncryptString)
	data := string(job.Data)
	fmt.Println(data)
	r    := alg.Eencode(data)
	fmt.Println(r)
	return []byte(r),nil
}

func Decode(job *worker.Job)([]byte, error){
	alg := algorithm.NewAlg()
	alg.SetKeyt(EncryptString)
	data := string(job.Data)
	r    := alg.Ddecode(data)
	return []byte(r),nil
}
