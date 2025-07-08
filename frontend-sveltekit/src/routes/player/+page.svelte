<script>
    import PlayerView from "$lib/components/PlayerView.svelte";

</script>
<!-- <script lang="ts">
    import QuizCard from "$lib/components/QuizCard.svelte";
    import { GameState, type Packet, PacketCode } from "$lib/model/net";
    import type { Quiz, QuizQuestion } from "$lib/model/quiz";
    import { NetService } from "$lib/services/net";

 

  let currentQuestion: QuizQuestion | null = null;

  let quizzes:Quiz[] = [];
  let code = "";
  let name = "";
  const msg = "";
  let state = -1;
  let host = false;
  let tick = 0;

  async function getQuizzes() {
    const response = await fetch("http://localhost:3000/api/quizzes");
    if (!response.ok) {
      console.error(response);
      return;
    }
    const json = await response.json();

    quizzes = json;
  }

  getQuizzes();

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
      case PacketCode.ChangeGameState: {
        let data = packet.data;
        console.log("Game State", data.state);
        state = data.state;
        break;
      }
      case PacketCode.StartGame: {
        console.log("Start Game");
        state = GameState.PlayState;
        break;
      }
      case PacketCode.Tick: {
        let data = packet.data;
        tick = data.tick
      }
    }
  });

  function startGame() {
    netService.sendPacket({
      code: PacketCode.StartGame,
    });
  }

  

  function connect() {
    netService.sendPacket({
      code: PacketCode.Connect,
      data: {
        name: name,
        code: code,
      },
    });
  }

  function hostQuiz(quiz) {
    host = true;
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

  {#if state === -1}
    <div class="flex gap-4 items-center p-2">
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
  {:else if state === GameState.LobbyState}
    {#if host}
      <button class="btn preset-filled" onclick={startGame}>StartGame</button>
      <p>HOST</p>
      <p>lobby State</p>
    {:else}
      <p>PLAYER</p>
      <p>you have successfully joined</p>
    {/if}
  {:else if state === GameState.PlayState}
    {#if host}
    <p>HOST</p>
      <p>Clock: {tick}</p>
      {#if currentQuestion != null}
        <div class="card p-4 m-2 preset-filled flex flex-col text-center gap-4">
          <h2 class="text-lg p-2">{currentQuestion.name}</h2>
          <div class="flex justify-around gap-4">
            {#each currentQuestion.choices as choice}
              <div class="btn preset-filled-secondary-50-950">
                {choice.name}
              </div>
            {/each}
          </div>
        </div>
      {/if}
    {:else}
    <p>PLAYER</p>
      <p>press correct answer</p>
    {/if}
  {/if}
</main>

<style>
</style> -->

<main>
  <h1>Quiz App</h1>

  <PlayerView/>
</main>
