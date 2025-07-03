<script lang="ts">
  import GetQuizzes from "./lib/GetQuizzes.svelte";
  import QuizCard from "./lib/QuizCard.svelte";

  let quizzes: { _id: string; name: string }[] = [];
  let code = "";
  let msg = "";

  async function getQuizzes() {
    let response = await fetch("http://localhost:3000/api/quizzes");
    if (!response.ok) {
      console.error(response);
      return;
    }
    let json = await response.json();

    quizzes = json;
    console.log(json);
  }

  function connect() {
    let websocket = new WebSocket("ws://localhost:3000/ws");
    websocket.onopen = () => {
      console.log("connection opened");
      websocket.send(`join: ${code}`);
    };
    websocket.onmessage = (event) => {
      console.log(event.data);
    };
  }

  function hostQuiz(quiz) {
    let websocket = new WebSocket("ws://localhost:3000/ws");
    websocket.onopen = () => {
      console.log("connection opened");
      websocket.send(`host: ${quiz.id}`);
    };
    websocket.onmessage = (event) => {
      msg = event.data;
    };
  }
</script>

<main>
  <h1>Quiz App</h1>

  <div class="flex gap-4 items-center p-2">
    <button class="btn preset-filled" onclick={getQuizzes}>Get Quizzes </button>
    <p>Message: {msg}</p>
  </div>

  <div class="flex gap-4 items-center p-2">
    <input
      class="input"
      type="text"
      placeholder="Game Code"
      bind:value={code}
    />
    <button class="btn preset-filled" onclick={connect}>Join Game </button>
  </div>

  {#each quizzes as quiz}
    <QuizCard {quiz} host={() => hostQuiz(quiz)} />
  {/each}
</main>

<style>
</style>
