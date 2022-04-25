using System;
using System.Diagnostics.CodeAnalysis;
using System.Security.Cryptography;
using System.Text;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// Calculate the md5 hash of a string.
		/// </summary>
		/// <remarks>In PHP's version, there are a second parameter which whould change
		/// the function output format to binary.</remarks>
		/// <param name="str">The string.</param>
		/// <returns>Returns the hash as a 32-character lowercases hexadecimal number.</returns>
		[SuppressMessage("Globalization", "CA1308:Normalize strings to uppercase", Justification = "PHP's md5 output are in lowercases")]
		public static string Md5(string str) {
			// In php, md5 of null are calculated as empty string, we decide to avoid that behaviour in this version
			if (str is null) throw new ArgumentNullException(nameof(str));
			
			byte[] hashBytes = Md5Binary(str);
			return Convert.ToHexString(hashBytes).ToLowerInvariant();
		}

		/// <summary>
		/// Calculate the md5 hash of a string.
		/// </summary>
		/// <param name="str"></param>
		/// <returns>Returns the hash as a 32 byte array</returns>
		[SuppressMessage("Security", "CA5351:Do Not Use Broken Cryptographic Algorithms", Justification = "PHP does have MD5 function")]
		public static byte[] Md5Binary(string str) {
			if (str is null) throw new ArgumentNullException(nameof(str));

			using var md5 = MD5.Create();
			return md5.ComputeHash(Encoding.UTF8.GetBytes(str));
		}
	}
}
