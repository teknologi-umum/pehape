using System;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// Returns the portion of string specified by the offset and length parameters.
		/// </summary>
		/// <remarks>
		/// - If offset is negative, the returned string will start at the offset position from the end of string.
		/// - If length is given and is negative, then that many characters will be omitted from the end of string.
		/// </remarks>
		/// <param name="str">The input string.</param>
		/// <param name="offset">Starting position.</param>
		/// <param name="length">(optional) length of the returned substring.</param>
		/// <returns>Extracted part of the input string, or an empty string.</returns>
		public static string Substr(string str, int offset, int? length = null) {
			if (str is null) throw new ArgumentNullException(nameof(str));
			if (offset > str.Length) return string.Empty;
			if (offset < 0 && Math.Abs(offset) > str.Length) return str;
			if (length is 0) return string.Empty;

			// offset is negative, we cut it from the right side (like python index)
			int startIndex = offset;
			if (offset < 0)
				startIndex = str.Length + offset;

			if (length == null)
				return str[startIndex..];

			/* positive length case */

			int cutOffLength = length.Value;
			if (cutOffLength >= 0)
				return str.Substring(startIndex, cutOffLength);

			/* when length is negative, chop off the substring starting from right */

			string temp = str[startIndex..];
			// too much character to omit
			if (temp.Length + cutOffLength < 0)
				return string.Empty;

			return temp.Remove(temp.Length + cutOffLength);
		}
	}
}
