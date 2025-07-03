<script lang="ts">

  export let quiz: { _id: string; name: string };

  export let host;
</script>

<div class="card p-4 m-2 preset-filled flex justify-between items-center">
  {quiz.name}
  <button class="btn preset-filled-secondary-50-950" on:click={host}>Host</button>
</div>
