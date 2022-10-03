using System.Globalization;
using System.Text;
using FluentAssertions;
using Pehape;
using Xunit;

namespace Tests.String {
	public class Md5Tests {
		//test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/md5.phpt
		[InlineData("", "d41d8cd98f00b204e9800998ecf8427e")]
		[InlineData("a", "0cc175b9c0f1b6a831c399e269772661")]
		[InlineData("abc", "900150983cd24fb0d6963f7d28e17f72")]
		[InlineData("message digest", "f96b697d7cb7938d525a2f31aaf161d0")]
		[InlineData("abcdefghijklmnopqrstuvwxyz", "c3fcd3d76192e4007dfb496cca67e13b")]
		[InlineData("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			"d174ab98d277d9f5a5611c2c9f419d9f")]
		[InlineData("12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"57edf4a22be3c955ac49da2e2107b67a")]

		//test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/md5_basic1.phpt
		[InlineData("apple", "1f3870be274f6c49b3e31a0c6728957f")]
		[Theory]
		public void ReturnCorrectResult(string input, string result) {
			PHP.Md5(input).Should().Be(result);
		}

		//test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/md5.phpt
		[InlineData("", "d41d8cd98f00b204e9800998ecf8427e")]
		[InlineData("a", "0cc175b9c0f1b6a831c399e269772661")]
		[InlineData("abc", "900150983cd24fb0d6963f7d28e17f72")]
		[InlineData("message digest", "f96b697d7cb7938d525a2f31aaf161d0")]
		[InlineData("abcdefghijklmnopqrstuvwxyz", "c3fcd3d76192e4007dfb496cca67e13b")]
		[InlineData("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			"d174ab98d277d9f5a5611c2c9f419d9f")]
		[InlineData("12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			"57edf4a22be3c955ac49da2e2107b67a")]

		//test cases from https://github.com/php/php-src/blob/master/ext/standard/tests/strings/md5_basic1.phpt
		[InlineData("apple", "1f3870be274f6c49b3e31a0c6728957f")]
		[Theory]
		public void ReturnRawCorrectResult(string input, string result) {
			var raw = PHP.Md5Raw(input);
			var sb = new StringBuilder();
			foreach (var t in raw) {
				sb.Append(t.ToString("X2", CultureInfo.InvariantCulture).ToLowerInvariant());
			}

			raw.Length.Should().Be(16);
			sb.ToString().Should().Be(result);
		}
	}
}