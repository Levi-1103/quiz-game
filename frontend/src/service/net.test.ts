import { expect, test } from 'vitest';
import { hello } from './net';



test('Hello', () => {
	let test = hello();

	expect(test).toEqual("hello");

});