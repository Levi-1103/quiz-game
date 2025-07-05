<script lang="ts">
  import QuizCard from "./lib/QuizCard.svelte";
  import { PacketCode, type Packet } from "./model/net";
  import type { QuizQuestion } from "./model/quiz";
  import { NetService } from "./service/net";

  let currentQuestion: QuizQuestion | null = null;

  let quizzes: { _id: string; name: string }[] = [];
  let code = "";
  let name = "";
  const msg = "";

  const netService = new NetService();
  netService.connect();
  netService.onPacket((packet: Packet) => {
    // eslint-disable-next-line no-console
    console.log(packet);

    switch (packet.code) {
      case PacketCode.QuestionShow: {
        currentQuestion = packet.data.question;
        break;
      }
    }
  });

  async function getQuizzes() {
    const response = await fetch("http://localhost:3000/api/quizzes");
    if (!response.ok) {
      console.error(response);
      return;
    }
    const json = await response.json();

    quizzes = json;
    // eslint-disable-next-line no-console
    console.log(json);
  }

  function connect() {
    netService.sendPacket({
      code: PacketCode.Connect,
      data: {
        name: name,
        code: code
      },
    });
  }

  function hostQuiz(quiz) {
    netService.sendPacket({
      code: PacketCode.Host,
      data: {
        quizId: quiz.id,
      },
    });
  }
</script>

<main>
  <h1>Quiz App</h1>

  <h1>Hello World</h1>

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
    <input
      class="input"
      type="text"
      placeholder="Player Name"
      bind:value={name}
    />
    <button class="btn preset-filled" onclick={connect}>Join Game </button>
  </div>

  {#each quizzes as quiz}
    <QuizCard {quiz} host={() => hostQuiz(quiz)} />
  {/each}

  
  {#if currentQuestion != null}
  <div class="card p-4 m-2 preset-filled flex flex-col text-center gap-4">
    <h2 class="text-lg p-2">{currentQuestion.name}</h2>
    <div class="flex justify-around gap-4">
    {#each currentQuestion.choices as choice}
      <button class="btn preset-filled-secondary-50-950">{choice.name}</button>
    {/each}
    </div>
    </div>
  {/if}
</main>

<style>
</style>
