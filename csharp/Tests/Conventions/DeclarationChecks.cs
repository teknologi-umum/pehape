using FluentAssertions;
using Pehape;
using System.Linq;
using Xunit;

namespace Tests.Conventions {
	public class DeclarationChecks {
		[Fact]
		public void EverythingIsDeclaredInPHPClass() {
			var exportedTypes = typeof(PHP).Assembly.GetTypes()
				.Where(type => type.IsPublic)
				.ToList();
			exportedTypes.Should().ContainSingle(type => type == typeof(PHP), because: "Everything should be declared in one single partial class: PHP");
		}
	}
}
