<script>
</script>

<header>
  <div class="corner">
    <a href="/">
      <img src="/api/proxy?publisher=cap&path=/logo.png" alt="logo" />
    </a>
  </div>

  <nav>
    <ul />
  </nav>

  <div class="corner">
    <!-- TODO put something else here? github link? -->
  </div>
</header>

<style>
  header {
    display: flex;
    justify-content: space-between;
    align-items: center;

    box-sizing: border-box;

    width: 100%;
    height: 40pt;

    border-bottom: 1px solid var(--accents-3);
  }

  header .corner {
    width: 3em;
    height: 3em;
  }

  header .corner a {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;
  }

  header .corner img {
    width: 2em;
    height: 2em;
    object-fit: contain;
  }

  header nav {
    display: flex;
    justify-content: center;
    --background: rgba(255, 255, 255, 0.7);
  }

  nav a {
    display: flex;
    height: 100%;
    align-items: center;
    padding: 0 1em;
    color: var(--heading-color);
    font-weight: 700;
    font-size: 0.8rem;
    text-transform: uppercase;
    letter-spacing: 10%;
    text-decoration: none;
    transition: color 0.2s linear;
  }

  a:hover {
    color: var(--accent-color);
  }
</style>
