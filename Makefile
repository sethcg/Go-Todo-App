run:
	@npx tailwindcss -i ./css/input.css -o ./css/output.css
	@templ generate
	@go mod tidy
	@go run main.go