<script>
  import { GetRandomImageUrl } from "../wailsjs/go/main/App.js";
  import { GetHeroList } from "../wailsjs/go/main/App.js";
  import { GetImageUrlsById } from "../wailsjs/go/main/App.js";

  let randomImageUrl = "";
  let heroes = [];
  let photos = [];
  let selectedHero;
  let showRandomPhoto = false;
  let showHeroPhotos = false;

  function init() {
    getHeroList();
  }

  init();

  function getRandomImageUrl() {
    showRandomPhoto = false;
    showHeroPhotos = false;
    GetRandomImageUrl().then((result) => (randomImageUrl = result));
    showRandomPhoto = true;
  }

  function getHeroList() {
    GetHeroList().then((result) => (heroes = result));
  }

  function getImageUrlsById() {
    init();
    showRandomPhoto = false;
    showHeroPhotos = false;
    GetImageUrlsById(selectedHero).then((result) => (photos = result));
    showHeroPhotos = true;
  }
</script>

<h3>Heroes API</h3>
<div>
  <button class="btn" on:click={getRandomImageUrl}>
    Fetch a hero randomly
  </button>
  Click on down arrow to select a breed
  <select bind:value={selectedHero}>
    {#each heroes as hero}
      <option value={hero}>
        {hero}
      </option>
    {/each}
  </select>
  <button class="btn" on:click={getImageUrlsById}>
    Fetch by this hero
  </button>
</div>
<br />
{#if showRandomPhoto}
  <img id="random-photo" src={randomImageUrl} alt="No dog found" />
{/if}
{#if showHeroPhotos}
  {#each photos as photo}
    <img id="hero-photos" src={photo} alt="No dog found" />
  {/each}
{/if}

<style>
  #random-photo {
    width: 600px;
    height: auto;
  }

  #hero-photos {
    width: 300px;
    height: auto;
  }

  .btn:focus {
    border-width: 3px;
  }
</style>