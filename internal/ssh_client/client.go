package ssh_client

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
)

type SSHCommand struct {
	Path   string
	Env    []string
	Stdin  io.Reader
	Stdout *bytes.Buffer
	Stderr io.Writer
}

type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
	Ctx    context.Context
	conn   *ssh.Client
}

var lock sync.Mutex

func (client *SSHClient) RunCommand(cmd *SSHCommand) error {
	var (
		session *ssh.Session
		err     error
	)

	if session, err = client.NewSession(); err != nil {
		return err
	}
	defer session.Close()

	if err = client.prepareCommand(session, cmd); err != nil {
		return err
	}
	err = session.Run(cmd.Path)
	if e, ok := err.(*ssh.ExitMissingError); ok {
		log.Printf("Exit code missing: %s\n", e.Error())
		lock.Lock()
		client.conn = nil
		lock.Unlock()
		return client.RunCommand(cmd)
	}
	return err
}

func (client *SSHClient) prepareCommand(session *ssh.Session, cmd *SSHCommand) error {
	for _, env := range cmd.Env {
		variable := strings.Split(env, "=")
		if len(variable) != 2 {
			continue
		}

		if err := session.Setenv(variable[0], variable[1]); err != nil {
			return err
		}
	}

	if cmd.Stdin != nil {
		stdin, err := session.StdinPipe()
		if err != nil {
			return fmt.Errorf("unable to setup stdin for session: %v", err)
		}
		go io.Copy(stdin, cmd.Stdin)
	}

	if cmd.Stdout != nil {
		stdout, err := session.StdoutPipe()
		if err != nil {
			return fmt.Errorf("unable to setup stdout for session: %v", err)
		}
		go io.Copy(bufio.NewWriterSize(cmd.Stdout, 10000), stdout)
	}

	if cmd.Stderr != nil {
		stderr, err := session.StderrPipe()
		if err != nil {
			return fmt.Errorf("unable to setup stderr for session: %v", err)
		}
		go io.Copy(cmd.Stderr, stderr)
	}

	return nil
}

func (client *SSHClient) NewSession() (*ssh.Session, error) {
	lock.Lock()
	defer lock.Unlock()
	if err := client.connect(); err != nil {
		return nil, err
	}
	session, err := client.conn.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %s", err)
	}
	/*
		modes := ssh.TerminalModes{
			ssh.ECHO:          0,     // disable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}

		if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
			session.Close()
			return nil, fmt.Errorf("request for pseudo terminal failed: %s", err)
		}
	*/
	return session, nil
}

func (client *SSHClient) connect() error {
	if client.conn == nil {
		log.Printf("Connect to: %s:%d\n", client.Host, client.Port)
		var err error
		client.conn, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
		if err != nil {
			return fmt.Errorf("failed to dial: %s", err)
		}
	}
	return nil
}

func (client *SSHClient) FileFromRemoteHost(localFile, targetFile string) error {
	var (
		session *ssh.Session
		err     error
	)

	if session, err = client.NewSession(); err != nil {
		return err
	}
	defer session.Close()
	iw, err := session.StdinPipe()
	if err != nil {
		log.Println("Failed to create input pipe: " + err.Error())
		return err
	}
	or, err := session.StdoutPipe()
	if err != nil {
		log.Println("Failed to create output pipe: " + err.Error())
		return err
	}
	src, srcErr := os.Create(localFile)
	if srcErr != nil {
		log.Println("Failed to create source file: " + srcErr.Error())
		return srcErr
	}
	go func() {
		fmt.Fprint(iw, "\x00")
		r, err := gzip.NewReader(or)
		if err != nil {
			log.Println("Failed to read gzip: " + err.Error())
			return
		}
		if n, err := io.Copy(src, r); err != nil {
			fmt.Println(n)
			fmt.Fprint(iw, "\x02")
			return
		}
		fmt.Fprint(iw, "\x00")
	}()
	cmd := fmt.Sprintf("xzcat %s |sed -re 's/(.*) requestBody=\".*\" (serveTime.*)/\\1 \\2/' |gzip -qc ", targetFile)
	if !strings.HasSuffix(targetFile, ".xz") {
		cmd = cmd[2:]
	}
	return session.Run(cmd)
}
