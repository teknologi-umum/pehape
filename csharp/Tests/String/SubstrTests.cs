using FluentAssertions;
using Pehape;
using System;
using Xunit;

namespace Tests.String {
	public class SubstrTest {

		// basic offset test
		[InlineData("abcdef", 2, "cdef")]
		[InlineData("abcdef", 0, "abcdef")]

		// negative offset
		[InlineData("abcdef", -1, "f")]
		[InlineData("abcdef", -2, "ef")]
		[InlineData("abcdef", -6, "abcdef")]

		// If string is less than offset characters long, an empty string will be returned.
		[InlineData("abcdef", 10, "")]

		// however.. if it's negative length, it will return the original string
		[InlineData("abcdef", -7, "abcdef")]
		[InlineData("abcdef", -10, "abcdef")]

		// null characters..
		[InlineData("abc\x0xy\x0z", 2, "c\x0xy\x0z")]
		[InlineData("Iñtërnâtiônàlizætiøn", 3, "ërnâtiônàlizætiøn")]

		// empty case
		[InlineData("", 0, "")]
		[InlineData("", 1, "")]

		[Theory]
		public void ReturnSubstringFromGivenOffsetPosition(string inputString, int offset, string expectedOutput) {
			PHP.Substr(inputString, offset).Should().Be(expectedOutput);
		}

		// basic length test
		[InlineData("abcdef", 0, 3, "abc")]

		// If length is given and is negative, then that many characters will be omitted from the end of string
		[InlineData("abcdef", 0, -2, "abcd")]
		[InlineData("abcdef", 3, -1, "de")]
		[InlineData("abcdef", 3, -10, "")]

		// both length and offset is negative
		[InlineData("abcdef", -3, -1, "de")]
		[InlineData("abcdef", -3, -2, "d")]
		[InlineData("abcdef", -3, -3, "")]


		// taken from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/substr.phpt
		// Testing for variations of start and length to point to same element
		[InlineData("abcde" , 2, -2, "c")]
		[InlineData("abcde", -3, -2, "c")]

		// Testing for start > truncation
		[InlineData("abcdef", 4, -4, "")]

		// empty case
		[InlineData("", 0, 0, "")]

		[Theory]
		public void ReturnSubstringWithOffsetAndLength(string inputString, int offset, int length, string expectedOutput) {
			PHP.Substr(inputString, offset, length).Should().Be(expectedOutput);
		}

		[InlineData(typeof(ArgumentNullException), null, 0, null)]
		[InlineData(typeof(ArgumentNullException), null, 0, 2)]
		[InlineData(typeof(ArgumentNullException), null, 5, 10)]
		[Theory]
		public void ThrowsExceptionWhenInvalidArgumentSupplied(Type exceptionType, string inputString, int offset, int? length) {
			new Action(() => PHP.Substr(inputString, offset, length)).Should().Throw<Exception>().Which.Should().BeOfType(exceptionType);
		}
	}
}
