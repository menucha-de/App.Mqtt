package swagger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"sync"

	"github.com/peramic/utils"

	loglib "github.com/peramic/logging"
)

const filename = "./conf/mqtt/config.json"
const dirname = "./conf/mqtt/"

type config struct {
	Security      Security      `json:"security"`
	AccessControl AccessControl `json:"accControl"`
}
type command struct {
	cmd *exec.Cmd
	mu  sync.Mutex
}

var configuration *config
var lg *loglib.Logger = loglib.GetLogger("mqtt")
var cmd command

//*exec.Cmd = exec.Command("/usr/sbin/mosquitto", "-c", "/etc/mosquitto/mosquitto.conf")

func init() {

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		os.MkdirAll(dirname, 0700)
	}
	if _, err := os.Stat("/var/log/mosquitto"); os.IsNotExist(err) {
		os.MkdirAll("/var/log/mosquitto", 0700)
	}
	if _, err := os.Stat("/var/lib/mosquitto"); os.IsNotExist(err) {
		os.MkdirAll("/var/lib/mosquitto", 0700)
	}
	user, err := user.Lookup("mosquitto")
	if err != nil {
		lg.WithError(err).Warning("User 'mosquitto' does not exists")
	} else {
		uid, _ := strconv.Atoi(user.Uid)
		gid, _ := strconv.Atoi(user.Gid)
		err = os.Chown("/var/log/mosquitto", uid, gid)
		if err != nil {
			lg.WithError(err).Warning("Failed to change mosquitto log ownership")
		}
		os.Chown("/var/lib/mosquitto", uid, gid)
		if err != nil {
			lg.WithError(err).Warning("Failed to change mosquitto lib ownership")
		}
	}

	f, err := os.Open(filename)
	cmd.cmd = exec.Command("/usr/sbin/mosquitto", "-c", "/etc/mosquitto/mosquitto.conf")
	// if we os.Open returns an error then handle it
	if err == nil {
		dec := json.NewDecoder(f)
		err = dec.Decode(&configuration)

		if err != nil {
			lg.WithError(err).Warning("Failed to parse config")
			configuration = new(config)
			configuration.AccessControl.Anonymous = true
		}

	} else {
		lg.WithError(err).Trace("Failed to read config")
		configuration = new(config)
		configuration.AccessControl.Anonymous = true
	}
	defer f.Close()

	restartservice()
}
func (c *config) serialize() {

	f, err := os.Create(filename)

	if err != nil {
		lg.WithError(err).Error("Failed to create or open configuration file")
	} else {
		enc := json.NewEncoder(f)
		enc.SetIndent("", " ")
		enc.Encode(c)
	}
	defer f.Close()
}
func writeConfig(c *config) error {
	file := "/etc/mosquitto/conf.d/mosquitto.conf"
	f, err := os.Create(file)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		return err
	}
	defer f.Close()

	if c.Security.SSL {
		fmt.Fprintln(f, "listener 8883")
		fmt.Fprintln(f, "cafile "+caCertificateFilePath)
		fmt.Fprintln(f, "certfile "+serverCertificateFile)
		fmt.Fprintln(f, "keyfile  "+serverKeyFile)
		if c.Security.ClientVerification {
			fmt.Fprintln(f, "require_certificate true")

			if utils.FileExists(revocationListFilePath) {
				fmt.Fprintln(f, "crlfile "+revocationListFilePath)
			}
		}

	} else {
		fmt.Fprintln(f, "listener 1883")
	}

	if !c.AccessControl.Anonymous {
		fmt.Fprintln(f, "allow_anonymous false")
	} else {
		fmt.Fprintln(f, "allow_anonymous true")
	}
	if utils.FileExists(passFilePath) {
		fmt.Fprintln(f, "password_file "+passFilePath)
	}

	if utils.FileExists(acControlFilePath) {
		fmt.Fprintln(f, "acl_file "+acControlFilePath)
	}
	return nil
}
func restartservice() error {
	if cmd.cmd.Process != nil {
		if err := cmd.cmd.Process.Kill(); err != nil {
			lg.WithError(err).Error("failed to kill mqtt process: ")
			//return err
		}
		//lg.Trace("mqtt process killed")
		cmd.mu.Lock()
		cmd.cmd = exec.Command("/usr/sbin/mosquitto", "-c", "/etc/mosquitto/mosquitto.conf")
		cmd.mu.Unlock()
	}
	cmdReader, err := cmd.cmd.StdoutPipe()
	if err != nil {
		lg.Error(err.Error())
		return err
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {

			line := scanner.Text()
			if strings.Contains(line, "Error") {
				lg.Error(line)
			} else {
				lg.Debug(line)
			}
		}

	}()

	//errs := make(chan error)
	if err := cmd.cmd.Start(); err != nil {
		lg.Error(err.Error())
		return err
	}

	go func() {
		cmd.mu.Lock()
		defer cmd.mu.Unlock()
		if err := cmd.cmd.Wait(); err != nil {
			lg.Trace(err.Error())
		}

	}()

	return nil
}
