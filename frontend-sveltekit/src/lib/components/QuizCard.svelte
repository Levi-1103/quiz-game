<script lang="ts">
  /* eslint-disable import/no-mutable-exports */
  export let quiz;

  export let host;
</script>

<div class="card p-4 m-2 preset-filled flex justify-between items-center">
  {quiz.name}
  <button class="btn preset-filled-secondary-50-950" onclick={host}>Host</button>
</div>
