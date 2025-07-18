<script lang="ts">
    import HostView from "$lib/components/host/HostView.svelte";

 
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

  <div class="">
    <HostView/>
  </div>
</main>
