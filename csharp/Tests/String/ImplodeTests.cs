using FluentAssertions;
using Pehape;
using Xunit;

namespace Tests.String {
	public class ImplodeTests {
		[Fact]
		public void CanImplodeArrayOfString() {
			var array = new[] { "Hello", "world" };
			PHP.Implode(" ", array).Should().Be("Hello world");
			PHP.Implode("  ", array).Should().Be("Hello  world");
			PHP.Implode("", array).Should().Be("Helloworld");
			PHP.Implode(null, array).Should().Be("Helloworld");
			PHP.Implode(" wkwkwk ", array).Should().Be("Hello wkwkwk world");
			PHP.Implode(' ', array).Should().Be("Hello world");
			PHP.Implode('\n', array).Should().Be("Hello\nworld");
		}

		[Fact]
		public void CanImplodeArrayOfInt() {
			var array = new[] { 1, 3, 2, 4, 5 };
			PHP.Implode(":. ", array).Should().Be("1:. 3:. 2:. 4:. 5");
		}

		record AnyObject(int X);

		[Fact]
		public void CanImplodeArrayOfAnyObject() {
			AnyObject[] array = new[] { new AnyObject(5), new AnyObject(10) };
			PHP.Implode("-", array).Should().Be("AnyObject { X = 5 }-AnyObject { X = 10 }");
		}

		[Fact]
		public void CanImplodeEmptyArray() {
			var array = new string[] { };
			PHP.Implode(" ", array).Should().Be("");
		}

		[Fact]
		public void CanImplodeNestedArrayOfString() {
			var array = new[] { new[] { "Hello", "world" } };
			PHP.Implode(" ", array).Should().Be("Hello world");
			PHP.Implode("  ", array).Should().Be("Hello  world");
			PHP.Implode("", array).Should().Be("Helloworld");
			PHP.Implode(null, array).Should().Be("Helloworld");
			PHP.Implode(" wkwkwk ", array).Should().Be("Hello wkwkwk world");
			PHP.Implode(' ', array).Should().Be("Hello world");
			PHP.Implode('\n', array).Should().Be("Hello\nworld");
		}

		[Fact]
		public void CanImplodeNestedArrayOfInt() {
			var array = new[] { new[] { 1, 2, 3, 4, 5 } };
			PHP.Implode("+", array).Should().Be("1+2+3+4+5");
		}

		[Fact]
		public void CanImplodeNestedArrayOfAnyObject() {
			var array = new[] { new[] { new AnyObject(5), new AnyObject(10) } };
			PHP.Implode("-", array).Should().Be("AnyObject { X = 5 }-AnyObject { X = 10 }");
		}

		[Fact]
		public void CanImplodeNestedEmptyArray() {
			var array = new[] { new string[] { } };
			PHP.Implode(" ", array).Should().Be("");
		}
	}
}