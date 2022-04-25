using System;
using FluentAssertions;
using Pehape;
using Xunit;

namespace Tests.String {
	public class Md5Test {
		[InlineData("Hello World", "b10a8db164e0754105b7a99be72e3fe5")]
		[InlineData("Hello World!", "ed076287532e86365e841e92bfc50d8c")]

		[InlineData("12345678901234567890123456789012345678901234567890123456789012345678901234567890", "57edf4a22be3c955ac49da2e2107b67a")]
		[InlineData("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "d174ab98d277d9f5a5611c2c9f419d9f")]

		[InlineData("", "d41d8cd98f00b204e9800998ecf8427e")]

		[InlineData("Iñtërnâtiônàlizætiøn", "e5e628206e73b1ae69b37fc69762a1e1")]
		[InlineData("اختبار md 5 آجا", "af16b693491fc21bc80f157021c3bb71")]

		[Theory]
		public void ReturnValidMd5Hash(string inputString, string expectedOutput) {
			PHP.Md5(inputString).Should().Be(expectedOutput);
		}

		[InlineData(typeof(ArgumentNullException), null)]
		[Theory]
		public void ThrowsExceptionWhenInvalidArgumentSupplied(Type exceptionType, string inputString) {
			new Action(() => PHP.Md5(inputString)).Should().Throw<Exception>().Which.Should().BeOfType(exceptionType);
		}
	}
}