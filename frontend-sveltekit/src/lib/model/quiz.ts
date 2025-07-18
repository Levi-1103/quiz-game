export interface Quiz {
  id: string;
  name: string;
  questions: QuizQuestion[];
}

export interface Player {
  id: string;
  name: string;
}

export interface QuizQuestion {
  id: string;
  name: string;
  choices: QuizChoice[];
}
export interface QuizChoice {
  id: string;
  name: string;
  correct: boolean;
}

export const COLORS = ["bg-primary-950", "bg-secondary-950", "bg-tertiary-950", "bg-warning-950"];
