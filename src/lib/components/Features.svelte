<script>
  const features = [
    {
      icon: 'fas fa-database',
      title: 'Comprehensive Database',
      description: 'All files organized by device model, manufacturer, and codename so you never have to guess which file is right for your device.'
    },
    {
      icon: 'fas fa-book',
      title: 'Step-by-Step Guides',
      description: 'Clear instructions tailored to your specific device model, making the rooting and flashing process straightforward.'
    },
    {
      icon: 'fas fa-download',
      title: 'Fast Downloads',
      description: 'Direct downloads for all firmware files, tools, and utilities from our high-speed servers - no more hunting across forums.'
    },
    {
      icon: 'fas fa-users',
      title: 'Active Community',
      description: 'Get help from our community of Android enthusiasts and experts for any issues you encounter.'
    },
    {
      icon: 'fas fa-tools',
      title: 'Troubleshooting Tools',
      description: 'Built-in tools to diagnose and fix common issues that occur during the customization process.'
    },
    {
      icon: 'fas fa-bell',
      title: 'Update Notifications',
      description: 'Get notified when new ROMs, kernels, or tools become available for your specific device.'
    }
  ];
</script>

<section class="features" id="features">
  <div class="container">
    <div class="section-title">
      <h2>Why Choose AndroidCustomizer</h2>
      <p>We streamline the Android customization process, saving you days of research and troubleshooting</p>
    </div>
    <div class="features-grid">
      {#each features as feature}
        <div class="feature-card">
          <div class="feature-icon">
            <i class={feature.icon}></i>
          </div>
          <h3>{feature.title}</h3>
          <p>{feature.description}</p>
        </div>
      {/each}
    </div>
  </div>
</section>

<style>
  .features {
    padding: 5rem 0;
    background-color: white;
  }

  .features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
  }

  .feature-card {
    background-color: var(--light-color);
    padding: 2rem;
    border-radius: 10px;
    box-shadow: 0 5px 15px rgba(0,0,0,0.1);
    transition: transform 0.3s;
  }

  .feature-card:hover {
    transform: translateY(-10px);
  }

  .feature-icon {
    font-size: 2.5rem;
    color: var(--primary-color);
    margin-bottom: 1.5rem;
  }

  .feature-card h3 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
  }

  .feature-card p {
    color: var(--gray-color);
  }
</style>
