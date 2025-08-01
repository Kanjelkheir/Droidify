<header>
  <div class="container header-container">
    <div class="logo" role="banner">
      <i class="fas fa-mobile-alt" aria-hidden="true"></i>
      <span>AndroidCustomizer</span>
    </div>

    <nav role="navigation" aria-label="Primary Navigation">
      <ul>
        <li><a href="https://droididfy.netlify.app/Home">Home</a></li>
        <li><a href="https://droididfy.netlify.app/Features">Features</a></li>
        <li><a href="https://droididfy.netlify.app/Devices">Devices</a></li>
        <li><a href="https://droididfy.netlify.app/Guides">Guides</a></li>
        <li><a href="https://droididfy.netlify.app/Community">Community</a></li>
        <li><a href="https://droididfy.netlify.app/Contact">Contact</a></li>
      </ul>
    </nav>
  </div>
</header>

<style>
  header {
    background-color: var(--dark-color);
    color: white;
    padding: 1rem 0;
    position: sticky;
    top: 0;
    z-index: 100;
    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  }

  .header-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .logo {
    font-size: 1.5rem;
    font-weight: bold;
    display: flex;
    align-items: center;
  }

  .logo i {
    color: var(--accent-color);
    margin-right: 0.5rem;
  }

  nav ul {
    display: flex;
    list-style: none;
    margin: 0;
    padding: 0;
  }

  nav ul li {
    margin-left: 1.5rem;
  }

  nav ul li:first-child {
    margin-left: 0;
  }

  nav ul li a {
    color: white;
    text-decoration: none;
    transition: color 0.3s;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    min-width: 48px;
    min-height: 48px;
    padding: 0 0.5rem;
    font-size: 1rem;
    user-select: none;
  }

  nav ul li a:hover,
  nav ul li a:focus {
    color: var(--accent-color);
    outline: none;
  }

</style>
