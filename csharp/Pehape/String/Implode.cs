using System.Collections.Generic;
using System.Linq;

namespace Pehape {
	// ReSharper disable once InconsistentNaming
	public static partial class PHP {
		/// <summary>
		/// The Implode function joins array elements with a string.
		/// </summary>
		/// <typeparam name="T">The type of members of array.</typeparam>
		/// <param name="separator">Specifies what to put between the array elements.</param>
		/// <param name="array">The array to join to a string.</param>
		/// <returns>A string from elements of an array</returns>
		public static string Implode<T>(string? separator, IEnumerable<T> array) => string.Join(separator, array);

		/// <summary>
		/// The Implode function joins array elements with a string.
		/// </summary>
		/// <typeparam name="T">The type of members of array.</typeparam>
		/// <param name="separator">Specifies what to put between the array elements.</param>
		/// <param name="array">The array to join to a string.</param>
		/// <returns>A string from elements of an array</returns>
		public static string Implode<T>(string? separator, IEnumerable<T[]> array) =>
			string.Join(separator, array.SelectMany(x => x));

		/// <summary>
		/// The Implode function joins array of elements with a string.
		/// </summary>
		/// <typeparam name="T">The type of members of array.</typeparam>
		/// <param name="separator">Specifies what to put between the array elements.</param>
		/// <param name="array">The array to join to a string.</param>
		/// <returns>A string from elements of an array</returns>
		public static string Implode<T>(char separator, IEnumerable<T> array) => string.Join(separator, array);

		/// <summary>
		/// The Implode function joins array of elements with a string.
		/// </summary>
		/// <typeparam name="T">The type of members of array.</typeparam>
		/// <param name="separator">Specifies what to put between the array elements.</param>
		/// <param name="array">The array to join to a string.</param>
		/// <returns>A string from elements of an array</returns>
		public static string Implode<T>(char separator, IEnumerable<T[]> array) =>
			string.Join(separator, array.SelectMany(x => x));
	}
}
