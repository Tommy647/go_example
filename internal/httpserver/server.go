package httpserver

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	_db "github.com/Tommy647/go_example/internal/db"
	"github.com/Tommy647/go_example/internal/dbgreeter"
	"github.com/Tommy647/go_example/internal/greeter"
	"github.com/Tommy647/go_example/internal/jwt"
)

// HelloResponse to http requests // @todo: fixme
type HelloResponse struct {
	Response []string `json:"response" xml:"response"`
	Error    error    `json:"error,omitempty" xml:"error"`
}

// TextEncode formatted text
func (h HelloResponse) TextEncode(buf *bytes.Buffer) error {
	if _, err := buf.WriteString("Responses:\n" + strings.Join(h.Response, "\n")); err != nil {
		return err
	}
	if h.Error != nil {
		if _, err := buf.WriteString("Errors:\n" + h.Error.Error() + "\n"); err != nil {
			return err
		}
	}
	return nil
}

// GreetProvider something that greets
type GreetProvider interface {
	Greet(context.Context, string) string
}

// CoffeeProvider something that provides a CoffeeGreet
type CoffeeProvider interface {
	CoffeeGreet(context.Context, string) string
}

// ErrNotAuthorised without jwt
var ErrNotAuthorised = errors.New("not authorise: jwt required")

// HandleHello as a http request
func HandleHello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello http request")
		// try to get our user from the request context, this was added in the middleware from the jwt token
		c := jwt.GetUser(r.Context())
		// if the user is nil, perform the default greeting
		if c == nil {
			// middleware should have caught this, but better safe than sorry
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(ErrNotAuthorised.Error()))
			return
		}
		// get a greeter instance, based on the headers passed in, defaults to the basic greeter
		g, err := getGreeter(r.Header)
		if err != nil {
			log.Println("database", err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		resp := HelloResponse{
			Response: make([]string, 2, len(c.Roles)), //nolint:gomnd // not a magic number
		}

		resp.Response[0] = g.Greet(r.Context(), c.UserName)
		resp.Response[1] = g.Greet(r.Context(), c.Subject)
		// perform custom greetings based on the user calling this request (data from the jwt)
		for i := range c.Roles {
			resp.Response = append(resp.Response, g.Greet(r.Context(), c.Roles[i]))
		}

		data, contentType, err := formatResponse(r.Header.Get("Accept"), resp)
		if err != nil {
			log.Println("encoding", err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", contentType)
		if _, err := data.WriteTo(w); err != nil {
			log.Println("write", err.Error())
		}
	})
}

func HandleCoffee() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("coffee http request")

		g, err := getCoffeeGreeter(r.Header)
		if err != nil {
			log.Println("database", err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		query := r.URL.Query()
		_, _ = w.Write([]byte(g.CoffeeGreet(r.Context(), query.Get("drink"))))
	})
}

// formatResponse depending on the requested 'Accept' Header values
func formatResponse(format string, resp HelloResponse) (*bytes.Buffer, string, error) {
	log.Println("Accept:", format)
	buf := &bytes.Buffer{}
	switch format {
	case "application/xml":
		if err := xml.NewEncoder(buf).Encode(resp); err != nil {
			return nil, format, errors.Wrap(err, "xml")
		}
		return buf, format, nil
	case "application/json":
		if err := json.NewEncoder(buf).Encode(resp); err != nil {
			return nil, format, errors.Wrap(err, "json")
		}
	default:
		format = "application/text"
		if err := resp.TextEncode(buf); err != nil {
			return nil, format, errors.Wrap(err, "text")
		}
	}
	log.Println("Content-Type", format)
	return buf, format, nil
}

// getGreeter based on a header field
func getGreeter(h http.Header) (GreetProvider, error) {
	header := h.Get("X-Greeter")
	if strings.EqualFold(header, "DB") {
		log.Println(">>>>> testing db access")
		db, err := _db.NewConnection()
		if err != nil {
			return nil, err
		}
		return dbgreeter.New(db), nil
	}
	log.Println(">>>>> testing NO db access")
	return greeter.New(), nil
}

// getCoffeeGreeter based on a header field
func getCoffeeGreeter(h http.Header) (CoffeeProvider, error) {
	header := h.Get("X-Greeter")
	if strings.EqualFold(header, "DB") {
		db, err := _db.NewConnection()
		if err != nil {
			log.Println("error from accessing db")
			return nil, err
		}
		log.Println("trying to get a coffee from db")
		return dbgreeter.New(db), nil
	}
	return greeter.New(), nil
}
