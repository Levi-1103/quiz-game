<script lang="ts">
  import GetQuizzes from "./lib/GetQuizzes.svelte";

  async function getQuizzes() {
    let response = await fetch("http://localhost:3000/api/quizzes");
    if (!response.ok) {
      console.error(response);
      return;
    }
    let json = await response.json();

    console.log(json);
  }

  function connect() {
    let websocket = new WebSocket("ws://localhost:3000/ws");
    websocket.onopen = ()=> {
      console.log("connection opened")
      websocket.send("HELLO FROM FRONTEND")
    };
    websocket.onmessage = (event) => {
      console.log(event.data)
    }
  }
</script>

<main>
  <h1>Vite + Svelte</h1>

  <button
    class="btn preset-filled"
    onclick={getQuizzes}
    >Get Quizzes
  </button>

  <button
    class="btn preset-filled"
    onclick={connect}
    >Connect to Server
  </button>
</main>

<style>
</style>
