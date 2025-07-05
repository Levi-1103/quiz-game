import type { QuizQuestion } from "./quiz";

export enum PacketCode {
  Connect = "connect",
  Host = "host",
  QuestionShow = "question",
}

interface ConnectPacketData {
  name: string;
}

interface HostGamePacketData {
  quizId: string;
}

interface QuestionShowPacketData {
  question: QuizQuestion;
}

export type Packet
  = | { code: PacketCode.Connect; data: ConnectPacketData }
    | { code: PacketCode.Host; data: HostGamePacketData }
    | { code: PacketCode.QuestionShow; data: QuestionShowPacketData };
