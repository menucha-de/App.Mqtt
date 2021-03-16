/*
 * Mqtts API
 *
 * Mqtts service
 *
 * API version: 1.0.0
 */
package swagger

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
	"os/exec"

	"net/http"
	"os"

	utils "github.com/menucha-de/utils"

	"golang.org/x/crypto/pkcs12"
)

var password string
var acControlFilePath string = "/etc/mosquitto/ca_certificates/aclfile"
var acControlFilePathEx string = "/opt/menucha-de/examples/aclfile.example"
var caCertificateFilePath string = "/etc/mosquitto/ca_certificates/cacertfile.crt"
var passFilePath string = "/etc/mosquitto/ca_certificates/passfile"
var passFilePathEx string = "/opt/menucha-de/examples/pskfile.example"
var serverCertificateFilePath string = "/opt/menucha-de/conf/servercertfile"
var revocationListFilePath string = "/etc/mosquitto/ca_certificates/serverrevlistfile"
var serverKeyFile string = "/etc/mosquitto/certs/key.pem"
var serverCertificateFile string = "/etc/mosquitto/certs/cert.pem"

func accessControlGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	s := configuration.AccessControl

	var err = json.NewEncoder(w).Encode(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func accessControlSet(w http.ResponseWriter, r *http.Request) {
	var access *AccessControl

	err := utils.DecodeJSONBody(w, r, &access)
	if err != nil {
		var mr *utils.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			lg.WithError(err).Error("Failed to get access control")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	configuration.AccessControl = *access
	configuration.serialize()
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteACFile(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(acControlFilePath)
	if err != nil {
		lg.WithError(err).Error("Failed to Delete Access Control File")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteCAFile(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(caCertificateFilePath)
	if err != nil {
		lg.WithError(err).Error("Failed to Delete Access Control File")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if configuration.Security.SSL {
		configuration.Security.SSL = false
		configuration.Security.ClientVerification = false
		configuration.serialize()
		err = writeConfig(configuration)
		if err != nil {
			lg.WithError(err).Error("Failed to apply configuration")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err := restartservice()
		if err != nil {
			lg.WithError(err).Error("Failed to apply configuration")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func deletePassFile(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(passFilePath)
	if err != nil {
		lg.WithError(err).Error("Failed to Delete Password File")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !configuration.AccessControl.Anonymous {
		configuration.AccessControl.Anonymous = true
		configuration.serialize()

	}
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteRevFile(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(revocationListFilePath)
	if err != nil {
		lg.WithError(err).Error("Failed to Delete Access Control File")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteServFile(w http.ResponseWriter, r *http.Request) {
	err := os.Remove(serverCertificateFile)
	if err != nil {
		lg.WithError(err).Error("Failed to Delete Server Cerificate File")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.Remove(serverCertificateFilePath)
	if err != nil {
		lg.WithError(err).Error("Failed to Delete Server Cerificate File")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.Remove(serverKeyFile)
	if err != nil {
		lg.WithError(err).Error("Failed to Delete Server Key File")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if configuration.Security.SSL {
		configuration.Security.SSL = false
		configuration.Security.ClientVerification = false
		configuration.serialize()
		err = writeConfig(configuration)
		if err != nil {
			lg.WithError(err).Error("Failed to apply configuration")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err := restartservice()
		if err != nil {
			lg.WithError(err).Error("Failed to apply configuration")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func getACFile(w http.ResponseWriter, r *http.Request) {
	file := acControlFilePath
	if !utils.FileExists(file) {
		file = acControlFilePathEx
	}
	http.ServeFile(w, r, file)
}

func getCertFile(w http.ResponseWriter, r *http.Request) {
	if !utils.FileExists(caCertificateFilePath) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getPassFile(w http.ResponseWriter, r *http.Request) {
	file := passFilePath
	if !utils.FileExists(file) {
		file = passFilePathEx
	}
	http.ServeFile(w, r, file)
}

func getRevFile(w http.ResponseWriter, r *http.Request) {
	if !utils.FileExists(revocationListFilePath) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getServFile(w http.ResponseWriter, r *http.Request) {
	if !utils.FileExists(serverCertificateFilePath) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func headACFile(w http.ResponseWriter, r *http.Request) {
	if !utils.FileExists(acControlFilePath) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func headPassFile(w http.ResponseWriter, r *http.Request) {
	if !utils.FileExists(passFilePath) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func putACFile(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(acControlFilePath)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func putCertFile(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(caCertificateFilePath)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if configuration.Security.SSL {
		err = restartservice()
		if err != nil {
			lg.WithError(err).Error("Failed to apply configuration")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)

}

func putPassFile(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(passFilePath)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//todo mosquito password

	cmdp := exec.Command("mosquitto_passwd", "-U", passFilePath)
	if err := cmdp.Start(); err != nil {
		lg.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := cmdp.Wait(); err != nil {
		lg.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func putPassPhrase(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	n, err := b.ReadFrom(r.Body)
	if err != nil || n == 0 {
		http.Error(w, "Could not read passphrase value", http.StatusBadRequest)
		return
	}
	password = b.String()
	w.WriteHeader(http.StatusNoContent)

}

func putRevFile(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(revocationListFilePath)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func putServFile(w http.ResponseWriter, r *http.Request) {
	if password == "" {
		http.Error(w, "Please upload first the PassPhrase", http.StatusInternalServerError)
		return
	}
	f, err := os.Create(serverCertificateFilePath)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = f.Write(body)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = keyStore(body)
	if err != nil {
		lg.WithError(err).Error(err.Error())
		os.Remove(serverCertificateFilePath)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	password = ""
	if configuration.Security.SSL {
		err = restartservice()
	}
	if err != nil {
		lg.WithError(err).Error("Could not start mqtt service")
		os.Remove(serverCertificateFilePath)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func securityGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	s := configuration.Security

	var err = json.NewEncoder(w).Encode(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func securitySet(w http.ResponseWriter, r *http.Request) {
	var security *Security
	err := utils.DecodeJSONBody(w, r, &security)
	if err != nil {
		var mr *utils.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			lg.WithError(err).Error("Failed to get access control")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	err = validateSSl(*security)
	if err != nil {
		lg.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	configuration.Security = *security
	configuration.serialize()
	err = writeConfig(configuration)
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = restartservice()
	if err != nil {
		lg.WithError(err).Error("Failed to apply configuration")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func keyStore(body []byte) error {

	privateKey, certificate, err := pkcs12.Decode(body, password)
	if err != nil {
		lg.Error(err.Error())
		return err
	}

	if err := verify(certificate); err != nil {
		lg.Error(err.Error())
		return err
	}

	priv, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		err = errors.New("expected RSA private key type")
		lg.Error(err.Error())
		return err
	}

	keyFile, err := os.Create(serverKeyFile)
	if err != nil {
		lg.Error(err.Error())
		return err
	}
	defer keyFile.Close()
	err = pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	if err != nil {
		lg.Error(err.Error())
		return err
	}

	certFile, err := os.Create(serverCertificateFile)
	if err != nil {
		lg.Error(err.Error())
		return err
	}
	defer certFile.Close()
	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certificate.Raw})
	if err != nil {
		lg.Error(err.Error())
		return err
	}
	return nil
}
func validateSSl(security Security) error {

	if security.SSL {
		if !utils.FileExists(caCertificateFilePath) {
			return errors.New("CA Certificate not uploaded")
		}
		if !utils.FileExists(serverCertificateFile) {
			return errors.New("Server Certificate not uploaded")
		}
		if !utils.FileExists(serverKeyFile) {
			return errors.New("Server Key not uploaded")
		}
	}
	return nil
}
func verify(cert *x509.Certificate) error {
	_, err := cert.Verify(x509.VerifyOptions{})
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case x509.CertificateInvalidError:
		switch e.Reason {
		case x509.Expired:
			return errors.New("certificate has expired or is not yet valid")
		default:
			return err
		}
	case x509.UnknownAuthorityError:
		// Apple cert isn't in the cert pool
		// ignoring this error
		return nil
	default:
		return err
	}
}
