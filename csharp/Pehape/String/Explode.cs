using System;
using System.Linq;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// The Explode function breaks a string into an array.
		/// </summary>
		/// <param name="separator">Specifies where to break the string.</param>
		/// <param name="str">The string to split.</param>
		/// <param name="limit">Specifies the number of array elements to return.
		/// If limit is greater than 0, then an array with a maximum of limit
		/// elements is returned. If limit is less than 0 then an array excluding
		/// the last -limit elements is returned. If limit is 0 then an array with
		/// one element is returned.</param>
		/// <returns>An array of strings.</returns>
		public static string[] Explode(string? separator, string str, int? limit = null) {
			if (separator is "") throw new ArgumentException("Separator cannot be empty.", nameof(separator));
			ArgumentNullException.ThrowIfNull(str);
			return limit switch {
				null => str.Split(separator),
				> 0 => str.Split(separator).Take(limit.Value).ToArray(),
				< 0 => str.Split(separator).SkipLast(-limit.Value).ToArray(),
				_ => str.Split(separator).Take(1).ToArray()
			};
		}

		/// <summary>
		/// The Explode function breaks a string into an array.
		/// </summary>
		/// <param name="separator">Specifies where to break the string.</param>
		/// <param name="str">The string to split.</param>
		/// <param name="limit">Specifies the number of array elements to return.
		/// If limit is greater than 0, then an array with a maximum of limit
		/// elements is returned. If limit is less than 0 then an array excluding
		/// the last -limit elements is returned. If limit is 0 then an array with
		/// one element is returned.</param>
		/// <returns>An array of strings.</returns>
		public static string[] Explode(char separator, string str, int? limit = null) {
			ArgumentNullException.ThrowIfNull(str);
			return limit switch {
				null => str.Split(separator),
				> 0 => str.Split(separator).Take(limit.Value).ToArray(),
				< 0 => str.Split(separator).SkipLast(-limit.Value).ToArray(),
				_ => str.Split(separator).Take(1).ToArray()
			};
		}
	}
}
