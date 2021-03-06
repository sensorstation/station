package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type WSServer struct {
	c websocket.Conn
}

type KeyVal struct {
	K string
	V interface{}
}

var (
	WServ WSServer
)

// ServeHTTP
func (ws WSServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Println("Warning Cors Header to '*'")
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:       []string{"echo"},
		InsecureSkipVerify: true, // Take care of CORS
		// OriginPatterns: ["*"],
	})

	if err != nil {
		log.Println("ERROR ", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "houston, we have a problem")

	log.Println("Wait a minute...")
	tQ := time.Tick(time.Second)

	cb := func(q Quote) {
		c1 := c
		q1 := q

		log.Println("ws [IN] message: ", q1)
		err = wsjson.Write(r.Context(), c1, q1)
		if err != nil {
			log.Println("ERROR: ", err)
		}
	}
	quoteCallbacks[c] = cb
	defer func() { delete(quoteCallbacks, c) }()

	go func() {
		running := true
		for running {

			select {
			case now := <-tQ:
				t := NewTimeMsg(now)

				if config.Debug {
					log.Printf("Sending time %q", t)
				}

				err = wsjson.Write(r.Context(), c, t)
				if err != nil {
					log.Println("ERROR: ", err)
					running = false
				}
				tf := KeyVal{ K: "tempf", V: 88 }
				sl := KeyVal{ K: "soil", V: .49 }
				lt := KeyVal{ K: "light", V: .62 }
				hu := KeyVal{ K: "humid", V: .12 }
				err = wsjson.Write(r.Context(), c, tf)
				err = wsjson.Write(r.Context(), c, sl)
				err = wsjson.Write(r.Context(), c, lt)
				err = wsjson.Write(r.Context(), c, hu)
			}
		}
	}()

	for {
		data := make([]byte, 8192)
		_, data, err := c.Read(r.Context())
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			log.Println("ws Closed")
			return
		}
		if err != nil {
			log.Println("ERROR: reading websocket ", err)
			return
		}
		log.Printf("incoming: %s", string(data))
	}

}

func echo(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}

	err = w.Close()
	return err
}
