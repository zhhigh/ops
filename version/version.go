package version

import(
	"fmt"
	"github.com/mikespook/gearman-go/worker"

)

const (
	Ver = "V0.12"
)
/*----------------------------------------------------------------------------
[V0.1 2013-11-20]具备ssh2通道的能力，通过该通道，可以登陆到任何ssh2 server执行命令
                 1、SSH2()  多指令无数据返回，返回空
                 2、增加获取worker版本的、打印版本的能力
                 perl:my $result_v = $client->do_task("Ver","get version");
[V0.11 2013-11-20]具备ssh2通道的能力，通过该通道，可以登陆到任何ssh2 server执行命令,单指令可返回数据
                 1、SSH2()  多指令无数据返回，返回空
                 2、SSH21() 单指令，返回指令执行后的数据
[V0.12 2012-11-23]增加对密码的加减密能力，并放入gearman worker
                  1、opsworker.Encrypt
                  2、opsworker.Decode
                  my $result_sys = $client->do_task("Encrypt","redis");
                  my $result_pwd = $client->do_task("Decode",$$result_sys);
[V0.13  beat]增加对worker函数的配置加载
             全局变量、常量的加载
             mongo库的err返回值重新编写
             日志的打印及写入日志文件
-----------------------------------------------------------------------------*/

/*---------------------------BUG---------------------------------------------

---------------------------------------------------------------------------*/
func Version(job *worker.Job)([]byte, error){

	fmt.Printf("%v\n",string(job.Data))
	data := []byte("golang worker version:"+Ver)
	return data, nil
}


