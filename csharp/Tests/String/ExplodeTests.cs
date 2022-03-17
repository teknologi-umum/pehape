using FluentAssertions;
using Pehape;
using System;
using Xunit;

namespace Tests.String {
	public class ExplodeTests {
		[Fact]
		public void CanExplodeSpaceSeparatedStringBySpace() {
			string str = "Hello pehape world";
			PHP.Explode(" ", str).Should().ContainInOrder("Hello", "pehape", "world");
			PHP.Explode(" ", str, 0).Should().ContainInOrder("Hello");
			PHP.Explode(" ", str, 1).Should().ContainInOrder("Hello");
			PHP.Explode(" ", str, 5).Should().ContainInOrder("Hello", "pehape", "world");
			PHP.Explode(" ", str, -1).Should().ContainInOrder("Hello", "pehape");
			PHP.Explode(" ", str, -3).Should().ContainInOrder();
			PHP.Explode(' ', str).Should().ContainInOrder("Hello", "pehape", "world");
			PHP.Explode(' ', str, 0).Should().ContainInOrder("Hello");
			PHP.Explode(' ', str, 1).Should().ContainInOrder("Hello");
			PHP.Explode(' ', str, 5).Should().ContainInOrder("Hello", "pehape", "world");
			PHP.Explode(' ', str, -1).Should().ContainInOrder("Hello", "pehape");
			PHP.Explode(' ', str, -3).Should().ContainInOrder();
		}

		[Fact]
		public void CanExplodeDoubleSpaceSeparatedStringBySpace() {
			string str = "Hello  pehape  world";
			PHP.Explode(" ", str).Should().ContainInOrder("Hello", "", "pehape", "", "world");
		}

		[Fact]
		public void CanExplodeSpaceSeparatedStringByDoubleSpace() {
			string str = "Hello pehape world";
			PHP.Explode("  ", str).Should().ContainInOrder("Hello pehape world");
		}

		[Fact]
		public void ExplodeSpaceSeparatedStringByEmptyStringThrowsException() {
			string str = "Hello pehape world";
			new Action(() => PHP.Explode("", str)).Should().Throw<ArgumentException>().Which.ParamName.Should().Be("separator");
		}
	}
}
