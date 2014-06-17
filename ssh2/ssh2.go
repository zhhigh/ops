/*
ssh2 package 
Created by zhhigh
Date:2013-11-10
ssh2 source code :http://godoc.org/code.google.com/p/go.crypto/ssh#ClientConn.NewSession
*/

package ssh2
import(
    "bytes"
    "code.google.com/p/go.crypto/ssh"
    "log"
    "fmt"
)

type SSH2 struct{
	client   *ssh.ClientConn
	session  *ssh.Session
    server   string
    username string
    password clientPassword
}

type TerminalModes map[uint8]uint32
type clientPassword string

func (p clientPassword) Password(user string) (string, error) {
    return string(p), nil
}
const (
    VINTR         = 1
    VQUIT         = 2
    VERASE        = 3
    VKILL         = 4
    VEOF          = 5
    VEOL          = 6
    VEOL2         = 7
    VSTART        = 8
    VSTOP         = 9
    VSUSP         = 10
    VDSUSP        = 11
    VREPRINT      = 12
    VWERASE       = 13
    VLNEXT        = 14
    VFLUSH        = 15
    VSWTCH        = 16
    VSTATUS       = 17
    VDISCARD      = 18
    IGNPAR        = 30
    PARMRK        = 31
    INPCK         = 32
    ISTRIP        = 33
    INLCR         = 34
    IGNCR         = 35
    ICRNL         = 36
    IUCLC         = 37
    IXON          = 38
    IXANY         = 39
    IXOFF         = 40
    IMAXBEL       = 41
    ISIG          = 50
    ICANON        = 51
    XCASE         = 52
    ECHO          = 53
    ECHOE         = 54
    ECHOK         = 55
    ECHONL        = 56
    NOFLSH        = 57
    TOSTOP        = 58
    IEXTEN        = 59
    ECHOCTL       = 60
    ECHOKE        = 61
    PENDIN        = 62
    OPOST         = 70
    OLCUC         = 71
    ONLCR         = 72
    OCRNL         = 73
    ONOCR         = 74
    ONLRET        = 75
    CS7           = 90
    CS8           = 91
    PARENB        = 92
    PARODD        = 93
    TTY_OP_ISPEED = 128
    TTY_OP_OSPEED = 129
)

func NewSSH2() *SSH2{
  return &SSH2{}
}

func (s *SSH2)InitCfg(server string,username string,password string){
    s.server = server
    s.username = username
    s.password = clientPassword(password)
}


func (s *SSH2)Connect(){
	config := &ssh.ClientConfig{
        User: s.username,
        Auth: []ssh.ClientAuth{
            // ClientAuthPassword wraps a ClientPassword implementation
            // in a type that implements ClientAuth.
            ssh.ClientAuthPassword(s.password),
        },
    }
    fmt.Println("%v\n",s.password)
    fmt.Println("config----:",config)
	fmt.Println(s.server)
    client, err := ssh.Dial("tcp", s.server, config)

	fmt.Println(client)
    if err != nil {
        panic("Failed to dial: " + err.Error())
    }
    s.client = client
    fmt.Printf("client:%v\n",s.client)
}

func (s *SSH2)Exec(command string) string{
    session, err := s.client.NewSession()
    if err != nil {
        log.Fatalf("unable to create session: %s", err)
    }

    s.session = session
    // Set up terminal modes
    modes := ssh.TerminalModes{
        ECHO:          0,     // disable echoing
        TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    // Request pseudo terminal
    if err := s.session.RequestPty("xterm", 80, 40, modes); err != nil {
        log.Fatalf("request for pseudo terminal failed: %s", err)
    }
 
    var b bytes.Buffer
    s.session.Stdout = &b
    if err := session.Run(command); err != nil {
        panic("Failed to run: " + err.Error())
    }
    //fmt.Println(b.String())
    return b.String()
}


func (s *SSH2)Stop(){
	s.session.Close()
	s.client.Close()
}
