# Eesti-inglise homograafilised virvasõnad

```sh
$ go run . --help
```

Jooksutamise järjekord:
1. `ekipull`
1. `ekiprocess`
1. `ppap`
1. `fm`
1. `fs`

Väljundid: `data/`:
 - `compare.txt` võrdlus eelmise tööga
 - `matches/`
   - `0.complog` koondlogi
   - `comp_s0.json` koondväljund
   - `comp_s4_paradigms/` käänete kaupa filtreerimisjärgsed sõnad, millest on käsitsi mõned läbi käidud
   - `human_ok.json` leiud
