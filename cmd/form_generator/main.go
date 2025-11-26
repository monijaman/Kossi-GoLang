package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Specification struct {
	ID    int
	Key   string
	Value string
}

func main() {
	specifications := []Specification{
		{ID: 1, Key: "Fuel Efficiency (km/l)", Value: ""},
		{ID: 2, Key: "Engine Displacement (cc)", Value: ""},
		{ID: 3, Key: "Max Power (hp)", Value: ""},
		// Add more specifications as needed
	}

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			for _, spec := range specifications {
				spec.Value = r.FormValue(spec.Key)
			}
			fmt.Fprintf(w, "Form submitted successfully!\n")
			for _, spec := range specifications {
				fmt.Fprintf(w, "%s: %s\n", spec.Key, spec.Value)
			}
			return
		}

		tmpl := template.Must(template.New("form").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Car Specifications</title>
			</head>
			<body>
				<h1>Car Specifications Form</h1>
				<form method="POST">
					{{range .}}
						<div>
							<label for="{{.Key}}">{{.Key}}</label>
							<input type="text" id="{{.Key}}" name="{{.Key}}" value="{{.Value}}">
						</div>
					{{end}}
					<button type="submit">Submit</button>
				</form>
			</body>
			</html>
		`))

		tmpl.Execute(w, specifications)
	})

	http.ListenAndServe(":8080", nil)
}
