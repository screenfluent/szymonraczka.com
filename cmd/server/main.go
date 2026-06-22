package main

import (
	"io"
	"io/fs"
	"log"
	"net/http"
	"time"

	"szymonraczka.com/web"
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


	// Wyciągamy zawartość wbudowanego folderu web/static,
	// żeby fileServer traktował ten folder jako swój główny (root).
	staticFiles, err := fs.Sub(web.Files, "static")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServerFS(staticFiles)

	// StripPrefix odcina "/static/" z przedlądarki, żeby ścieżka np. "site.css"
	// idealnie pasowała do tego, co fileServer widzi u siebie na wierzchu.
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /now", nowPage)
	mux.HandleFunc("GET /sixteenth-attempt", sixteenthAttemptPost)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<!doctype html>
    <html lang="en">
    <head>
       <meta charset="utf-8">
       <meta name="viewport" content="width=device-width, initial-scale=1">
       <link rel="stylesheet" href="/static/site.css">
       <title>Szymon Raczka</title>
    </head>
    <body>
       <header>
          <a href="/">Szymon Raczka</a>
       </header>
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
             <h2>Pages</h2>
             <ul>
                <li>
                   <a href="/now">Now</a>
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
       <meta name="viewport" content="width=device-width, initial-scale=1">
       <link rel="stylesheet" href="/static/site.css">
       <title>Sixteenth Attempt</title>
    </head>
    <body>
       <header>
          <a href="/">Szymon Raczka</a>
       </header>
       <main>
          <article>
             <h1>Sixteenth Attempt</h1>
          </article>
       </main>
    </body>
    </html>`)
}

func nowPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<!doctype html>
        <html lang="en">
        <head>
           <meta charset="utf-8">
           <meta name="viewport" content="width=device-width, initial-scale=1">
           <link rel="stylesheet" href="/static/site.css">
           <title>now</title>
        </head>
        <body>
           <header>
              <a href="/">Szymon Raczka</a>
           </header>
           <main>
              <article>
                 <h1>now</h1>
                 <p>Current focus goes here.</p>
              </article>
           </main>
        </body>
        </html>`)
}
