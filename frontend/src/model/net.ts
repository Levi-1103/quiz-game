import type { QuizQuestion } from "./quiz";

export enum PacketCode {
  Connect = "connect",
  Host = "host",
  QuestionShow = "question",
  ChangeGameState = "state",
}

export enum GameState {
  LobbyState,
  PlayState,
  RevealState,
  EndState,
}

interface ConnectPacketData {
  name: string;
  code: string;
}

interface HostGamePacketData {
  quizId: string;
}

interface QuestionShowPacketData {
  question: QuizQuestion;
}

interface ChangeGameStateData {
  state: GameState;
}

export type Packet
  = | { code: PacketCode.Connect; data: ConnectPacketData }
    | { code: PacketCode.Host; data: HostGamePacketData }
    | { code: PacketCode.QuestionShow; data: QuestionShowPacketData }
    | { code: PacketCode.ChangeGameState; data: ChangeGameStateData };
