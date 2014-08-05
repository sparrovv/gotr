package main

import "os/exec"

type Player struct {
	Name string
}

func (p *Player) Play(path string) (output []byte, err error) {
	cmd := exec.Command("sh", "-l", "-c", p.Name+" "+path)
	output, err = cmd.Output()
	return
}
