<script>
  const steps = [
    {
      number: 1,
      title: 'Find Your Device',
      description: 'Enter your device model or select from our comprehensive database of Android devices.'
    },
    {
      number: 2,
      title: 'Choose Your ROM',
      description: 'Browse through verified custom ROMs, kernels, and mods specifically compatible with your device.'
    },
    {
      number: 3,
      title: 'Download Resources',
      description: 'Get all necessary files in one place - ROMs, recovery images, drivers, and tools.'
    },
    {
      number: 4,
      title: 'Follow Guide',
      description: 'Use our detailed step-by-step guide customized for your specific device model.'
    }
  ];
</script>

<section class="how-it-works" id="guide">
  <div class="container">
    <div class="section-title">
      <h2>How It Works</h2>
      <p>Our platform simplifies the Android customization process down to a few easy steps</p>
    </div>
    <div class="steps">
      {#each steps as step}
        <div class="step">
          <div class="step-number">{step.number}</div>
          <h3>{step.title}</h3>
          <p>{step.description}</p>
        </div>
      {/each}
    </div>
  </div>
</section>

<style>
  .how-it-works {
    padding: 5rem 0;
    background-color: var(--light-color);
  }

  .steps {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
    margin-top: 3rem;
  }

  .step {
    text-align: center;
    position: relative;
  }

  .step-number {
    background-color: var(--primary-color);
    color: white;
    width: 50px;
    height: 50px;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 1.5rem;
    font-weight: bold;
    margin: 0 auto 1.5rem;
  }

  .step h3 {
    margin-bottom: 1rem;
  }

  .step p {
    color: var(--gray-color);
  }
</style>
