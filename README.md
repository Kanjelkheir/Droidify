Droidify
React

📝 Description

Droidify is your one-stop platform for unleashing the full potential of your Android device! Tired of the limitations imposed by your stock ROM? Dive into a world of custom ROMs, recovery images, and essential tools, all meticulously organized and readily available for countless devices. We've streamlined the often-complex process of Android customization, providing device-specific guides and direct downloads to empower you to transform your Android experience in minutes, not days. Droidify's user-friendly React-based interface ensures a seamless and intuitive experience as you explore a vast database of resources to breathe new life into your device.

✨ Features

🗄️ Database
🛠️ Tech Stack

⚛️ React
📦 Key Dependencies

react: ^19.1.1
react-dom: ^19.1.1
🚀 Run Commands

dev: npm run dev
build: npm run build
lint: npm run lint
preview: npm run preview
Run: go run .
Build: go build
📁 Project Structure

.
├── backend
│   ├── Dockerfile
│   ├── droidify
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── agent
│   │   │   ├── agent.go
│   │   │   └── agent_test.go
│   │   └── database
│   │       ├── db.go
│   │       ├── devices.sql.go
│   │       └── models.go
│   ├── main.go
│   ├── sql
│   │   ├── queries
│   │   │   └── devices.sql
│   │   └── schema
│   │       ├── 001_devices.sql
│   │       ├── 002_remove_id.sql
│   │       └── 003_make_model_primary_key.sql
│   └── sqlc.yaml
├── bun.lock
├── eslint.config.js
├── index.html
├── package.json
├── public
│   └── favicon.svg
├── src
│   ├── App.css
│   ├── App.jsx
│   ├── app.css
│   ├── app.d.ts
│   ├── app.html
│   ├── assets
│   │   └── favicon.svg
│   ├── components
│   │   ├── CTA.css
│   │   ├── CTA.jsx
│   │   ├── Features.css
│   │   ├── Features.jsx
│   │   ├── Footer.css
│   │   ├── Footer.jsx
│   │   ├── Header.css
│   │   ├── Header.jsx
│   │   ├── Hero.css
│   │   ├── Hero.jsx
│   │   ├── HowItWorks.css
│   │   ├── HowItWorks.jsx
│   │   ├── PopularDevices.css
│   │   └── PopularDevices.jsx
│   ├── index.css
│   └── main.jsx
├── static
│   └── favicon.ico
├── vite.config.js
└── vite.config.ts
🛠️ Development Setup

Node.js/JavaScript Setup

Install Node.js (v18+ recommended)
Install dependencies: npm install or yarn install
Start development server: (Check scripts in package.json, e.g., npm run dev)
👥 Contributing

Contributions are welcome! Here's how you can help:

Fork the repository
Clone your fork: git clone https://github.com/eliekh05/Droidify/.git
Create a new branch: git checkout -b feature/your-feature
Commit your changes: git commit -am 'Add some feature'
Push to your branch: git push origin feature/your-feature
Open a pull request
Please ensure your code follows the project's style guidelines and includes tests where applicable.
