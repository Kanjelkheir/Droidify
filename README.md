# AndroidCustomizer Svelte App

This is a Svelte implementation of the AndroidCustomizer website, a platform for Android device customization resources.

## Description

AndroidCustomizer is a one-stop solution for Android customization. The platform provides all the resources, tools, and guides users need for any Android device - ready to use in minutes, not days.

## Features

- Comprehensive database of Android ROMs, kernels, and mods
- Device-specific guides for rooting and customization
- Fast direct downloads for firmware files and tools
- Active community support
- Troubleshooting tools and resources
- Update notifications for new ROMs and tools

## Netlify Status

[![Netlify Status](https://api.netlify.com/api/v1/badges/955f13f5-90a3-4b8e-8d43-e47b1ed1e3d6/deploy-status)](https://app.netlify.com/projects/droidify/deploys)

## Getting Started

### Prerequisites

- Node.js (version 14 or higher recommended)
- npm or yarn package manager

### Installation

1. Clone this repository
   ```bash
   git clone https://github.com/eliekh05/android-customizer.git
   cd android-customizer
   ```

2. Install dependencies
   ```bash
   npm install
   # or
   yarn install
   ```

3. Start the development server
   ```bash
   npm run dev
   # or
   yarn dev
   ```

4. Build for production
   ```bash
   npm run build
   # or
   yarn build
   ```

## Technologies Used

- [Svelte](https://svelte.dev/) - Frontend framework
- [Rollup](https://rollupjs.org/) - Module bundler
- [FontAwesome](https://fontawesome.com/) - Icon library
- [Google Fonts](https://fonts.google.com/) - Web fonts

## Converting from HTML to Svelte

This project is a conversion of a static HTML website to a component-based Svelte application. The conversion process involved:

1. Breaking down the HTML into reusable components
2. Converting inline JavaScript to Svelte reactive declarations and event handlers
3. Scoping CSS styles to their respective components
4. Implementing data flow between components

## License

This project is licensed under the MIT License - see the LICENSE file for details.
