using System;

namespace Pehape {
	// ReSharper disable once InconsistentNaming
	public static partial class PHP {
		/// <summary>
		/// The Levenshtein function returns the Levenshtein distance between two strings.
		/// </summary>
		/// <param name="string1">First string to compare.</param>
		/// <param name="string2">Second string to compare.</param>
		/// <param name="insertCost">The cost of inserting a character.</param>
		/// <param name="replaceCost">The cost of replacing a character.</param>
		/// <param name="deleteCost">The cost of deleting a character.</param>
		/// <returns>The Levenshtein distance between two argument strings.</returns>
		/// <exception cref="ArgumentNullException">String1 or string2 is null.</exception>
		/// <exception cref="ArgumentException">InsertCost, replaceCost, or deleteCost is negative.</exception>
		public static int Levenshtein(string string1, string string2, int insertCost = 1, int replaceCost = 1, int deleteCost = 1) {
			// Ported from https://github.com/php/php-src/blob/master/ext/standard/levenshtein.c

			ArgumentNullException.ThrowIfNull(string1);
			ArgumentNullException.ThrowIfNull(string2);
			if (insertCost < 0) throw new ArgumentException("InsertCost cannot be negative.", nameof(insertCost));
			if (replaceCost < 0) throw new ArgumentException("ReplaceCost cannot be negative.", nameof(replaceCost));
			if (deleteCost < 0) throw new ArgumentException("DeleteCost cannot be negative.", nameof(deleteCost));

			// Early return if either of the string is empty
			if (string1.Length == 0) return string2.Length * insertCost;
			if (string2.Length == 0) return string1.Length * deleteCost;

			Span<int> p1 = stackalloc int[string2.Length + 1];
			Span<int> p2 = stackalloc int[string2.Length + 1];
			var swap = false;

			for (var i2 = 0; i2 <= string2.Length; i2++) p1[i2] = i2 * insertCost;
			foreach (var c in string1) {
				(swap ? p1 : p2)[0] = (swap ? p2 : p1)[0] + deleteCost;

				for (var i2 = 0; i2 < string2.Length; i2++) {
					var c0 = (swap ? p2 : p1)[i2] + (c == string2[i2] ? 0 : replaceCost);
					var c1 = (swap ? p2 : p1)[i2 + 1] + deleteCost;
					if (c1 < c0) c0 = c1;
					var c2 = (swap ? p1 : p2)[i2] + insertCost;
					if (c2 < c0) c0 = c2;
					(swap ? p1 : p2)[i2 + 1] = c0;
				}
				swap = !swap;
			}

			return (swap ? p2 : p1)[string2.Length];
		}
	}
}
