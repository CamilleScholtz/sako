package main

import (
	"regexp"
)

// TODO: Possibly use an universal model.
type settingsModel struct {
	Template string
	Sidebar  Sidebar
	Config   Config
}

/*func settings(w http.ResponseWriter, r *http.Request) {
	// Handle POST requests.
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		d, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		f, err := ioutil.ReadFile(path.Join(d, "runtime", "config.toml"))
		if err != nil {
			log.Fatal(err)
		}

		f = replaceConfig(f, "currency", r.FormValue("currency"))

		if err := ioutil.WriteFile(path.Join(d, "runtime", "config.toml"), f,
			0644); err != nil {
			log.Fatal(err)
		}

		// Re-parse the config file.
		if err := parseConfig(); err != nil {
			log.Fatal(err)
		}
	}

	sb, err := sidebar()
	if err != nil {
		log.Print(err)
	}

	c, err := cryptoCompare()
	if err != nil {
		log.Print(err)
	}

	model := settingsModel{"settings", sb, c, config}

	t, err := template.ParseFiles(
		"static/html/layout.html",
		"static/html/head.html",
		"static/html/sidebar.html",
		"static/html/settings.html",
		"static/html/settings.js",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}*/

func replaceConfig(f []byte, k, v string) []byte {
	re := regexp.MustCompile(
		"(?m)(^[[:space:]]*" + k +
			"[[:space:]]*=[[:space:]]*\")[[:upper:]]*(\".*)")
	f = re.ReplaceAll(f, []byte("${1}"+v+"${2}"))

	return f
}
