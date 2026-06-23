package main

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"

	"szymonraczka.com/web"
)

type pageData struct {
	Title string
	Body  template.HTML
}

var pageTemplate = template.Must(template.New("page").Parse(`<!doctype html>
    <html lang="en">
    <head>
       <meta charset="utf-8">
       <meta name="viewport" content="width=device-width, initial-scale=1">
       <link rel="stylesheet" href="/static/site.css">
       <title>{{.Title}}</title>
       <script
        defer
        src="https://assets.onedollarstats.com/stonks.js"
       ></script>
    </head>
    <body>
       <header>
          <a href="/">Szymon Rączka</a>
       </header>
       <main>
          {{.Body}}
       </main>
    </body>
    </html>`))

func renderPage(w http.ResponseWriter, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := pageTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

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
	mux.HandleFunc("GET /about", aboutPage)
	mux.HandleFunc("GET /now", nowPage)
	mux.HandleFunc("GET /sixteenth-attempt", sixteenthAttemptPost)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	renderPage(w, pageData{
		Title: "Home - szymonraczka.com",
		Body: template.HTML(`<section>
   <p>Hey, I'm Szymon Rączka.</p>

   <h1>This is my personal website.</h1>

   <p>
      I am building my online home: a place for notes, web projects, experiments, and things I learn along the way.
   </p>

   <p>
      Right now I am keeping it simple. This website is built in Go and will grow slowly as I write, build, and publish more.
   </p>
</section>

        <section>
           <h2>Projects</h2>
           <ul>
              <li>
                 <a href="/">szymonraczka.com</a>
                 <time>June 2026</time>
                 <p>A tiny Go website coded by hand from scratch as a learning project.</p>
              </li>
           </ul>
        </section>

        <section>
           <h2>Pages</h2>
           <ul>
              <li>
                 <a href="/about">About</a>
              </li>
              <li>
                 <a href="/now">Now</a>
              </li>
           </ul>
        </section>

        <section>
           <h2>Posts</h2>
           <ul>
              <li>
                 <p>First post coming soon</p>
              </li>
           </ul>
        </section>`),
	})
}

func sixteenthAttemptPost(w http.ResponseWriter, r *http.Request) {
	renderPage(w, pageData{
		Title: "Sixteenth Attempt",
		Body: template.HTML(`<article>
       <h1>Sixteenth Attempt</h1>
    </article>`),
	})
}

func nowPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, pageData{
		Title: "Now - szymonraczka.com",
		Body: template.HTML(`<article>
   <h1>Now</h1>

   <p>
      I am learning Go by building this website from scratch.
   </p>

   <p>
      I use AI as a teacher and reviewer, but I type the code myself so I actually understand what I am building.
   </p>

   <p>
      My current focus is small software, clear writing, and rebuilding a slower, more sustainable creative practice.
   </p>
</article>`),
	})
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, pageData{
		Title: "About - szymonraczka.com",
		Body: template.HTML(`<article>
       <h1>About Me</h1>
       <p>
          I'm Szymon Rączka. I make small web projects and write about learning,
          rebuilding, and making things on the internet.
       </p>

       <p>
          In the past I created Handimania, a media project that reached more than 165 million YouTube views. These days I am focused on quieter, smaller, more sustainable work.
       </p>

    </article>`),
	})
}
