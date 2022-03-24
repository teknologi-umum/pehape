using FluentAssertions;
using Pehape;
using System;
using Xunit;

namespace Tests.String {
	public class StrRot13Tests {
		// Source: https://github.com/php/php-src/blob/master/ext/standard/tests/strings/str_rot13_basic.phpt

		// Basic test
		[InlineData("Hello, World!", "Uryyb, Jbeyq!")]
		[InlineData("str_rot13() tests starting", "fge_ebg13() grfgf fgnegvat")]
		[InlineData("abcdefghijklmnopqrstuvwxyz", "nopqrstuvwxyzabcdefghijklm")]
		[InlineData("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "NOPQRSTUVWXYZABCDEFGHIJKLM")]

		// Ensure numeric characters are left untouched
		[InlineData("0123456789", "0123456789")]

		// Ensure non-alphabetic characters are left untouched";
		[InlineData("!%^&*()_-+={}[]:;@~#<,>.?", "!%^&*()_-+={}[]:;@~#<,>.?")]

		// Additional tests

		// Empty string left untouched
		[InlineData("", "")]

		// Whitespaces should left untouched
		[InlineData("    \r\n", "    \r\n")]

		// Non latin alphabet
		[InlineData("地獄、お元気ですか XD", "地獄、お元気ですか KQ")]
		[InlineData("drs الجحيم كيف حالك drs", "qef الجحيم كيف حالك qef")]

		[Theory]
		public void ReturnRot13ShiftedString(string inputString, string expectedOutput) {
			PHP.StrRot13(inputString).Should().Be(expectedOutput);
		}

		[InlineData(typeof(ArgumentNullException), null)]
		[Theory]
		public void ThrowsExceptionWhenInvalidArgumentSupplied(Type exceptionType, string str) {
			new Action(() => PHP.StrRot13(str)).Should().Throw<Exception>().Which.Should().BeOfType(exceptionType);
		}
	}
}
