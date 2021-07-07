<script>
  let test = [
    { _id: "35873", title: "amatör-3", parent_id: "35106", is_parent: false },
  ];

  import { onMount } from "svelte";
  import { session } from "$app/stores";
  import Player from "$lib/Player/index.svelte";

  onMount(() => {
    if ($session.navPath.length === 0)
      // updateNav([{ _id: 2, title: "Kitaplar", is_parent: true }]);
      updateNav(test);
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
  <nav class:active={!$session.solution}>
    <ul class="book-sections">
      {#each $session.navItems as item}
        <li
          class:active={item.active}
          on:click={() =>
            item.solution_id
              ? ($session.solution = {
                  ...item,
                  pdf: `/api/proxy?publisher=${
                    $session.publisher
                  }&path=${item.swf.replace(".swf", ".pdf")}`,
                  xml: `/api/proxy?publisher=${$session.publisher}&path=${item.xml}`,
                  audio: `/api/proxy?publisher=${$session.publisher}&path=${item.audio}`,
                })
              : updateNav([...$session.navPath, item])}
        >
          {item.title}
        </li>
      {/each}
    </ul>
  </nav>

  <main class:active={$session.solution}>
    <Player solution={$session.solution} />
    <div class="prev-next">
      <button
        on:click={() => {
          let currentIndex = $session.navItems.findIndex(
            (item) => item.solution_id === $session.solution.solution_id
          );
          let item = $session.navItems[currentIndex - 1];
          if (!item) return;
          $session.solution = {
            ...item,
            pdf: `/api/proxy?publisher=${
              $session.publisher
            }&path=${item.swf.replace(".swf", ".pdf")}`,
            xml: `/api/proxy?publisher=${$session.publisher}&path=${item.xml}`,
            audio: `/api/proxy?publisher=${$session.publisher}&path=${item.audio}`,
          };
        }}
      >
        Önceki
      </button>
      <button
        on:click={() => {
          let currentIndex = $session.navItems.findIndex(
            (item) => item.solution_id === $session.solution.solution_id
          );
          let item = $session.navItems[currentIndex + 1];
          if (!item) return;
          $session.solution = {
            ...item,
            pdf: `/api/proxy?publisher=${
              $session.publisher
            }&path=${item.swf.replace(".swf", ".pdf")}`,
            xml: `/api/proxy?publisher=${$session.publisher}&path=${item.xml}`,
            audio: `/api/proxy?publisher=${$session.publisher}&path=${item.audio}`,
          };
        }}
      >
        Sonraki
      </button>
    </div>
  </main>
</div>

<style>
  .index {
    display: flex;
  }

  nav {
    display: none;
    flex-direction: column;

    width: 100%;
    min-height: calc(100vh - 40pt);

    border-right: 1px solid var(--accents-2);
  }
  nav.active {
    display: flex;
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
    display: none;
    justify-content: space-between;
    flex-direction: column;

    width: 100%;
    min-height: calc(100vh - var(--header-height));

    background-color: #f9f9f9;
  }
  main.active {
    display: flex;
  }
  main .prev-next {
    display: flex;
    flex-shrink: 0;

    border-top: 1px solid var(--accents-3);
  }
  main .prev-next button {
    width: 50%;
    height: var(--header-height);

    margin: 0;
    padding: 0;
    border: 0;
    background-color: white;

    border: 1px solid var(--accents-2);
    border-width: 0 1px 0 1px;
  }

  @media only screen and (min-width: 480px) {
    nav {
      display: flex;
      width: 24%;
    }
    main {
      display: flex;
      width: 76%;
    }
  }
</style>
