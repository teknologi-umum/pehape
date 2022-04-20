using System;
using FluentAssertions;
using Pehape;
using Xunit;

namespace Tests.String
{
	public class StrRepeatTest {
		// https://github.com/php/php-src/blob/master/ext/standard/tests/strings/str_repeat.phpt
		[InlineData("ha", 2,"haha")]
		[InlineData("foo", 4,"foofoofoofoo")]
		[InlineData("foo", 0, "")]
		[InlineData("%0", 3, "%0%0%0")]
		[InlineData("\0", 2, "\0\0")]
		[InlineData("\n\t", 2, "\n\t\n\t")]
		[InlineData("1.23", 3, "1.231.231.23")]
		[InlineData("", 0, "")]
		[InlineData("", 1, "")]
		[InlineData("", 4, "")]
		[InlineData(" ", 0, "")]
		[InlineData(" ", 4, "    ")]
		[InlineData(" _ ", 3, " _  _  _ ")]
		[InlineData("終わり", 3, "終わり終わり終わり")]

		[Theory]
		public void ReturnReversedString(string stringToRepeat, int numberOfRepetition, string expectedOutput) {
			PHP.StrRepeat(stringToRepeat, numberOfRepetition).Should().Be(expectedOutput);
		}

		[InlineData(typeof(ArgumentNullException), null, 3)]
		[InlineData(typeof(ArgumentOutOfRangeException), "oke", -1)]

		[Theory]
		public void ThrowsExceptionWhenInvalidArgumentSupplied(Type exceptionType, string stringToRepeat, int times) {
			new Action(() => PHP.StrRepeat(stringToRepeat, times))
				.Should().Throw<Exception>().Which.Should().BeOfType(exceptionType);
		}
	}
}