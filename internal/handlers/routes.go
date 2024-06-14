package handlers

import (
	"net/http"
	"path/filepath"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	// add a css file to route
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", h.home)
	mux.HandleFunc("/benefits", h.benefits)
	mux.HandleFunc("/tariffs", h.tariffs)
	mux.HandleFunc("/list", h.list)
	mux.HandleFunc("/add_user", h.newUser)
	mux.HandleFunc("/usermenu", h.usermenu)
	return mux
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

// func rateLimit(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// here
// 		next.ServeHTTP(w, r)
// 	}
// }

// позже изучить

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	// проверяем заголовки
// 	if r.Header.Get("Upgrade") != "websocket" {
// 		return
// 	}
// 	if r.Header.Get("Connection") != "Upgrade" {
// 		return
// 	}
// 	k := r.Header.Get("Sec-Websocket-Key")
// 	if k == "" {
// 		return
// 	}

// 	// вычисляем ответ
// 	sum := k + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
// 	hash := sha1.Sum([]byte(sum))
// 	str := base64.StdEncoding.EncodeToString(hash[:])

// 	// Берем под контроль соединение https://pkg.go.dev/net/http#Hijacker
// 	hj, ok := w.(http.Hijacker)
// 	if !ok {
// 		return
// 	}
// 	conn, bufrw, err := hj.Hijack()
// 	if err != nil {
// 		return
// 	}
// 	defer conn.Close()

// 	// формируем ответ
// 	bufrw.WriteString("HTTP/1.1 101 Switching Protocols\r\n")
// 	bufrw.WriteString("Upgrade: websocket\r\n")
// 	bufrw.WriteString("Connection: Upgrade\r\n")
// 	bufrw.WriteString("Sec-Websocket-Accept: " + str + "\r\n\r\n")
// 	bufrw.Flush()

// 	// выводим все, что пришло от клиента
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := bufrw.Read(buf)
// 		if err != nil {
// 			return
// 		}
// 		fmt.Println(buf[:n])
// 	}
// }
