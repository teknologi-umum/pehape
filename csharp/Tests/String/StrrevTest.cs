using System;
using FluentAssertions;
using Pehape;
using Xunit;

namespace Tests.String
{
	public class StrrevTest {
		[InlineData("hello", "olleh")]
		[InlineData("", "")]
		[InlineData("Hello", "olleH")]
		[InlineData(" a b c d ", " d c b a ")]
		[InlineData("   ", "   ")]
		[InlineData("123454321", "123454321")]
		[InlineData("☆❤world", "dlrow❤☆")]
		[InlineData("トマト", "トマト")]
		[InlineData("!!!", "!!!")]

		// https://stackoverflow.com/questions/228038/best-way-to-reverse-a-string/15111719#15111719
		[InlineData("Les Mise\u0301rables", "selbare\u0301siM seL")]

		// https://mathiasbynens.be/notes/javascript-unicode
		[InlineData("ma\u006E\u0303ana", "ana\u006E\u0303am")]

		[Theory]
		public void ReturnReversedString(string stringToCheck, string expectedOutput) {
			PHP.Strrev(stringToCheck).Should().Be(expectedOutput);
		}

		[InlineData(typeof(ArgumentNullException), null)]
		[Theory]
		public void ThrowsExceptionWhenInvalidArgumentSupplied(Type exceptionType, string str) {
			new Action(() => PHP.Strrev(str)).Should().Throw<Exception>().Which.Should().BeOfType(exceptionType);
		}
	}
}