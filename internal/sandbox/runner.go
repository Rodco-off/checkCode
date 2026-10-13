package sandbox

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/moby/moby/api/pkg/stdcopy"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

type Runner struct {
	cli         *client.Client
	containerID string
}

func NewRunner(cont context.Context) (*Runner, error) {
	cli, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("docker client %w", err)
	}

	resp, err := cli.ContainerCreate(cont, client.ContainerCreateOptions{
		Config: &container.Config{
			Image: "golang:1.21-alpine",
			Cmd:   []string{"sleep", "infinity"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("контейнер создан с %w", err)
	}

	if _, err := cli.ContainerStart(cont, resp.ID, client.ContainerStartOptions{}); err != nil {
		return nil, fmt.Errorf("контейнер запустился с %w", err)
	}

	return &Runner{cli: cli, containerID: resp.ID}, nil
}

func (run *Runner) Run(cont context.Context, userCode, input string) (string, error) {
	content, err := GenerativeMainFile(userCode)
	if err != nil {
		return "", fmt.Errorf("generative: %w", err)
	}

	tarBuf, err := makeTar("main.go", content)
	if err != nil {
		return "", fmt.Errorf("tar: %w", err)
	}

	_, err = run.cli.CopyToContainer(cont, run.containerID, client.CopyToContainerOptions{DestinationPath: "/tmp", Content: tarBuf})
	if err != nil {
		return "", fmt.Errorf("copy: %w", err)
	}

	execCfg := client.ExecCreateOptions{
		Cmd:          []string{"go", "run", "/tmp/main.go", input},
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir:   "/tmp",
	}

	execResp, err := run.cli.ExecCreate(cont, run.containerID, execCfg)
	if err != nil {
		return "", fmt.Errorf("exec create: %w", err)
	}

	attach, err := run.cli.ExecAttach(cont, execResp.ID, client.ExecAttachOptions{})
	if err != nil {
		return "", fmt.Errorf("exec attach: %w", err)
	}
	defer attach.Close()

	var stdout, stderr bytes.Buffer
	if _, err := stdcopy.StdCopy(&stdout, &stderr, attach.Reader); err != nil {
		return "", fmt.Errorf("read: %w", err)
	}

	return stdout.String() + stderr.String(), nil

}

func (r *Runner) ContainerID() string {
	return r.containerID
}

func makeTar(filename, content string) (io.Reader, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	header := &tar.Header{
		Name: filename,
		Mode: 0644,
		Size: int64(len(content)),
	}
	if err := tw.WriteHeader(header); err != nil {
		return nil, err
	}
	if _, err := tw.Write([]byte(content)); err != nil {
		return nil, err
	}
	if err := tw.Close(); err != nil {
		return nil, err
	}

	return &buf, nil
}
