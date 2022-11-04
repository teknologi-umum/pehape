import {echo} from "./echo";

describe("echo", () => {
  it("should log the string parameter", () => {
    const consoleSpy = jest.spyOn(console, "log");
    echo("Hello, World!");
    expect(consoleSpy.mock.calls.length).toStrictEqual(1); // Expect the console.log function to be called 1 time
    expect(consoleSpy.mock.calls[0].length).toStrictEqual(1); // Expect the console.log function to be called using exactly 1 parameter
    expect(consoleSpy.mock.calls[0][0]).toStrictEqual("Hello, World!"); // Expect the parameter to be "Hello, World!"
    consoleSpy.mockClear();
  });

  it("should work with multiple parameters", () => {
    const consoleSpy = jest.spyOn(console, "log");
    echo("Hello, World!", "Good Morning!", "I Am Atomic!"); 
    expect(consoleSpy.mock.calls.length).toStrictEqual(1); // Expect the console.log function to be called 1 time
    expect(consoleSpy.mock.calls[0].length).toStrictEqual(3); // Expect the console.log function to be called using exactly 3 parameters
    expect(consoleSpy.mock.calls[0][0]).toStrictEqual("Hello, World!"); // Expect the first parameter to be "Hello, World!" 
    expect(consoleSpy.mock.calls[0][1]).toStrictEqual("Good Morning!"); // Expect the second parameter to be "Good Morning!" 
    expect(consoleSpy.mock.calls[0][2]).toStrictEqual("I Am Atomic!"); // Expect the third parameter to be "I Am Atomic!" 
    consoleSpy.mockClear();
  });
});