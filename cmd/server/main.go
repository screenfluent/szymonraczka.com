package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      newHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func newHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /sixteenth-attempt", sixteenthAttemptPost)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<!doctype html>
    <html lang="en">
    <head>
       <meta charset="utf-8">
       <title>Szymon Raczka</title>
    </head>
    <body>
       <main>
          <p>hello, nice to see you here.</p>

          <h1>I am Szymon</h1>

          <section>
             <h2>Projects</h2>
             <ul>
                <li>
                   <a href="/">This Website</a>
                   <time>June 2026</time>
                </li>
             </ul>
          </section>

          <section>
             <h2>Posts</h2>
             <ul>
                <li>
                   <a href="/sixteenth-attempt">Sixteenth Attempt</a>
                </li>
             </ul>
          </section>
       </main>
    </body>
    </html>`)
}

func sixteenthAttemptPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<!doctype html>
    <html lang="en">
    <head>
       <meta charset="utf-8">
       <title>Sixteenth Attempt</title>
    </head>
    <body>
       <main>
          <article>
             <h1>Sixteenth Attempt</h1>
          </article>
       </main>
    </body>
    </html>`)
}
