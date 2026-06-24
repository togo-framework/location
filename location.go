// Package location is a togo plugin: a localization dataset — countries (ISO
// 3166-1 + currency + dial code + a representative timezone), languages (ISO
// 639-1), and IANA timezones — with lookup helpers and a REST API. An optional
// cities/areas dataset can be loaded at runtime via LoadCities.
package location

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/togo-framework/togo"
)

// Country is an ISO 3166-1 country with currency, calling code and a default tz.
type Country struct {
	Code     string `json:"code"`     // ISO 3166-1 alpha-2
	Name     string `json:"name"`
	Currency string `json:"currency"` // ISO 4217
	Dial     string `json:"dial"`     // E.164 calling code
	Region   string `json:"region"`
	Timezone string `json:"timezone"` // representative IANA zone
}

// Language is an ISO 639-1 language.
type Language struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Native string `json:"native"`
}

// City is a place inside a country (loaded via LoadCities).
type City struct {
	Name     string  `json:"name"`
	Country  string  `json:"country"`
	Region   string  `json:"region"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

var (
	citiesMu sync.RWMutex
	cities   = map[string][]City{}
)

// LoadCities registers a cities dataset (country code → cities). Call from an
// app or a richer location-data plugin to extend the built-in dataset.
func LoadCities(data map[string][]City) {
	citiesMu.Lock()
	defer citiesMu.Unlock()
	for c, list := range data {
		cities[strings.ToUpper(c)] = append(cities[strings.ToUpper(c)], list...)
	}
}

func init() {
	togo.RegisterProviderFunc("location", togo.PriorityService, func(k *togo.Kernel) error {
		k.Set("location", struct{}{})
		mount(k.Router)
		return nil
	})
}

// Countries returns all known countries (sorted by name).
func Countries() []Country { return countries }

// Country looks up a country by ISO 3166-1 alpha-2 code (case-insensitive).
func CountryByCode(code string) (Country, bool) {
	code = strings.ToUpper(code)
	for _, c := range countries {
		if c.Code == code {
			return c, true
		}
	}
	return Country{}, false
}

// Languages returns the ISO 639-1 language list.
func Languages() []Language { return languages }

// Timezones returns the IANA timezone list.
func Timezones() []string { return timezones }

// CitiesByCountry returns cities for a country code (empty unless LoadCities ran).
func CitiesByCountry(code string) []City {
	citiesMu.RLock()
	defer citiesMu.RUnlock()
	return cities[strings.ToUpper(code)]
}

func mount(r chi.Router) {
	r.Route("/api/location", func(r chi.Router) {
		r.Get("/countries", func(w http.ResponseWriter, req *http.Request) { writeJSON(w, countries) })
		r.Get("/countries/{code}", func(w http.ResponseWriter, req *http.Request) {
			if c, ok := CountryByCode(chi.URLParam(req, "code")); ok {
				writeJSON(w, c)
			} else {
				http.Error(w, "country not found", http.StatusNotFound)
			}
		})
		r.Get("/countries/{code}/cities", func(w http.ResponseWriter, req *http.Request) {
			writeJSON(w, CitiesByCountry(chi.URLParam(req, "code")))
		})
		r.Get("/languages", func(w http.ResponseWriter, req *http.Request) { writeJSON(w, languages) })
		r.Get("/timezones", func(w http.ResponseWriter, req *http.Request) { writeJSON(w, timezones) })
	})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func init() { sort.Slice(countries, func(i, j int) bool { return countries[i].Name < countries[j].Name }) }

var _ = context.Background
