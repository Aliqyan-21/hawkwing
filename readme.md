# **Hawkwing 🦅**

**A Lightweight and Minimalist Web Framework for Go**

Hawkwing is a blazing-fast, minimalist web framework for developers who value simplicity, performance, and control. Inspired by framework like Flask, it focuses on providing essential tools for building modern web applications without unnecessary bloat.

---

## 🚀 **What Makes Hawkwing Different?**

- **Minimalist:** Focused on core features you need, nothing more.
- **Blazing Fast:** Leverages Go's native power for optimal performance..
- **Elegant API:** Easy-to-use and intuitive for a smooth development experience.
- **Highly Customizable:** Middleware and routing designed for maximum flexibility.

---

## ✨ **Key Features**

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

This snippet demonstrates how easy it is to get started with Hawkwing:

```go
package main

import (
    "github.com/aliqyan-21/hawkwing"
)

func main() {
    app := hawkwing.Init()

    // Load templates and static files
    hawkwing.LoadTemplates("templates")
    app.LoadStatic("/static", "static")

    // Define routes
    app.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
        hawkwing.RenderHTML(w, "index.html", map[string]interface{}{
            "title": "Welcome to Hawkwing!",
        })
    })

    // Start server (localhost only)
    hawkwing.Start("localhost", "8080")

    // Publicly accessible server
    // hawkwing.Start("0.0.0.0", "8080")
}
```

📂 Features in Detail

1. Routing:
   Define routes that accept dynamic parameters and utilize middleware for enhanced functionality.

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

2. Template Rendering:
   Hawkwing simplifies template rendering and includes hot-reloading for efficient development.

   Example:

   > Store your templates anywhere and simply reference the path during initialization:

   ```go
   hawking.LoadTemplates("templates")
   hawking.RenderHTML(w, "index.html", data)
   ```

3. Middleware:
   Leverage built-in middleware or create custom solutions for tasks like logging or authentication:

   Example:

   ```go
   r.AddRoute("GET", "/secure", secureHandler, AuthMiddleware)
   ```

4. Hot-Reloading:
   Experience automatic reloading for templates and static files whenever changes are made, eliminating the need for server restarts.

## Catty Cuteness Voting

> A WebApp made with hawkwing

##### [CattyCuteness](https://github.com/aliqyan-21/cattycuteness)

## Future Implementations

I am actively working on enhancing Hawkwing with exciting new features:

- **CLI Tool:** A powerful command-line interface to scaffold projects, generate routes, and manage templates.
- **Database Integration:** Seamless integration with databases like SQLite for data persistence.
- **Customizable Middleware Pipeline:** Globally register middleware (e.g., logging, CORS) that applies to all routes, simplifying configuration.
- **And more!**

## 🌍 Community and Contributions

Hawkwing is an open-source project and welcomes your contributions! Here's how you can get involved:

- Explore open issues or create new ones to report bugs or request features.
- Share your feedback and suggest improvements for a better framework.

📜 License
Hawkwing is distributed under the Apache License. Refer to the LICENSE file for details.

🌟 Acknowledgements

- Built upon the simplicity and power of Go.
- Inspired by the minimalist approach of frameworks like Flask.

> 🦅 Soar to New Heights with Hawkwing!
