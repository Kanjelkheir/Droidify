<script>
  import Header from '$lib/components/Header.svelte';
  import Hero from '$lib/components/Hero.svelte';
  import Features from '$lib/components/Features.svelte';
  import HowItWorks from '$lib/components/HowItWorks.svelte';
  import PopularDevices from '$lib/components/PopularDevices.svelte';
  import CTA from '$lib/components/CTA.svelte';
  import Footer from '$lib/components/Footer.svelte';
</script>

<svelte:head>
  <title>AndroidCustomizer - One-Stop Solution for Android Customization</title>
  <link
    rel="stylesheet"
    href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css"
  />
  <link
    href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;700&display=swap"
    rel="stylesheet"
  />
</svelte:head>

<div class="app">
  <Header />
  <main>
    <Hero />
    <Features />
    <HowItWorks />
    <PopularDevices />
    <CTA />
  </main>
  <Footer />
</div>

<style>
  :global(:root) {
    --primary-color: #4285f4;
    --secondary-color: #34a853;
    --accent-color: #fbbc05;
    --accent-color-hover: #e0a800;
    --danger-color: #ea4335;
    --dark-color: #333;
    --light-color: #f8f9fa;
    --gray-color: #6c757d;
  }

  :global(*) {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: "Roboto", sans-serif;
  }

  :global(body) {
    background-color: var(--light-color);
    color: var(--dark-color);
    width: 100%;
    overflow-x: hidden;
  }

  :global(.container) {
    width: 90%;
    max-width: 1200px;
    margin: 0 auto;
  }

  .app {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    width: 100%;
  }

  main {
    display: flex;
    flex-direction: column;
    flex: 1;
  }
</style>
