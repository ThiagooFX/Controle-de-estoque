package services

import (
	"api/internal/repository"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ua-parser/uap-go/uaparser"
)

func GetIP(r *http.Request) string {
	// IP via proxy
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return strings.Split(ip, ",")[0]
	}
	// IP direto
	ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	return ip
}

func Addlogs(r *http.Request, action string) error {
	uaString := r.UserAgent()
	parser := uaparser.NewFromSaved()
	client := parser.Parse(uaString)

	ocorredAt := time.Now()
	user := "Thiago"
	userEmail := "dev.tfx11@gmail.com"
	dispositivo := client.Device.Family
	ip := GetIP(r)

	_, err := repository.DB.Exec(`
		INSERT INTO logs (action, ocorred_at, user, user_email, device, ip)
		VALUES (?, ?, ?, ?, ?, ?)`,
		action, ocorredAt.Format(time.RFC3339), user, userEmail, dispositivo, ip,
	)
	return err
}
