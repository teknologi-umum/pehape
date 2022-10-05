using System;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// Determine if a string contains a given substring.
		/// </summary>
		/// <remarks>
		/// This function only exist starting from PHP version 8.
		/// </remarks>
		/// <param name="haystack">The string to search in.</param>
		/// <param name="needle">The substring to search for in the haystack.</param>
		/// <returns>Returns true if needle is in haystack, false otherwise.</returns>
		public static bool StrContains(string haystack, string needle) {
			if (haystack is null) throw new ArgumentNullException(nameof(haystack));
			if (needle is null) throw new ArgumentNullException(nameof(needle));

			return haystack.Contains(needle, StringComparison.InvariantCulture);
		}
	}
}
