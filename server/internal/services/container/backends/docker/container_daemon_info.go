/*
*

  - @license

  - Copyright (c) The Authors (see the AUTHORS file)
    *

  - This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).

  - You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
*/
package dockerbackend

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log/slog"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/docker/docker/client"
	container "github.com/openorch/openorch/server/internal/services/container/types"
	"github.com/pkg/errors"

	"github.com/openorch/openorch/sdk/go/logger"
)

func (d *DockerBackend) DaemonInfo(container.DaemonInfoRequest) (*container.DaemonInfoResponse, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	inf, err := d.client.Info(context.Background())
	// even on windows, we want a docker daemon that can run linux containers
	// as our containers are linux ones
	if err == nil && inf.OSType == "linux" {
		ret := &container.DaemonInfoResponse{
			Available: true,
		}
		if d.dockerHost != "" && d.dockerPort != 0 {
			addr := fmt.Sprintf("%v:%v", d.dockerHost, d.dockerPort)
			ret.Address = &addr
		} else if d.dockerHost != "" {
			ret.Address = &d.dockerHost
		}

		return ret, nil
	}

	ip, port, err := d.tryFixDockerAddress()
	if err != nil {
		logger.Warn(
			"Cannot find Docker address",
			slog.String("error", err.Error()),
		)
		return &container.DaemonInfoResponse{
			Available: false,
		}, nil
	}
	logger.Info(
		"Found Docker address",
		slog.String("ip", ip),
		slog.Int("port", port),
	)

	d.dockerHost = ip
	d.dockerPort = port

	daemonAddress := fmt.Sprintf("%v:%v", ip, port)
	return &container.DaemonInfoResponse{
		Available: true,
		Address:   &daemonAddress,
	}, nil
}

func (d *DockerBackend) tryFixDockerAddress() (ip string, port int, err error) {
	dockerTcpPort := 2375

	switch runtime.GOOS {
	case "darwin":
		host, err := getLimaAddress()
		if err != nil {
			return "", 0, errors.Wrap(err, "getting lima address")
		}

		if host != "" {
			newDockerClient, err := client.NewClientWithOpts(
				client.WithAPIVersionNegotiation(),
				client.WithHost(fmt.Sprintf("tcp://%s", host)),
				// ??, see dind below but is it needed here
				client.WithVersion("1.44"),
			)
			if err != nil {
				return "", 0, errors.Wrap(
					err,
					"error creating new Docker client",
				)
			}

			inf, err := newDockerClient.Info(context.Background())
			if err != nil {
				return "", 0, errors.Wrap(
					err,
					"error pinging Docker with new address",
				)
			}
			if inf.OSType != "linux" {
				return "", 0, errors.Wrap(
					err,
					fmt.Sprintf(
						"docker os type is not linux but '%v'",
						inf.OSType,
					),
				)
			}

			d.client = newDockerClient
			port := 0
			if strings.Contains(host, ":") {
				parts := strings.Split(host, ":")
				host = parts[0]
				port64, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					return "", 0, errors.Wrap(
						err,
						"error converting port to int",
					)
				}
				port = int(port64)
			}
			return host, port, nil
		}
	case "windows":
		ipAddress, err := getWslIpAddress()
		if err != nil {
			return "", 0, errors.Wrap(err, "error getting WSL IP address")
		}

		if ipAddress != "" {
			newDockerClient, err := client.NewClientWithOpts(
				client.WithAPIVersionNegotiation(),
				client.WithHost(
					fmt.Sprintf("tcp://%s:%v", ipAddress, dockerTcpPort),
				),
				// I think dind is not up to date to use the current version 1.45, win breaks if we do
				// not specify this
				client.WithVersion("1.44"),
			)
			if err != nil {
				return "", 0, errors.Wrap(
					err,
					"error creating new Docker client",
				)
			}

			_, err = newDockerClient.Ping(context.Background())
			if err != nil {
				return "", 0, errors.Wrap(
					err,
					"error pinging Docker with new address",
				)
			}

			d.client = newDockerClient

			return ipAddress, 2375, nil
		}
	}

	return "", 0, fmt.Errorf(
		"runtime '%v' has no docker address fix",
		runtime.GOOS,
	)
}

func getLimaAddress() (string, error) {
	cmd := exec.Command(
		"limactl",
		"list",
		"docker",
		"--format",
		`unix://{{.Dir}}/sock/docker.sock`,
	)

	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "failed to execute lima docker command")
	}

	dockerHost := strings.TrimSpace(string(output))

	return dockerHost, nil
}

func getWslIpAddress() (string, error) {
	if runtime.GOOS != "windows" {
		return "", fmt.Errorf("not a Windows system")
	}

	cmd := exec.Command("wsl", "ip", "addr", "show", "eth0")

	// this doesn't seem to work to fix the UTF8 issue but I'll still leave it here
	cmd.Env = append(cmd.Env, "WSL_UTF8=1")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.Wrap(
			err,
			fmt.Sprintf("wsl output: '%v'", out.String()),
		)
	}
	output := out.Bytes()
	var decodedOutput string

	isUtf16, littleEndian := detectUtf16(output)
	if isUtf16 {
		decodedOutput, err = utf16ToUtf8(output, littleEndian)
		if err != nil {
			return "", fmt.Errorf("error decoding UTF-16 output: %v", err)
		}
	} else {
		decodedOutput = string(output)
	}

	re := regexp.MustCompile(`inet\s+(\d+\.\d+\.\d+\.\d+)/`)
	ipAddressMatch := re.FindStringSubmatch(decodedOutput)
	if len(ipAddressMatch) > 1 {
		return ipAddressMatch[1], nil
	}
	return "", fmt.Errorf("IP address not found in output")
}

func detectUtf16(b []byte) (bool, bool) {
	if len(b) < 2 {
		return false, false
	}

	// Check for BOM
	if b[0] == 0xFE && b[1] == 0xFF {
		return true, false // UTF-16 BE
	}
	if b[0] == 0xFF && b[1] == 0xFE {
		return true, true // UTF-16 LE
	}

	// Heuristic: check for alternating null bytes in even positions
	nullCount := 0
	for i := 1; i < len(b); i += 2 {
		if b[i] == 0 {
			nullCount++
		}
	}
	if nullCount > len(
		b,
	)/4 { // More than 25% of the bytes in odd positions are nulls
		return true, true // Assuming little-endian if no BOM but pattern matches
	}

	return false, false
}

func utf16ToUtf8(b []byte, littleEndian bool) (string, error) {
	if len(b)%2 != 0 {
		return "", fmt.Errorf("input byte slice has odd length")
	}

	u16s := make([]uint16, len(b)/2)
	if littleEndian {
		for i := range u16s {
			u16s[i] = binary.LittleEndian.Uint16(b[2*i:])
		}
	} else {
		for i := range u16s {
			u16s[i] = binary.BigEndian.Uint16(b[2*i:])
		}
	}
	u8s := make([]byte, 0, len(u16s)*2)
	for _, r := range utf16.Decode(u16s) {
		buf := make([]byte, 4)
		n := utf8.EncodeRune(buf, r)
		u8s = append(u8s, buf[:n]...)
	}

	return string(u8s), nil
}
