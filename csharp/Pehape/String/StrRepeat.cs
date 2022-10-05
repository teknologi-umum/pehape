using System;
using System.Linq;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// Repeat a string.
		/// </summary>
		/// <remarks>
		/// Times has to be greater than or equal to 0. If the times is set to 0,
		/// the function will return an empty string.
		/// </remarks>
		/// <param name="str">The string to be repeated.</param>
		/// <param name="times">Number of time the string string should be repeated.</param>
		/// <returns>Returns the repeated string.</returns>
		public static string StrRepeat(string str, int times) {
			if (str is null) throw new ArgumentNullException(nameof(str));
			// when	`times` are negative, php would return an empty string but print a warning
			if (times < 0) throw new ArgumentOutOfRangeException(nameof(times), "parameter times cannot be negative");

			if (times == 0) return string.Empty;
			return string.Concat(Enumerable.Repeat(str, times));
		}
	}
}
