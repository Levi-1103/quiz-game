<script lang="ts">
  import QuizCard from "./lib/QuizCard.svelte";
  import { NetService } from "./service/net";

  const netService = new NetService();
  netService.connect();
  netService.onPacket((packet: any) => {
    console.log(packet);
  });
  
  let quizzes: { _id: string; name: string }[] = [];
  let code = "";
  const msg = "";

  async function getQuizzes() {
    const response = await fetch("http://localhost:3000/api/quizzes");
    if (!response.ok) {
      console.error(response);
      return;
    }
    const json = await response.json();

    quizzes = json;
    console.log(json);
  }

  function connect() {
    netService.sendPacket({
      id: 0,
      code: "1234",
      name: "testSlop",
    });
  }

  function hostQuiz(quiz) {
    netService.sendPacket({
      id: 1,
      quizId: quiz.id,
    });
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
    <QuizCard quiz={quiz} host={() => hostQuiz(quiz)} />
  {/each}
</main>

<style>
</style>
