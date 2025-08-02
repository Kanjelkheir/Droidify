<script>
  const footerColumns = [
    {
      title: "Resources",
      links: [
        { text: "Custom Rom", url: "https://droidify.netlify.app/Custom_Rom" },
        { text: "Recovery Images", url: "https://droidify.netlify.app/Recovery_Images" },
        { text: "Kernels", url: "https://droidify.netlify.app/Kernels" },
        { text: "Root Tools", url: "https://droidify.netlify.app/Root_Tools" },
        { text: "Mods & Tweaks", url: "https://droidify.netlify.app/Mods_and_Tweaks" }
      ],
    },
    {
      title: "Guides",
      links: [
        { text: "Rooting Guide", url: "https://droidify.netlify.app/Rooting_Guide" },
        { text: "Custom ROM Installation", url: "https://droidify.netlify.app/Custom_ROM_Installation" },
        { text: "Backup & Restore", url: "https://droidify.netlify.app/Backup_and_Restore" },
        { text: "Troubleshooting", url: "https://droidify.netlify.app/Troubleshooting" },
        { text: "FAQ", url: "https://droidify.netlify.app/FAQ" }
      ],
    },
    {
      title: "Company",
      links: [
        { text: "About Us", url: "https://droidify.netlify.app/About_Us" },
        { text: "Our Team", url: "https://droidify.netlify.app/Our_Team" },
        { text: "Privacy Policy", url: "https://droidify.netlify.app/Privacy_Policy" },
        { text: "Terms of Service", url: "https://droidify.netlify.app/Terms_of_Service" },
        { text: "Contact Us", url: "https://droidify.netlify.app/Contact_Us" }
      ],
    },
  ];

  const socialLinks = [
    { icon: "fab fa-facebook-f", url: "https://www.facebook.com/eliekh05", label: "Facebook" },
    { icon: "fab fa-x-twitter", url: "https://x.com/eliekh05", label: "Twitter" },
    { icon: "fab fa-github", url: "https://github.com/eliekh05/droidify", label: "GitHub" },
    { icon: "fab fa-telegram", url: "https://t.me/eliekh05", label: "Telegram" },
    { icon: "fab fa-whatsapp", url: "https://wa.me/+96179140038", label: "WhatsApp" }
  ];
</script>

<footer id="contact" role="contentinfo" aria-label="Footer">
  <div class="container">
    <div class="footer-content">
      <div class="footer-column">
        <h3 id="footer-brand">AndroidCustomizer</h3>
        <p>
          Making Android customization accessible to everyone by
          eliminating the confusion and wasted time.
        </p>
        <div class="social-links">
          {#each socialLinks as social}
            <a href={social.url} aria-label={social.label} target="_blank" rel="noopener noreferrer" class="social-link">
              <i class={social.icon} aria-hidden="true"></i>
            </a>
          {/each}
        </div>
      </div>

      {#each footerColumns as column}
        <div class="footer-column" role="region" aria-labelledby={"footer-" + column.title.toLowerCase().replace(/\s+/g, '-')}>
          <h3 id={"footer-" + column.title.toLowerCase().replace(/\s+/g, '-')}>{column.title}</h3>
          <ul>
            {#each column.links as link}
              <li>
                <a href={link.url} class="footer-link" target="_blank" rel="noopener noreferrer">{link.text}</a>
              </li>
            {/each}
          </ul>
        </div>
      {/each}
    </div>

    <div class="footer-bottom">
      <p>&copy; {new Date().getFullYear()} AndroidCustomizer. All rights reserved.</p>
    </div>
  </div>
</footer>

<style>
  footer {
    background-color: var(--dark-color);
    color: white;
    padding: 3rem 0 1.5rem;
  }

  .footer-content {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 2rem;
    margin-bottom: 2rem;
  }

  .footer-column h3 {
    margin-bottom: 1.5rem;
    position: relative;
  }

  .footer-column h3::after {
    content: "";
    position: absolute;
    left: 0;
    bottom: -0.5rem;
    width: 50px;
    height: 2px;
    background-color: var(--accent-color);
  }

  .footer-column ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  .footer-column ul li {
    margin-bottom: 0.8rem;
  }

  .footer-link {
    color: #ccc;
    text-decoration: none;
    transition: color 0.3s;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    min-width: 48px;
    min-height: 48px;
    padding: 0 0.5rem;
    user-select: none;
  }

  .footer-link:hover,
  .footer-link:focus {
    color: white;
    outline: none;
  }

  .social-links {
    display: flex;
    gap: 1rem;
    margin-top: 1rem;
  }

  .social-link {
    color: white;
    background-color: rgba(255, 255, 255, 0.1);
    width: 48px;
    height: 48px;
    border-radius: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    transition: background-color 0.3s;
  }

  .social-link:hover,
  .social-link:focus {
    background-color: var(--primary-color);
    outline: none;
  }

  .footer-bottom {
    text-align: center;
    padding-top: 2rem;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
  }
</style>
