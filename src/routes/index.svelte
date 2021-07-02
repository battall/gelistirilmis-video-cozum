<script>
  import { onMount } from "svelte";
  import { session } from "$app/stores";
  import Player from "$lib/Player/index.svelte";

  onMount(() => {
    console.log("INDEX MOUNT", $session);
    if ($session.navPath.length === 0)
      updateNav([{ _id: 2, title: "Kitaplar", is_parent: true }]);
  });

  let updateNav = (navPath) => {
    $session.navPath = navPath;

    let path = navPath[navPath.length - 1];
    fetch(
      `/api/${path.is_parent ? "sections" : "solutions"}?publisher=${
        $session.publisher
      }&parent_id=${path._id}`
    )
      .then((res) => res.json())
      .then((res) => ($session.navItems = res));
  };
</script>

<svelte:head>
  <title>Video Çözüm</title>
</svelte:head>

<div class="index">
  <nav>
    <ul class="book-sections">
      {#each $session.navItems as item}
        <li
          class:active={item.active}
          on:click={() =>
            item.solution_id
              ? ($session.solution = {
                  swf: `/api/proxy?publisher=${$session.publisher}&path=${item.swf}`,
                })
              : updateNav([...$session.navPath, item])}
        >
          {item.title}
        </li>
      {/each}
    </ul>
  </nav>

  <main>
    <Player solution={$session.solution} />
  </main>
</div>

<style>
  .index {
    display: flex;
  }

  nav {
    display: flex;
    flex-direction: column;

    width: 24%;
    min-height: calc(100vh - 40pt);

    border-right: 1px solid var(--accents-2);
  }

  nav .book-sections {
    width: 100%;
    max-height: calc(100vh - 40pt);

    margin: 0;
    padding: 0;

    overflow-y: scroll;
    list-style: none;
  }
  nav .book-sections li {
    display: flex;
    align-items: center;
    box-sizing: border-box;

    padding: 16pt 2vw 16pt 2vw;

    cursor: pointer;
  }
  nav .book-sections li.active {
    background-color: rgba(0, 0, 0, 0.1);
  }
  nav .book-sections li:hover {
    background-color: rgba(0, 0, 0, 0.05);
  }

  main {
    display: flex;

    width: 76%;
    min-height: calc(100vh - 40pt);

    background-color: #f9f9f9;
  }
</style>
