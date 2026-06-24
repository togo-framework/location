<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/location</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/location"><img src="https://pkg.go.dev/badge/github.com/togo-framework/location.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/location
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
