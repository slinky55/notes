<script lang="ts">
  import type { PageData } from './$types';
  import { goto } from "$app/navigation";

  export let data: PageData;

  $: user = data.profile.user;
  $: notes = data.profile.notes;
</script>

<div class="pageContainer">
  <h1>{user.name}</h1>

  <div class="noteView">
    {#each notes as note}
      <div class="note">
        <h2>{note.title}</h2>
        <p>{note.body}</p>
      </div>
    {/each}
  </div>

  <button on:click={() => goto("/new")}>New Note</button>
</div>

<style>
  h1 {
    margin-top: 5rem;
  }

  .pageContainer {
    display: flex;
    flex-direction: column;
    align-items: center;

    width: 100vw;
    height: 100vh;
  }

  .noteView {
    margin-top: 10rem;

    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    grid-gap: 1rem;
  }

  .note {
    border: 1px solid #ccc;
    padding: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    background: #f2ffb4;
  }
</style>