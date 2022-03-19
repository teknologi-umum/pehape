using FluentAssertions;
using Pehape;
using System;
using Xunit;

namespace Tests.String {
	public class LevenshteinTests {
		// equal
		[InlineData(0, "12345", "12345")]

		// first string empty
		[InlineData(3, "", "xyz")]

		// second string empty
		[InlineData(3, "xyz", "")]

		// both empty
		[InlineData(0, "", "")]
		[InlineData(0, "", "", 10, 10, 10)]

		// 1 character 
		[InlineData(1, "1", "2")]

		// 2 character swapped
		[InlineData(2, "12", "21")]

		// inexpensive deletion
		[InlineData(2, "2121", "11", 2)]

		// expensive deletion
		[InlineData(10, "2121", "11", 2, 1, 5)]

		// inexpensive insertion
		[InlineData(2, "11", "2121")]

		// expensive insertion
		[InlineData(10, "11", "2121", 5)]

		// expensive replacement
		[InlineData(3, "111", "121", 2, 3, 2)]

		// very expensive replacement
		[InlineData(4, "111", "121", 2, 9, 2)]

		// test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_error_conditions.phpt
		// levenshtein no longer has a maximum string length limit
		[InlineData(254, "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsu", "A")]
		[InlineData(255, "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuv", "A")]
		[InlineData(254, "A", "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsu")]
		[InlineData(255, "A", "AbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrstuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuvwxyzAbcdefghijklmnopqrtsuv")]

		// test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_bug_16473.phpt
		[InlineData(2, "a", "bc")]
		[InlineData(2, "xa", "xbc")]
		[InlineData(2, "xax", "xbcx")]
		[InlineData(2, "ax", "bcx")]

		// test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_bug_6562.phpt
		[InlineData(1, "debug", "debugg")]
		[InlineData(1, "ddebug", "debug")]
		[InlineData(2, "debbbug", "debug")]
		[InlineData(1, "debugging", "debuging")]

		// test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/levenshtein_bug_7368.phpt
		[InlineData(2, "13458", "12345")]
		[InlineData(2, "1345", "1234")]

		// theory
		[Theory]
		public void ReturnsCorrectDistance(int expectedDistance, string string1, string string2, int insertCost = 1, int replaceCost = 1, int deleteCost = 1) {
			var distance = PHP.Levenshtein(string1, string2, insertCost, replaceCost, deleteCost);
			distance.Should().Be(expectedDistance);
		}

		// string1 is null
		[InlineData(typeof(ArgumentNullException), null, "")]

		// string2 is null
		[InlineData(typeof(ArgumentNullException), "", null)]

		// negative cost
		[InlineData(typeof(ArgumentException), "", "", -1, 0, 0)]
		[InlineData(typeof(ArgumentException), "", "", 0, -1, 0)]
		[InlineData(typeof(ArgumentException), "", "", 0, 0, -1)]

		// theory
		[Theory]
		public void ThrowsExceptionWhenInvalidArgumentSupplied(Type exceptionType, string string1, string string2, int insertCost = 1, int replaceCost = 1, int deleteCost = 1) {
			new Action(() => PHP.Levenshtein(string1, string2, insertCost, replaceCost, deleteCost)).Should().Throw<Exception>().Which.Should().BeOfType(exceptionType);
		}
	}
}
