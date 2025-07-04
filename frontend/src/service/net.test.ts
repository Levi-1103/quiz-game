import { expect, it } from "vitest";
import { hello } from "./net";

it("hello", () => {
  const test = hello();

  expect(test).toEqual("hello");
});
