using System;
using System.Text;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// Performs the ROT13 encoding on the string argument and returns the resulting string.
		/// </summary>
		/// <remarks>
		/// The ROT13 encoding simply shifts every letter by 13 places in the alphabet while 
		/// leaving non-alpha characters untouched. Encoding and decoding are done by the same 
		/// function, passing an encoded string as argument will return the original version.
		/// </remarks>
		/// <param name="str">The input string.</param>
		/// <returns>Returns the ROT13 version of the given string.</returns>
		public static string StrRot13(string str) {
			if (str is null) throw new ArgumentNullException(nameof(str));

			var buffer = new StringBuilder();
			foreach (char ch in str) {
				buffer.Append((char)(ch switch {
					>= 'a' and <= 'z' => ch > 'm' ? ch - 13 : ch + 13,
					>= 'A' and <= 'Z' => ch > 'M' ? ch - 13 : ch + 13,
					_ => ch
				}));
			}
			return buffer.ToString();
		}
	}
}
