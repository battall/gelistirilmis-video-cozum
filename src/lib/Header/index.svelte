<script>
</script>

<header>
  <div class="corner">
    <a href="/">
      <img src="/api/proxy?publisher=cap&path=/logo.png" alt="logo" />
    </a>
  </div>

  <nav>
    <ul class="breadcrumb">
      <li><a href="/">Main</a> →</li>
      <li><a href="/products/">Products</a> →</li>
      <li><a href="/products/dishwashers/">Dishwashers</a> →</li>
      <li><a>Second hand</a></li>
    </ul>
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

  header nav .breadcrumb {
    display: inline-flex;
    list-style: none;
    font-size: initial;
    color: #270215;
    padding: 0;
    margin: 0;
  }

  header nav .breadcrumb li {
    padding: 0;
    margin: 0;
  }

  header nav .breadcrumb > li + li {
    margin-left: 5px;
  }

  header nav .breadcrumb > li a[href] {
    text-decoration: none;
    position: relative;
    font-weight: 400;
    color: gray;
    z-index: 0;
  }

  header nav .breadcrumb > li a[href]:hover,
  header nav .breadcrumb > li a[href]:focus {
    color: #c41c87;
    outline: none;
  }

  header nav .breadcrumb > li a:not([href]) {
    font-weight: bold;
    color: #2c363f;
  }
</style>
