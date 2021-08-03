package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"graphgl-demo/graph/model"
	"io/ioutil"
	"log"
	"net/http"
)


type RootHandler func(http.ResponseWriter, *http.Request) error

func (re *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	println(r.Body)

}

type CustomServer struct {
		transports []graphql.Transport
		exec       *executor.Executor
	}

func (s *CustomServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("loan request", r.Body)
	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
		return
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))

	fmt.Printf("loan parsed", rdr1)

	defer func() {
		if err := recover(); err != nil {
			err := s.exec.PresentRecoveredError(r.Context(), err)
			resp := &graphql.Response{Errors: []*gqlerror.Error{err}}
			b, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write(b)
		}
	}()

	r = r.WithContext(graphql.StartOperationTrace(r.Context()))

	transport := s.getTransport(r)
	if transport == nil {
		fmt.Printf("transport error")
		return
	}

	transport.Do(w, r, s.exec)
}

func New(es graphql.ExecutableSchema) *CustomServer {
	return &CustomServer{
		exec: executor.New(es),
	}
}


func (s *CustomServer) AddTransport(transport graphql.Transport) {
	s.transports = append(s.transports, transport)
}

func (s *CustomServer) getTransport(r *http.Request) graphql.Transport {
	for _, t := range s.transports {
		if t.Supports(r) {
			return t
		}
	}
	return nil
}

func (s *CustomServer) CustomHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


		fmt.Printf("loan request", r.Body)

		var newLoan model.NewLoan
		if err := json.NewDecoder(r.Body).Decode(&newLoan); err != nil {
			println("error in parsing" , r.Body)
		}

		fmt.Printf("loan parsed", newLoan)

		defer func() {
			if err := recover(); err != nil {
				err := s.exec.PresentRecoveredError(r.Context(), err)
				resp := &graphql.Response{Errors: []*gqlerror.Error{err}}
				b, _ := json.Marshal(resp)
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write(b)
			}
		}()

		r = r.WithContext(graphql.StartOperationTrace(r.Context()))

		transport := s.getTransport(r)
		if transport == nil {
			fmt.Printf("transport error")
			return
		}

		transport.Do(w, r, s.exec)
	})
}



