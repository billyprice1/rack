package client

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	Id      string    `json:"id"`
	App     string    `json:"app"`
	Command string    `json:"command"`
	Host    string    `json:"host"`
	Image   string    `json:"image"`
	Name    string    `json:"name"`
	Ports   []string  `json:"ports"`
	Release string    `json:"release"`
	Cpu     float64   `json:"cpu"`
	Memory  float64   `json:"memory"`
	Started time.Time `json:"started"`
}

type Processes []Process

func (c *Client) GetProcesses(app string) (Processes, error) {
	var processes Processes

	err := c.Get(fmt.Sprintf("/apps/%s/processes", app), &processes)

	if err != nil {
		return nil, err
	}

	return processes, nil
}

func (c *Client) RunProcessAttached(app, process, command string, in io.Reader, out io.Writer) (int, error) {
	r, w := io.Pipe()

	defer r.Close()
	defer w.Close()

	ch := make(chan int)

	go copyWithExit(out, r, ch)

	err := c.Stream(fmt.Sprintf("/apps/%s/processes/%s/run", app, process), map[string]string{"Command": command}, in, w)

	if err != nil {
		return 0, err
	}

	code := <-ch

	return code, nil
}

func (c *Client) RunProcessDetached(app, process, command string) error {
	var success interface{}

	params := map[string]string{
		"command": command,
	}

	err := c.Post(fmt.Sprintf("/apps/%s/processes/%s/run", app, process), params, &success)

	if err != nil {
		return err
	}

	return nil
}
func (c *Client) StopProcess(app, id string) (*Process, error) {
	var process Process

	err := c.Delete(fmt.Sprintf("/apps/%s/processes/%s", app, id), &process)

	if err != nil {
		return nil, err
	}

	return &process, nil
}

func copyWithExit(w io.Writer, r io.Reader, ch chan int) {
	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)

		if err != nil {
			break
		}

		if s := string(buf[0:n]); strings.HasPrefix(s, "F1E49A85-0AD7-4AEF-A618-C249C6E6568D:") {
			code, _ := strconv.Atoi(s[37:])
			ch <- code
			return
		}

		_, err = w.Write(buf[0:n])

		if err != nil {
			break
		}
	}
}