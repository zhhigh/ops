/*
Created by zhhigh
Date : 2013-11-11
*/
package ssh2

import (
	//"os"
	//"log"
	//"time"
	"strings"
	//"github.com/mikespook/golib/signal"
	"github.com/mikespook/gearman-go/worker"
	"github.com/robfig/config"
	"fmt"
)


type SshInfo struct{
	Host  string
	User  string
	Pwd   string
	Port  string
}



