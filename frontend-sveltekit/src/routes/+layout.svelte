<script>
  let { children } = $props();
  import "../app.css";
</script>

<main class="m-auto p-7 md:mx-96">
  {@render children()}
</main>
