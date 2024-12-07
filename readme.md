# **Hawkwing 🦅**

**A Lightweight and Minimalist Web Framework for Go**

Hawkwing is a blazing-fast, minimalist web framework for developers who value simplicity, performance, and control. Inspired by framework like Flask, it focuses on providing essential tools for building modern web applications without unnecessary bloat.

---

## 🚀 **Why Hawkwing?**

- **Minimalist:** Focused on core features you need, nothing more.
- **Efficient:** Built with Go's native power and performance.
- **Developer-Friendly:** Easy-to-use APIs, hot-reloading for templates and static files.
- **Customizable:** Middleware and routing designed to suit any web application.

---

## ✨ **Features**

- **Routing with Dynamic Parameters**  
  Add clean and expressive routes with support for dynamic parameters.
- **Middleware Support**  
  Compose reusable middleware for logging, authentication, and more.
- **Template Rendering**  
  Render HTML templates seamlessly using `html/template`, with hot-reloading for real-time development.
- **Static File Serving**  
  Serve static assets like CSS, JS, and images with automatic reload on changes.
- **Modular and Simple API**  
  Modular design lets you build only what you need.

---

## 🛠️ **Getting Started**

### **Installation**

```bash
go get github.com/aliqyan-21/hawkwing
```

## Example Usage

Here’s a quick demo to show how easy it is to get started:

```go
package main

import (
	"github.com/aliqyan-21/hawkwing"
)

func main() {
	app := hawking.Init()

	// Load templates and static files
	hawking.LoadTemplates("templates")
	app.LoadStatic("/static", "static")

	// Add routes
	app.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		hawking.RenderHTML(w, "index.html", map[string]interface{}{
			"title": "Welcome to Hawkwing!",
		})
	})

	// Start the server
	hawking.Start(":5000", r)
}
```

📂 Features in Detail

1. Routing Define routes with dynamic parameters and middleware.
   Example:

   ```go
   r.AddRoute("GET", "/user/:id", userHandler)
   ```

   Dynamic parameters like :id are automatically extracted and passed via context.
   Example:

   ```go
   r.AddRoute("GET", "/users/:name", func(w http.ResponseWriter, r *http.Request) {
   	params := r.GetRouteParams(r)
   	name := params["name"]
   	fmt.Fprintf(w, "Hello, %s!", name)
   })
   ```

2. Template Rendering
   Hawkwing simplifies rendering templates while supporting hot-reloading for development.

   Example:

   > you can store your templates in any folder anywehere just mention the path in the function LoadTemplates(path)

   ```go
   hawking.LoadTemplates("templates")
   hawking.RenderHTML(w, "index.html", data)
   ```

3. Middleware

Use built-in middleware or create your own for logging, authentication, and more.

Example:

```go
r.AddRoute("GET", "/secure", secureHandler, authMiddleware)
```

4. Hot-Reloading
   Automatically reload templates and static files when changes are detected. Thus no need to restart the server

## 🌍 Community and Contributions

Hawkwing is open-source and welcomes contributions! To get involved:

- Browse open issues or submit new ones.
- Share your feedback and ideas.

📜 License
Hawkwing is licensed under the Apache License. See LICENSE for details.

🌟 Acknowledgements

- Built with Go's simplicity and power.
- Inspired by minimalist frameworks like Flask in python.

> 🦅 Soar to New Heights with Hawkwing!
