using System;
using FluentAssertions;
using Pehape;
using Xunit;

namespace Tests.String
{
	public class StrContainsTest {
		[InlineData("hello world", "hello", true)]
		[InlineData("hello world", "hello!", false)]
		[InlineData("hellO world", "hello", false)]
		[InlineData("hello world", "", true)]

		// https://github.com/php/php-src/blob/master/ext/standard/tests/strings/str_contains.phpt

		[InlineData("test string", "t s", true)]
		[InlineData("tEst", "test", false)]
		[InlineData("", "", true)]
		[InlineData("", "a", false)]
		[InlineData("a      a", "   ", true)]

		[InlineData("\\\\a", "\\a", true)]
		[InlineData("例の単語", "の単", true)]
		[InlineData("アネバンゲット", "の単", false)]

		[Theory]
		public void CheckIfHaystackContainNeedle(string stringToCheck, string stringToLookFor, bool expectedOutput) {
			PHP.StrContains(stringToCheck, stringToLookFor).Should().Be(expectedOutput);
		}

		[InlineData(typeof(ArgumentNullException), null, "b")]
		[InlineData(typeof(ArgumentNullException), "a", null)]
		[Theory]
		public void ThrowsExceptionWhenInvalidArgumentSupplied(Type exceptionType, string haystack, string needle) {
			new Action(() => PHP.StrContains(haystack, needle)).Should().Throw<Exception>().Which.Should().BeOfType(exceptionType);
		}
	}
}