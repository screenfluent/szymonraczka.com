# Roadmap: szymonraczka.com

Cel: szybko opublikować prostą osobistą stronę marketingową, która pomaga budować zaufanie do Szymona i jego głównych projektów. Nauka Go jest ważna, ale drugorzędna wobec publikacji.

To repo ma być najpierw małą, działającą stroną. Dopiero po publikacji może stać się pobocznym playgroundem do Go, architektury i eksperymentów z treścią.

Ważna część wizji: samo budowanie strony, uczenie się Go i używanie LLM jako wsparcia w nauce jest też materiałem na treści. Projekt ma tworzyć pętlę: buduję → uczę się → porządkuję wiedzę → publikuję notatkę lub post → wracam do budowy.

## Zasada główna

Jeśli publikacja i nauka ciągną w różne strony, wybieramy publikację.

Uczymy się tyle Go, ile potrzeba, żeby rozumieć aktualny slice i móc bezpiecznie iść dalej. Nie próbujemy opanować całej architektury blog engine przed launchem.

## Content loop

Strona nie jest tylko efektem nauki. Jest też miejscem, gdzie ta nauka może zostać pokazana.

Dobre treści dla tego projektu:

- krótkie podsumowania tego, co właśnie zostało zrozumiane
- notatki z budowania strony
- refleksje o pracy z LLM jako partnerem w nauce
- decyzje projektowe opisane prostym językiem
- posty łączące osobistą historię, technologię i sposób myślenia

Reguła: treść może powstać szybko i ręcznie. Nie trzeba najpierw budować systemu notatek, CMS-a ani Markdown pipeline.

## Non-goals przed publikacją

Nie budować teraz:

- CMS-a
- SQLite
- panelu admina
- auth/session/cookies
- dynamicznego systemu postów
- feedów RSS/Atom
- rozbudowanego Markdown pipeline
- frameworka do bespoke pages/posts
- abstrakcji „na zapas”

Te rzeczy mogą wrócić po publikacji, jeśli realna potrzeba treści je uzasadni.

## Obecny kierunek produktu

Minimalna wersja do publikacji powinna mieć:

1. `/` — homepage z krótkim przedstawieniem, linkami do stron, postów i projektów
2. `/about` — osobista historia, wiarygodność, kontekst „kim jestem”
3. `/now` — aktualny fokus
4. `/sixteenth-attempt` — pierwszy zwykły, ręcznie napisany post
5. co najmniej jedną krótką notatkę/post o procesie nauki lub budowania z LLM
6. wspólny HTML shell
7. minimalny CSS wystarczający do czytelności
8. prostą instrukcję uruchomienia/build/deploy

## M0 — domknąć aktualny diff

Najbliższy cel: repo ma być czyste.

Aktualnie template refactor jest rozpoczęty i powinien zostać domknięty osobnym commitem albo świadomie cofnięty. Nie zaczynać nowego app-code slice, dopóki ten diff nie jest zamknięty.

Done gdy:

- `gofmt` wykonany na zmienionych plikach Go
- `gopls check` przechodzi
- `go test ./...` przechodzi
- `go vet ./...` przechodzi
- `staticcheck ./...` przechodzi
- `go build ./...` przechodzi
- `git diff --check` przechodzi
- commit zrobiony
- `git status --short` czysty albo zawiera tylko świadomie zostawione notatki

## M1 — publishable static site

Cel: strona jest prosta, ale gotowa do pokazania ludziom.

Kolejne małe slices:

1. commit template refactor
2. zastąpić placeholder homepage realnym copy
3. dodać `/about`
4. zastąpić placeholder `/now` realnym tekstem
5. napisać pierwszą wersję `/sixteenth-attempt`
6. dodać krótką notatkę/post o procesie budowania i nauki z LLM
7. lekko poprawić CSS pod czytelność mobile/desktop
8. dodać krótkie instrukcje build/run/deploy

Done gdy:

- nie ma placeholderów typu „Current focus goes here”
- wszystkie publiczne linki działają
- tekst jest wystarczająco dobry do pierwszej publikacji
- strona jest czytelna na telefonie i desktopie
- pełna weryfikacja Go przechodzi

## M2 — launch support

Cel: publikacja bez wchodzenia w ciężką architekturę.

Możliwe slices:

1. wybrać deployment target
2. dodać minimalne production docs
3. dodać podstawowe `<title>` i opisy stron
4. dodać Open Graph tylko jeśli strona ma być udostępniana społecznościowo od razu
5. dodać `robots.txt` / sitemap tylko jeśli URL-e są stabilne

## M3 — playground po publikacji

Dopiero po launchu wracają tematy edukacyjne:

- Markdown z plików
- lista postów generowana z folderu
- RSS/Atom
- SQLite, jeśli pojawi się prawdziwa potrzeba persistence
- bespoke interactive storytelling
- lepsze testy HTTP
- notatki architektoniczne / ADR-y

## Rytm pracy

Każda sesja powinna mieć lekki, powtarzalny rytm:

1. sprawdź stan repo
2. wybierz jeden mały slice
3. zdecyduj, czy slice jest kodowy, treściowy, stylistyczny, czy dokumentacyjny
4. dowieź slice do widocznego końca
5. zweryfikuj tylko tyle, ile ma sens dla danego typu zmiany
6. zapisz krótki WIP.co todo o tym, co faktycznie zostało zrobione
7. zrób commit albo jasno powiedz, dlaczego jeszcze nie commitujemy

Typy slices:

- **code** — zmiana zachowania aplikacji, testy i pełniejsza weryfikacja Go
- **content** — tekst strony, copy, notatka, post; testy tylko jeśli zmieniają routing lub strukturę
- **style** — CSS i czytelność; preferuj szybki browser/manual check
- **docs/process** — roadmapa, instrukcje, decyzje; wystarczy diff review i `git diff --check`

WIP.co:

- oficjalne API docs: https://wip.co/api
- wpisujemy tylko rzeczy zrobione, nie plany
- wpis ma być krótki i konkretny
- można używać hashtagu projektu, np. `#szymonraczka`
- API jest opcjonalne; jeśli używane, klucz trzymamy tylko w zmiennej środowiskowej, nigdy w repo
- domyślnie użytkownik może też wkleić wpis ręcznie, jeśli API byłoby tarciem

Przykład wpisu:

```txt
Defined publish-first learning roadmap for my personal site #szymonraczka
```

Przykład API:

```sh
curl -X POST https://api.wip.co/v1/todos \
  -H "Authorization: Bearer $WIP_API_KEY" \
  -d "body=Defined publish-first learning roadmap for my personal site #szymonraczka"
```

## Reguły dla przyszłych LLM sesji

Na starcie sesji:

1. przeczytaj `AGENTS.md`
2. przeczytaj `ROADMAP.md`
3. przeczytaj `CONTEXT.md`
4. uruchom `git status --short`
5. jeśli app code jest dirty, najpierw domknij/reviewuj istniejący diff

Podczas pracy:

- preferuj język polski w wyjaśnieniach
- trzymaj slices małe, najlepiej do zamknięcia w jednej sesji
- przy niskiej energii użytkownika wybieraj zadania copy/content/review zamiast ciężkich refactorów
- oddzielaj: Go mechanism vs standard library mechanism vs project decision
- nie kopiuj `../llm-generated/` 1:1
- nie dodawaj architektury bez aktualnej potrzeby treści

Przed uznaniem slice za zakończony:

- wypisz zmienione pliki
- wypisz komendy, które przeszły
- powiedz, czego nie sprawdzono
- zrób krótki understanding check, jeśli slice uczył ważnego mechanizmu
- zaproponuj commit message
