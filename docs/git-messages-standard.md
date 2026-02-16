# Hvorfor standardisere?
For å opprettholde oversikten over tid ved hjelp av enkel kategorisering av commit-meldinger. 

# Commit-typer
| type | beskrivelse |
|------|------------|
| `feat` | Ny feature du legger til i applikasjonen |
| `fix` | En bugfix |
| `style` | Feature og oppdateringer relatert til styling |
| `refactor` | Refaktorering av en spesifikk seksjon av kodebasen |
| `test` | Alt relatert til testing |
| `docs` | Alt relatert til dokumentasjon |
| `chore` | Vedlikehold av kode |

Om en av disse ikke passer til din spesifikke commit, velg den som passer best. 

Eksempel: `git commit -m "feat: logg inn funksjon" -m "Logger bruker inn ved hjelp av autentisering av bruker i database for så å videresende brukeren til hjemmesiden for bruker."`

eller via *nano*:
```md
feat: logg inn funksjon

Logger bruker inn ved hjelp av autentisering av bruker i database for så å videresende brukeren til hjemmesiden for bruker.
```

# Metoder
Nedenfor er metoder for hvordan man kan strukturere enkle commit-meldinger, men legge til en lengre beskrivelse om ønskelig.
## Editor metoden
Konfigurer standard editor med kommando:
`git config --global core.editor nano`

Da vil nano-editoren bli commit-melding editoren din med kommando: `git commit`

Format: (i dette tilfellet i nano):
```md
<Subject>

<Description>
```
**NB!!** Det er veldig viktig at det er en *newline* mellom subjekt og description.

## CLI-metode
`git commit -m "<Subject>" -m "<Description>"`

---
Mer detaljert info finnes [her](https://www.freecodecamp.org/news/writing-good-commit-messages-a-practical-guide/).