# location — localization dataset for togo

A togo plugin that ships a **localization dataset** — countries (ISO 3166-1 +
currency + dial code + a representative timezone), languages (ISO 639-1), and IANA
timezones — with lookup helpers and a REST API. An optional cities/areas dataset is
loaded at runtime via `location.LoadCities`.

```bash
togo install togo-framework/location
```

## REST API

| Method | Path |
|---|---|
| `GET` | `/api/location/countries` |
| `GET` | `/api/location/countries/{code}` |
| `GET` | `/api/location/countries/{code}/cities` |
| `GET` | `/api/location/languages` |
| `GET` | `/api/location/timezones` |

## Go API

```go
all := location.Countries()
eg, _ := location.CountryByCode("EG")      // {Code:EG, Currency:EGP, Dial:+20, Timezone:Africa/Cairo}
langs := location.Languages()
zones := location.Timezones()
// extend with your own cities/areas dataset:
location.LoadCities(map[string][]location.City{"EG": {{Name:"Cairo", Lat:30.04, Lng:31.23}}})
cairo := location.CitiesByCountry("EG")
```

The built-in set covers the common countries/languages/timezones and is extensible —
append in a fork or load a richer cities/areas dataset via `LoadCities`.

MIT
