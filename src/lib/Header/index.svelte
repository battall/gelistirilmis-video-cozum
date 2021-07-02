<script>
  import { session } from "$app/stores";
</script>

<header>
  <div class="corner">
    <a href="/">
      <img
        src="/api/proxy?publisher={$session.publisher}&path=/logo.png"
        alt={$session.publisher}
      />
    </a>
  </div>

  <nav>
    <ul class="breadcrumb">
      {#each $session.navPath as path, i}
        {#if $session.navPath.length - 1 === i}
          <li>{path.title}</li>
        {:else}
          <li><a>{path.title}</a></li>
        {/if}
      {/each}
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

  .breadcrumb {
    margin: 0;
    padding: 2em;
    list-style-type: none;
    color: #333;
  }
  .breadcrumb li {
    display: inline-block;
    position: relative;
    overflow: hidden;
    margin: 0;
    padding-right: 2em;
  }
  .breadcrumb li:after {
    content: ">";
    position: absolute;
    display: inline-block;
    right: 0;
    width: 2em;
    text-align: center;

    background: white;
    background: linear-gradient(
      90deg,
      rgba(255, 255, 255, 0.4) 0%,
      rgba(255, 255, 255, 1) 35%
    );
    padding-left: 1em;
  }
  .breadcrumb li:last-child {
    font-weight: bold;
  }
  .breadcrumb li:last-child:after {
    content: "";
  }

  .breadcrumb a {
    text-decoration: none;
    display: inline-block;
    color: #333;
    white-space: nowrap;
    cursor: pointer;

    max-width: 2em;
    transition: max-width 500ms ease-in-out;
  }
  .breadcrumb a:hover {
    color: #c41c87;
    text-decoration: underline;
    outline: none;
  }
  .breadcrumb a:hover,
  .breadcrumb a:focus,
  .breadcrumb li:hover a {
    max-width: 1000px;
  }

  .breadcrumb li:hover:after {
    padding-left: 0em;
    background: transparent;
  }
</style>
