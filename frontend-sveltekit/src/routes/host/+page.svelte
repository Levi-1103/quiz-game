<script lang="ts">
    import HostQuizListView from "$lib/components/HostQuizListView.svelte";
    import HostView from "$lib/components/HostView.svelte";
  import { GameState, type Packet, PacketCode } from "$lib/model/net";
  import type { Quiz, QuizQuestion } from "$lib/model/quiz";
  import { NetService } from "$lib/services/net";

  let code = "";
  let name = "";
  const msg = "";
  let state = -1;
  let host = false;
  let tick = 0;

  const netService = new NetService();
  netService.connect();
  netService.onPacket((packet: Packet) => {
    // eslint-disable-next-line no-console
    console.log(packet);

    //   switch (packet.code) {
    //     case PacketCode.QuestionShow: {
    //       currentQuestion = packet.data.question;
    //       break;
    //     }
    //     case PacketCode.ChangeGameState: {
    //       let data = packet.data;
    //       console.log("Game State", data.state);
    //       state = data.state;
    //       break;
    //     }
    //     case PacketCode.StartGame: {
    //       console.log("Start Game");
    //       state = GameState.PlayState;
    //       break;
    //     }
    //     case PacketCode.Tick: {
    //       let data = packet.data;
    //       tick = data.tick
    //     }
    //   }
  });

  function startGame() {
    netService.sendPacket({
      code: PacketCode.StartGame,
    });
  }

  // function connect() {
  //   netService.sendPacket({
  //     code: PacketCode.Connect,
  //     data: {
  //       name: name,
  //       code: code,
  //     },
  //   });
  // }

  // function hostQuiz(quiz) {
  //   host = true;
  //   netService.sendPacket({
  //     code: PacketCode.Host,
  //     data: {
  //       quizId: quiz.id,
  //     },
  //   });
  // }
</script>

<main>
  <h1>Quiz App</h1>

  <HostView/>

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
      <button class="btn preset-filled" onclick={undefined}>Join Game </button>
    </div>
  {:else if state === GameState.LobbyState}
    {#if host}
      <button class="btn preset-filled" onclick={startGame}>StartGame</button>
      <p>HOST</p>
      <p>lobby State</p>
    {/if}
  {:else if state === GameState.PlayState}
    {#if host}
      <p>HOST</p>
      <p>Clock: {tick}</p>
      <!-- {#if currentQuestion != null}
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
      {/if} -->
    {/if}
  {/if}
</main>
